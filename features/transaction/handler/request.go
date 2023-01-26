package handler

type TransactionRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	SubTotal string `json:"subtotal" form:"subtotal"`
}
