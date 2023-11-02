package runner

import (
	"context"
	"fmt"
	"testing"

	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	"github.com/kyverno/chainsaw/pkg/client"
	"github.com/kyverno/chainsaw/pkg/discovery"
	runnerclient "github.com/kyverno/chainsaw/pkg/runner/client"
	"github.com/kyverno/chainsaw/pkg/runner/logging"
	"github.com/kyverno/chainsaw/pkg/runner/names"
	"github.com/kyverno/chainsaw/pkg/runner/summary"
	"github.com/kyverno/kyverno/ext/output/color"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/utils/clock"
)

type testsRunner struct {
	config  v1alpha1.ConfigurationSpec
	client  client.Client
	clock   clock.PassiveClock
	summary *summary.Summary
}

type testRunner struct {
	config  v1alpha1.ConfigurationSpec
	client  client.Client
	clock   clock.PassiveClock
	summary *summary.Summary
}

func (r *testRunner) runTest(t *testing.T, test discovery.Test) {
	t.Helper()
	if r.summary != nil {
		t.Cleanup(func() {
			if t.Skipped() {
				r.summary.IncSkipped()
			} else {
				if t.Failed() {
					r.summary.IncFailed()
				} else {
					r.summary.IncPassed()
				}
			}
		})
	}
	if test.Spec.Concurrent == nil || *test.Spec.Concurrent {
		t.Parallel()
	}
	if test.Spec.Skip != nil && *test.Spec.Skip {
		t.SkipNow()
	}
}

func (r *testsRunner) runTests(t *testing.T, tests ...discovery.Test) {
	t.Helper()
	logger := logging.NewLogger(t, r.clock, t.Name(), "@main")
	// ctx := Context{
	// 	clock: r.clock,
	// 	clientFactory: func(logger logging.Logger) client.Client {
	// 		return runnerclient.New(logger, r.client)
	// 	},
	// }
	fmt.Println(r.config.Namespace)
	if r.config.Namespace != "" {
		c := runnerclient.New(logger, r.client)
		namespace := client.Namespace(r.config.Namespace)
		if err := c.Get(context.Background(), client.ObjectKey(&namespace), namespace.DeepCopy()); err != nil {
			if !errors.IsNotFound(err) {
				// Get doesn't log
				logger.Log("GET   ", color.BoldRed, err)
				t.FailNow()
			}
			t.Cleanup(func() {
				// TODO: wait ?
				if err := c.Delete(context.Background(), &namespace); err != nil {
					t.FailNow()
				}
			})
			if err := c.Create(context.Background(), namespace.DeepCopy()); err != nil {
				t.FailNow()
			}
		}
		// 	ctx.namespacer = namespacer.New(c, r.config.Namespace)
	}
	for _, test := range tests {
		name, err := names.Test(r.config, test)
		if err != nil {
			logger.Log("INTERN", color.BoldRed, err)
			t.FailNow()
		}
		t.Run(name, func(t *testing.T) {
			t.Helper()
			runner := &testRunner{
				config:  r.config,
				client:  r.client,
				clock:   r.clock,
				summary: r.summary,
			}
			runner.runTest(t, test)
		})
	}

	// test := tests[i]
	// size := 0
	// for i, step := range test.Spec.Steps {
	// 	name := step.Name
	// 	if name == "" {
	// 		name = fmt.Sprintf("step-%d", i+1)
	// 	}
	// 	if size < len(name) {
	// 		size = len(name)
	// 	}
	// }
	// t.Run(name, func(t *testing.T) {
	// 	t.Helper()
	// 	if r.summary != nil {
	// 		t.Cleanup(func() {
	// 			if t.Skipped() {
	// 				r.summary.IncSkipped()
	// 			} else {
	// 				if t.Failed() {
	// 					r.summary.IncFailed()
	// 				} else {
	// 					r.summary.IncPassed()
	// 				}
	// 			}
	// 		})
	// 	}
	// 	if test.Spec.Concurrent == nil || *test.Spec.Concurrent {
	// 		t.Parallel()
	// 	}
	// 	if test.Spec.Skip != nil && *test.Spec.Skip {
	// 		t.SkipNow()
	// 	}
	// 	beginLogger := logging.NewLogger(t, r.clock, test.Name, fmt.Sprintf("%-*s", size, "@begin"))
	// 	// cleanLogger := logging.NewLogger(t, clock, test.Name, fmt.Sprintf("%-*s", size, "@clean"))
	// 	ctx := ctx
	// 	if test.Spec.Namespace != "" {
	// 		namespace := client.Namespace(test.Spec.Namespace)
	// 		c := ctx.clientFactory(beginLogger)
	// 		if err := c.Get(context.Background(), client.ObjectKey(&namespace), namespace.DeepCopy()); err != nil {
	// 			if !errors.IsNotFound(err) {
	// 				// Get doesn't log
	// 				beginLogger.Log("GET   ", color.BoldRed, err)
	// 				t.FailNow()
	// 			}
	// 			if err := c.Create(context.Background(), namespace.DeepCopy()); err != nil {
	// 				t.FailNow()
	// 			}
	// 			// TODO
	// 			t.Cleanup(func() {
	// 				// if err := operations.Delete(context.Background(), cleanLogger, &namespace, ctx.clientFactory(cleanLogger)); err != nil {
	// 				// 	t.FailNow()
	// 				// }
	// 			})
	// 		}
	// 		ctx.namespacer = namespacer.New(c, test.Spec.Namespace)
	// 	}
	// 	if ctx.namespacer == nil {
	// 		namespace := client.PetNamespace()
	// 		if err := ctx.clientFactory(beginLogger).Create(context.Background(), namespace.DeepCopy()); err != nil {
	// 			t.FailNow()
	// 		}
	// 		t.Cleanup(func() {
	// 			// TODO
	// 			// if err := operations.Delete(context.Background(), cleanLogger, &namespace, ctx.clientFactory(cleanLogger)); err != nil {
	// 			// 	t.FailNow()
	// 			// }
	// 		})
	// 		ctx.namespacer = namespacer.New(r.client, namespace.Name)
	// 	}
	// 	runTest(t, ctx, r.config, test, size)
	// })
}
