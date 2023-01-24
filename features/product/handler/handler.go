package handler

import (
	"ecommerce/features/product"

	"github.com/labstack/echo/v4"
)

type handlerController struct {
	srv product.ProductService
}

func New(ps product.ProductService) product.ProductHandler {
	return &handlerController{
		srv: ps,
	}
}

// AddProduct implements product.ProductHandler
func (*handlerController) AddProduct() echo.HandlerFunc {
	panic("unimplemented")
}

// AllProduct implements product.ProductHandler
func (*handlerController) AllProduct() echo.HandlerFunc {
	panic("unimplemented")
}

// Delete implements product.ProductHandler
func (*handlerController) Delete() echo.HandlerFunc {
	panic("unimplemented")
}

// EditProduct implements product.ProductHandler
func (*handlerController) EditProduct() echo.HandlerFunc {
	panic("unimplemented")
}

// ProductDetail implements product.ProductHandler
func (*handlerController) ProductDetail() echo.HandlerFunc {
	panic("unimplemented")
}
