// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/contracts/login-attempts-repository.go
//
// Generated by this command:
//
//	mockgen -source=internal/app/contracts/login-attempts-repository.go -destination=test/mocks/login-attempts-repository-mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	entities "github.com/henrique998/go-auth/internal/app/entities"
	gomock "go.uber.org/mock/gomock"
)

// MockLoginAttemptsRepository is a mock of LoginAttemptsRepository interface.
type MockLoginAttemptsRepository struct {
	ctrl     *gomock.Controller
	recorder *MockLoginAttemptsRepositoryMockRecorder
}

// MockLoginAttemptsRepositoryMockRecorder is the mock recorder for MockLoginAttemptsRepository.
type MockLoginAttemptsRepositoryMockRecorder struct {
	mock *MockLoginAttemptsRepository
}

// NewMockLoginAttemptsRepository creates a new mock instance.
func NewMockLoginAttemptsRepository(ctrl *gomock.Controller) *MockLoginAttemptsRepository {
	mock := &MockLoginAttemptsRepository{ctrl: ctrl}
	mock.recorder = &MockLoginAttemptsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoginAttemptsRepository) EXPECT() *MockLoginAttemptsRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockLoginAttemptsRepository) Create(la entities.LoginAttempt) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", la)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockLoginAttemptsRepositoryMockRecorder) Create(la any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockLoginAttemptsRepository)(nil).Create), la)
}
