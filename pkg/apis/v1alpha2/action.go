package v1alpha2

// import (
// 	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
// 	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
// )

// // FileReference represents a file reference.
// type FileReference struct {
// 	// File is the path to the referenced file. This can be a direct path to a file
// 	// or an expression that matches multiple files, such as "manifest/*.yaml" for all YAML
// 	// files within the "manifest" directory.
// 	File string `json:"file,omitempty"`
// }

// // ResourceReference represents a file reference or resource.
// type ResourceReference struct {
// 	// Resource provides a resource to be applied.
// 	// +kubebuilder:validation:XEmbeddedResource
// 	// +kubebuilder:pruning:PreserveUnknownFields
// 	// +optional
// 	Resource *unstructured.Unstructured `json:"resource,omitempty"`
// }

// type FileOrResourceReference struct {
// 	FileReference     `json:",inline"`
// 	ResourceReference `json:",inline"`
// }

// type ActionTimeout struct {
// 	// Timeout for the operation. Overrides the global timeout set in the Configuration.
// 	// +optional
// 	Timeout *metav1.Duration `json:"timeout,omitempty"`
// }

// type ActionExpectations struct {
// 	// Expect defines a list of matched checks to validate the operation outcome.
// 	// +optional
// 	Expect []Expectation `json:"expect,omitempty"`
// }

// // Apply represents a set of configurations or resources that
// // should be applied during testing.
// type Apply struct {
// 	// ActionTimeout specifies the action timeout.
// 	ActionTimeout `json:",inline"`

// 	// FileRefOrResource provides a reference to the resources to be applied.
// 	FileOrResourceReference `json:",inline"`

// 	// DryRun determines whether the file should be applied in dry run mode.
// 	// +optional
// 	DryRun *bool `json:"dryRun,omitempty"`

// 	ActionExpectations `json:",inline"`
// }

// // // Assert represents a test condition that is expected to hold true
// // // during the testing process.
// // type Assert struct {
// // 	// Timeout for the operation. Overrides the global timeout set in the Configuration.
// // 	// +optional
// // 	Timeout *metav1.Duration `json:"timeout,omitempty"`

// // 	// Bindings defines additional binding key/values.
// // 	// +optional
// // 	Bindings []Binding `json:"bindings,omitempty"`

// // 	// Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).
// // 	// +optional
// // 	Cluster string `json:"cluster,omitempty"`

// // 	// Clusters holds a registry to clusters to support multi-cluster tests.
// // 	// +optional
// // 	Clusters Clusters `json:"clusters,omitempty"`

// // 	// FileRefOrAssert provides a reference to the assertion.
// // 	FileRefOrCheck `json:",inline"`

// // 	// Template determines whether resources should be considered for templating.
// // 	// +optional
// // 	Template *bool `json:"template,omitempty"`
// // }

// // // Command describes a command to run as a part of a test step.
// // type Command struct {
// // 	// Timeout for the operation. Overrides the global timeout set in the Configuration.
// // 	// +optional
// // 	Timeout *metav1.Duration `json:"timeout,omitempty"`

// // 	// Bindings defines additional binding key/values.
// // 	// +optional
// // 	Bindings []Binding `json:"bindings,omitempty"`

// // 	// Outputs defines output bindings.
// // 	// +optional
// // 	Outputs []Output `json:"outputs,omitempty"`

// // 	// Env defines additional environment variables.
// // 	// +optional
// // 	Env []Binding `json:"env,omitempty"`

// // 	// Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).
// // 	// +optional
// // 	Cluster string `json:"cluster,omitempty"`

// // 	// Clusters holds a registry to clusters to support multi-cluster tests.
// // 	// +optional
// // 	Clusters Clusters `json:"clusters,omitempty"`

// // 	// Entrypoint is the command entry point to run.
// // 	Entrypoint string `json:"entrypoint"`

// // 	// Args is the command arguments.
// // 	// +optional
// // 	Args []string `json:"args,omitempty"`

// // 	// SkipLogOutput removes the output from the command. Useful for sensitive logs or to reduce noise.
// // 	// +optional
// // 	SkipLogOutput bool `json:"skipLogOutput,omitempty"`

// // 	// Check is an assertion tree to validate the operation outcome.
// // 	// +optional
// // 	Check *Check `json:"check,omitempty"`
// // }

// // // Create represents a set of resources that should be created.
// // // If a resource already exists in the cluster it will fail.
// // type Create struct {
// // 	// Timeout for the operation. Overrides the global timeout set in the Configuration.
// // 	// +optional
// // 	Timeout *metav1.Duration `json:"timeout,omitempty"`

// // 	// Bindings defines additional binding key/values.
// // 	// +optional
// // 	Bindings []Binding `json:"bindings,omitempty"`

// // 	// Outputs defines output bindings.
// // 	// +optional
// // 	Outputs []Output `json:"outputs,omitempty"`

// // 	// Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).
// // 	// +optional
// // 	Cluster string `json:"cluster,omitempty"`

// // 	// Clusters holds a registry to clusters to support multi-cluster tests.
// // 	// +optional
// // 	Clusters Clusters `json:"clusters,omitempty"`

// // 	// FileRefOrResource provides a reference to the file containing the resources to be created.
// // 	FileRefOrResource `json:",inline"`

// // 	// Template determines whether resources should be considered for templating.
// // 	// +optional
// // 	Template *bool `json:"template,omitempty"`

// // 	// DryRun determines whether the file should be applied in dry run mode.
// // 	// +optional
// // 	DryRun *bool `json:"dryRun,omitempty"`

// // 	// Expect defines a list of matched checks to validate the operation outcome.
// // 	// +optional
// // 	Expect []Expectation `json:"expect,omitempty"`
// // }

// // // Delete is a reference to an object that should be deleted
// // type Delete struct {
// // 	// Timeout for the operation. Overrides the global timeout set in the Configuration.
// // 	// +optional
// // 	Timeout *metav1.Duration `json:"timeout,omitempty"`

// // 	// Bindings defines additional binding key/values.
// // 	// +optional
// // 	Bindings []Binding `json:"bindings,omitempty"`

// // 	// Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).
// // 	// +optional
// // 	Cluster string `json:"cluster,omitempty"`

// // 	// Clusters holds a registry to clusters to support multi-cluster tests.
// // 	// +optional
// // 	Clusters Clusters `json:"clusters,omitempty"`

// // 	// Template determines whether resources should be considered for templating.
// // 	// +optional
// // 	Template *bool `json:"template,omitempty"`

// // 	// File is the path to the referenced file. This can be a direct path to a file
// // 	// or an expression that matches multiple files, such as "manifest/*.yaml" for all YAML
// // 	// files within the "manifest" directory.
// // 	// +optional
// // 	File string `json:"file,omitempty"`

// // 	// Ref determines objects to be deleted.
// // 	// +optional
// // 	Ref *ObjectReference `json:"ref,omitempty"`

// // 	// Expect defines a list of matched checks to validate the operation outcome.
// // 	// +optional
// // 	Expect []Expectation `json:"expect,omitempty"`

// // 	// DeletionPropagationPolicy decides if a deletion will propagate to the dependents of
// // 	// the object, and how the garbage collector will handle the propagation.
// // 	// Overrides the deletion propagation policy set in the Configuration, the Test and the TestStep.
// // 	// +optional
// // 	// +kubebuilder:validation:Enum:=Orphan;Background;Foreground
// // 	DeletionPropagationPolicy *metav1.DeletionPropagation `json:"deletionPropagationPolicy,omitempty"`
// // }

// // // Describe defines how to describe resources.
// // type Describe struct {
// // 	// Timeout for the operation. Overrides the global timeout set in the Configuration.
// // 	// +optional
// // 	Timeout *metav1.Duration `json:"timeout,omitempty"`

// // 	// Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).
// // 	// +optional
// // 	Cluster string `json:"cluster,omitempty"`

// // 	// Clusters holds a registry to clusters to support multi-cluster tests.
// // 	// +optional
// // 	Clusters Clusters `json:"clusters,omitempty"`

// // 	// ResourceReference referenced resource type.
// // 	ResourceReference `json:",inline"`

// // 	// ObjectLabelsSelector determines the selection process of referenced objects.
// // 	// +optional
// // 	ObjectLabelsSelector `json:",inline"`

// // 	// Show Events indicates whether to include related events.
// // 	// +optional
// // 	ShowEvents *bool `json:"showEvents,omitempty"`
// // }

// // // Error represents an anticipated error condition that may arise during testing.
// // // Instead of treating such an error as a test failure, it acknowledges it as expected.
// // type Error struct {
// // 	// Timeout for the operation. Overrides the global timeout set in the Configuration.
// // 	// +optional
// // 	Timeout *metav1.Duration `json:"timeout,omitempty"`

// // 	// Bindings defines additional binding key/values.
// // 	// +optional
// // 	Bindings []Binding `json:"bindings,omitempty"`

// // 	// Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).
// // 	// +optional
// // 	Cluster string `json:"cluster,omitempty"`

// // 	// Clusters holds a registry to clusters to support multi-cluster tests.
// // 	// +optional
// // 	Clusters Clusters `json:"clusters,omitempty"`

// // 	// FileRefOrAssert provides a reference to the expected error.
// // 	FileRefOrCheck `json:",inline"`

// // 	// Template determines whether resources should be considered for templating.
// // 	// +optional
// // 	Template *bool `json:"template,omitempty"`
// // }

// // // Events defines how to collect events.
// // type Events struct {
// // 	// Timeout for the operation. Overrides the global timeout set in the Configuration.
// // 	// +optional
// // 	Timeout *metav1.Duration `json:"timeout,omitempty"`

// // 	// Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).
// // 	// +optional
// // 	Cluster string `json:"cluster,omitempty"`

// // 	// Clusters holds a registry to clusters to support multi-cluster tests.
// // 	// +optional
// // 	Clusters Clusters `json:"clusters,omitempty"`

// // 	// ObjectLabelsSelector determines the selection process of referenced objects.
// // 	// +optional
// // 	ObjectLabelsSelector `json:",inline"`

// // 	// Format determines the output format (json or yaml).
// // 	// +optional
// // 	Format Format `json:"format,omitempty"`
// // }

// // // Get defines how to get resources.
// // type Get struct {
// // 	// Timeout for the operation. Overrides the global timeout set in the Configuration.
// // 	// +optional
// // 	Timeout *metav1.Duration `json:"timeout,omitempty"`

// // 	// Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).
// // 	// +optional
// // 	Cluster string `json:"cluster,omitempty"`

// // 	// Clusters holds a registry to clusters to support multi-cluster tests.
// // 	// +optional
// // 	Clusters Clusters `json:"clusters,omitempty"`

// // 	// ResourceReference referenced resource type.
// // 	ResourceReference `json:",inline"`

// // 	// ObjectLabelsSelector determines the selection process of referenced objects.
// // 	// +optional
// // 	ObjectLabelsSelector `json:",inline"`

// // 	// Format determines the output format (json or yaml).
// // 	// +optional
// // 	Format Format `json:"format,omitempty"`
// // }

// // // Patch represents a set of resources that should be patched.
// // // If a resource doesn't exist yet in the cluster it will fail.
// // type Patch struct {
// // 	// Timeout for the operation. Overrides the global timeout set in the Configuration.
// // 	// +optional
// // 	Timeout *metav1.Duration `json:"timeout,omitempty"`

// // 	// Bindings defines additional binding key/values.
// // 	// +optional
// // 	Bindings []Binding `json:"bindings,omitempty"`

// // 	// Outputs defines output bindings.
// // 	// +optional
// // 	Outputs []Output `json:"outputs,omitempty"`

// // 	// Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).
// // 	// +optional
// // 	Cluster string `json:"cluster,omitempty"`

// // 	// Clusters holds a registry to clusters to support multi-cluster tests.
// // 	// +optional
// // 	Clusters Clusters `json:"clusters,omitempty"`

// // 	// FileRefOrResource provides a reference to the file containing the resources to be patched.
// // 	FileRefOrResource `json:",inline"`

// // 	// Template determines whether resources should be considered for templating.
// // 	// +optional
// // 	Template *bool `json:"template,omitempty"`

// // 	// DryRun determines whether the file should be applied in dry run mode.
// // 	// +optional
// // 	DryRun *bool `json:"dryRun,omitempty"`

// // 	// Expect defines a list of matched checks to validate the operation outcome.
// // 	// +optional
// // 	Expect []Expectation `json:"expect,omitempty"`
// // }

// // // PodLogs defines how to collect pod logs.
// // type PodLogs struct {
// // 	// Timeout for the operation. Overrides the global timeout set in the Configuration.
// // 	// +optional
// // 	Timeout *metav1.Duration `json:"timeout,omitempty"`

// // 	// Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).
// // 	// +optional
// // 	Cluster string `json:"cluster,omitempty"`

// // 	// Clusters holds a registry to clusters to support multi-cluster tests.
// // 	// +optional
// // 	Clusters Clusters `json:"clusters,omitempty"`

// // 	// ObjectLabelsSelector determines the selection process of referenced objects.
// // 	// +optional
// // 	ObjectLabelsSelector `json:",inline"`

// // 	// Container in pod to get logs from else --all-containers is used.
// // 	// +optional
// // 	Container string `json:"container,omitempty"`

// // 	// Tail is the number of last lines to collect from pods. If omitted or zero,
// // 	// then the default is 10 if you use a selector, or -1 (all) if you use a pod name.
// // 	// This matches default behavior of `kubectl logs`.
// // 	// +optional
// // 	Tail *int `json:"tail,omitempty"`
// // }

// // // Script describes a script to run as a part of a test step.
// // type Script struct {
// // 	// Timeout for the operation. Overrides the global timeout set in the Configuration.
// // 	// +optional
// // 	Timeout *metav1.Duration `json:"timeout,omitempty"`

// // 	// Bindings defines additional binding key/values.
// // 	// +optional
// // 	Bindings []Binding `json:"bindings,omitempty"`

// // 	// Outputs defines output bindings.
// // 	// +optional
// // 	Outputs []Output `json:"outputs,omitempty"`

// // 	// Env defines additional environment variables.
// // 	// +optional
// // 	Env []Binding `json:"env,omitempty"`

// // 	// Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).
// // 	// +optional
// // 	Cluster string `json:"cluster,omitempty"`

// // 	// Clusters holds a registry to clusters to support multi-cluster tests.
// // 	// +optional
// // 	Clusters Clusters `json:"clusters,omitempty"`

// // 	// Content defines a shell script (run with "sh -c ...").
// // 	// +optional
// // 	Content string `json:"content,omitempty"`

// // 	// SkipLogOutput removes the output from the command. Useful for sensitive logs or to reduce noise.
// // 	// +optional
// // 	SkipLogOutput bool `json:"skipLogOutput,omitempty"`

// // 	// Check is an assertion tree to validate the operation outcome.
// // 	// +optional
// // 	Check *Check `json:"check,omitempty"`
// // }

// // // Sleep represents a duration while nothing happens.
// // type Sleep struct {
// // 	// Duration is the delay used for sleeping.
// // 	Duration metav1.Duration `json:"duration"`
// // }

// // // Update represents a set of resources that should be updated.
// // // If a resource does not exist in the cluster it will fail.
// // type Update struct {
// // 	// Timeout for the operation. Overrides the global timeout set in the Configuration.
// // 	// +optional
// // 	Timeout *metav1.Duration `json:"timeout,omitempty"`

// // 	// Bindings defines additional binding key/values.
// // 	// +optional
// // 	Bindings []Binding `json:"bindings,omitempty"`

// // 	// Outputs defines output bindings.
// // 	// +optional
// // 	Outputs []Output `json:"outputs,omitempty"`

// // 	// Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).
// // 	// +optional
// // 	Cluster string `json:"cluster,omitempty"`

// // 	// Clusters holds a registry to clusters to support multi-cluster tests.
// // 	// +optional
// // 	Clusters Clusters `json:"clusters,omitempty"`

// // 	// FileRefOrResource provides a reference to the file containing the resources to be created.
// // 	FileRefOrResource `json:",inline"`

// // 	// Template determines whether resources should be considered for templating.
// // 	// +optional
// // 	Template *bool `json:"template,omitempty"`

// // 	// DryRun determines whether the file should be applied in dry run mode.
// // 	// +optional
// // 	DryRun *bool `json:"dryRun,omitempty"`

// // 	// Expect defines a list of matched checks to validate the operation outcome.
// // 	// +optional
// // 	Expect []Expectation `json:"expect,omitempty"`
// // }

// // // Wait specifies how to perform wait operations on resources.
// // type Wait struct {
// // 	// Timeout for the operation. Specifies how long to wait for the condition to be met before timing out.
// // 	// +optional
// // 	Timeout *metav1.Duration `json:"timeout,omitempty"`

// // 	// Cluster defines the target cluster where the wait operation will be performed (default cluster will be used if not specified).
// // 	// +optional
// // 	Cluster string `json:"cluster,omitempty"`

// // 	// Clusters holds a registry to clusters to support multi-cluster tests.
// // 	// +optional
// // 	Clusters Clusters `json:"clusters,omitempty"`

// // 	// ResourceReference referenced resource type.
// // 	ResourceReference `json:",inline"`

// // 	// ObjectLabelsSelector determines the selection process of referenced objects.
// // 	// +optional
// // 	ObjectLabelsSelector `json:",inline"`

// // 	// For specifies the condition to wait for.
// // 	For `json:"for"`

// // 	// Format determines the output format (json or yaml).
// // 	// +optional
// // 	Format Format `json:"format,omitempty"`
// // }
