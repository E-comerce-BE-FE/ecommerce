package handler

import (
	"ecommerce/features/cart"
)

type AddToCartRequest struct {
	ProductID uint `json:"product_id" form:"product_id"`
	Qty       int  `json:"qty" form:"qty"`
	Amount    int  `json:"amount" form:"amount"`
}
type EditCartRequest struct {
	Qty    int `json:"qty" form:"qty"`
	Amount int `json:"amount" form:"amount"`
}

func RequestToCore(dataCart interface{}) *cart.Core {
	res := cart.Core{}
	switch dataCart.(type) {
	case AddToCartRequest:
		cnv := dataCart.(AddToCartRequest)
		res.Qty = cnv.Qty
		res.Amount = cnv.Amount
	case EditCartRequest:
		cnv := dataCart.(EditCartRequest)
		res.Qty = cnv.Qty
		res.Amount = cnv.Amount
	default:
		return nil
	}
	return &res

}
