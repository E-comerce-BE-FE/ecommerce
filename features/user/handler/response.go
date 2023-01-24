package handler

import (
	"ecommerce/features/user"
	"net/http"
	"strings"
)

type UserReponse struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Profilepicture string `json:"profilepicture"`
}

func ToResponse(data user.Core) UserReponse {
	return UserReponse{
		ID:             data.ID,
		Name:           data.Name,
		Profilepicture: data.Profilepicture,
	}
}

type UpdateUserResp struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Address        string `json:"address"`
	Password       string `json:"password"`
	Profilepicture string `json:"profilepicture"`
}

func PPToResponse(data user.Core) UpdateUserResp {
	return UpdateUserResp{
		Name:           data.Name,
		Email:          data.Email,
		Phone:          data.Phone,
		Address:        data.Address,
		Password:       data.Password,
		Profilepicture: data.Profilepicture,
	}
}

func PrintSuccessReponse(code int, message string, data ...interface{}) (int, interface{}) {
	resp := map[string]interface{}{}
	if len(data) < 2 {
		resp["data"] = (data[0])
	} else {
		resp["data"] = (data[0])
		resp["token"] = data[1].(string)
	}

	if message != "" {
		resp["message"] = message
	}

	return code, resp
}

func PrintErrorResponse(msg string) (int, interface{}) {
	resp := map[string]interface{}{}
	code := -1
	if msg != "" {
		resp["message"] = msg
	}

	if strings.Contains(msg, "server") {
		code = http.StatusInternalServerError
	} else if strings.Contains(msg, "format") {
		code = http.StatusBadRequest
	} else if strings.Contains(msg, "not found") {
		code = http.StatusNotFound
	}

	return code, resp
}
