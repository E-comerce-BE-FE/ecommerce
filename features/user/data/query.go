package data

import (
	"ecommerce/features/user"
	"errors"
	"log"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserData {
	return &userQuery{
		db: db,
	}
}

func (uq *userQuery) Login(email string) (user.Core, error) {
	if email == "" {
		log.Println("data empty, query error")
		return user.Core{}, errors.New("email not allowed empty")
	}
	res := User{}
	if err := uq.db.Where("email = ?", email).First(&res).Error; err != nil {
		log.Println("login query error", err.Error())
		return user.Core{}, errors.New("data not found")
	}

	return ToCore(res), nil
}

func (uq *userQuery) Register(newUser user.Core) (user.Core, error) {
	if newUser.Email == "" || newUser.Password == "" {
		log.Println("data empty")
		return user.Core{}, errors.New("email or password not allowed empty")
	}
	dupEmail := CoreToData(newUser)
	err := uq.db.Where("email = ?", newUser.Email).First(&dupEmail).Error
	if err == nil {
		log.Println("duplicated")
		return user.Core{}, errors.New("email duplicated")
	}

	cnv := CoreToData(newUser)
	err = uq.db.Create(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return user.Core{}, errors.New("server error")
	}

	newUser.ID = cnv.ID
	return newUser, nil
}

func (uq *userQuery) Profile(userID uint) (interface{}, error) {
	res := User{}
	if err := uq.db.Where("id=?", userID).Preload("Product").First(&res).Error; err != nil {
		log.Println("Get By ID query error", err.Error())
		return user.Core{}, errors.New("query error, problem with server")
	}

	result := make(map[string]interface{})
	result["id"] = res.ID
	result["name"] = res.Name
	// result["email"] = res.Email
	// result["phone"] = res.Phone
	// result["address"] = res.Address
	result["user_image"] = res.UserImage
	result["product"] = make([]map[string]interface{}, len(res.Product))

	for k, v := range res.Product {
		data := make(map[string]interface{})
		data["product_id"] = v.ID
		data["product_name"] = v.ProductName
		data["stock"] = v.Stock
		data["price"] = v.Price
		data["description"] = v.Description
		data["product_image"] = v.ProductImage
		result["product"].([]map[string]interface{})[k] = data
	}

	return result, nil
}

func (uq *userQuery) Update(id uint, updateData user.Core) (user.Core, error) {
	cnv := CoreToData(updateData)
	qry := uq.db.Model(&User{}).Where("id = ?", id).Updates(&cnv)

	affrows := qry.RowsAffected
	if affrows == 0 {
		log.Println("no rows affected")
		return user.Core{}, errors.New("no data updated")
	}

	err := qry.Error
	if err != nil {
		log.Println("update user query error", err.Error())
		return user.Core{}, err
	}

	return ToCore(cnv), nil
}

func (uq *userQuery) Delete(id uint) error {
	qry := uq.db.Delete(&User{}, id)
	err := qry.Error

	affrows := qry.RowsAffected
	if affrows == 0 {
		log.Println("no rows affected")
		return errors.New("no data deleted")
	}

	if err != nil {
		log.Println("delete query error")
		return errors.New("cannot delete data, data not found")
	}

	return nil
}
