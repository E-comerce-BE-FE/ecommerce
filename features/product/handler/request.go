package handler

import (
	"ecommerce/features/product"
	"mime/multipart"
)

type AddProductRequest struct {
	ProductName string `json:"product_name" form:"product_name"`
	Price       int    `json:"price" form:"price"`
	Stock       int    `json:"stock" form:"stock"`
	FileHeader  multipart.FileHeader
}
type EditContentRequest struct {
	Content string `json:"content" form:"content"`
}

func RequestToCore(dataProduct interface{}) *product.Core {
	res := product.Core{}
	switch dataProduct.(type) {
	case AddProductRequest:
		cnv := dataProduct.(AddProductRequest)
		res.ProductName = cnv.ProductName
		res.Price = cnv.Price
		res.Stock = cnv.Stock
	// case EditProductRequest:
	// 	cnv := dataProduct.(EditProductRequest)
	// 	res.Product = cnv.Product
	default:
		return nil
	}
	return &res

}
