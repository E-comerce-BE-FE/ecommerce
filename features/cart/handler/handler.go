package handler

import (
	"ecommerce/features/cart"
	"net/http"

	"github.com/labstack/echo/v4"
)

type cartController struct {
	srv cart.CartService
}

func New(cs cart.CartService) cart.CartHandler {
	return &cartController{
		srv: cs,
	}
}

// AddToCart implements cart.CartHandler
func (cc *cartController) AddToCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := AddToCartRequest{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "wrong input format"})
		}
		res, err := cc.srv.AddToCart(c.Get("user"), input.ProductID, *RequestToCore(input))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    res,
			"message": "success add to cart",
		})
	}
}

// CartList implements cart.CartHandler
func (cc *cartController) CartList() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := cc.srv.CartList(c.Get("user"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"message": "success",
		})
	}
}

// Delete implements cart.CartHandler
func (cc *cartController) Delete() echo.HandlerFunc {
	panic("unimplemented")
}

// UpdateQty implements cart.CartHandler
func (cc *cartController) UpdateQty() echo.HandlerFunc {
	panic("unimplemented")
}
