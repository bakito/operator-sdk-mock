# Mock functionality for the operator SDK

[![Build Status](https://travis-ci.org/bakito/operator-sdk-mock.svg?branch=master)](https://travis-ci.org/bakito/operator-sdk-mock)

## Usage
This module provides an interface over the [Opererator SDK](https://github.com/operator-framework/operator-sdk) to allow writing mock tests by use of [gomock](https://github.com/golang/mock).

### Handler

Create a handler that accesses the operator sdk functions via the sdk operator SDK implementation.

```go

import (
    "github.com/operator-framework/operator-sdk/pkg/sdk"
    sdkIf "github.com/bakito/operator-sdk-mock/pkg/sdk"
)

type Handler struct {
    sdk sdkIf.SdkInterface
}

func (h *Handler) Handle(ctx context.Context, event sdk.Event) error {
    h.sdkInterface.Update(event.Object)
    return nil
}

```

## Testing

See the [examples](./exmple/pkg/stub).

```go

import (
    "testing"

    "github.com/golang/mock/gomock"
    "github.com/operator-framework/operator-sdk/pkg/sdk"
    mock "github.com/bakito/operator-sdk-mock/pkg/sdk/mock_sdk"
)

func Test_Handle_Delete(t *testing.T) {

    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockSdk := mock.NewMockSdkInterface(ctrl)

    handler := Handler{
        sdk:       mockSdk,
    }

    pp := &v1.ProjectProvisioning{}

    mockSdk.EXPECT().List(gomock.Any(), gomock.Any(), gomock.Any()).Do(func(namespace string, into sdk.Object, opts sdk.ListOption) {
        ...
    })

    mockSdk.EXPECT().Delete(ns)

    handler.Handle(nil, sdk.Event{
        Deleted: true,
        Object:  pp,
    })
}

```


## Generate mock
```bash
mockgen -source=pkg/sdk/interface.go -destination=pkg/sdk/mock_sdk/interface_mock.go 
```
