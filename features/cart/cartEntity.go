package cart

import (
	"github.com/labstack/echo/v4"
)

type Core struct {
	ID           uint   `json:"id"`
	ProductName  string `json:"product_name"`
	ProductImage string `json:"product_image"`
	Seller       string `json:"seller"`
	Qty          int    `json:"quantity"`
	Amount       int    `json:"amount"`
	User         User   `json:"user"`
}

type User struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
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
