package services

import (
	"ecommerce/features/product"
	"ecommerce/helper"
	"errors"
	"log"
	"mime/multipart"

	uuid "github.com/satori/go.uuid"
)

type productServiceCase struct {
	qry product.ProductData
}

func New(pd product.ProductData) product.ProductService {
	return &productServiceCase{
		qry: pd,
	}
}

// AddProduct implements product.ProductService
func (psc *productServiceCase) AddProduct(token interface{}, formHeader multipart.FileHeader, newProduct product.Core) (product.Core, error) {
	userID := helper.ExtractToken(token)
	//image proses
	if formHeader.Size != 0 {
		if formHeader.Size > 500000 {
			return product.Core{}, errors.New("size error")
		}
		fileName := uuid.NewV4().String()
		formHeader.Filename = fileName + formHeader.Filename[(len(formHeader.Filename)-5):len(formHeader.Filename)]
		src, err := formHeader.Open()
		if err != nil {
			return product.Core{}, errors.New("error open formheader")
		}
		defer src.Close()
		uploadURL, err := helper.UploadToS3(formHeader.Filename, src)
		if err != nil {
			return product.Core{}, errors.New("cannot upload to s3 server error")
		}
		newProduct.ProductImage = uploadURL
	}
	res, err := psc.qry.AddProduct(uint(userID), newProduct)
	if err != nil {
		return product.Core{}, errors.New("query error cannot, problem with server")
	}
	return res, nil
}

// AllProduct implements product.ProductService
func (psc *productServiceCase) AllProduct() ([]product.Core, error) {
	res, err := psc.qry.AllProduct()
	if err != nil {
		log.Println("query error")
		return []product.Core{}, errors.New("server error")
	}
	return res, nil
}

// Delete implements product.ProductService
func (psc *productServiceCase) Delete(token interface{}, productID uint) error {
	panic("unimplemented")
}

// EditProduct implements product.ProductService
func (psc *productServiceCase) EditProduct(token interface{}, formHeader multipart.FileHeader, productID uint, editedProduct product.Core) (product.Core, error) {
	panic("unimplemented")
}

// ProductDetail implements product.ProductService
func (psc *productServiceCase) ProductDetail(productID uint) (product.Core, error) {
	panic("unimplemented")
}
