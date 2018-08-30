package stub

import (
	"context"

	"encoding/json"

	sdkIf "github.com/bakito/operator-sdk-mock/pkg/sdk"
	"github.com/operator-framework/operator-sdk/pkg/sdk"
	k8sv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func NewHandler() sdk.Handler {
	return &Handler{
		sdkInterface: &sdkIf.OperatorSDK{}}
}

type Handler struct {
	sdkInterface sdkIf.SdkInterface
}

func (h *Handler) Handle(ctx context.Context, event sdk.Event) error {
	switch eventNs := event.Object.(type) {
	case *k8sv1.Namespace:
		testNs := &k8sv1.Namespace{}

		switch eventNs.ObjectMeta.Name {

		case "Get":
			testNs.ObjectMeta.Name = "From Get"

			h.sdkInterface.Get(testNs)
		case "Create":
			testNs.ObjectMeta.Name = "From Create"

			h.sdkInterface.Create(testNs)

		case "Delete":
			testNs.ObjectMeta.Name = "From Delete"

			h.sdkInterface.Delete(testNs)

		case "Update":
			testNs.ObjectMeta.Name = "From Update"

			h.sdkInterface.Update(testNs)

		case "Patch":

			nsPatch := &k8sv1.Namespace{

				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						"test": "test",
					},
				},
			}

			bytes, _ := json.Marshal(nsPatch)
			h.sdkInterface.Patch(nsPatch, types.StrategicMergePatchType, bytes)

		case "List":
			namespaceList := &k8sv1.NamespaceList{
				TypeMeta: metav1.TypeMeta{
					Kind:       "Namespace",
					APIVersion: "v1",
				},
			}

			h.sdkInterface.List(metav1.NamespaceAll, namespaceList)
		}
	}
	return nil
}
