package handler

import "ecommerce/features/product"

// import "ecommerce/features/product"

type CoreContent struct {
	ID           uint   `json:"id_content" from:"id"`
	Content      string `validate:"required" json:"content" from:"content"`
	ContentImage string `json:"content_image" from:"content_image"`
	CreateAt     string `json:"create_at" from:"create_at"`
	NumbComment  uint   `json:"comment" from:"comment"`
	Users        CoreUser
	Comment      []CommentCore
}
type CoreUser struct {
	ID       uint   `json:"id_user" from:"id"`
	UserName string `json:"username" from:"username"`
	Name     string `json:"name" from:"name"`
	Image    string `json:"profilepicture" from:"profilepicture"`
}
type CommentCore struct {
	ID       uint   `json:"id_comment" from:"id"`
	UserName string `json:"username" from:"username"`
	Comment  string `json:"comment" from:"comment"`
	Content  CoreContent
}

type AllContent struct {
	ID           uint   `json:"id_content" from:"id_content"`
	Content      string `validate:"required" json:"content" from:"content"`
	ContentImage string `json:"content_image" from:"content_image"`
	CreateAt     string `json:"create_at" from:"create_at"`
	Users        CoreUser
	NumbComment  uint `json:"comment" from:"comment"`
}

// func AllProductResponse(data product.Core) AllContent {
// 	return AllContent{
// 		ID:           data.ID,
// 		Content:      data.Content,
// 		ContentImage: data.ContentImage,
// 		CreateAt:     data.CreateAt,
// 		NumbComment:  data.NumbComment,
// 		Users:        CoreUser(data.Users),
// 	}
// }

type GetAllResp struct {
	ID            uint   `json:"id"`
	ProductName   string `json:"product_name"`
	Address       string `json:"address"`
	Price         int    `json:"price"`
	Description   string `json:"description"`
	ProductImages string `json:"product_images"`
}

func ToResp(data product.Core) GetAllResp {
	return GetAllResp{
		ID:            data.ID,
		ProductName:   data.ProductName,
		Address:       data.User.Address,
		Price:         data.Price,
		Description:   data.Description,
		ProductImages: data.ProductImage,
	}
}

func GetAllProductResp(data []product.Core) []GetAllResp {
	res := []GetAllResp{}
	for _, v := range data {
		res = append(res, ToResp(v))
	}
	return res
}

type Search struct {
	ID           uint   `json:"id"`
	ProductName  string `json:"product_name"`
	ProductImage string `json:"product_image"`
	Price        int    `json:"price"`
	Stock        int    `json:"stock"`
	Description  string `json:"description"`
}

func SearchResponse(data product.Core) Search {
	return Search{
		ID:           data.ID,
		ProductName:  data.ProductName,
		ProductImage: data.ProductImage,
		Price:        data.Price,
		Stock:        data.Stock,
		Description:  data.Description,
	}
}
