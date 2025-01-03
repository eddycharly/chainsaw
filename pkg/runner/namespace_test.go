package runner

import (
	"context"
	"testing"

	"github.com/kyverno/chainsaw/pkg/apis"
	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	"github.com/kyverno/chainsaw/pkg/loaders/config"
	enginecontext "github.com/kyverno/chainsaw/pkg/runner/context"
	"github.com/kyverno/kyverno-json/pkg/core/compilers"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/ptr"
)

func Test_buildNamespace(t *testing.T) {
	config, err := config.DefaultConfiguration()
	assert.NoError(t, err)
	tc, err := enginecontext.InitContext(config.Spec, nil, nil)
	assert.NoError(t, err)
	tests := []struct {
		name       string
		compilers  compilers.Compilers
		nsName     string
		nsTemplate *v1alpha1.Projection
		tc         enginecontext.TestContext
		want       *corev1.Namespace
		wantErr    bool
	}{{
		name:   "simple",
		nsName: "foo",
		want: &corev1.Namespace{
			TypeMeta: metav1.TypeMeta{
				APIVersion: corev1.SchemeGroupVersion.String(),
				Kind:       "Namespace",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "foo",
			},
		},
	}, {
		name:      "with template",
		compilers: apis.DefaultCompilers,
		tc:        tc.WithBinding("bar", "bar"),
		nsName:    "foo",
		nsTemplate: ptr.To(v1alpha1.NewProjection(map[string]any{
			"metadata": map[string]any{
				"labels": map[string]any{
					"foo": "bar",
				},
			},
		})),
		want: &corev1.Namespace{
			TypeMeta: metav1.TypeMeta{
				APIVersion: corev1.SchemeGroupVersion.String(),
				Kind:       "Namespace",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "foo",
				Labels: map[string]string{
					"foo": "bar",
				},
			},
		},
	}, {
		name:       "with nil template",
		compilers:  apis.DefaultCompilers,
		tc:         tc.WithBinding("bar", "bar"),
		nsName:     "foo",
		nsTemplate: ptr.To(v1alpha1.NewProjection(nil)),
		want: &corev1.Namespace{
			TypeMeta: metav1.TypeMeta{
				APIVersion: corev1.SchemeGroupVersion.String(),
				Kind:       "Namespace",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "foo",
			},
		},
	}, {
		name:      "with template",
		compilers: apis.DefaultCompilers,
		tc:        tc.WithBinding("bar", "bar"),
		nsName:    "foo",
		nsTemplate: ptr.To(v1alpha1.NewProjection(map[string]any{
			"metadata": map[string]any{
				"labels": map[string]any{
					"foo": "($bar)",
				},
			},
		})),
		want: &corev1.Namespace{
			TypeMeta: metav1.TypeMeta{
				APIVersion: corev1.SchemeGroupVersion.String(),
				Kind:       "Namespace",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "foo",
				Labels: map[string]string{
					"foo": "bar",
				},
			},
		},
	}, {
		name:      "with bad template",
		compilers: apis.DefaultCompilers,
		tc:        tc.WithBinding("bar", "bar"),
		nsName:    "foo",
		nsTemplate: ptr.To(v1alpha1.NewProjection(map[string]any{
			"metadata": map[string]any{
				"labels": map[string]any{
					"foo": "($flop)",
				},
			},
		})),
		wantErr: true,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := buildNamespace(context.Background(), tt.compilers, tt.nsName, tt.nsTemplate, tt.tc)
			assert.Equal(t, tt.want, got)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
