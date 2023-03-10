package services

import (
	"ecommerce/features/cart"
	"ecommerce/helper"
	"errors"
	"log"
)

type cartServiceCase struct {
	qry cart.CartData
}

func New(cd cart.CartData) cart.CartService {
	return &cartServiceCase{
		qry: cd,
	}
}

// AddToCart implements cart.CartService
func (csc *cartServiceCase) AddToCart(token interface{}, productID uint, addToCart cart.Core) (cart.Core, error) {
	userID := helper.ExtractToken(token)
	res, err := csc.qry.AddToCart(uint(userID), productID, addToCart)
	if err != nil {
		log.Println("query error", err.Error())
		return cart.Core{}, errors.New("query error, problem with server")
	}
	return res, nil
}

// CartList implements cart.CartService
func (csc *cartServiceCase) CartList(token interface{}) ([]cart.Core, error) {
	userID := helper.ExtractToken(token)
	res, err := csc.qry.CartList(uint(userID))
	if err != nil {
		log.Println("query error", err.Error())
		return []cart.Core{}, errors.New("query error, problem with server")
	}
	return res, nil
}

// Delete implements cart.CartService
func (csc *cartServiceCase) Delete(token interface{}, cartID uint) error {
	userID := helper.ExtractToken(token)
	err := csc.qry.Delete(uint(userID), cartID)

	if err != nil {
		log.Println("query error", err.Error())
		return errors.New("query error, problem with server")
	}
	return nil
}

// UpdateQty implements cart.CartService
func (csc *cartServiceCase) UpdateQty(token interface{}, cartID uint, quantity int) (cart.Core, error) {
	userID := helper.ExtractToken(token)
	res, err := csc.qry.UpdateQty(uint(userID), cartID, quantity)
	if err != nil {
		log.Println("query error", err.Error())
		return cart.Core{}, errors.New("query error, problem with server")
	}
	return res, nil
}

// CartResult implements cart.CartService
func (csc *cartServiceCase) CartResult(token interface{}) ([]cart.Core, error) {
	userID := helper.ExtractToken(token)
	res, err := csc.qry.CartResult(uint(userID))
	if err != nil {
		log.Println("query error", err.Error())
		return []cart.Core{}, errors.New("query error, problem with server")
	}
	return res, nil
}
