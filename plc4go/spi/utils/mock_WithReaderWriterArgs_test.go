/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// Code generated by mockery v2.32.4. DO NOT EDIT.

package utils

import mock "github.com/stretchr/testify/mock"

// MockWithReaderWriterArgs is an autogenerated mock type for the WithReaderWriterArgs type
type MockWithReaderWriterArgs struct {
	mock.Mock
}

type MockWithReaderWriterArgs_Expecter struct {
	mock *mock.Mock
}

func (_m *MockWithReaderWriterArgs) EXPECT() *MockWithReaderWriterArgs_Expecter {
	return &MockWithReaderWriterArgs_Expecter{mock: &_m.Mock}
}

// isReaderArgs provides a mock function with given fields:
func (_m *MockWithReaderWriterArgs) isReaderArgs() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockWithReaderWriterArgs_isReaderArgs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'isReaderArgs'
type MockWithReaderWriterArgs_isReaderArgs_Call struct {
	*mock.Call
}

// isReaderArgs is a helper method to define mock.On call
func (_e *MockWithReaderWriterArgs_Expecter) isReaderArgs() *MockWithReaderWriterArgs_isReaderArgs_Call {
	return &MockWithReaderWriterArgs_isReaderArgs_Call{Call: _e.mock.On("isReaderArgs")}
}

func (_c *MockWithReaderWriterArgs_isReaderArgs_Call) Run(run func()) *MockWithReaderWriterArgs_isReaderArgs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWithReaderWriterArgs_isReaderArgs_Call) Return(_a0 bool) *MockWithReaderWriterArgs_isReaderArgs_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWithReaderWriterArgs_isReaderArgs_Call) RunAndReturn(run func() bool) *MockWithReaderWriterArgs_isReaderArgs_Call {
	_c.Call.Return(run)
	return _c
}

// isWriterArgs provides a mock function with given fields:
func (_m *MockWithReaderWriterArgs) isWriterArgs() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockWithReaderWriterArgs_isWriterArgs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'isWriterArgs'
type MockWithReaderWriterArgs_isWriterArgs_Call struct {
	*mock.Call
}

// isWriterArgs is a helper method to define mock.On call
func (_e *MockWithReaderWriterArgs_Expecter) isWriterArgs() *MockWithReaderWriterArgs_isWriterArgs_Call {
	return &MockWithReaderWriterArgs_isWriterArgs_Call{Call: _e.mock.On("isWriterArgs")}
}

func (_c *MockWithReaderWriterArgs_isWriterArgs_Call) Run(run func()) *MockWithReaderWriterArgs_isWriterArgs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWithReaderWriterArgs_isWriterArgs_Call) Return(_a0 bool) *MockWithReaderWriterArgs_isWriterArgs_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWithReaderWriterArgs_isWriterArgs_Call) RunAndReturn(run func() bool) *MockWithReaderWriterArgs_isWriterArgs_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockWithReaderWriterArgs creates a new instance of MockWithReaderWriterArgs. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockWithReaderWriterArgs(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockWithReaderWriterArgs {
	mock := &MockWithReaderWriterArgs{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
