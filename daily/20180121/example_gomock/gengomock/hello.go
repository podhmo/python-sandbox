// Code generated by MockGen. DO NOT EDIT.
// Source: hello/hello.go

// Package gengomock is a generated GoMock package.
package gengomock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockHello is a mock of Hello interface
type MockHello struct {
	ctrl     *gomock.Controller
	recorder *MockHelloMockRecorder
}

// MockHelloMockRecorder is the mock recorder for MockHello
type MockHelloMockRecorder struct {
	mock *MockHello
}

// NewMockHello creates a new mock instance
func NewMockHello(ctrl *gomock.Controller) *MockHello {
	mock := &MockHello{ctrl: ctrl}
	mock.recorder = &MockHelloMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHello) EXPECT() *MockHelloMockRecorder {
	return m.recorder
}

// Hello mocks base method
func (m *MockHello) Hello(s string) {
	m.ctrl.Call(m, "Hello", s)
}

// Hello indicates an expected call of Hello
func (mr *MockHelloMockRecorder) Hello(s interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hello", reflect.TypeOf((*MockHello)(nil).Hello), s)
}