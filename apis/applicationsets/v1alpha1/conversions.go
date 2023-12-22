package v1alpha1

import (
	argocdv1alpha1 "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	extv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// Converter helps to convert ArgoCD types to api types of this provider and vise-versa
// goverter:converter
// goverter:useZeroValueOnPointerInconsistency
// goverter:ignoreUnexported
// goverter:extend ExtV1JSONToRuntimeRawExtension
// +k8s:deepcopy-gen=false
type Converter interface {
	ToArgoApplicationSpec(in *ApplicationSetParameters) *argocdv1alpha1.ApplicationSetSpec
}

// ExtV1JSONToRuntimeRawExtension converts an extv1.JSON into a
// *runtime.RawExtension.
func ExtV1JSONToRuntimeRawExtension(in extv1.JSON) *runtime.RawExtension {
	return &runtime.RawExtension{
		Raw: in.Raw,
	}
}
