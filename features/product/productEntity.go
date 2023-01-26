package product

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID           uint     `json:"id"`
	ProductName  string   `json:"product_name"`
	ProductImage string   `json:"product_image"`
	Price        int      `json:"price"`
	Stock        int      `json:"stock"`
	Description  string   `json:"description"`
	User         UserCore `json:"user"`
}

type UserCore struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	UserImage string `json:"user_image"`
	Address   string `json:"address"`
}

type ProductHandler interface {
	AddProduct() echo.HandlerFunc
	EditProduct() echo.HandlerFunc
	Delete() echo.HandlerFunc
	AllProduct() echo.HandlerFunc
	ProductDetail() echo.HandlerFunc
}

type ProductService interface {
	AddProduct(token interface{}, formHeader multipart.FileHeader, newProduct Core) (Core, error)
	EditProduct(token interface{}, formHeader multipart.FileHeader, productID uint, editedProduct Core) (Core, error)
	Delete(token interface{}, productID uint) error
	AllProduct() ([]Core, error)
	ProductDetail(productID uint) (Core, error)
}

type ProductData interface {
	AddProduct(userID uint, newProduct Core) (Core, error)
	EditProduct(userID uint, productID uint, editedProduct Core) (Core, error)
	Delete(userID uint, productID uint) error
	AllProduct() ([]Core, error)
	ProductDetail(productID uint) (Core, error)
}
