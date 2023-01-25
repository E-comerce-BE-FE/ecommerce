package data

import (
	"ecommerce/features/cart"
	"errors"
	"log"

	"gorm.io/gorm"
)

type cartQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) cart.CartData {
	return &cartQuery{
		db: db,
	}
}

// AddToCart implements cart.CartData
func (cq *cartQuery) AddToCart(userID uint, productID uint, add cart.Core) (cart.Core, error) {
	// cari data dari product berdasarkan IDnya
	prd := Product{}
	err := cq.db.Where("id=?", productID).First(&prd).Error
	if err != nil {
		log.Println("query error", err.Error())
		return cart.Core{}, errors.New("server error")
	}
	cnv := CoreToData(add)
	cnv.ProductId = prd.ID
	cnv.UserId = userID
	cnv.Qty = 1
	cnv.Amount = prd.Price
	err = cq.db.Create(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return cart.Core{}, errors.New("server error")
	}
	result := DataToCore(cnv)
	return result, nil
}

// CartList implements cart.CartData
func (cq *cartQuery) CartList(userID uint) ([]cart.Core, error) {
	res := []Cart{}
	err := cq.db.Where("user_id = ?", userID).Find(&res).Error
	if err != nil {
		log.Println("query error", err.Error())
		return []cart.Core{}, errors.New("server error")
	}

	result := []cart.Core{}
	for i := 0; i < len(res); i++ {
		result = append(result, DataToCore(res[i]))
		// cari data user berdasarkan cart user_id
		user := User{}
		err = cq.db.Where("id = ?", res[i].UserId).First(&user).Error
		if err != nil {
			log.Println("query error", err.Error())
			return []cart.Core{}, errors.New("server error")
		}
		// cari data product berdasarkan cart product_id
		prd := Product{}
		err = cq.db.Where("id = ?", res[i].ProductId).First(&prd).Error
		if err != nil {
			log.Println("query error", err.Error())
			return []cart.Core{}, errors.New("server error")
		}
		result[i].Seller = user.Name
		result[i].ProductName = prd.ProductName
		result[i].ProductImage = prd.ProductImage
	}
	return result, nil

}

// Delete implements cart.CartData
func (cq *cartQuery) Delete() error {
	panic("unimplemented")
}

// UpdateQty implements cart.CartData
func (cq *cartQuery) UpdateQty(userID uint, productID uint, quantity int) (cart.Core, error) {
	panic("unimplemented")
}
