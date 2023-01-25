package data

import (
	"ecommerce/features/cart"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	Qty       int
	Amount    int
	UserId    uint
	ProductId uint
}

type Product struct {
	gorm.Model
	ProductName  string
	ProductImage string
	Stock        int
	Price        int
	UserId       uint
}

type User struct {
	gorm.Model
	Name      string
	Address   string
	UserImage string
}

func DataToCore(data Cart) cart.Core {
	return cart.Core{
		ID:     data.ID,
		Qty:    data.Qty,
		Amount: data.Amount,
	}
}

func CoreToData(core cart.Core) Cart {
	return Cart{
		Model:  gorm.Model{ID: core.ID},
		Qty:    core.Qty,
		Amount: core.Qty,
	}
}
