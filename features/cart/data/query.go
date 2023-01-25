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
func (cq *cartQuery) Delete(userID uint, cartID uint) error {
	data := Cart{}
	qry := cq.db.Where("id = ? and user_id = ?", cartID, userID).Delete(&data)

	affrows := qry.RowsAffected
	if affrows <= 0 {
		log.Println("no rows affected")
		return errors.New("no cart deleted")
	}

	err := qry.Error
	if err != nil {
		log.Println("delete query error", err.Error())
		return errors.New("delete data fail")
	}

	return nil
}

// UpdateQty implements cart.CartData
func (cq *cartQuery) UpdateQty(userID uint, cartID uint, quantity int) (cart.Core, error) {
	crt := Cart{}
	err := cq.db.Where("id = ?", cartID).First(&crt).Error
	if err != nil {
		log.Println("select query error", err.Error())
		return cart.Core{}, errors.New("select data fail")
	}

	prd := Product{}
	err = cq.db.Where("id = ?", crt.ProductId).First(&prd).Error
	if err != nil {
		log.Println("select query error", err.Error())
		return cart.Core{}, errors.New("select data fail")
	}

	res := Cart{}
	res.Qty = quantity
	res.UserId = userID
	res.Amount = prd.Price * quantity
	qry := cq.db.Where("id = ?", cartID).Updates(&res)

	affrows := qry.RowsAffected
	if affrows <= 0 {
		log.Println("no rows affected")
		return cart.Core{}, errors.New("no cart updated")
	}

	err = qry.Error
	if err != nil {
		log.Println("update query error", err.Error())
		return cart.Core{}, errors.New("update data fail")
	}

	return DataToCore(res), nil
}
