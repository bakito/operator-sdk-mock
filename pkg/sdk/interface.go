package stub

import (
	"k8s.io/apimachinery/pkg/types"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
)

type SdkInterface interface {
	Delete(object sdk.Object) error
	Update(object sdk.Object) error
	Create(object sdk.Object) error
	Get(object sdk.Object) error
	Patch(object sdk.Object, pt types.PatchType, patch []byte) error
	List(namespace string, into sdk.Object, opts sdk.ListOption) error
}

type OperatorSDK struct {
}

func (s *OperatorSDK) Delete(object sdk.Object) error {
	return sdk.Delete(object)
}

func (s *OperatorSDK) Update(object sdk.Object) error {
	return sdk.Update(object)
}

func (s *OperatorSDK) Create(object sdk.Object) error {
	return sdk.Create(object)
}
func (s *OperatorSDK) Get(object sdk.Object) error {
	return sdk.Get(object)
}
func (s *OperatorSDK) Patch(object sdk.Object, pt types.PatchType, patch []byte) error {
	return sdk.Patch(object, pt, patch)
}

func (s *OperatorSDK) List(namespace string, into sdk.Object, opts ...sdk.ListOption) error {
	return sdk.List(namespace, into, opts...)
}
