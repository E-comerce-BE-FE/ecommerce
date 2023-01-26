package cart

import (
	"github.com/labstack/echo/v4"
)

type Core struct {
	ID           uint
	ProductName  string
	ProductImage string
	Seller       string
	Qty          int
	Amount       int
	User         User
}

type User struct {
	ID      uint
	Name    string
	Email   string
	Phone   string
	Address string
}

type CartHandler interface {
	AddToCart() echo.HandlerFunc
	CartList() echo.HandlerFunc
	UpdateQty() echo.HandlerFunc
	Delete() echo.HandlerFunc
	CartResult() echo.HandlerFunc
}

type CartService interface {
	AddToCart(token interface{}, productID uint, addToCart Core) (Core, error)
	CartList(token interface{}) ([]Core, error)
	UpdateQty(token interface{}, cartID uint, quantity int) (Core, error)
	Delete(token interface{}, cartID uint) error
	CartResult(token interface{}) ([]Core, error)
}

type CartData interface {
	AddToCart(userID uint, productID uint, add Core) (Core, error)
	CartList(userID uint) ([]Core, error)
	UpdateQty(userID uint, cartID uint, quantity int) (Core, error)
	Delete(userID uint, cartID uint) error
	CartResult(userID uint) ([]Core, error)
}
