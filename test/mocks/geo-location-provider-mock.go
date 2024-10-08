// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/contracts/geo-location-provider.go
//
// Generated by this command:
//
//	mockgen -source=internal/app/contracts/geo-location-provider.go -destination=test/mocks/geo-location-provider-mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockGeoLocationProvider is a mock of GeoLocationProvider interface.
type MockGeoLocationProvider struct {
	ctrl     *gomock.Controller
	recorder *MockGeoLocationProviderMockRecorder
}

// MockGeoLocationProviderMockRecorder is the mock recorder for MockGeoLocationProvider.
type MockGeoLocationProviderMockRecorder struct {
	mock *MockGeoLocationProvider
}

// NewMockGeoLocationProvider creates a new mock instance.
func NewMockGeoLocationProvider(ctrl *gomock.Controller) *MockGeoLocationProvider {
	mock := &MockGeoLocationProvider{ctrl: ctrl}
	mock.recorder = &MockGeoLocationProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGeoLocationProvider) EXPECT() *MockGeoLocationProviderMockRecorder {
	return m.recorder
}

// GetInfo mocks base method.
func (m *MockGeoLocationProvider) GetInfo(ip string) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfo", ip)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetInfo indicates an expected call of GetInfo.
func (mr *MockGeoLocationProviderMockRecorder) GetInfo(ip any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfo", reflect.TypeOf((*MockGeoLocationProvider)(nil).GetInfo), ip)
}
