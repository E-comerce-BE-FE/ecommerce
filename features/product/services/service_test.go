package services

import (
	"ecommerce/features/product"
	"ecommerce/helper"
	"ecommerce/mocks"
	"errors"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestAddProduct(t *testing.T) {
	data := mocks.NewProductData(t)
	input := product.Core{ID: 1, ProductName: "Kitkat", Price: 10000, Stock: 5}
	resData := product.Core{ID: 1, ProductName: "Kitkat", Price: 10000, Stock: 5}

	t.Run("Success add product", func(t *testing.T) {
		data.On("AddProduct", 1, input).Return(resData, nil).Once()

		// srv := New(data)
		// _, token := helper.GenerateToken(1)
		// useToken := token.(*jwt.Token)
		// useToken.Valid = true
		// res, err := srv.AddProduct(useToken,)
	})
}

func TestAllProduct(t *testing.T) {
	data := mocks.NewProductData(t)
	resData := []product.Core{
		{
			ID:          1,
			ProductName: "Kitkat",
			Price:       10000,
			Stock:       1,
			Description: "biskuit yang dilapisi dengan coklat lembut",
		},
	}

	t.Run("Succes show product", func(t *testing.T) {
		data.On("AllProduct").Return(resData, nil).Once()

		srv := New(data)
		res, err := srv.AllProduct()
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		data.AssertExpectations(t)
	})

	t.Run("Trouble in server", func(t *testing.T) {
		data.On("AllProduct").Return([]product.Core{}, errors.New("server error")).Once()

		srv := New(data)
		res, err := srv.AllProduct()
		assert.NotNil(t, err)
		assert.Empty(t, res)
		assert.ErrorContains(t, err, "error")
		data.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	data := mocks.NewProductData(t)

	t.Run("Success delete data", func(t *testing.T) {
		data.On("Delete", uint(1), uint(1)).Return(nil).Once()

		srv := New(data)
		_, token := helper.GenerateToken(1)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		err := srv.Delete(useToken, 1)
		assert.Nil(t, err)
		data.AssertExpectations(t)
	})

	t.Run("Trouble in server", func(t *testing.T) {
		data.On("Delete", uint(1), uint(1)).Return(errors.New("server error")).Once()

		srv := New(data)
		_, token := helper.GenerateToken(1)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		err := srv.Delete(useToken, 1)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "error")
		data.AssertExpectations(t)
	})
}

func TestEditProduct(t *testing.T) {
	data := mocks.NewProductData(t)
	input := product.Core{ID: 1, ProductName: "Kitkat", Price: 10000, Stock: 3}
	resData := product.Core{ID: 1, ProductName: "Tehbotol", Price: 5000, Stock: 1}

	t.Run("Success update data", func(t *testing.T) {
		data.On("EditProduct", uint(1), uint(1), input).Return(resData, nil).Once()

		// srv := New(data)
		// _, token := helper.GenerateToken(1)
		// useToken := token.(*jwt.Token)
		// useToken.Valid = true
		// res, err := srv.EditProduct(useToken)
	})
}

func TestProductDetail(t *testing.T) {
	data := mocks.NewProductData(t)
	resData := product.Core{ID: 1, ProductName: "Tehbotol", Price: 5000, Stock: 1}

	t.Run("Success show product detail", func(t *testing.T) {
		data.On("ProductDetail", uint(1)).Return(resData, nil).Once()

		srv := New(data)
		_, token := helper.GenerateToken(1)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		res, err := srv.ProductDetail(1)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		data.AssertExpectations(t)
	})

	t.Run("Trouble in server", func(t *testing.T) {
		data.On("ProductDetail", uint(1)).Return(product.Core{}, errors.New("server error")).Once()

		srv := New(data)
		_, token := helper.GenerateToken(1)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		res, err := srv.ProductDetail(1)
		assert.NotNil(t, err)
		assert.Empty(t, res)
		assert.ErrorContains(t, err, "error")
		data.AssertExpectations(t)
	})
}
