package stub

import (
	"testing"

	mock "github.com/bakito/operator-sdk-mock/pkg/sdk/mock_sdk"
	"github.com/golang/mock/gomock"
	"github.com/operator-framework/operator-sdk/pkg/sdk"
	k8sv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Test_Handle_Get(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSdk := mock.NewMockSdkInterface(ctrl)

	handler := Handler{
		sdkInterface: mockSdk,
	}

	ns := &k8sv1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "Get",
		},
	}

	mockSdk.EXPECT().Get(gomock.Any())

	handler.Handle(nil, sdk.Event{
		Deleted: true,
		Object:  ns,
	})
}

func Test_Handle_Create(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSdk := mock.NewMockSdkInterface(ctrl)

	handler := Handler{
		sdkInterface: mockSdk,
	}

	ns := &k8sv1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "Create",
		},
	}

	mockSdk.EXPECT().Create(gomock.Any())

	handler.Handle(nil, sdk.Event{
		Deleted: true,
		Object:  ns,
	})
}
func Test_Handle_Delete(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSdk := mock.NewMockSdkInterface(ctrl)

	handler := Handler{
		sdkInterface: mockSdk,
	}

	ns := &k8sv1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "Delete",
		},
	}

	mockSdk.EXPECT().Delete(gomock.Any())

	handler.Handle(nil, sdk.Event{
		Deleted: true,
		Object:  ns,
	})
}
func Test_Handle_Update(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSdk := mock.NewMockSdkInterface(ctrl)

	handler := Handler{
		sdkInterface: mockSdk,
	}

	ns := &k8sv1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "Update",
		},
	}

	mockSdk.EXPECT().Update(gomock.Any())

	handler.Handle(nil, sdk.Event{
		Deleted: true,
		Object:  ns,
	})
}
func Test_Handle_List(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSdk := mock.NewMockSdkInterface(ctrl)

	handler := Handler{
		sdkInterface: mockSdk,
	}

	ns := &k8sv1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "List",
		},
	}

	mockSdk.EXPECT().List(gomock.Any(), gomock.Any(), gomock.Any())

	handler.Handle(nil, sdk.Event{
		Deleted: true,
		Object:  ns,
	})
}
func Test_Handle_Patch(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSdk := mock.NewMockSdkInterface(ctrl)

	handler := Handler{
		sdkInterface: mockSdk,
	}

	ns := &k8sv1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "Patch",
		},
	}

	mockSdk.EXPECT().Patch(gomock.Any(), gomock.Any(), gomock.Any())

	handler.Handle(nil, sdk.Event{
		Deleted: true,
		Object:  ns,
	})
}
