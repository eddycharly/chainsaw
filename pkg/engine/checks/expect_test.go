package checks

import (
	"context"
	"testing"

	"github.com/jmespath-community/go-jmespath/pkg/binding"
	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestExpectations(t *testing.T) {
	tests := []struct {
		name     string
		obj      unstructured.Unstructured
		bindings binding.Bindings
		expect   []v1alpha1.Expectation
		want     bool
		wantErr  bool
	}{{
		name:     "nil",
		obj:      unstructured.Unstructured{},
		bindings: nil,
		expect:   nil,
		want:     false,
		wantErr:  false,
	}, {
		name:     "empty",
		obj:      unstructured.Unstructured{},
		bindings: nil,
		expect:   []v1alpha1.Expectation{},
		want:     false,
		wantErr:  false,
	}, {
		name: "no match",
		obj: unstructured.Unstructured{
			Object: map[string]any{
				"foo": "bar",
			},
		},
		bindings: nil,
		expect: []v1alpha1.Expectation{{
			Match: &v1alpha1.Check{
				Value: map[string]any{
					"foo": "baz",
				},
			},
			Check: v1alpha1.Check{
				Value: map[string]any{
					"foo": "bar",
				},
			},
		}},
		want:    false,
		wantErr: false,
	}, {
		name: "match",
		obj: unstructured.Unstructured{
			Object: map[string]any{
				"foo": "bar",
			},
		},
		bindings: nil,
		expect: []v1alpha1.Expectation{{
			Match: &v1alpha1.Check{
				Value: map[string]any{
					"foo": "bar",
				},
			},
			Check: v1alpha1.Check{
				Value: map[string]any{
					"foo": "bar",
				},
			},
		}},
		want:    true,
		wantErr: false,
	}, {
		name: "check err",
		obj: unstructured.Unstructured{
			Object: map[string]any{
				"foo": "bar",
			},
		},
		bindings: nil,
		expect: []v1alpha1.Expectation{{
			Match: &v1alpha1.Check{
				Value: map[string]any{
					"foo": "bar",
				},
			},
			Check: v1alpha1.Check{
				Value: map[string]any{
					"(foo())": "bar",
				},
			},
		}},
		want:    true,
		wantErr: true,
	}, {
		name: "match err",
		obj: unstructured.Unstructured{
			Object: map[string]any{
				"foo": "bar",
			},
		},
		bindings: nil,
		expect: []v1alpha1.Expectation{{
			Match: &v1alpha1.Check{
				Value: map[string]any{
					"(foo())": "bar",
				},
			},
			Check: v1alpha1.Check{
				Value: map[string]any{
					"foo": "bar",
				},
			},
		}},
		want:    true,
		wantErr: true,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Expect(context.TODO(), tt.obj, tt.bindings, tt.expect...)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
