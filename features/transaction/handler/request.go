package handler

type TransactionRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	SubTotal int64  `json:"sub_total" form:"sub_total"`
}

type UpdateTransRequest struct {
	CodeTrans string `json:"code_transaction" form:"code_transaction"`
}
