package mutate

import (
	"context"
	"fmt"
	"reflect"

	"github.com/kyverno/chainsaw/pkg/apis"
	"github.com/kyverno/kyverno-json/pkg/core/compilers"
	reflectutils "github.com/kyverno/kyverno-json/pkg/utils/reflect"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

func Parse(ctx context.Context, mutation any) Mutation {
	switch reflectutils.GetKind(mutation) {
	case reflect.Slice:
		node := sliceNode{}
		valueOf := reflect.ValueOf(mutation)
		for i := 0; i < valueOf.Len(); i++ {
			node = append(node, Parse(ctx, valueOf.Index(i).Interface()))
		}
		return node
	case reflect.Map:
		node := mapNode{}
		iter := reflect.ValueOf(mutation).MapRange()
		for iter.Next() {
			node[iter.Key().Interface()] = Parse(ctx, iter.Value().Interface())
		}
		return node
	default:
		return &scalarNode{rhs: mutation}
	}
}

// mapNode is the mutation type represented by a map.
// it is responsible for projecting the analysed resource and passing the result to the descendant
type mapNode map[any]Mutation

func (n mapNode) mutate(ctx context.Context, c compilers.Compilers, path *field.Path, value any, bindings apis.Bindings) (any, error) {
	out := map[any]any{}
	for k, v := range n {
		var projection any
		if value != nil {
			mapValue := reflect.ValueOf(value).MapIndex(reflect.ValueOf(k))
			if mapValue.IsValid() {
				projection = mapValue.Interface()
			}
		}
		inner, err := v.mutate(ctx, c, path.Child(fmt.Sprint(k)), projection, bindings)
		if err != nil {
			return nil, err
		}
		out[k] = inner
	}
	return out, nil
}

// sliceNode is the mutation type represented by a slice.
// it first compares the length of the analysed resource with the length of the descendants.
// if lengths match all descendants are evaluated with their corresponding items.
type sliceNode []Mutation

func (n sliceNode) mutate(ctx context.Context, c compilers.Compilers, path *field.Path, value any, bindings apis.Bindings) (any, error) {
	if value != nil && reflectutils.GetKind(value) != reflect.Slice && reflectutils.GetKind(value) != reflect.Array {
		return nil, field.TypeInvalid(path, value, "expected a slice or array")
	} else {
		var out []any
		for i := range n {
			var projection any
			if value != nil {
				valueOf := reflect.ValueOf(value)
				if i < valueOf.Len() {
					projection = valueOf.Index(i).Interface()
				}
			}
			inner, err := n[i].mutate(ctx, c, path.Index(i), projection, bindings)
			if err != nil {
				return nil, err
			}
			out = append(out, inner)
		}
		return out, nil
	}
}

// scalarNode is a terminal type of mutation.
// it receives a value and compares it with an expected value.
// the expected value can be the result of an expression.
type scalarNode struct {
	rhs any
}

func (n *scalarNode) mutate(ctx context.Context, c compilers.Compilers, path *field.Path, value any, bindings apis.Bindings) (any, error) {
	rhs := n.rhs
	expression := parseExpression(ctx, rhs)
	// we only project if the expression uses the engine syntax
	// this is to avoid the case where the value is a map and the RHS is a string
	if expression != nil {
		if compiler := c.Compiler(expression.Engine); compiler != nil {
			projected, err := compilers.Execute(expression.Statement, value, bindings, compiler)
			if err != nil {
				return nil, field.InternalError(path, err)
			}
			rhs = projected
		}
	}
	return rhs, nil
}
