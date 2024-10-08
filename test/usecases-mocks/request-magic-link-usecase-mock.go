// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/usecases/session-usecases.go
//
// Generated by this command:
//
//	mockgen -source=internal/app/usecases/session-usecases.go -destination=test/usecases-mocks/request-magic-link-usecase-mock.go -package=usecasesmocks
//

// Package usecasesmocks is a generated GoMock package.
package usecasesmocks

import (
	reflect "reflect"

	errors "github.com/henrique998/go-auth/internal/app/errors"
	gomock "go.uber.org/mock/gomock"
)

// MockRequestMagicLinkUseCase is a mock of RequestMagicLinkUseCase interface.
type MockRequestMagicLinkUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockRequestMagicLinkUseCaseMockRecorder
}

// MockRequestMagicLinkUseCaseMockRecorder is the mock recorder for MockRequestMagicLinkUseCase.
type MockRequestMagicLinkUseCaseMockRecorder struct {
	mock *MockRequestMagicLinkUseCase
}

// NewMockRequestMagicLinkUseCase creates a new mock instance.
func NewMockRequestMagicLinkUseCase(ctrl *gomock.Controller) *MockRequestMagicLinkUseCase {
	mock := &MockRequestMagicLinkUseCase{ctrl: ctrl}
	mock.recorder = &MockRequestMagicLinkUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRequestMagicLinkUseCase) EXPECT() *MockRequestMagicLinkUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockRequestMagicLinkUseCase) Execute(email string) errors.AppErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", email)
	ret0, _ := ret[0].(errors.AppErr)
	return ret0
}

// Execute indicates an expected call of Execute.
func (mr *MockRequestMagicLinkUseCaseMockRecorder) Execute(email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockRequestMagicLinkUseCase)(nil).Execute), email)
}
