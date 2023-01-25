package services

import (
	"ecommerce/features/cart"
	"ecommerce/helper"
	"ecommerce/mocks"
	"errors"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestAddToCart(t *testing.T) {
	data := mocks.NewCartData(t)
	input := cart.Core{ID: 1, ProductName: "Freshtea", Seller: "griffin", Qty: 1, Amount: 5000}
	resData := cart.Core{ID: 1, ProductName: "Freshtea", Seller: "griffin", Qty: 1, Amount: 5000}

	t.Run("Success add cart", func(t *testing.T) {
		data.On("AddToCart", uint(1), uint(1), input).Return(resData, nil).Once()

		srv := New(data)
		_, token := helper.GenerateToken(1)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		res, err := srv.AddToCart(useToken, 1, input)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		assert.Equal(t, resData.ProductName, res.ProductName)
		data.AssertExpectations(t)
	})

	t.Run("Cart not found", func(t *testing.T) {
		data.On("AddToCart", uint(5), uint(1), input).Return(cart.Core{}, errors.New("data not found")).Once()

		srv := New(data)
		_, token := helper.GenerateToken(5)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		res, err := srv.AddToCart(useToken, 1, input)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "error")
		data.AssertExpectations(t)
	})

	t.Run("Trouble in server", func(t *testing.T) {
		data.On("AddToCart", uint(1), uint(1), input).Return(cart.Core{}, errors.New("server error")).Once()

		srv := New(data)
		_, token := helper.GenerateToken(1)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		res, err := srv.AddToCart(useToken, 1, input)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "error")
		data.AssertExpectations(t)
	})

	t.Run("JWT not valid", func(t *testing.T) {
		data.On("AddToCart", uint(1), uint(1), input).Return(cart.Core{}, errors.New("jwt not valid")).Once()

		srv := New(data)
		_, token := helper.GenerateToken(1)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		res, err := srv.AddToCart(useToken, 1, input)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "error")
		data.AssertExpectations(t)
	})
}

func TestCartList(t *testing.T) {
	data := mocks.NewCartData(t)
	resData := []cart.Core{
		{
			ID:          1,
			ProductName: "Freshtea",
			Seller:      "griffin",
			Qty:         1,
			Amount:      5000,
		},
	}

	t.Run("Success show cart", func(t *testing.T) {
		data.On("CartList", uint(1)).Return(resData, nil).Once()

		srv := New(data)
		_, token := helper.GenerateToken(1)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		res, err := srv.CartList(useToken)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		data.AssertExpectations(t)
	})

	t.Run("Data not found", func(t *testing.T) {
		data.On("CartList", uint(1)).Return([]cart.Core{}, errors.New("data not found")).Once()

		srv := New(data)
		_, token := helper.GenerateToken(1)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		res, err := srv.CartList(useToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "error")
		assert.Empty(t, res)
		data.AssertExpectations(t)
	})

	t.Run("Trouble in server", func(t *testing.T) {
		data.On("CartList", uint(1)).Return([]cart.Core{}, errors.New("server error")).Once()

		srv := New(data)
		_, token := helper.GenerateToken(1)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		res, err := srv.CartList(useToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "error")
		assert.Empty(t, res)
		data.AssertExpectations(t)
	})

	t.Run("JWT not valid", func(t *testing.T) {
		data.On("CartList", uint(1)).Return([]cart.Core{}, errors.New("jwt not valid")).Once()

		srv := New(data)
		_, token := helper.GenerateToken(1)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		res, err := srv.CartList(useToken)
		assert.NotNil(t, err)
		assert.Empty(t, res)
		assert.ErrorContains(t, err, "error")
		data.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	data := mocks.NewCartData(t)

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

	t.Run("Data not found", func(t *testing.T) {
		data.On("Delete", uint(5), uint(1)).Return(errors.New("data not found")).Once()

		srv := New(data)
		_, token := helper.GenerateToken(5)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		err := srv.Delete(useToken, 1)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "error")
		data.AssertExpectations(t)
	})

}

func TestUpdateQty(t *testing.T) {
	data := mocks.NewCartData(t)
	qty := 1
	resData := cart.Core{Qty: 2}

	t.Run("Success update data", func(t *testing.T) {
		data.On("UpdateQty", uint(1), uint(1), qty).Return(resData, nil).Once()

		srv := New(data)
		_, token := helper.GenerateToken(1)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		res, err := srv.UpdateQty(useToken, 1, qty)
		assert.Nil(t, err)
		assert.Equal(t, resData.Qty, res.Qty)
		data.AssertExpectations(t)
	})

	t.Run("Data not found", func(t *testing.T) {
		data.On("UpdateQty", uint(1), uint(5), qty).Return(cart.Core{}, errors.New("data not found")).Once()

		srv := New(data)
		_, token := helper.GenerateToken(1)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		res, err := srv.UpdateQty(useToken, 5, qty)
		assert.NotNil(t, err)
		assert.Equal(t, 0, res.Qty)
		assert.ErrorContains(t, err, "error")
		data.AssertExpectations(t)
	})

	t.Run("Tourble in server", func(t *testing.T) {
		data.On("UpdateQty", uint(1), uint(1), qty).Return(cart.Core{}, errors.New("server error")).Once()

		srv := New(data)
		_, token := helper.GenerateToken(1)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		res, err := srv.UpdateQty(useToken, 1, qty)
		assert.NotNil(t, err)
		assert.Equal(t, 0, res.Qty)
		assert.ErrorContains(t, err, "error")
		data.AssertExpectations(t)
	})
}
