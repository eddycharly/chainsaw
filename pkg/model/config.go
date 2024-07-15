package model

import (
	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	"github.com/kyverno/chainsaw/pkg/apis/v1alpha2"
)

type Configuration struct {
	Cleanup    *v1alpha2.CleanupOptions
	Clusters   v1alpha2.Clusters
	Deletion   v1alpha2.DeletionOptions
	Discovery  v1alpha2.DiscoveryOptions
	Error      *v1alpha2.ErrorOptions
	Execution  *v1alpha2.ExecutionOptions
	Namespace  *v1alpha2.NamespaceOptions
	Report     *v1alpha2.ReportOptions
	Templating v1alpha2.TemplatingOptions
	Timeouts   v1alpha2.Timeouts
}

func From_v1alpha1(config v1alpha1.Configuration) Configuration {
	var c Configuration
	return c
}

func From_v1alpha2(config v1alpha2.Configuration) Configuration {
	var c Configuration
	return c
}
