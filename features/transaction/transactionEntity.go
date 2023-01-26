package transaction

import "github.com/labstack/echo/v4"

type Core struct {
	ID                uint
	TotalProduct      int
	SubTotal          int
	TransactionDate   string
	Status            string
	TransactionName   string
	TransactionCode   string
	PaymentLink       string
	TransactionDetail TransactionDetail
}
type TransactionDetail struct {
}

type TransactionHandler interface {
	CreateTransaction() echo.HandlerFunc
	UpdateTransaction() echo.HandlerFunc
	TransactionHistory() echo.HandlerFunc
	// CancelTransaction() echo.HandlerFunc
	// TransactionDetail() echo.HandlerFunc
}

type TransactionService interface {
	CreateTransaction(token interface{}, paymentLink string, codeTrans string) (Core, error)
	UpdateTransaction(codeTrans string) (Core, error)
	TransactionHistory(token interface{}) ([]Core, error)
	// CancelTransaction() (Core, error)
	// TransactionDetail()
}

type TransactionData interface {
	CreateTransaction(userID uint, paymentLink string, codeTrans string) (Core, error)
	UpdateTransaction(codeTrans string) (Core, error)
	TransactionHistory(userID uint) ([]Core, error)
	// CancelTransaction() (Core, error)
	// TransactionDetail()
}
