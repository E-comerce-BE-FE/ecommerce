package transaction

import "github.com/labstack/echo/v4"

type Core struct {
	ID              uint   `json:"transaction_id"`
	TotalProduct    int    `json:"total_product"`
	SubTotal        int    `json:"sub_total"`
	TransactionDate string `json:"create_at"`
	Status          string `json:"status"`
	TransactionName string `json:"transaction_name"`
	TransactionCode string `json:"transaction_code"`
	PaymentLink     string `json:"payment_link"`
}

type TransactionHandler interface {
	CreateTransaction() echo.HandlerFunc
	UpdateTransaction() echo.HandlerFunc
	TransactionHistory() echo.HandlerFunc
	CancelTransaction() echo.HandlerFunc
	TransactionDetail() echo.HandlerFunc
}

type TransactionService interface {
	CreateTransaction(token interface{}, paymentLink string, codeTrans string) (Core, error)
	UpdateTransaction(codeTrans string) (Core, error)
	TransactionHistory(token interface{}) ([]Core, error)
	CancelTransaction(token interface{}, transactionID uint) error
	TransactionDetail(token interface{}, transactionID uint) (interface{}, error)
}

type TransactionData interface {
	CreateTransaction(userID uint, paymentLink string, codeTrans string) (Core, error)
	UpdateTransaction(codeTrans string) (Core, error)
	TransactionHistory(userID uint) ([]Core, error)
	CancelTransaction(userID uint, transactionID uint) error
	TransactionDetail(userID uint, transactionID uint) (interface{}, error)
}
