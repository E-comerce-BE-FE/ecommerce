package handler

import (
	"ecommerce/features/transaction"
	"net/http"

	"github.com/labstack/echo/v4"
)

type transactionController struct {
	srv transaction.TransactionService
}

func New(ts transaction.TransactionService) transaction.TransactionHandler {
	return &transactionController{
		srv: ts,
	}
}

// CreateTransaction implements transaction.TransactionHandler
func (tc *transactionController) CreateTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := TransactionRequest{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "wrong input format"})
		}
		// log.Println(snapResp.RedirectURL)
		// return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})

		res, err := tc.srv.CreateTransaction(c.Get("user"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    res,
			"message": "transaction success created",
		})
	}
}

// UpdateTransaction implements transaction.TransactionHandler
func (*transactionController) UpdateTransaction() echo.HandlerFunc {
	panic("unimplemented")
}
