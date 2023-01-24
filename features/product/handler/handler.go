package handler

import (
	"ecommerce/features/product"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type handlerController struct {
	srv product.ProductService
}

func New(ps product.ProductService) product.ProductHandler {
	return &handlerController{
		srv: ps,
	}
}

// AddProduct implements product.ProductHandler
func (hc *handlerController) AddProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := AddProductRequest{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "wrong input format"})
		}
		// log.Println(input.FileHeader)
		// return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "wrong input format"})
		//proses cek ada gambar atau tidak
		checkFile, _, _ := c.Request().FormFile("file")
		//cek file kalau ada isi nya
		if checkFile != nil {
			formHeader, err := c.FormFile("file")
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Select a file to upload"})
			}
			input.FileHeader = *formHeader
		}
		res, err := hc.srv.AddProduct(c.Get("user"), input.FileHeader, *RequestToCore(input))
		if err != nil {
			if strings.Contains(err.Error(), "type") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "only jpg or png file can be upload"})
			} else if strings.Contains(err.Error(), "size") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "max file size is 500KB"})
			} else {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
			}
		}
		log.Println(res)
		return c.JSON(http.StatusCreated, map[string]interface{}{
			// "data":    res,
			"message": "success create content",
		})
	}
}

// AllProduct implements product.ProductHandler
func (hc *handlerController) AllProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := hc.srv.AllProduct()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
		}
		// result := []AllContent{}
		// for i := 0; i < len(res); i++ {
		// 	result = append(result, AllContentResponse(res[i]))
		// }
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"message": "success",
		})
	}
}

// Delete implements product.ProductHandler
func (hc *handlerController) Delete() echo.HandlerFunc {
	panic("unimplemented")
}

// EditProduct implements product.ProductHandler
func (hc *handlerController) EditProduct() echo.HandlerFunc {
	panic("unimplemented")
}

// ProductDetail implements product.ProductHandler
func (hc *handlerController) ProductDetail() echo.HandlerFunc {
	panic("unimplemented")
}
