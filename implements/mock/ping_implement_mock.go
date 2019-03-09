// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/tozastation/gRPC-Training-Golang/interfaces/rpc/ping (interfaces: CheckClient)

// Package mock_ping is a generated GoMock package.
package mock_ping

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	ping "github.com/tozastation/gRPC-Training-Golang/interfaces/rpc/ping"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockCheckClient is a mock of CheckClient interface
type MockCheckClient struct {
	ctrl     *gomock.Controller
	recorder *MockCheckClientMockRecorder
}

// MockCheckClientMockRecorder is the mock recorder for MockCheckClient
type MockCheckClientMockRecorder struct {
	mock *MockCheckClient
}

// NewMockCheckClient creates a new mock instance
func NewMockCheckClient(ctrl *gomock.Controller) *MockCheckClient {
	mock := &MockCheckClient{ctrl: ctrl}
	mock.recorder = &MockCheckClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCheckClient) EXPECT() *MockCheckClientMockRecorder {
	return m.recorder
}

// Ping mocks base method
func (m *MockCheckClient) Ping(arg0 context.Context, arg1 *ping.Empty, arg2 ...grpc.CallOption) (*ping.Pong, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Ping", varargs...)
	ret0, _ := ret[0].(*ping.Pong)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Ping indicates an expected call of Ping
func (mr *MockCheckClientMockRecorder) Ping(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockCheckClient)(nil).Ping), varargs...)
}
