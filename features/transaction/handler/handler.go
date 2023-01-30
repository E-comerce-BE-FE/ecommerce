package handler

import (
	"ecommerce/features/transaction"
	"ecommerce/helper"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
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
		var s = snap.Client{}
		s.New("YOUR-SERVER-KEY", midtrans.Sandbox)
		// Use to midtrans.Production if you want Production Environment (accept real transaction).
		// 2. Initiate Snap request param
		orderID := helper.GenerateRandomString()
		orderID = "GROUP-3-ORDER-ID-" + orderID
		req := &snap.Request{
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  orderID,
				GrossAmt: int64(input.SubTotal),
			},
			CreditCard: &snap.CreditCardDetails{
				Secure: true,
			},
			CustomerDetail: &midtrans.CustomerDetails{
				FName: input.Name,
				// LName: "",
				Email:    input.Email,
				Phone:    input.Phone,
				ShipAddr: &midtrans.CustomerAddress{Address: input.Address},
			},
		}
		snapResp, _ := s.CreateTransaction(req)
		if snapResp.RedirectURL == "" {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Payment Error, error when creating transaction payment to midtrans"})
		}
		// log.Println("URL:", snapResp.RedirectURL)
		// log.Println("Error:", snapResp.ErrorMessages)
		// log.Println("Status:", snapResp.StatusCode)
		// log.Println("Token:", snapResp.Token)
		// return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
		res, err := tc.srv.CreateTransaction(c.Get("user"), snapResp.RedirectURL, orderID)
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
func (tc *transactionController) UpdateTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := UpdateTransRequest{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "wrong input format"})
		}
		//get status
		var z = coreapi.Client{}
		z.New("YOUR-SERVER-KEY", midtrans.Sandbox)
		cekStatus, _ := z.CheckTransaction(input.CodeTrans)
		log.Println(cekStatus.TransactionStatus)
		log.Println(cekStatus.StatusMessage)
		if cekStatus.TransactionStatus != "settlement" {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "payment not complete, please complete the payment first"})
		}

		result, err := tc.srv.UpdateTransaction(input.CodeTrans)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    result,
			"message": "payment success updated",
		})
	}
}

// TransactionHistory implements transaction.TransactionHandler
func (tc *transactionController) TransactionHistory() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := tc.srv.TransactionHistory(c.Get("user"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"message": "show all transaction success",
		})
	}
}

// CancelTransaction implements transaction.TransactionHandler
func (tc *transactionController) CancelTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		trsID := c.Param("id")
		transactionID, _ := strconv.Atoi(trsID)
		err := tc.srv.CancelTransaction(c.Get("user"), uint(transactionID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{"message": "transaction success canceled"})
	}
}

// TransactionDetail implements transaction.TransactionHandler
func (tc *transactionController) TransactionDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		trsID := c.Param("id")
		transactionID, _ := strconv.Atoi(trsID)
		res, err := tc.srv.TransactionDetail(c.Get("user"), uint(transactionID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"message": "show transaction item success",
		})
	}
}
