package transaction

import "github.com/labstack/echo/v4"

type Core struct {
	ID              uint
	TotalProduct    int
	SubTotal        int
	TransactionDate string
	Status          string
	TransactionName string
	TransactionCode string
}

type TransactionHandler interface {
	CreateTransaction() echo.HandlerFunc
	UpdateTransaction() echo.HandlerFunc
	// CancelTransaction() echo.HandlerFunc
}

type TransactionService interface {
	CreateTransaction(token interface{}) (Core, error)
	UpdateTransaction() (Core, error)
	// CancelTransaction() (Core,error)
}

type TransactionData interface {
	CreateTransaction(userID uint) (Core, error)
	UpdateTransaction() (Core, error)
	// CancelTransaction() (Core,error)
}
