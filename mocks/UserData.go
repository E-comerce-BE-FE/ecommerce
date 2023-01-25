// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	user "ecommerce/features/user"

	mock "github.com/stretchr/testify/mock"
)

// UserData is an autogenerated mock type for the UserData type
type UserData struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *UserData) Delete(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Login provides a mock function with given fields: email
func (_m *UserData) Login(email string) (user.Core, error) {
	ret := _m.Called(email)

	var r0 user.Core
	if rf, ok := ret.Get(0).(func(string) user.Core); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(user.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Profile provides a mock function with given fields:
func (_m *UserData) Profile() (interface{}, error) {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: newUser
func (_m *UserData) Register(newUser user.Core) (user.Core, error) {
	ret := _m.Called(newUser)

	var r0 user.Core
	if rf, ok := ret.Get(0).(func(user.Core) user.Core); ok {
		r0 = rf(newUser)
	} else {
		r0 = ret.Get(0).(user.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(user.Core) error); ok {
		r1 = rf(newUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, updateData
func (_m *UserData) Update(id uint, updateData user.Core) (user.Core, error) {
	ret := _m.Called(id, updateData)

	var r0 user.Core
	if rf, ok := ret.Get(0).(func(uint, user.Core) user.Core); ok {
		r0 = rf(id, updateData)
	} else {
		r0 = ret.Get(0).(user.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, user.Core) error); ok {
		r1 = rf(id, updateData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserData interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserData creates a new instance of UserData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserData(t mockConstructorTestingTNewUserData) *UserData {
	mock := &UserData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
