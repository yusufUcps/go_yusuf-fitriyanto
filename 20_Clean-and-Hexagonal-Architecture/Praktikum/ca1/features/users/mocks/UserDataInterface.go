// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	users "restEcho1/features/users"

	mock "github.com/stretchr/testify/mock"
)

// UserDataInterface is an autogenerated mock type for the UserDataInterface type
type UserDataInterface struct {
	mock.Mock
}

// GetAllUsers provides a mock function with given fields:
func (_m *UserDataInterface) GetAllUsers() ([]users.User, error) {
	ret := _m.Called()

	var r0 []users.User
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]users.User, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []users.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]users.User)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: newData
func (_m *UserDataInterface) Insert(newData users.User) (*users.User, error) {
	ret := _m.Called(newData)

	var r0 *users.User
	var r1 error
	if rf, ok := ret.Get(0).(func(users.User) (*users.User, error)); ok {
		return rf(newData)
	}
	if rf, ok := ret.Get(0).(func(users.User) *users.User); ok {
		r0 = rf(newData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.User)
		}
	}

	if rf, ok := ret.Get(1).(func(users.User) error); ok {
		r1 = rf(newData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: hp, password
func (_m *UserDataInterface) Login(hp string, password string) (*users.User, error) {
	ret := _m.Called(hp, password)

	var r0 *users.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*users.User, error)); ok {
		return rf(hp, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) *users.User); ok {
		r0 = rf(hp, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(hp, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserDataInterface creates a new instance of UserDataInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserDataInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserDataInterface {
	mock := &UserDataInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
