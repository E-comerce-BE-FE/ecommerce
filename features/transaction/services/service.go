package services

import (
	"ecommerce/features/transaction"
	"ecommerce/helper"
	"errors"
	"log"
)

type transactionServiceCase struct {
	qry transaction.TransactionData
}

func New(td transaction.TransactionData) transaction.TransactionService {
	return &transactionServiceCase{
		qry: td,
	}
}

// CreateTransaction implements transaction.TransactionService
func (tsc *transactionServiceCase) CreateTransaction(token interface{}, paymentLink string, codeTrans string) (transaction.Core, error) {
	userID := helper.ExtractToken(token)
	res, err := tsc.qry.CreateTransaction(uint(userID), paymentLink, codeTrans)
	if err != nil {
		if err != nil {
			log.Println("query error", err.Error())
			return transaction.Core{}, errors.New("query error, problem with server")
		}
	}
	return res, nil
}

// UpdateTransaction implements transaction.TransactionService
func (tsc *transactionServiceCase) UpdateTransaction(codeTrans string) (transaction.Core, error) {
	res, err := tsc.qry.UpdateTransaction(codeTrans)
	if err != nil {
		log.Println("query error", err.Error())
		return transaction.Core{}, errors.New("query error, problem with server")
	}
	return res, nil
}

// TransactionHistory implements transaction.TransactionService
func (tsc *transactionServiceCase) TransactionHistory(token interface{}) ([]transaction.Core, error) {
	userID := helper.ExtractToken(token)
	res, err := tsc.qry.TransactionHistory(uint(userID))
	if err != nil {
		if err != nil {
			log.Println("query error", err.Error())
			return []transaction.Core{}, errors.New("query error, problem with server")
		}
	}
	return res, nil
}
