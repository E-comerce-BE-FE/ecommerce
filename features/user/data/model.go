package data

import (
	"ecommerce/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string
	Email          string
	Phone          string
	Address        string
	Password       string
	Profilepicture string
	Product        Product
}

type Product struct {
	gorm.Model
	ProductName  string
	ProductImage string
	Stock        int
	Price        int
	Description  string
}

func ToCore(data User) user.Core {
	return user.Core{
		ID:             data.ID,
		Name:           data.Name,
		Email:          data.Email,
		Phone:          data.Phone,
		Address:        data.Address,
		Password:       data.Password,
		Profilepicture: data.Profilepicture,
	}
}

func CoreToData(data user.Core) User {
	return User{
		Model:          gorm.Model{ID: data.ID},
		Name:           data.Name,
		Email:          data.Email,
		Phone:          data.Phone,
		Address:        data.Address,
		Password:       data.Password,
		Profilepicture: data.Profilepicture,
	}
}
