package v1alpha2

// // OperationBase defines common elements to all operations.
// type OperationBase struct {
// 	// Description contains a description of the operation.
// 	// +optional
// 	Description string `json:"description,omitempty"`

// 	// ContinueOnError determines whether a test should continue or not in case the operation was not successful.
// 	// Even if the test continues executing, it will still be reported as failed.
// 	// +optional
// 	ContinueOnError *bool `json:"continueOnError,omitempty"`

// 	// Bindings defines additional binding key/values.
// 	// +optional
// 	Bindings []Binding `json:"bindings,omitempty"`

// 	// Outputs defines output bindings.
// 	// +optional
// 	Outputs []Output `json:"outputs,omitempty"`

// 	// Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).
// 	// +optional
// 	Cluster string `json:"cluster,omitempty"`

// 	// Clusters holds a registry to clusters to support multi-cluster tests.
// 	// +optional
// 	Clusters Clusters `json:"clusters,omitempty"`
// }

// // Operation defines a single operation, only one action is permitted for a given operation.
// type Operation struct {
// 	// OperationBase defines common elements to all operations.
// 	// +optional
// 	OperationBase `json:",inline"`
// }
