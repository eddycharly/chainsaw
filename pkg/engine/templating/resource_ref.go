package templating

import (
	"context"

	"github.com/kyverno/chainsaw/pkg/apis"
	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	"github.com/kyverno/kyverno-json/pkg/core/compilers"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func ResourceRef(ctx context.Context, c compilers.Compilers, obj *unstructured.Unstructured, bindings apis.Bindings) error {
	if obj == nil {
		return nil
	}
	apiVersion := obj.GetAPIVersion()
	kind := obj.GetKind()
	// this is not a valid resource (non resource assertion maybe ?)
	if apiVersion == "" || kind == "" {
		return nil
	}
	var temp unstructured.Unstructured
	temp.SetAPIVersion(apiVersion)
	temp.SetKind(kind)
	temp.SetName(obj.GetName())
	temp.SetNamespace(obj.GetNamespace())
	temp.SetLabels(obj.GetLabels())
	template := v1alpha1.NewProjection(temp.UnstructuredContent())
	if merged, err := TemplateAndMerge(ctx, c, temp, bindings, template); err != nil {
		return err
	} else {
		temp = merged
	}
	obj.SetAPIVersion(temp.GetAPIVersion())
	obj.SetKind(temp.GetKind())
	obj.SetName(temp.GetName())
	obj.SetNamespace(temp.GetNamespace())
	obj.SetLabels(temp.GetLabels())
	return nil
}
