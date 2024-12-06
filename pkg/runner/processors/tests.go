package processors

import (
	"context"
	"fmt"

	"github.com/kyverno/chainsaw/pkg/cleanup/cleaner"
	"github.com/kyverno/chainsaw/pkg/discovery"
	"github.com/kyverno/chainsaw/pkg/engine"
	"github.com/kyverno/chainsaw/pkg/engine/logging"
	"github.com/kyverno/chainsaw/pkg/engine/namespacer"
	"github.com/kyverno/chainsaw/pkg/model"
	"github.com/kyverno/chainsaw/pkg/runner/failer"
	"github.com/kyverno/chainsaw/pkg/runner/names"
	"github.com/kyverno/chainsaw/pkg/testing"
	"github.com/kyverno/pkg/ext/output/color"
	"k8s.io/utils/clock"
)

type TestsProcessor interface {
	Run(context.Context, engine.Context, ...discovery.Test)
}

func NewTestsProcessor(config model.Configuration, clock clock.PassiveClock) TestsProcessor {
	return &testsProcessor{
		config: config,
		clock:  clock,
	}
}

type testsProcessor struct {
	config model.Configuration
	clock  clock.PassiveClock
}

func (p *testsProcessor) Run(ctx context.Context, tc engine.Context, tests ...discovery.Test) {
	t := testing.FromContext(ctx)
	// setup namespace
	timeouts := tc.Timeouts()
	mainCleaner := cleaner.New(timeouts.Cleanup.Duration, nil, tc.DeletionPropagation())
	t.Cleanup(func() {
		if !mainCleaner.Empty() {
			logging.Log(ctx, logging.Cleanup, logging.BeginStatus, color.BoldFgCyan)
			defer func() {
				logging.Log(ctx, logging.Cleanup, logging.EndStatus, color.BoldFgCyan)
			}()
			for _, err := range mainCleaner.Run(ctx, nil) {
				logging.Log(ctx, logging.Cleanup, logging.ErrorStatus, color.BoldRed, logging.ErrSection(err))
				failer.Fail(ctx)
			}
		}
	})
	var nspacer namespacer.Namespacer
	if p.config.Namespace.Name != "" {
		var nsCleaner cleaner.CleanerCollector
		if !tc.SkipDelete() {
			nsCleaner = mainCleaner
		}
		compilers := tc.Compilers()
		if p.config.Namespace.Compiler != nil {
			compilers = compilers.WithDefaultCompiler(string(*p.config.Namespace.Compiler))
		}
		namespaceData := namespaceData{
			cleaner:   nsCleaner,
			compilers: compilers,
			name:      p.config.Namespace.Name,
			template:  p.config.Namespace.Template,
		}
		nsTc, namespace, err := setupNamespace(ctx, tc, namespaceData)
		if err != nil {
			logging.Log(ctx, logging.Internal, logging.ErrorStatus, color.BoldRed, logging.ErrSection(err))
			tc.IncFailed()
			failer.FailNow(ctx)
		}
		tc = nsTc
		if namespace != nil {
			nspacer = namespacer.New(namespace.GetName())
		}
	}
	// loop through tests
	for i := range tests {
		test := tests[i]
		name, err := names.Test(tc.FullName(), test)
		if err != nil {
			logging.Log(ctx, logging.Internal, logging.ErrorStatus, color.BoldRed, logging.ErrSection(err))
			tc.IncFailed()
			failer.FailNow(ctx)
		}
		// compute test scenarios
		scenarios := applyScenarios(test)
		// loop through test scenarios
		for s := range scenarios {
			test := scenarios[s]
			// run each test scenario in a separate T
			t.Run(name, func(t *testing.T) {
				t.Helper()
				ctx := testing.IntoContext(ctx, t)
				size := len("@chainsaw")
				for i, step := range test.Test.Spec.Steps {
					name := step.Name
					if name == "" {
						name = fmt.Sprintf("step-%d", i+1)
					}
					if size < len(name) {
						size = len(name)
					}
				}
				ctx = logging.IntoContext(ctx, logging.NewLogger(t, p.clock, test.Test.Name, fmt.Sprintf("%-*s", size, "@chainsaw")))
				info := TestInfo{
					Id:         i + 1,
					ScenarioId: s + 1,
					Metadata:   test.Test.ObjectMeta,
				}
				tc := tc.WithBinding(ctx, "test", info)
				t.Cleanup(func() {
					if t.Skipped() {
						tc.IncSkipped()
					} else {
						if t.Failed() {
							tc.IncFailed()
						} else {
							tc.IncPassed()
						}
					}
				})
				// TODO: move into each test processor
				if test.Test.Spec.Concurrent == nil || *test.Test.Spec.Concurrent {
					t.Parallel()
				}
				if test.Test.Spec.Skip != nil && *test.Test.Spec.Skip {
					t.SkipNow()
				}
				if test.Test.Spec.FailFast != nil {
					tc = tc.WithFailFast(ctx, *test.Test.Spec.FailFast)
				}
				if tc.FailFast() && tc.Failed() > 0 {
					t.SkipNow()
				}
				processor := p.createTestProcessor(test, size)
				processor.Run(ctx, nspacer, tc)
			})
		}
	}
}

func (p *testsProcessor) createTestProcessor(test discovery.Test, size int) TestProcessor {
	return NewTestProcessor(
		test,
		size,
		p.clock,
		p.config.Namespace.Template,
		p.config.Namespace.Compiler,
	)
}
