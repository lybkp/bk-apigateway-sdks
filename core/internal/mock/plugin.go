/**
 * TencentBlueKing is pleased to support the open source community by
 * making 蓝鲸智云-蓝鲸 PaaS 平台(BlueKing-PaaS) available.
 * Copyright (C) 2017 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by MockGen. DO NOT EDIT.
// Source: gopkg.in/h2non/gentleman.v2/plugin (interfaces: Plugin)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	context "gopkg.in/h2non/gentleman.v2/context"
)

// MockPlugin is a mock of Plugin interface.
type MockPlugin struct {
	ctrl     *gomock.Controller
	recorder *MockPluginMockRecorder
}

// MockPluginMockRecorder is the mock recorder for MockPlugin.
type MockPluginMockRecorder struct {
	mock *MockPlugin
}

// NewMockPlugin creates a new mock instance.
func NewMockPlugin(ctrl *gomock.Controller) *MockPlugin {
	mock := &MockPlugin{ctrl: ctrl}
	mock.recorder = &MockPluginMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPlugin) EXPECT() *MockPluginMockRecorder {
	return m.recorder
}

// Disable mocks base method.
func (m *MockPlugin) Disable() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Disable")
}

// Disable indicates an expected call of Disable.
func (mr *MockPluginMockRecorder) Disable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disable", reflect.TypeOf((*MockPlugin)(nil).Disable))
}

// Disabled mocks base method.
func (m *MockPlugin) Disabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Disabled indicates an expected call of Disabled.
func (mr *MockPluginMockRecorder) Disabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disabled", reflect.TypeOf((*MockPlugin)(nil).Disabled))
}

// Enable mocks base method.
func (m *MockPlugin) Enable() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Enable")
}

// Enable indicates an expected call of Enable.
func (mr *MockPluginMockRecorder) Enable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enable", reflect.TypeOf((*MockPlugin)(nil).Enable))
}

// Exec mocks base method.
func (m *MockPlugin) Exec(arg0 string, arg1 *context.Context, arg2 context.Handler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Exec", arg0, arg1, arg2)
}

// Exec indicates an expected call of Exec.
func (mr *MockPluginMockRecorder) Exec(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"Exec",
		reflect.TypeOf((*MockPlugin)(nil).Exec),
		arg0,
		arg1,
		arg2,
	)
}

// Remove mocks base method.
func (m *MockPlugin) Remove() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Remove")
}

// Remove indicates an expected call of Remove.
func (mr *MockPluginMockRecorder) Remove() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockPlugin)(nil).Remove))
}

// Removed mocks base method.
func (m *MockPlugin) Removed() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Removed")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Removed indicates an expected call of Removed.
func (mr *MockPluginMockRecorder) Removed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Removed", reflect.TypeOf((*MockPlugin)(nil).Removed))
}
