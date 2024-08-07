//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha2

import (
	v1alpha1 "github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionCheck) DeepCopyInto(out *ActionCheck) {
	*out = *in
	if in.Check != nil {
		in, out := &in.Check, &out.Check
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionCheck.
func (in *ActionCheck) DeepCopy() *ActionCheck {
	if in == nil {
		return nil
	}
	out := new(ActionCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionCheckRef) DeepCopyInto(out *ActionCheckRef) {
	*out = *in
	out.FileRef = in.FileRef
	if in.Check != nil {
		in, out := &in.Check, &out.Check
		*out = (*in).DeepCopy()
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionCheckRef.
func (in *ActionCheckRef) DeepCopy() *ActionCheckRef {
	if in == nil {
		return nil
	}
	out := new(ActionCheckRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionDryRun) DeepCopyInto(out *ActionDryRun) {
	*out = *in
	if in.DryRun != nil {
		in, out := &in.DryRun, &out.DryRun
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionDryRun.
func (in *ActionDryRun) DeepCopy() *ActionDryRun {
	if in == nil {
		return nil
	}
	out := new(ActionDryRun)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionEnv) DeepCopyInto(out *ActionEnv) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1alpha1.Binding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionEnv.
func (in *ActionEnv) DeepCopy() *ActionEnv {
	if in == nil {
		return nil
	}
	out := new(ActionEnv)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionExpectations) DeepCopyInto(out *ActionExpectations) {
	*out = *in
	if in.Expect != nil {
		in, out := &in.Expect, &out.Expect
		*out = make([]v1alpha1.Expectation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionExpectations.
func (in *ActionExpectations) DeepCopy() *ActionExpectations {
	if in == nil {
		return nil
	}
	out := new(ActionExpectations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionFormat) DeepCopyInto(out *ActionFormat) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionFormat.
func (in *ActionFormat) DeepCopy() *ActionFormat {
	if in == nil {
		return nil
	}
	out := new(ActionFormat)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionInlineResource) DeepCopyInto(out *ActionInlineResource) {
	*out = *in
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionInlineResource.
func (in *ActionInlineResource) DeepCopy() *ActionInlineResource {
	if in == nil {
		return nil
	}
	out := new(ActionInlineResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionObject) DeepCopyInto(out *ActionObject) {
	*out = *in
	out.ObjectType = in.ObjectType
	out.ActionObjectSelector = in.ActionObjectSelector
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionObject.
func (in *ActionObject) DeepCopy() *ActionObject {
	if in == nil {
		return nil
	}
	out := new(ActionObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionObjectSelector) DeepCopyInto(out *ActionObjectSelector) {
	*out = *in
	out.ObjectName = in.ObjectName
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionObjectSelector.
func (in *ActionObjectSelector) DeepCopy() *ActionObjectSelector {
	if in == nil {
		return nil
	}
	out := new(ActionObjectSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionResourceRef) DeepCopyInto(out *ActionResourceRef) {
	*out = *in
	out.FileRef = in.FileRef
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = (*in).DeepCopy()
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionResourceRef.
func (in *ActionResourceRef) DeepCopy() *ActionResourceRef {
	if in == nil {
		return nil
	}
	out := new(ActionResourceRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionTimeout) DeepCopyInto(out *ActionTimeout) {
	*out = *in
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionTimeout.
func (in *ActionTimeout) DeepCopy() *ActionTimeout {
	if in == nil {
		return nil
	}
	out := new(ActionTimeout)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Apply) DeepCopyInto(out *Apply) {
	*out = *in
	in.ActionDryRun.DeepCopyInto(&out.ActionDryRun)
	in.ActionExpectations.DeepCopyInto(&out.ActionExpectations)
	in.ActionResourceRef.DeepCopyInto(&out.ActionResourceRef)
	in.ActionTimeout.DeepCopyInto(&out.ActionTimeout)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Apply.
func (in *Apply) DeepCopy() *Apply {
	if in == nil {
		return nil
	}
	out := new(Apply)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Assert) DeepCopyInto(out *Assert) {
	*out = *in
	in.ActionCheckRef.DeepCopyInto(&out.ActionCheckRef)
	in.ActionTimeout.DeepCopyInto(&out.ActionTimeout)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Assert.
func (in *Assert) DeepCopy() *Assert {
	if in == nil {
		return nil
	}
	out := new(Assert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CleanupOptions) DeepCopyInto(out *CleanupOptions) {
	*out = *in
	if in.DelayBeforeCleanup != nil {
		in, out := &in.DelayBeforeCleanup, &out.DelayBeforeCleanup
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CleanupOptions.
func (in *CleanupOptions) DeepCopy() *CleanupOptions {
	if in == nil {
		return nil
	}
	out := new(CleanupOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Command) DeepCopyInto(out *Command) {
	*out = *in
	in.ActionCheck.DeepCopyInto(&out.ActionCheck)
	in.ActionEnv.DeepCopyInto(&out.ActionEnv)
	in.ActionTimeout.DeepCopyInto(&out.ActionTimeout)
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Command.
func (in *Command) DeepCopy() *Command {
	if in == nil {
		return nil
	}
	out := new(Command)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Configuration) DeepCopyInto(out *Configuration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Configuration.
func (in *Configuration) DeepCopy() *Configuration {
	if in == nil {
		return nil
	}
	out := new(Configuration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Configuration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationSpec) DeepCopyInto(out *ConfigurationSpec) {
	*out = *in
	in.Cleanup.DeepCopyInto(&out.Cleanup)
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make(v1alpha1.Clusters, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Deletion = in.Deletion
	out.Discovery = in.Discovery
	in.Error.DeepCopyInto(&out.Error)
	in.Execution.DeepCopyInto(&out.Execution)
	in.Namespace.DeepCopyInto(&out.Namespace)
	if in.Report != nil {
		in, out := &in.Report, &out.Report
		*out = new(ReportOptions)
		**out = **in
	}
	out.Templating = in.Templating
	out.Timeouts = in.Timeouts
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationSpec.
func (in *ConfigurationSpec) DeepCopy() *ConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(ConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Create) DeepCopyInto(out *Create) {
	*out = *in
	in.ActionDryRun.DeepCopyInto(&out.ActionDryRun)
	in.ActionExpectations.DeepCopyInto(&out.ActionExpectations)
	in.ActionResourceRef.DeepCopyInto(&out.ActionResourceRef)
	in.ActionTimeout.DeepCopyInto(&out.ActionTimeout)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Create.
func (in *Create) DeepCopy() *Create {
	if in == nil {
		return nil
	}
	out := new(Create)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Delete) DeepCopyInto(out *Delete) {
	*out = *in
	in.ActionExpectations.DeepCopyInto(&out.ActionExpectations)
	in.ActionTimeout.DeepCopyInto(&out.ActionTimeout)
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(bool)
		**out = **in
	}
	if in.Ref != nil {
		in, out := &in.Ref, &out.Ref
		*out = new(ObjectReference)
		(*in).DeepCopyInto(*out)
	}
	if in.DeletionPropagationPolicy != nil {
		in, out := &in.DeletionPropagationPolicy, &out.DeletionPropagationPolicy
		*out = new(v1.DeletionPropagation)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Delete.
func (in *Delete) DeepCopy() *Delete {
	if in == nil {
		return nil
	}
	out := new(Delete)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeletionOptions) DeepCopyInto(out *DeletionOptions) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeletionOptions.
func (in *DeletionOptions) DeepCopy() *DeletionOptions {
	if in == nil {
		return nil
	}
	out := new(DeletionOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Describe) DeepCopyInto(out *Describe) {
	*out = *in
	out.ActionObject = in.ActionObject
	in.ActionTimeout.DeepCopyInto(&out.ActionTimeout)
	if in.ShowEvents != nil {
		in, out := &in.ShowEvents, &out.ShowEvents
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Describe.
func (in *Describe) DeepCopy() *Describe {
	if in == nil {
		return nil
	}
	out := new(Describe)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveryOptions) DeepCopyInto(out *DiscoveryOptions) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveryOptions.
func (in *DiscoveryOptions) DeepCopy() *DiscoveryOptions {
	if in == nil {
		return nil
	}
	out := new(DiscoveryOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Error) DeepCopyInto(out *Error) {
	*out = *in
	in.ActionCheckRef.DeepCopyInto(&out.ActionCheckRef)
	in.ActionTimeout.DeepCopyInto(&out.ActionTimeout)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Error.
func (in *Error) DeepCopy() *Error {
	if in == nil {
		return nil
	}
	out := new(Error)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ErrorOptions) DeepCopyInto(out *ErrorOptions) {
	*out = *in
	if in.Catch != nil {
		in, out := &in.Catch, &out.Catch
		*out = make([]v1alpha1.CatchFinally, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrorOptions.
func (in *ErrorOptions) DeepCopy() *ErrorOptions {
	if in == nil {
		return nil
	}
	out := new(ErrorOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Events) DeepCopyInto(out *Events) {
	*out = *in
	out.ActionFormat = in.ActionFormat
	out.ActionObjectSelector = in.ActionObjectSelector
	in.ActionTimeout.DeepCopyInto(&out.ActionTimeout)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Events.
func (in *Events) DeepCopy() *Events {
	if in == nil {
		return nil
	}
	out := new(Events)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExecutionOptions) DeepCopyInto(out *ExecutionOptions) {
	*out = *in
	if in.Parallel != nil {
		in, out := &in.Parallel, &out.Parallel
		*out = new(int)
		**out = **in
	}
	if in.RepeatCount != nil {
		in, out := &in.RepeatCount, &out.RepeatCount
		*out = new(int)
		**out = **in
	}
	if in.ForceTerminationGracePeriod != nil {
		in, out := &in.ForceTerminationGracePeriod, &out.ForceTerminationGracePeriod
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExecutionOptions.
func (in *ExecutionOptions) DeepCopy() *ExecutionOptions {
	if in == nil {
		return nil
	}
	out := new(ExecutionOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileRef) DeepCopyInto(out *FileRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileRef.
func (in *FileRef) DeepCopy() *FileRef {
	if in == nil {
		return nil
	}
	out := new(FileRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Get) DeepCopyInto(out *Get) {
	*out = *in
	out.ActionFormat = in.ActionFormat
	out.ActionObject = in.ActionObject
	in.ActionTimeout.DeepCopyInto(&out.ActionTimeout)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Get.
func (in *Get) DeepCopy() *Get {
	if in == nil {
		return nil
	}
	out := new(Get)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceOptions) DeepCopyInto(out *NamespaceOptions) {
	*out = *in
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceOptions.
func (in *NamespaceOptions) DeepCopy() *NamespaceOptions {
	if in == nil {
		return nil
	}
	out := new(NamespaceOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectReference) DeepCopyInto(out *ObjectReference) {
	*out = *in
	out.ObjectType = in.ObjectType
	out.ObjectName = in.ObjectName
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectReference.
func (in *ObjectReference) DeepCopy() *ObjectReference {
	if in == nil {
		return nil
	}
	out := new(ObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Operation) DeepCopyInto(out *Operation) {
	*out = *in
	in.OperationAction.DeepCopyInto(&out.OperationAction)
	in.OperationBindings.DeepCopyInto(&out.OperationBindings)
	in.OperationClusters.DeepCopyInto(&out.OperationClusters)
	in.OperationOutputs.DeepCopyInto(&out.OperationOutputs)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Operation.
func (in *Operation) DeepCopy() *Operation {
	if in == nil {
		return nil
	}
	out := new(Operation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperationAction) DeepCopyInto(out *OperationAction) {
	*out = *in
	if in.Apply != nil {
		in, out := &in.Apply, &out.Apply
		*out = new(Apply)
		(*in).DeepCopyInto(*out)
	}
	if in.Assert != nil {
		in, out := &in.Assert, &out.Assert
		*out = new(Assert)
		(*in).DeepCopyInto(*out)
	}
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = new(Command)
		(*in).DeepCopyInto(*out)
	}
	if in.Create != nil {
		in, out := &in.Create, &out.Create
		*out = new(Create)
		(*in).DeepCopyInto(*out)
	}
	if in.Delete != nil {
		in, out := &in.Delete, &out.Delete
		*out = new(Delete)
		(*in).DeepCopyInto(*out)
	}
	if in.Describe != nil {
		in, out := &in.Describe, &out.Describe
		*out = new(Describe)
		(*in).DeepCopyInto(*out)
	}
	if in.Error != nil {
		in, out := &in.Error, &out.Error
		*out = new(Error)
		(*in).DeepCopyInto(*out)
	}
	if in.Events != nil {
		in, out := &in.Events, &out.Events
		*out = new(Events)
		(*in).DeepCopyInto(*out)
	}
	if in.Get != nil {
		in, out := &in.Get, &out.Get
		*out = new(Get)
		(*in).DeepCopyInto(*out)
	}
	if in.Patch != nil {
		in, out := &in.Patch, &out.Patch
		*out = new(Patch)
		(*in).DeepCopyInto(*out)
	}
	if in.PodLogs != nil {
		in, out := &in.PodLogs, &out.PodLogs
		*out = new(PodLogs)
		(*in).DeepCopyInto(*out)
	}
	if in.Script != nil {
		in, out := &in.Script, &out.Script
		*out = new(Script)
		(*in).DeepCopyInto(*out)
	}
	if in.Sleep != nil {
		in, out := &in.Sleep, &out.Sleep
		*out = new(Sleep)
		**out = **in
	}
	if in.Update != nil {
		in, out := &in.Update, &out.Update
		*out = new(Update)
		(*in).DeepCopyInto(*out)
	}
	if in.Wait != nil {
		in, out := &in.Wait, &out.Wait
		*out = new(Wait)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperationAction.
func (in *OperationAction) DeepCopy() *OperationAction {
	if in == nil {
		return nil
	}
	out := new(OperationAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperationBindings) DeepCopyInto(out *OperationBindings) {
	*out = *in
	if in.Bindings != nil {
		in, out := &in.Bindings, &out.Bindings
		*out = make([]v1alpha1.Binding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperationBindings.
func (in *OperationBindings) DeepCopy() *OperationBindings {
	if in == nil {
		return nil
	}
	out := new(OperationBindings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperationClusters) DeepCopyInto(out *OperationClusters) {
	*out = *in
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make(v1alpha1.Clusters, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperationClusters.
func (in *OperationClusters) DeepCopy() *OperationClusters {
	if in == nil {
		return nil
	}
	out := new(OperationClusters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperationOutputs) DeepCopyInto(out *OperationOutputs) {
	*out = *in
	if in.Outputs != nil {
		in, out := &in.Outputs, &out.Outputs
		*out = make([]v1alpha1.Output, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperationOutputs.
func (in *OperationOutputs) DeepCopy() *OperationOutputs {
	if in == nil {
		return nil
	}
	out := new(OperationOutputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Patch) DeepCopyInto(out *Patch) {
	*out = *in
	in.ActionDryRun.DeepCopyInto(&out.ActionDryRun)
	in.ActionExpectations.DeepCopyInto(&out.ActionExpectations)
	in.ActionResourceRef.DeepCopyInto(&out.ActionResourceRef)
	in.ActionTimeout.DeepCopyInto(&out.ActionTimeout)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Patch.
func (in *Patch) DeepCopy() *Patch {
	if in == nil {
		return nil
	}
	out := new(Patch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodLogs) DeepCopyInto(out *PodLogs) {
	*out = *in
	out.ActionObjectSelector = in.ActionObjectSelector
	in.ActionTimeout.DeepCopyInto(&out.ActionTimeout)
	if in.Tail != nil {
		in, out := &in.Tail, &out.Tail
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodLogs.
func (in *PodLogs) DeepCopy() *PodLogs {
	if in == nil {
		return nil
	}
	out := new(PodLogs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportOptions) DeepCopyInto(out *ReportOptions) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportOptions.
func (in *ReportOptions) DeepCopy() *ReportOptions {
	if in == nil {
		return nil
	}
	out := new(ReportOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Script) DeepCopyInto(out *Script) {
	*out = *in
	in.ActionCheck.DeepCopyInto(&out.ActionCheck)
	in.ActionEnv.DeepCopyInto(&out.ActionEnv)
	in.ActionTimeout.DeepCopyInto(&out.ActionTimeout)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Script.
func (in *Script) DeepCopy() *Script {
	if in == nil {
		return nil
	}
	out := new(Script)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sleep) DeepCopyInto(out *Sleep) {
	*out = *in
	out.Duration = in.Duration
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sleep.
func (in *Sleep) DeepCopy() *Sleep {
	if in == nil {
		return nil
	}
	out := new(Sleep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplatingOptions) DeepCopyInto(out *TemplatingOptions) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplatingOptions.
func (in *TemplatingOptions) DeepCopy() *TemplatingOptions {
	if in == nil {
		return nil
	}
	out := new(TemplatingOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Test) DeepCopyInto(out *Test) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Test.
func (in *Test) DeepCopy() *Test {
	if in == nil {
		return nil
	}
	out := new(Test)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Test) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestExecutionOptions) DeepCopyInto(out *TestExecutionOptions) {
	*out = *in
	if in.TerminationGracePeriod != nil {
		in, out := &in.TerminationGracePeriod, &out.TerminationGracePeriod
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestExecutionOptions.
func (in *TestExecutionOptions) DeepCopy() *TestExecutionOptions {
	if in == nil {
		return nil
	}
	out := new(TestExecutionOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestSpec) DeepCopyInto(out *TestSpec) {
	*out = *in
	in.Cleanup.DeepCopyInto(&out.Cleanup)
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make(v1alpha1.Clusters, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Execution.DeepCopyInto(&out.Execution)
	if in.Bindings != nil {
		in, out := &in.Bindings, &out.Bindings
		*out = make([]v1alpha1.Binding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Deletion = in.Deletion
	in.Error.DeepCopyInto(&out.Error)
	in.Namespace.DeepCopyInto(&out.Namespace)
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]TestStep, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Templating = in.Templating
	in.Timeouts.DeepCopyInto(&out.Timeouts)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestSpec.
func (in *TestSpec) DeepCopy() *TestSpec {
	if in == nil {
		return nil
	}
	out := new(TestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestStep) DeepCopyInto(out *TestStep) {
	*out = *in
	in.TestStepSpec.DeepCopyInto(&out.TestStepSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestStep.
func (in *TestStep) DeepCopy() *TestStep {
	if in == nil {
		return nil
	}
	out := new(TestStep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestStepSpec) DeepCopyInto(out *TestStepSpec) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(v1alpha1.Timeouts)
		(*in).DeepCopyInto(*out)
	}
	if in.DeletionPropagationPolicy != nil {
		in, out := &in.DeletionPropagationPolicy, &out.DeletionPropagationPolicy
		*out = new(v1.DeletionPropagation)
		**out = **in
	}
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make(v1alpha1.Clusters, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SkipDelete != nil {
		in, out := &in.SkipDelete, &out.SkipDelete
		*out = new(bool)
		**out = **in
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(bool)
		**out = **in
	}
	if in.Bindings != nil {
		in, out := &in.Bindings, &out.Bindings
		*out = make([]v1alpha1.Binding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Try != nil {
		in, out := &in.Try, &out.Try
		*out = make([]TryOperation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Catch != nil {
		in, out := &in.Catch, &out.Catch
		*out = make([]Operation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Finally != nil {
		in, out := &in.Finally, &out.Finally
		*out = make([]Operation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Cleanup != nil {
		in, out := &in.Cleanup, &out.Cleanup
		*out = make([]Operation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestStepSpec.
func (in *TestStepSpec) DeepCopy() *TestStepSpec {
	if in == nil {
		return nil
	}
	out := new(TestStepSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TryOperation) DeepCopyInto(out *TryOperation) {
	*out = *in
	in.Operation.DeepCopyInto(&out.Operation)
	if in.ContinueOnError != nil {
		in, out := &in.ContinueOnError, &out.ContinueOnError
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TryOperation.
func (in *TryOperation) DeepCopy() *TryOperation {
	if in == nil {
		return nil
	}
	out := new(TryOperation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Update) DeepCopyInto(out *Update) {
	*out = *in
	in.ActionDryRun.DeepCopyInto(&out.ActionDryRun)
	in.ActionExpectations.DeepCopyInto(&out.ActionExpectations)
	in.ActionResourceRef.DeepCopyInto(&out.ActionResourceRef)
	in.ActionTimeout.DeepCopyInto(&out.ActionTimeout)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Update.
func (in *Update) DeepCopy() *Update {
	if in == nil {
		return nil
	}
	out := new(Update)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Wait) DeepCopyInto(out *Wait) {
	*out = *in
	in.ActionTimeout.DeepCopyInto(&out.ActionTimeout)
	out.ActionFormat = in.ActionFormat
	out.ActionObject = in.ActionObject
	in.WaitFor.DeepCopyInto(&out.WaitFor)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Wait.
func (in *Wait) DeepCopy() *Wait {
	if in == nil {
		return nil
	}
	out := new(Wait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WaitFor) DeepCopyInto(out *WaitFor) {
	*out = *in
	if in.Deletion != nil {
		in, out := &in.Deletion, &out.Deletion
		*out = new(WaitForDeletion)
		**out = **in
	}
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(WaitForCondition)
		(*in).DeepCopyInto(*out)
	}
	if in.JsonPath != nil {
		in, out := &in.JsonPath, &out.JsonPath
		*out = new(WaitForJsonPath)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WaitFor.
func (in *WaitFor) DeepCopy() *WaitFor {
	if in == nil {
		return nil
	}
	out := new(WaitFor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WaitForCondition) DeepCopyInto(out *WaitForCondition) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WaitForCondition.
func (in *WaitForCondition) DeepCopy() *WaitForCondition {
	if in == nil {
		return nil
	}
	out := new(WaitForCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WaitForDeletion) DeepCopyInto(out *WaitForDeletion) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WaitForDeletion.
func (in *WaitForDeletion) DeepCopy() *WaitForDeletion {
	if in == nil {
		return nil
	}
	out := new(WaitForDeletion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WaitForJsonPath) DeepCopyInto(out *WaitForJsonPath) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WaitForJsonPath.
func (in *WaitForJsonPath) DeepCopy() *WaitForJsonPath {
	if in == nil {
		return nil
	}
	out := new(WaitForJsonPath)
	in.DeepCopyInto(out)
	return out
}
