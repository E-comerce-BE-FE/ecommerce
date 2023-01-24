package data

import (
	"ecommerce/features/product"
	"errors"
	"log"

	"gorm.io/gorm"
)

type productQry struct {
	db *gorm.DB
}

func New(db *gorm.DB) product.ProductData {
	return &productQry{
		db: db,
	}
}

// AddProduct implements product.ProductData
func (pq *productQry) AddProduct(userID uint, newProduct product.Core) (product.Core, error) {
	data := CoreToData(newProduct)
	data.UserId = userID
	err := pq.db.Create(&data).Error
	if err != nil {
		log.Println("query error", err.Error())
		return product.Core{}, errors.New("server error")
	}
	newProduct.ID = data.ID
	return newProduct, nil
}

// AllProduct implements product.ProductData
func (pq *productQry) AllProduct() ([]product.Core, error) {
	data := []Product{}
	err := pq.db.Preload("User").Find(&data).Error
	if err != nil {
		log.Println("query error", err.Error())
		return []product.Core{}, errors.New("server error")
	}
	result := []product.Core{}
	for i := 0; i < len(data); i++ {
		result = append(result, DataToCore(data[i]))
		result[i].User.ID = data[i].User.ID
		result[i].User.Name = data[i].User.Name
		result[i].User.Address = data[i].User.Address
	}
	return result, nil
}

// Delete implements product.ProductData
func (pq *productQry) Delete(userID uint, productID uint) error {
	//cek apakah content yang akan dihapus milik user yang akan menghapus
	vld := Product{}
	err := pq.db.Where("user_id=? AND id=?", userID, productID).First(&vld).Error
	if err != nil {
		log.Println("product not found", err.Error())
		return errors.New("product cannot deleted")
	}
	// ok hapus
	qry := pq.db.Delete(&Product{}, productID)
	rowAffect := qry.RowsAffected
	if rowAffect <= 0 {
		log.Println("no data processed")
		return errors.New("no product has delete")
	}
	err = qry.Error
	if err != nil {
		log.Println("delete query error", err.Error())
		return errors.New("delete product fail")
	}
	return nil
}

// EditProduct implements product.ProductData
func (pq *productQry) EditProduct(userID uint, productID uint, editedProduct product.Core) (product.Core, error) {
	vld := Product{}
	err := pq.db.Where("user_id=? AND id=?", userID, productID).First(&vld).Error
	if err != nil {
		log.Println("product not found", err.Error())
		return product.Core{}, errors.New("product cannot edited")
	}
	res := Product{}
	qry := pq.db.Where("user_id=? AND id=?", userID, productID).Updates(&res)
	if qry.RowsAffected <= 0 {
		log.Println("update error : no rows affected")
		return product.Core{}, errors.New("update error : no rows updated")
	}
	err = qry.Error
	if err != nil {
		log.Println("update error")
		return product.Core{}, errors.New("query error,update fail")
	}
	return editedProduct, nil
}

// ProductDetail implements product.ProductData
func (pq *productQry) ProductDetail(productID uint) (product.Core, error) {
	data := Product{}
	err := pq.db.Where("id = ?", productID).Preload("User").First(&data).Error
	if err != nil {
		log.Println("query error", err.Error())
		return product.Core{}, errors.New("server error")
	}
	result := product.Core{}
	result = DataToCore(data)
	result.User.ID = data.User.ID
	result.User.Name = data.User.Name
	result.User.Address = data.User.Address
	result.User.UserImage = data.User.UserImage

	return result, nil
}
