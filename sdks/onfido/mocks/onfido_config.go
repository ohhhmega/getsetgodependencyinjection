// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// OnfidoConfig is an autogenerated mock type for the OnfidoConfig type
type OnfidoConfig struct {
	mock.Mock
}

// GetOnfidoAuthToken provides a mock function with given fields:
func (_m *OnfidoConfig) GetOnfidoAuthToken() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetOnfidoEndpoint provides a mock function with given fields:
func (_m *OnfidoConfig) GetOnfidoEndpoint() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewOnfidoConfig interface {
	mock.TestingT
	Cleanup(func())
}

// NewOnfidoConfig creates a new instance of OnfidoConfig. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOnfidoConfig(t mockConstructorTestingTNewOnfidoConfig) *OnfidoConfig {
	mock := &OnfidoConfig{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
