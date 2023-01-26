package services

import (
	"ecommerce/features/transaction"
	"ecommerce/helper"
	"ecommerce/mocks"
	"errors"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	data := mocks.NewTransactionData(t)
	pl := "paymentlink"
	tc := "trscode"
	resData := transaction.Core{ID: 1, SubTotal: 100000, TransactionCode: "trscode", Status: "pending"}

	t.Run("Success create transaction", func(t *testing.T) {
		data.On("CreateTransaction", uint(1), pl, tc).Return(resData, nil).Once()

		srv := New(data)
		_, token := helper.GenerateToken(1)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		res, err := srv.CreateTransaction(useToken, pl, tc)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		data.AssertExpectations(t)
	})

	t.Run("Server error", func(t *testing.T) {
		data.On("CreateTransaction", uint(1), pl, tc).Return(transaction.Core{}, errors.New("server error")).Once()

		srv := New(data)
		_, token := helper.GenerateToken(1)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		res, err := srv.CreateTransaction(useToken, pl, tc)
		assert.NotNil(t, err)
		assert.NotEqual(t, resData.ID, res.ID)
		assert.ErrorContains(t, err, "error")
		data.AssertExpectations(t)
	})
}

func TestUpdateTransaction(t *testing.T) {
	data := mocks.NewTransactionData(t)
	tc := "trscode"
	resData := transaction.Core{Status: "success"}

	t.Run("Success update status transaction", func(t *testing.T) {
		data.On("UpdateTransaction", tc).Return(resData, nil).Once()

		srv := New(data)
		res, err := srv.UpdateTransaction(tc)
		assert.Nil(t, err)
		assert.Equal(t, resData.TransactionCode, res.TransactionCode)
		data.AssertExpectations(t)
	})

	t.Run("Server error", func(t *testing.T) {
		data.On("UpdateTransaction", tc).Return(transaction.Core{}, errors.New("server error")).Once()

		srv := New(data)
		res, err := srv.UpdateTransaction(tc)
		assert.NotNil(t, err)
		assert.Empty(t, res.TransactionCode)
		assert.ErrorContains(t, err, "error")
		data.AssertExpectations(t)
	})
}

func TestTransactionHistory(t *testing.T) {
	data := mocks.NewTransactionData(t)
	resData := []transaction.Core{
		{
			ID:              1,
			TotalProduct:    5,
			SubTotal:        100000,
			TransactionCode: "trscode",
		},
	}

	t.Run("Success show history transaction", func(t *testing.T) {
		data.On("TransactionHistory", uint(1)).Return(resData, nil).Once()

		srv := New(data)
		_, token := helper.GenerateToken(1)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		res, err := srv.TransactionHistory(useToken)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		data.AssertExpectations(t)
	})

	t.Run("Server error", func(t *testing.T) {
		data.On("TransactionHistory", uint(1)).Return([]transaction.Core{}, errors.New("server error")).Once()

		srv := New(data)
		_, token := helper.GenerateToken(1)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		res, err := srv.TransactionHistory(useToken)
		assert.NotNil(t, err)
		assert.Empty(t, res)
		data.AssertExpectations(t)
	})
}

func TestCancelTransaction(t *testing.T) {
	data := mocks.NewTransactionData(t)
	t.Run("Success Cancel", func(t *testing.T) {
		data.On("CancelTransaction", uint(1), uint(1)).Return(nil).Once()
		srv := New(data)
		_, token := helper.GenerateToken(1)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		err := srv.CancelTransaction(useToken, uint(1))
		assert.Nil(t, err)
		data.AssertExpectations(t)
	})
	t.Run("Cancel Transaction Fail", func(t *testing.T) {
		data.On("CancelTransaction", uint(3), uint(1)).Return(errors.New("server error")).Once()
		srv := New(data)
		_, token := helper.GenerateToken(3)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		err := srv.CancelTransaction(useToken, uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		data.AssertExpectations(t)
	})
}

func TestTransactionDetail(t *testing.T) {
	data := mocks.NewTransactionData(t)
	resData := transaction.Core{ID: 1, TransactionName: "product-shoping", SubTotal: 35000}
	t.Run("Success show transaction detail", func(t *testing.T) {
		data.On("TransactionDetail", uint(1), uint(1)).Return(resData, nil).Once()
		srv := New(data)
		_, token := helper.GenerateToken(1)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		res, err := srv.TransactionDetail(useToken, uint(1))
		assert.Nil(t, err)
		assert.Equal(t, res, resData)
		data.AssertExpectations(t)
	})
	t.Run("Show Transaction detail Fail", func(t *testing.T) {
		data.On("TransactionDetail", uint(1), uint(1)).Return(transaction.Core{}, errors.New("server error")).Once()
		srv := New(data)
		_, token := helper.GenerateToken(1)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		res, err := srv.TransactionDetail(useToken, uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, transaction.Core{}, res)
		data.AssertExpectations(t)
	})
}
