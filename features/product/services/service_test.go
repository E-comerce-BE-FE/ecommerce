package services

import (
	"ecommerce/features/product"
	"ecommerce/helper"
	"ecommerce/mocks"
	"errors"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestAddProduct(t *testing.T) {
	data := mocks.NewProductData(t)
	filePath := filepath.Join("..", "..", "..", "ERD.png")
	// imageFalse, _ := os.Open(filePath)
	// imageFalseCnv := &multipart.FileHeader{
	// 	Filename: imageFalse.Name(),
	// }
	imageTrue, err := os.Open(filePath)
	if err != nil {
		log.Println(err.Error())
	}
	imageTrueCnv := &multipart.FileHeader{
		Filename: imageTrue.Name(),
	}

	input := product.Core{ID: 1, ProductName: "Kitkat", Price: 10000, Stock: 5, ProductImage: imageTrueCnv.Filename}
	resData := product.Core{ID: 1, ProductName: "Kitkat", Price: 10000, Stock: 5, ProductImage: imageTrueCnv.Filename}

	t.Run("Success add product", func(t *testing.T) {
		data.On("AddProduct", uint(1), input).Return(resData, nil).Once()

		srv := New(data)
		_, token := helper.GenerateToken(1)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		res, err := srv.AddProduct(useToken, *imageTrueCnv, input)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		data.AssertExpectations(t)
	})

	t.Run("Trouble in server", func(t *testing.T) {
		data.On("AddProduct", uint(1), input).Return(product.Core{}, errors.New("server error")).Once()

		srv := New(data)
		_, token := helper.GenerateToken(1)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		res, err := srv.AddProduct(useToken, *imageTrueCnv, input)
		assert.NotNil(t, err)
		assert.NotEqual(t, resData.ID, res.ID)
		data.AssertExpectations(t)
	})

	t.Run("JWT not valid", func(t *testing.T) {
		data.On("AddProduct", uint(1), input).Return(product.Core{}, errors.New("jwt invalid")).Once()

		srv := New(data)
		_, token := helper.GenerateToken(1)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		res, err := srv.AddProduct(useToken, *imageTrueCnv, input)
		assert.NotNil(t, err)
		assert.NotEqual(t, resData.ID, res.ID)
		assert.ErrorContains(t, err, "error")
		data.AssertExpectations(t)
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

	t.Run("Data not found", func(t *testing.T) {
		data.On("Delete", uint(1), uint(2)).Return(errors.New("product cannot deleted")).Once()

		srv := New(data)
		_, token := helper.GenerateToken(1)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		err := srv.Delete(useToken, 2)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not allowed")
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
	filePath := filepath.Join("..", "..", "..", "ERD.png")
	// imageFalse, _ := os.Open(filePath)
	// imageFalseCnv := &multipart.FileHeader{
	// 	Filename: imageFalse.Name(),
	// }
	imageTrue, err := os.Open(filePath)
	if err != nil {
		log.Println(err.Error())
	}
	imageTrueCnv := &multipart.FileHeader{
		Filename: imageTrue.Name(),
	}

	input := product.Core{ID: 1, ProductName: "Kitkat", Price: 10000, Stock: 3, ProductImage: "ERD.png"}
	resData := product.Core{ID: 1, ProductName: "Tehbotol", Price: 5000, Stock: 1, ProductImage: imageTrueCnv.Filename}

	t.Run("Success update data", func(t *testing.T) {
		data.On("EditProduct", uint(1), uint(1), input).Return(resData, nil).Once()

		srv := New(data)
		_, token := helper.GenerateToken(1)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		res, err := srv.EditProduct(useToken, *imageTrueCnv, 1, input)
		assert.Nil(t, err)
		assert.NotEqual(t, input.ProductName, res.ProductName)
		data.AssertExpectations(t)
	})

	t.Run("Trouble in server", func(t *testing.T) {
		data.On("EditProduct", uint(1), uint(1), input).Return(product.Core{}, errors.New("server error")).Once()

		srv := New(data)
		_, token := helper.GenerateToken(1)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		res, err := srv.EditProduct(useToken, *imageTrueCnv, 1, input)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "error")
		data.AssertExpectations(t)
	})

	t.Run("JWT not valid", func(t *testing.T) {
		data.On("AddProduct", uint(1), input).Return(product.Core{}, errors.New("jwt invalid")).Once()

		srv := New(data)
		_, token := helper.GenerateToken(1)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		res, err := srv.AddProduct(useToken, *imageTrueCnv, input)
		assert.NotNil(t, err)
		assert.NotEqual(t, resData.ID, res.ID)
		assert.ErrorContains(t, err, "error")
		data.AssertExpectations(t)
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

func TestSearching(t *testing.T) {
	repo := mocks.NewProductData(t)
	resData := []product.Core{{ID: 1, ProductName: "Lifebuoy", Price: 4000}}
	t.Run("success Found", func(t *testing.T) {
		repo.On("Searching", "eko").Return(resData, nil)
		srv := New(repo)
		res, err := srv.Searching("eko")
		assert.Nil(t, err)
		assert.Equal(t, resData[0].ProductName, res[0].ProductName)
		repo.AssertExpectations(t)
	})
	t.Run("Not found", func(t *testing.T) {
		repo.On("Searching", "").Return([]product.Core{}, errors.New("no user found"))
		srv := New(repo)
		res, err := srv.Searching("")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "error")
		assert.Equal(t, []product.Core{}, res)
		repo.AssertExpectations(t)
	})
}
