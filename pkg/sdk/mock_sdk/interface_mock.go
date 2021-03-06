// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/sdk/interface.go

// Package mock_stub is a generated GoMock package.
package mock_stub

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sdk "github.com/operator-framework/operator-sdk/pkg/sdk"
	types "k8s.io/apimachinery/pkg/types"
)

// MockSdkInterface is a mock of SdkInterface interface
type MockSdkInterface struct {
	ctrl     *gomock.Controller
	recorder *MockSdkInterfaceMockRecorder
}

// MockSdkInterfaceMockRecorder is the mock recorder for MockSdkInterface
type MockSdkInterfaceMockRecorder struct {
	mock *MockSdkInterface
}

// NewMockSdkInterface creates a new mock instance
func NewMockSdkInterface(ctrl *gomock.Controller) *MockSdkInterface {
	mock := &MockSdkInterface{ctrl: ctrl}
	mock.recorder = &MockSdkInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSdkInterface) EXPECT() *MockSdkInterfaceMockRecorder {
	return m.recorder
}

// Delete mocks base method
func (m *MockSdkInterface) Delete(object sdk.Object) error {
	ret := m.ctrl.Call(m, "Delete", object)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockSdkInterfaceMockRecorder) Delete(object interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSdkInterface)(nil).Delete), object)
}

// Update mocks base method
func (m *MockSdkInterface) Update(object sdk.Object) error {
	ret := m.ctrl.Call(m, "Update", object)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockSdkInterfaceMockRecorder) Update(object interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSdkInterface)(nil).Update), object)
}

// Create mocks base method
func (m *MockSdkInterface) Create(object sdk.Object) error {
	ret := m.ctrl.Call(m, "Create", object)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockSdkInterfaceMockRecorder) Create(object interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSdkInterface)(nil).Create), object)
}

// Get mocks base method
func (m *MockSdkInterface) Get(object sdk.Object) error {
	ret := m.ctrl.Call(m, "Get", object)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockSdkInterfaceMockRecorder) Get(object interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSdkInterface)(nil).Get), object)
}

// Patch mocks base method
func (m *MockSdkInterface) Patch(object sdk.Object, pt types.PatchType, patch []byte) error {
	ret := m.ctrl.Call(m, "Patch", object, pt, patch)
	ret0, _ := ret[0].(error)
	return ret0
}

// Patch indicates an expected call of Patch
func (mr *MockSdkInterfaceMockRecorder) Patch(object, pt, patch interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockSdkInterface)(nil).Patch), object, pt, patch)
}

// List mocks base method
func (m *MockSdkInterface) List(namespace string, into sdk.Object, opts ...sdk.ListOption) error {
	ret := m.ctrl.Call(m, "List", namespace, into, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// List indicates an expected call of List
func (mr *MockSdkInterfaceMockRecorder) List(namespace, into, opts interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSdkInterface)(nil).List), namespace, into, opts)
}
