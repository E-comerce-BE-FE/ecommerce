// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	product "ecommerce/features/product"

	mock "github.com/stretchr/testify/mock"
)

// ProductData is an autogenerated mock type for the ProductData type
type ProductData struct {
	mock.Mock
}

// AddProduct provides a mock function with given fields: userID, newProduct
func (_m *ProductData) AddProduct(userID uint, newProduct product.Core) (product.Core, error) {
	ret := _m.Called(userID, newProduct)

	var r0 product.Core
	if rf, ok := ret.Get(0).(func(uint, product.Core) product.Core); ok {
		r0 = rf(userID, newProduct)
	} else {
		r0 = ret.Get(0).(product.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, product.Core) error); ok {
		r1 = rf(userID, newProduct)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AllProduct provides a mock function with given fields:
func (_m *ProductData) AllProduct() ([]product.Core, error) {
	ret := _m.Called()

	var r0 []product.Core
	if rf, ok := ret.Get(0).(func() []product.Core); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]product.Core)
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

// Delete provides a mock function with given fields: userID, productID
func (_m *ProductData) Delete(userID uint, productID uint) error {
	ret := _m.Called(userID, productID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint) error); ok {
		r0 = rf(userID, productID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EditProduct provides a mock function with given fields: userID, productID, editedProduct
func (_m *ProductData) EditProduct(userID uint, productID uint, editedProduct product.Core) (product.Core, error) {
	ret := _m.Called(userID, productID, editedProduct)

	var r0 product.Core
	if rf, ok := ret.Get(0).(func(uint, uint, product.Core) product.Core); ok {
		r0 = rf(userID, productID, editedProduct)
	} else {
		r0 = ret.Get(0).(product.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, uint, product.Core) error); ok {
		r1 = rf(userID, productID, editedProduct)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProductDetail provides a mock function with given fields: productID
func (_m *ProductData) ProductDetail(productID uint) (product.Core, error) {
	ret := _m.Called(productID)

	var r0 product.Core
	if rf, ok := ret.Get(0).(func(uint) product.Core); ok {
		r0 = rf(productID)
	} else {
		r0 = ret.Get(0).(product.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(productID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewProductData interface {
	mock.TestingT
	Cleanup(func())
}

// NewProductData creates a new instance of ProductData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProductData(t mockConstructorTestingTNewProductData) *ProductData {
	mock := &ProductData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
