// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/weijiangan/bruno-test/brunotest (interfaces: AppClient,App_QueryClient)

// Package mock_brunotest is a generated GoMock package.
package mock_brunotest

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	brunotest "github.com/weijiangan/bruno-test/brunotest"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	reflect "reflect"
)

// MockAppClient is a mock of AppClient interface
type MockAppClient struct {
	ctrl     *gomock.Controller
	recorder *MockAppClientMockRecorder
}

// MockAppClientMockRecorder is the mock recorder for MockAppClient
type MockAppClientMockRecorder struct {
	mock *MockAppClient
}

// NewMockAppClient creates a new mock instance
func NewMockAppClient(ctrl *gomock.Controller) *MockAppClient {
	mock := &MockAppClient{ctrl: ctrl}
	mock.recorder = &MockAppClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAppClient) EXPECT() *MockAppClientMockRecorder {
	return m.recorder
}

// Query mocks base method
func (m *MockAppClient) Query(arg0 context.Context, arg1 *brunotest.QueryParam, arg2 ...grpc.CallOption) (brunotest.App_QueryClient, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(brunotest.App_QueryClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query
func (mr *MockAppClientMockRecorder) Query(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockAppClient)(nil).Query), varargs...)
}

// Send mocks base method
func (m *MockAppClient) Send(arg0 context.Context, arg1 *brunotest.AuditEvent, arg2 ...grpc.CallOption) (*brunotest.Response, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Send", varargs...)
	ret0, _ := ret[0].(*brunotest.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Send indicates an expected call of Send
func (mr *MockAppClientMockRecorder) Send(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockAppClient)(nil).Send), varargs...)
}

// MockApp_QueryClient is a mock of App_QueryClient interface
type MockApp_QueryClient struct {
	ctrl     *gomock.Controller
	recorder *MockApp_QueryClientMockRecorder
}

// MockApp_QueryClientMockRecorder is the mock recorder for MockApp_QueryClient
type MockApp_QueryClientMockRecorder struct {
	mock *MockApp_QueryClient
}

// NewMockApp_QueryClient creates a new mock instance
func NewMockApp_QueryClient(ctrl *gomock.Controller) *MockApp_QueryClient {
	mock := &MockApp_QueryClient{ctrl: ctrl}
	mock.recorder = &MockApp_QueryClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApp_QueryClient) EXPECT() *MockApp_QueryClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method
func (m *MockApp_QueryClient) CloseSend() error {
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockApp_QueryClientMockRecorder) CloseSend() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockApp_QueryClient)(nil).CloseSend))
}

// Context mocks base method
func (m *MockApp_QueryClient) Context() context.Context {
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockApp_QueryClientMockRecorder) Context() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockApp_QueryClient)(nil).Context))
}

// Header mocks base method
func (m *MockApp_QueryClient) Header() (metadata.MD, error) {
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *MockApp_QueryClientMockRecorder) Header() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockApp_QueryClient)(nil).Header))
}

// Recv mocks base method
func (m *MockApp_QueryClient) Recv() (*brunotest.AuditEvent, error) {
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*brunotest.AuditEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockApp_QueryClientMockRecorder) Recv() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockApp_QueryClient)(nil).Recv))
}

// RecvMsg mocks base method
func (m *MockApp_QueryClient) RecvMsg(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockApp_QueryClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockApp_QueryClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method
func (m *MockApp_QueryClient) SendMsg(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockApp_QueryClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockApp_QueryClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method
func (m *MockApp_QueryClient) Trailer() metadata.MD {
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *MockApp_QueryClientMockRecorder) Trailer() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockApp_QueryClient)(nil).Trailer))
}
