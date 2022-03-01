// Code generated by MockGen. DO NOT EDIT.
// Source: config.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	define "github.com/TencentBlueKing/bk-apigateway-sdks/bkapi-client-core/define"
	gomock "github.com/golang/mock/gomock"
)

// MockClientConfig is a mock of ClientConfig interface.
type MockClientConfig struct {
	ctrl     *gomock.Controller
	recorder *MockClientConfigMockRecorder
}

// MockClientConfigMockRecorder is the mock recorder for MockClientConfig.
type MockClientConfigMockRecorder struct {
	mock *MockClientConfig
}

// NewMockClientConfig creates a new mock instance.
func NewMockClientConfig(ctrl *gomock.Controller) *MockClientConfig {
	mock := &MockClientConfig{ctrl: ctrl}
	mock.recorder = &MockClientConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientConfig) EXPECT() *MockClientConfigMockRecorder {
	return m.recorder
}

// GetAuthorizationHeaders mocks base method.
func (m *MockClientConfig) GetAuthorizationHeaders() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthorizationHeaders")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetAuthorizationHeaders indicates an expected call of GetAuthorizationHeaders.
func (mr *MockClientConfigMockRecorder) GetAuthorizationHeaders() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthorizationHeaders", reflect.TypeOf((*MockClientConfig)(nil).GetAuthorizationHeaders))
}

// GetUrl mocks base method.
func (m *MockClientConfig) GetUrl() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUrl")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetUrl indicates an expected call of GetUrl.
func (mr *MockClientConfigMockRecorder) GetUrl() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUrl", reflect.TypeOf((*MockClientConfig)(nil).GetUrl))
}

// MockClientConfigProvider is a mock of ClientConfigProvider interface.
type MockClientConfigProvider struct {
	ctrl     *gomock.Controller
	recorder *MockClientConfigProviderMockRecorder
}

// MockClientConfigProviderMockRecorder is the mock recorder for MockClientConfigProvider.
type MockClientConfigProviderMockRecorder struct {
	mock *MockClientConfigProvider
}

// NewMockClientConfigProvider creates a new mock instance.
func NewMockClientConfigProvider(ctrl *gomock.Controller) *MockClientConfigProvider {
	mock := &MockClientConfigProvider{ctrl: ctrl}
	mock.recorder = &MockClientConfigProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientConfigProvider) EXPECT() *MockClientConfigProviderMockRecorder {
	return m.recorder
}

// Config mocks base method.
func (m *MockClientConfigProvider) Config(apiName string) define.ClientConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config", apiName)
	ret0, _ := ret[0].(define.ClientConfig)
	return ret0
}

// Config indicates an expected call of Config.
func (mr *MockClientConfigProviderMockRecorder) Config(apiName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockClientConfigProvider)(nil).Config), apiName)
}
