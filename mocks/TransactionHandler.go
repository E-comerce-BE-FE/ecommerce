// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// TransactionHandler is an autogenerated mock type for the TransactionHandler type
type TransactionHandler struct {
	mock.Mock
}

// CreateTransaction provides a mock function with given fields:
func (_m *TransactionHandler) CreateTransaction() echo.HandlerFunc {
	ret := _m.Called()

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

// UpdateTransaction provides a mock function with given fields:
func (_m *TransactionHandler) UpdateTransaction() echo.HandlerFunc {
	ret := _m.Called()

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

type mockConstructorTestingTNewTransactionHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewTransactionHandler creates a new instance of TransactionHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransactionHandler(t mockConstructorTestingTNewTransactionHandler) *TransactionHandler {
	mock := &TransactionHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
