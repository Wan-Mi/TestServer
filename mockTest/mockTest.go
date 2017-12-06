// Code generated by MockGen. DO NOT EDIT.
// Source: mockTest/mockDemo.go

// Package mockTest is a generated GoMock package.
package mockTest

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockmockDemo is a mock of mockDemo interface
type MockmockDemo struct {
	ctrl     *gomock.Controller
	recorder *MockmockDemoMockRecorder
}

// MockmockDemoMockRecorder is the mock recorder for MockmockDemo
type MockmockDemoMockRecorder struct {
	mock *MockmockDemo
}

// NewMockmockDemo creates a new mock instance
func NewMockmockDemo(ctrl *gomock.Controller) *MockmockDemo {
	mock := &MockmockDemo{ctrl: ctrl}
	mock.recorder = &MockmockDemoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockmockDemo) EXPECT() *MockmockDemoMockRecorder {
	return m.recorder
}

// GetString mocks base method
func (m *MockmockDemo) GetString() string {
	ret := m.ctrl.Call(m, "GetString")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetString indicates an expected call of GetString
func (mr *MockmockDemoMockRecorder) GetString() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetString", reflect.TypeOf((*MockmockDemo)(nil).GetString))
}
