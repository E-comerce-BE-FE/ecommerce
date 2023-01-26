package handler

import "ecommerce/features/cart"

type CartResponse struct {
	ID           uint   `json:"id"`
	ProductName  string `json:"product_name"`
	Seller       string `json:"seller"`
	Qty          int    `json:"quantity"`
	Amount       int    `json:"amount"`
	ProductImage string `json:"product_image"`
}

func CToResponse(data cart.Core) CartResponse {
	return CartResponse{
		ID:           data.ID,
		ProductName:  data.ProductName,
		Seller:       data.Seller,
		Qty:          data.Qty,
		Amount:       data.Amount,
		ProductImage: data.ProductImage,
	}
}
