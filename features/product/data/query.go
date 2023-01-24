package data

import (
	"ecommerce/features/product"
	"errors"
	"log"

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
func (pq *productQry) AddProduct(userID uint, newProduct product.Core) (product.Core, error) {
	data := CoreToData(newProduct)
	data.UserId = userID
	err := pq.db.Create(&data).Error
	if err != nil {
		log.Println("query error", err.Error())
		return product.Core{}, errors.New("server error")
	}
	newProduct.ID = data.ID
	return newProduct, nil
}

// AllProduct implements product.ProductData
func (pq *productQry) AllProduct() ([]product.Core, error) {
	panic("unimplemented")
}

// Delete implements product.ProductData
func (pq *productQry) Delete(productID uint) error {
	panic("unimplemented")
}

// EditProduct implements product.ProductData
func (pq *productQry) EditProduct(userID uint, productID uint, editedProduct product.Core) (product.Core, error) {
	panic("unimplemented")
}

// ProductDetail implements product.ProductData
func (pq *productQry) ProductDetail(productID uint) (product.Core, error) {
	panic("unimplemented")
}
