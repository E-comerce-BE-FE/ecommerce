package services

import "ecommerce/features/product"

type productServiceCase struct {
	qry product.ProductData
}

func New(pd product.ProductData) product.ProductService {
	return &productServiceCase{
		qry: pd,
	}
}

// AddProduct implements product.ProductService
func (*productServiceCase) AddProduct(token interface{}, newProduct product.Core) (product.Core, error) {
	panic("unimplemented")
}

// AllProduct implements product.ProductService
func (*productServiceCase) AllProduct() ([]product.Core, error) {
	panic("unimplemented")
}

// Delete implements product.ProductService
func (*productServiceCase) Delete(token interface{}, productID uint) error {
	panic("unimplemented")
}

// EditProduct implements product.ProductService
func (*productServiceCase) EditProduct(token interface{}, productID uint, editedProduct product.Core) (product.Core, error) {
	panic("unimplemented")
}

// ProductDetail implements product.ProductService
func (*productServiceCase) ProductDetail(productID uint) (product.Core, error) {
	panic("unimplemented")
}
