package data

import (
	"ecommerce/features/product"

	"gorm.io/gorm"
)

type productQry struct {
	db *gorm.DB
}

func New(db *gorm.DB) product.ProductData {
	return &productQry{
		db: db,
	}
}

// AddProduct implements product.ProductData
func (*productQry) AddProduct(userID uint, newProduct product.Core) (product.Core, error) {
	panic("unimplemented")
}

// AllProduct implements product.ProductData
func (*productQry) AllProduct() ([]product.Core, error) {
	panic("unimplemented")
}

// Delete implements product.ProductData
func (*productQry) Delete(productID uint) error {
	panic("unimplemented")
}

// EditProduct implements product.ProductData
func (*productQry) EditProduct(userID uint, productID uint, editedProduct product.Core) (product.Core, error) {
	panic("unimplemented")
}

// ProductDetail implements product.ProductData
func (*productQry) ProductDetail(productID uint) (product.Core, error) {
	panic("unimplemented")
}
