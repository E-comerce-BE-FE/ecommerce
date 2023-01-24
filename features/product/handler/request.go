package handler

import (
	"ecommerce/features/product"
	"mime/multipart"
)

type AddProductRequest struct {
	ProductName string `json:"product_name" form:"product_name"`
	Price       int    `json:"price" form:"price"`
	Stock       int    `json:"stock" form:"stock"`
	Description string `json:"description" form:"description"`
	FileHeader  multipart.FileHeader
}
type EditProductRequest struct {
	ProductName string `json:"product_name" form:"product_name"`
	Price       int    `json:"price" form:"price"`
	Stock       int    `json:"stock" form:"stock"`
	Description string `json:"description" form:"description"`
	FileHeader  multipart.FileHeader
}

func RequestToCore(dataProduct interface{}) *product.Core {
	res := product.Core{}
	switch dataProduct.(type) {
	case AddProductRequest:
		cnv := dataProduct.(AddProductRequest)
		res.ProductName = cnv.ProductName
		res.Price = cnv.Price
		res.Stock = cnv.Stock
		res.Description = cnv.Description
	case EditProductRequest:
		cnv := dataProduct.(EditProductRequest)
		res.ProductName = cnv.ProductName
		res.Price = cnv.Price
		res.Stock = cnv.Stock
		res.Description = cnv.Description
	default:
		return nil
	}
	return &res

}
