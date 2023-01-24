package user

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID             uint
	Name           string
	Email          string
	Phone          string
	Address        string
	Password       string
	Profilepicture string
}

type UserHandler interface {
	Login() echo.HandlerFunc
	Register() echo.HandlerFunc
	Profile() echo.HandlerFunc
	Delete() echo.HandlerFunc
	Update() echo.HandlerFunc
}

type UserService interface {
	Login(email, password string) (string, Core, error)
	Register(newUser Core) (Core, error)
	Profile(token interface{}) (interface{}, error)
	Update(token interface{}, fileData multipart.FileHeader, updateData Core) (Core, error)
	Delete(token interface{}) error
}

type UserData interface {
	Login(email string) (Core, error)
	Register(newUser Core) (Core, error)
	Profile() (interface{}, error)
	Update(id uint, updateData Core) (Core, error)
	Delete(id uint) error
}
