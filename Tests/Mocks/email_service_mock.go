// Code generated by MockGen. DO NOT EDIT.
// Source: Domain/Interfaces/email_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockEmailService is a mock of EmailService interface.
type MockEmailService struct {
	ctrl     *gomock.Controller
	recorder *MockEmailServiceMockRecorder
}

// MockEmailServiceMockRecorder is the mock recorder for MockEmailService.
type MockEmailServiceMockRecorder struct {
	mock *MockEmailService
}

// NewMockEmailService creates a new mock instance.
func NewMockEmailService(ctrl *gomock.Controller) *MockEmailService {
	mock := &MockEmailService{ctrl: ctrl}
	mock.recorder = &MockEmailServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmailService) EXPECT() *MockEmailServiceMockRecorder {
	return m.recorder
}

// IsValidEmail mocks base method.
func (m *MockEmailService) IsValidEmail(email string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValidEmail", email)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsValidEmail indicates an expected call of IsValidEmail.
func (mr *MockEmailServiceMockRecorder) IsValidEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValidEmail", reflect.TypeOf((*MockEmailService)(nil).IsValidEmail), email)
}
