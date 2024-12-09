package runner

import (
	"context"
	"fmt"
	"time"

	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	"github.com/kyverno/chainsaw/pkg/discovery"
	"github.com/kyverno/chainsaw/pkg/engine/clusters"
	"github.com/kyverno/chainsaw/pkg/model"
	"github.com/kyverno/chainsaw/pkg/report"
	enginecontext "github.com/kyverno/chainsaw/pkg/runner/context"
	"github.com/kyverno/chainsaw/pkg/runner/internal"
	"github.com/kyverno/chainsaw/pkg/testing"
	"k8s.io/client-go/rest"
	"k8s.io/utils/clock"
)

type Runner interface {
	Run(context.Context,
		// In test context
		*rest.Config,
		model.Configuration,
		map[string]any,
		//
		...discovery.Test,
	) (model.SummaryResult, error)
}

func New(clock clock.PassiveClock, onFailure func()) Runner {
	return &runner{
		clock:     clock,
		onFailure: onFailure,
	}
}

type runner struct {
	clock     clock.PassiveClock
	onFailure func()
}

func (r *runner) Run(
	ctx context.Context,
	cfg *rest.Config,
	config model.Configuration,
	values map[string]any,
	tests ...discovery.Test,
) (model.SummaryResult, error) {
	return r.run(ctx, cfg, config, nil, values, tests...)
}

func (r *runner) run(
	ctx context.Context,
	cfg *rest.Config,
	config model.Configuration,
	m mainstart,
	values map[string]any,
	tests ...discovery.Test,
) (model.SummaryResult, error) {
	// sanity check
	if len(tests) == 0 {
		return nil, nil
	}
	// setup flags
	// TODO: should be done externally ?
	if err := internal.SetupFlags(config); err != nil {
		return nil, err
	}
	// setup context
	// TODO: should be done externally ?
	tc, err := setupTestContext(ctx, values, cfg, config)
	if err != nil {
		return nil, err
	}
	internalTests := []testing.InternalTest{{
		Name: "chainsaw",
		F: func(t *testing.T) {
			t.Helper()
			t.Parallel()
			// run tests
			r.runTests(ctx, t, config.Namespace, tc, tests...)
		},
	}}
	deps := &internal.TestDeps{}
	if m == nil {
		m = testing.MainStart(deps, internalTests, nil, nil, nil)
	}
	// m.Run() returns:
	// - 0 if everything went well
	// - 1 if some of the tests failed
	// - 2 if running the tests was not possible
	// In our case, we consider an error only when running the tests was not possible.
	// For now, the case where some of the tests failed will be covered by the summary.
	if code := m.Run(); code > 1 {
		return nil, fmt.Errorf("testing framework exited with non zero code %d", code)
	}
	tc.Report.EndTime = time.Now()
	if config.Report != nil && config.Report.Format != "" {
		if err := report.Save(tc.Report, config.Report.Format, config.Report.Path, config.Report.Name); err != nil {
			return tc, err
		}
	}
	return tc, nil
}

func (r *runner) onFail() {
	if r.onFailure != nil {
		r.onFailure()
	}
}

func setupTestContext(ctx context.Context, values any, restConfig *rest.Config, config model.Configuration) (enginecontext.TestContext, error) {
	tc := enginecontext.EmptyContext()
	// cleanup options
	tc = tc.WithSkipDelete(ctx, config.Cleanup.SkipDelete)
	if config.Cleanup.DelayBeforeCleanup != nil {
		tc = tc.WithDelayBeforeCleanup(ctx, &config.Cleanup.DelayBeforeCleanup.Duration)
	}
	// templating options
	tc = tc.WithTemplating(ctx, config.Templating.Enabled)
	if config.Templating.Compiler != nil {
		tc = tc.WithDefaultCompiler(string(*config.Templating.Compiler))
	}
	// discovery options
	tc = tc.WithFullName(ctx, config.Discovery.FullName)
	// execution options
	tc = tc.WithFailFast(ctx, config.Execution.FailFast)
	if config.Execution.ForceTerminationGracePeriod != nil {
		tc = tc.WithTerminationGrace(ctx, &config.Execution.ForceTerminationGracePeriod.Duration)
	}
	// deletion options
	tc = tc.WithDeletionPropagation(ctx, config.Deletion.Propagation)
	// error options
	tc = tc.WithCatch(ctx, config.Error.Catch...)
	// timeouts
	tc = tc.WithTimeouts(ctx, v1alpha1.Timeouts{
		Apply:   &config.Timeouts.Apply,
		Assert:  &config.Timeouts.Assert,
		Cleanup: &config.Timeouts.Cleanup,
		Delete:  &config.Timeouts.Delete,
		Error:   &config.Timeouts.Error,
		Exec:    &config.Timeouts.Exec,
	})
	// values
	tc = enginecontext.WithValues(ctx, tc, values)
	// clusters
	tc = enginecontext.WithClusters(ctx, tc, "", config.Clusters)
	if restConfig != nil {
		cluster, err := clusters.NewClusterFromConfig(restConfig)
		if err != nil {
			return tc, err
		}
		tc = tc.WithCluster(ctx, clusters.DefaultClient, cluster)
		return enginecontext.WithCurrentCluster(ctx, tc, clusters.DefaultClient)
	}
	return tc, nil
}
