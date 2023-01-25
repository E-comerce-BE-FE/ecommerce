package handler

import (
	"ecommerce/features/user"
	"mime/multipart"
)

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type RegisterRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	Password string `json:"password" form:"password"`
}

type UpdateRequest struct {
	Name       string `json:"name" form:"name"`
	Email      string `json:"email" form:"email"`
	Phone      string `json:"phone" form:"phone"`
	Address    string `json:"address" form:"address"`
	Password   string `json:"password" form:"password"`
	UserImage  string `json:"user_image" form:"user_image"`
	FileHeader multipart.FileHeader
}

func ReqToCore(data interface{}) *user.Core {
	res := user.Core{}

	switch data.(type) {
	case LoginRequest:
		cnv := data.(LoginRequest)
		res.Email = cnv.Email
		res.Password = cnv.Password
	case RegisterRequest:
		cnv := data.(RegisterRequest)
		res.Name = cnv.Name
		res.Email = cnv.Email
		res.Phone = cnv.Phone
		res.Address = cnv.Address
		res.Password = cnv.Password
	case UpdateRequest:
		cnv := data.(UpdateRequest)
		res.Name = cnv.Name
		res.Email = cnv.Email
		res.Phone = cnv.Phone
		res.Address = cnv.Address
		res.Password = cnv.Password
		res.UserImage = cnv.UserImage
	default:
		return nil
	}

	return &res
}
