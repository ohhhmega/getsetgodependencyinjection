// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// DepositConfig is an autogenerated mock type for the DepositConfig type
type DepositConfig struct {
	mock.Mock
}

// GetDepositAPIKey provides a mock function with given fields:
func (_m *DepositConfig) GetDepositAPIKey() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetDepositAuthToken provides a mock function with given fields:
func (_m *DepositConfig) GetDepositAuthToken() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetDepositEndpoint provides a mock function with given fields:
func (_m *DepositConfig) GetDepositEndpoint() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewDepositConfig interface {
	mock.TestingT
	Cleanup(func())
}

// NewDepositConfig creates a new instance of DepositConfig. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDepositConfig(t mockConstructorTestingTNewDepositConfig) *DepositConfig {
	mock := &DepositConfig{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
