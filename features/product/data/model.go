package data

import (
	"ecommerce/features/product"
	// "ecommerce/features/cart/data"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductName  string
	ProductImage string
	Stock        int
	Price        int
	Description  string
	UserId       uint
	User         User
}

type User struct {
	ID        uint
	Name      string
	UserImage string
	Address   string
}

func DataToCore(data Product) product.Core {
	return product.Core{
		ID:           data.ID,
		ProductName:  data.ProductName,
		ProductImage: data.ProductImage,
		Price:        data.Price,
		Stock:        data.Stock,
		Description:  data.Description,
		User:         product.UserCore{},
	}
}

func CoreToData(core product.Core) Product {
	return Product{
		Model:        gorm.Model{ID: core.ID},
		ProductName:  core.ProductName,
		ProductImage: core.ProductImage,
		Price:        core.Price,
		Stock:        core.Stock,
		Description:  core.Description,
		User:         User{},
	}
}
