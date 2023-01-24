package product

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID           uint
	ProductName  string
	ProductImage string
	Price        int
	Stock        int
	Description  string
	User         UserCore
}

type UserCore struct {
	ID        uint
	Name      string
	UserImage string
	Address   string
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
