package data

import (
	"ecommerce/features/transaction"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type transactionQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) transaction.TransactionData {
	return &transactionQuery{
		db: db,
	}
}

// CreateTransaction implements transaction.TransactionData
func (tq *transactionQuery) CreateTransaction(userID uint) (transaction.Core, error) {
	// ambil semua data dari cart user
	cart := []Cart{}
	err := tq.db.Where("user_id=?", userID).Find(&cart).Error
	if err != nil {
		log.Println("query error", err.Error())
		return transaction.Core{}, errors.New("server error")
	}
	// menjumlahkan quantity dan subtotal
	totalProduct := 0
	totalAmount := 0
	for i := 0; i < len(cart); i++ {
		totalProduct += cart[i].Qty
		totalAmount += cart[i].Amount
	}
	// persiapan create transaction
	trs := Transaction{}
	trs.TotalProduct = totalProduct
	trs.SubTotal = totalAmount
	trs.TransactionCode = "uhIGuiG-IUh987-SOiuh-c21312k3h"
	trs.Status = "pending"
	err = tq.db.Create(&trs).Error
	if err != nil {
		log.Println("query error", err.Error())
		return transaction.Core{}, errors.New("server error")
	}
	//copy cart ke transaksi item
	temp := []TransactionItem{}
	for i := 0; i < len(cart); i++ {
		temp = append(temp, CartToTI(cart[i]))
		temp[i].TransactionID = trs.ID
	}
	// batch create
	err = tq.db.Create(&temp).Error
	if err != nil {
		log.Println("query error", err.Error())
		return transaction.Core{}, errors.New("server error")
	}
	// batch delete cart
	err = tq.db.Where("user_id = ?", userID).Delete(&Cart{}).Error
	if err != nil {
		log.Println("query error", err.Error())
		return transaction.Core{}, errors.New("server error")
	}
	// output tampilan JSON
	result := transaction.Core{}
	result.ID = trs.ID
	result.Status = "pending"
	result.SubTotal = totalAmount
	result.TotalProduct = totalProduct
	result.TransactionCode = "uhIGuiG-IUh987-SOiuh-c21312k3h"
	result.TransactionDate = fmt.Sprintf("%d-%d-%d %d:%d:%d", trs.CreatedAt.Day(), trs.CreatedAt.Month(), trs.CreatedAt.Year(), trs.CreatedAt.Hour(), trs.CreatedAt.Minute(), trs.CreatedAt.Second())
	result.TransactionName = "product-shoping"
	return result, nil
}

// UpdateTransaction implements transaction.TransactionData
func (*transactionQuery) UpdateTransaction() (transaction.Core, error) {
	panic("unimplemented")
}
