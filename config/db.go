package config

import (
	"fmt"
	"log"

	// cmData "ecommerce/features/cart/data"
	cartData "ecommerce/features/cart/data"
	pData "ecommerce/features/product/data"
	transData "ecommerce/features/transaction/data"
	usrData "ecommerce/features/user/data"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(dc DBConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dc.DBUser, dc.DBPass, dc.DBHost, dc.DBPort, dc.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("database connection error : ", err.Error())
		return nil
	}

	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(usrData.User{})
	db.AutoMigrate(pData.Product{})
	db.AutoMigrate(cartData.Cart{})
	db.AutoMigrate(transData.Transaction{})
	db.AutoMigrate(transData.TransactionItem{})
}
