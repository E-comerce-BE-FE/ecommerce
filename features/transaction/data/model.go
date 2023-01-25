package data

import (
	"ecommerce/features/transaction"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	TotalProduct    int
	SubTotal        int
	TransactionCode string
	Status          string
}

type TransactionItem struct {
	gorm.Model
	ProductID     uint
	TransactionID uint
	Qty           int
	Amount        int
}

type Cart struct {
	gorm.Model
	Qty       int
	Amount    int
	UserId    uint
	ProductId uint
}

type Core struct {
	ID              uint
	TotalProduct    int
	SubTotal        int
	CreateAt        string
	Status          string
	TransactionName string
	TransactionCode string
}

func DataToCore(data Transaction) transaction.Core {
	return transaction.Core{
		ID:              data.ID,
		TotalProduct:    data.TotalProduct,
		SubTotal:        data.SubTotal,
		TransactionCode: data.TransactionCode,
		Status:          data.Status,
	}
}

func CoreToData(core transaction.Core) Transaction {
	return Transaction{
		Model:           gorm.Model{ID: core.ID},
		TotalProduct:    core.TotalProduct,
		SubTotal:        core.SubTotal,
		TransactionCode: core.TransactionCode,
		Status:          core.Status,
	}
}

// conversi cart ke transactionItem
func CartToTI(cart Cart) TransactionItem {
	return TransactionItem{
		ProductID: cart.ProductId,
		Qty:       cart.Qty,
		Amount:    cart.Amount,
	}
}
