// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/tozastation/gRPC-Training-Golang/interfaces/rpc/user (interfaces: UsersClient)

// Package mock_user is a generated GoMock package.
package mock_user

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	user "github.com/tozastation/gRPC-Training-Golang/interfaces/rpc/user"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockUsersClient is a mock of UsersClient interface
type MockUsersClient struct {
	ctrl     *gomock.Controller
	recorder *MockUsersClientMockRecorder
}

// MockUsersClientMockRecorder is the mock recorder for MockUsersClient
type MockUsersClientMockRecorder struct {
	mock *MockUsersClient
}

// NewMockUsersClient creates a new mock instance
func NewMockUsersClient(ctrl *gomock.Controller) *MockUsersClient {
	mock := &MockUsersClient{ctrl: ctrl}
	mock.recorder = &MockUsersClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUsersClient) EXPECT() *MockUsersClientMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockUsersClient) Get(arg0 context.Context, arg1 *user.GetRequest, arg2 ...grpc.CallOption) (*user.GetResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*user.GetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockUsersClientMockRecorder) Get(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUsersClient)(nil).Get), varargs...)
}

// Login mocks base method
func (m *MockUsersClient) Login(arg0 context.Context, arg1 *user.LoginRequest, arg2 ...grpc.CallOption) (*user.LoginResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Login", varargs...)
	ret0, _ := ret[0].(*user.LoginResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login
func (mr *MockUsersClientMockRecorder) Login(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUsersClient)(nil).Login), varargs...)
}

// Post mocks base method
func (m *MockUsersClient) Post(arg0 context.Context, arg1 *user.PostRequest, arg2 ...grpc.CallOption) (*user.PostResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Post", varargs...)
	ret0, _ := ret[0].(*user.PostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Post indicates an expected call of Post
func (mr *MockUsersClientMockRecorder) Post(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockUsersClient)(nil).Post), varargs...)
}
