package processors

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/jmespath-community/go-jmespath/pkg/binding"
	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	"github.com/kyverno/chainsaw/pkg/client"
	"github.com/kyverno/chainsaw/pkg/discovery"
	"github.com/kyverno/chainsaw/pkg/model"
	"github.com/kyverno/chainsaw/pkg/report"
	apibindings "github.com/kyverno/chainsaw/pkg/runner/bindings"
	"github.com/kyverno/chainsaw/pkg/runner/cleanup"
	"github.com/kyverno/chainsaw/pkg/runner/clusters"
	"github.com/kyverno/chainsaw/pkg/runner/failer"
	"github.com/kyverno/chainsaw/pkg/runner/logging"
	"github.com/kyverno/chainsaw/pkg/runner/mutate"
	"github.com/kyverno/chainsaw/pkg/runner/namespacer"
	"github.com/kyverno/chainsaw/pkg/runner/operations"
	opdelete "github.com/kyverno/chainsaw/pkg/runner/operations/delete"
	"github.com/kyverno/chainsaw/pkg/runner/summary"
	"github.com/kyverno/chainsaw/pkg/runner/timeout"
	"github.com/kyverno/chainsaw/pkg/testing"
	"github.com/kyverno/chainsaw/pkg/utils/kube"
	"github.com/kyverno/pkg/ext/output/color"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/clock"
)

type TestProcessor interface {
	Run(context.Context, model.TestContext, namespacer.Namespacer, discovery.Test)
}

func NewTestProcessor(
	clock clock.PassiveClock,
	summary *summary.Summary,
	report *report.TestReport,
	shouldFailFast *atomic.Bool,
) TestProcessor {
	return &testProcessor{
		clock:          clock,
		summary:        summary,
		report:         report,
		shouldFailFast: shouldFailFast,
	}
}

type testProcessor struct {
	clock          clock.PassiveClock
	summary        *summary.Summary
	report         *report.TestReport
	shouldFailFast *atomic.Bool
}

func (p *testProcessor) Run(ctx context.Context, tc model.TestContext, nspacer namespacer.Namespacer, test discovery.Test) {
	t := testing.FromContext(ctx)
	config := tc.Configuration()
	bindings := tc.Bindings()
	timeouts := config.Timeouts.Combine(test.Test.Spec.Timeouts)
	if p.report != nil {
		p.report.SetStartTime(time.Now())
		t.Cleanup(func() {
			p.report.SetEndTime(time.Now())
			if t.Failed() {
				p.report.Fail()
			}
			if t.Skipped() {
				p.report.Skip()
			}
		})
	}
	size := len("@cleanup")
	for i, step := range test.Test.Spec.Steps {
		name := step.Name
		if name == "" {
			name = fmt.Sprintf("step-%d", i+1)
		}
		if size < len(name) {
			size = len(name)
		}
	}
	if p.summary != nil {
		t.Cleanup(func() {
			if t.Skipped() {
				p.summary.IncSkipped()
			} else {
				if t.Failed() {
					p.summary.IncFailed()
				} else {
					p.summary.IncPassed()
				}
			}
		})
	}
	if test.Test.Spec.Concurrent == nil || *test.Test.Spec.Concurrent {
		t.Parallel()
	}
	if test.Test.Spec.Skip != nil && *test.Test.Spec.Skip {
		t.SkipNow()
	}
	if config.Execution.FailFast {
		if p.shouldFailFast.Load() {
			t.SkipNow()
		}
	}
	registeredClusters := clusters.Register(tc.Clusters(), test.BasePath, test.Test.Spec.Clusters)
	clusterConfig, clusterClient, err := registeredClusters.Resolve(false, test.Test.Spec.Cluster)
	if err != nil {
		logging.Log(ctx, logging.Internal, logging.ErrorStatus, color.BoldRed, logging.ErrSection(err))
		failer.FailNow(ctx)
	}
	bindings = apibindings.RegisterClusterBindings(ctx, bindings, clusterConfig, clusterClient)
	setupLogger := logging.NewLogger(t, p.clock, test.Test.Name, fmt.Sprintf("%-*s", size, "@setup"))
	cleanupLogger := logging.NewLogger(t, p.clock, test.Test.Name, fmt.Sprintf("%-*s", size, "@cleanup"))
	var namespace *corev1.Namespace
	if clusterClient != nil {
		if nspacer == nil || test.Test.Spec.Namespace != "" {
			var ns corev1.Namespace
			if test.Test.Spec.Namespace != "" {
				ns = kube.Namespace(test.Test.Spec.Namespace)
			} else {
				ns = kube.PetNamespace()
			}
			namespace = &ns
		}
		if namespace != nil {
			object := kube.ToUnstructured(namespace)
			bindings = apibindings.RegisterNamedBinding(ctx, bindings, "namespace", object.GetName())
			if test.Test.Spec.NamespaceTemplate != nil && test.Test.Spec.NamespaceTemplate.Value != nil {
				template := v1alpha1.Any{
					Value: test.Test.Spec.NamespaceTemplate.Value,
				}
				if merged, err := mutate.Merge(ctx, object, bindings, template); err != nil {
					failer.FailNow(ctx)
				} else {
					object = merged
				}
				bindings = apibindings.RegisterNamedBinding(ctx, bindings, "namespace", object.GetName())
			} else if config.Namespace.Template != nil && config.Namespace.Template.Value != nil {
				template := v1alpha1.Any{
					Value: config.Namespace.Template.Value,
				}
				if merged, err := mutate.Merge(ctx, object, bindings, template); err != nil {
					failer.FailNow(ctx)
				} else {
					object = merged
				}
				bindings = apibindings.RegisterNamedBinding(ctx, bindings, "namespace", object.GetName())
			}
			nspacer = namespacer.New(clusterClient, object.GetName())
			setupCtx := logging.IntoContext(ctx, setupLogger)
			cleanupCtx := logging.IntoContext(ctx, cleanupLogger)
			if err := clusterClient.Get(setupCtx, client.ObjectKey(&object), object.DeepCopy()); err != nil {
				if !errors.IsNotFound(err) {
					// Get doesn't log
					setupLogger.Log(logging.Get, logging.ErrorStatus, color.BoldRed, logging.ErrSection(err))
					failer.FailNow(ctx)
				}
				if !cleanup.Skip(config.Cleanup.SkipDelete, test.Test.Spec.SkipDelete, nil) {
					t.Cleanup(func() {
						operation := newOperation(
							model.OperationInfo{},
							false,
							timeout.Get(nil, timeouts.CleanupDuration()),
							func(ctx context.Context, bindings binding.Bindings) (operations.Operation, binding.Bindings, error) {
								bindings = apibindings.RegisterClusterBindings(ctx, bindings, clusterConfig, clusterClient)
								return opdelete.New(clusterClient, object, nspacer, false, metav1.DeletePropagationBackground), bindings, nil
							},
							nil,
						)
						operation.execute(cleanupCtx, bindings)
					})
				}
				if err := clusterClient.Create(logging.IntoContext(setupCtx, setupLogger), object.DeepCopy()); err != nil {
					failer.FailNow(ctx)
				}
			}
		}
	}
	if p.report != nil && nspacer != nil {
		p.report.SetNamespace(nspacer.GetNamespace())
	}
	bindings, err = apibindings.RegisterBindings(ctx, bindings, test.Test.Spec.Bindings...)
	if err != nil {
		logging.Log(ctx, logging.Internal, logging.ErrorStatus, color.BoldRed, logging.ErrSection(err))
		failer.FailNow(ctx)
	}
	for i, step := range test.Test.Spec.Steps {
		processor := p.createStepProcessor(
			config,
			nspacer,
			registeredClusters,
			test,
			step,
		)
		name := step.Name
		if name == "" {
			name = fmt.Sprintf("step-%d", i+1)
		}
		processor.Run(
			logging.IntoContext(ctx, logging.NewLogger(t, p.clock, test.Test.Name, fmt.Sprintf("%-*s", size, name))),
			apibindings.RegisterNamedBinding(ctx, bindings, "step", model.StepInfo{Id: i + 1}),
		)
	}
}

func (p *testProcessor) createStepProcessor(config model.Configuration, nspacer namespacer.Namespacer, clusters clusters.Registry, test discovery.Test, step v1alpha1.TestStep) StepProcessor {
	var report *report.StepReport
	if p.report != nil {
		report = p.report.ForStep(&step)
	}
	return NewStepProcessor(config, clusters, nspacer, p.clock, test, step, report)
}
