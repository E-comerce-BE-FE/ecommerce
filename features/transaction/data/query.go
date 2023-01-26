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
func (tq *transactionQuery) CreateTransaction(userID uint, paymentLink string, codeTrans string) (transaction.Core, error) {
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
	trs.PaymentLink = paymentLink
	trs.TransactionCode = codeTrans
	trs.UserId = userID
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
	result.PaymentLink = paymentLink
	result.TransactionCode = codeTrans
	result.TransactionDate = fmt.Sprintf("%d-%d-%d %d:%d:%d", trs.CreatedAt.Day(), trs.CreatedAt.Month(), trs.CreatedAt.Year(), trs.CreatedAt.Hour(), trs.CreatedAt.Minute(), trs.CreatedAt.Second())
	result.TransactionName = "product-shoping"
	return result, nil
}

// UpdateTransaction implements transaction.TransactionData
func (tq *transactionQuery) UpdateTransaction(codeTrans string) (transaction.Core, error) {
	trans := Transaction{}
	trans.Status = "success"
	qry := tq.db.Where("transaction_code = ?", codeTrans).Updates(&trans)
	if qry.RowsAffected <= 0 {
		log.Println("update error : no rows affected")
		return transaction.Core{}, errors.New("update error : no rows updated")
	}
	err := qry.Error
	if err != nil {
		log.Println("query error", err.Error())
		return transaction.Core{}, errors.New("server error")
	}
	data := Transaction{}
	err = tq.db.Where("transaction_code = ?", codeTrans).First(&data).Error
	if err != nil {
		log.Println("query error", err.Error())
		return transaction.Core{}, errors.New("server error")
	}
	result := DataToCore(data)
	result.TransactionDate = fmt.Sprintf("%d-%d-%d %d:%d:%d", data.CreatedAt.Day(), data.CreatedAt.Month(), data.CreatedAt.Year(), data.CreatedAt.Hour(), data.CreatedAt.Minute(), data.CreatedAt.Second())
	result.TransactionName = "product-shoping"
	return result, nil

}

// TransactionHistory implements transaction.TransactionData
func (tq *transactionQuery) TransactionHistory(userID uint) ([]transaction.Core, error) {
	res := []Transaction{}
	err := tq.db.Where("user_id=?", userID).Find(&res).Error
	if err != nil {
		log.Println("query error", err.Error())
		return []transaction.Core{}, errors.New("server error")
	}
	result := []transaction.Core{}
	for i := 0; i < len(res); i++ {
		result = append(result, DataToCore(res[i]))
		result[i].TransactionDate = fmt.Sprintf("%d-%d-%d %d:%d:%d", res[i].CreatedAt.Day(), res[i].CreatedAt.Month(), res[i].CreatedAt.Year(), res[i].CreatedAt.Hour(), res[i].CreatedAt.Minute(), res[i].CreatedAt.Second())
		result[i].TransactionName = "product-shoping"
	}
	return result, nil
}
