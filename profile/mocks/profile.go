// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	profile "bitbucket.org/junglee_games/getsetgo/profile"
	mock "github.com/stretchr/testify/mock"
)

// Profile is an autogenerated mock type for the Profile type
type Profile struct {
	mock.Mock
}

// GetUserByID provides a mock function with given fields: userId
func (_m *Profile) GetUserByID(userId int) (*profile.ProfileResponse, error) {
	ret := _m.Called(userId)

	var r0 *profile.ProfileResponse
	if rf, ok := ret.Get(0).(func(int) *profile.ProfileResponse); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*profile.ProfileResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewProfile interface {
	mock.TestingT
	Cleanup(func())
}

// NewProfile creates a new instance of Profile. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProfile(t mockConstructorTestingTNewProfile) *Profile {
	mock := &Profile{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}