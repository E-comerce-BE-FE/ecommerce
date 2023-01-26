package handler

import (
	"ecommerce/features/product"
	"log"
	"net/http"
	"strconv"
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
		// //debugging
		// log.Println(checkFile)
		// return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Select a file to upload"})
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
			"message": "success add product",
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
			"data":    GetAllProductResp(res),
			"message": "success show all product",
		})
	}
}

// Delete implements product.ProductHandler
func (hc *handlerController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		prdID := c.Param("id")
		productID, _ := strconv.Atoi(prdID)
		err := hc.srv.Delete(c.Get("user"), uint(productID))
		if err != nil {
			if strings.Contains(err.Error(), "not") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "you are not allowed delete other people product"})
			} else {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error, deleting product fail"})
			}
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete product",
		})
	}
}

// EditProduct implements product.ProductHandler
func (hc *handlerController) EditProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		prdID := c.Param("id")
		productID, _ := strconv.Atoi(prdID)
		input := EditProductRequest{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "wrong input format"})
		}
		// log.Println(input.FileHeader)
		// return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "wrong input format"})
		//proses cek ada gambar atau tidak
		checkFile, _, _ := c.Request().FormFile("file")
		// Debugging
		// log.Println(checkFile)
		// return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Select a file to upload"})
		//cek file kalau ada isi nya
		if checkFile != nil {
			formHeader, err := c.FormFile("file")
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Select a file to upload"})
			}
			input.FileHeader = *formHeader
		}
		res, err := hc.srv.EditProduct(c.Get("user"), input.FileHeader, uint(productID), *RequestToCore(input))
		if err != nil {
			if strings.Contains(err.Error(), "type") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "only jpg or png file can be insert"})
			} else if strings.Contains(err.Error(), "size") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "max file size is 500KB"})
			} else {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
			}
		}
		log.Println(res)
		return c.JSON(http.StatusCreated, map[string]interface{}{
			// "data":    res,
			"message": "success change product data",
		})
	}
}

// ProductDetail implements product.ProductHandler
func (hc *handlerController) ProductDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		prdID := c.Param("id")
		productID, _ := strconv.Atoi(prdID)
		res, err := hc.srv.ProductDetail(uint(productID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"message": "success",
		})
	}
}

// Searching implements product.ProductHandler
func (hc *handlerController) Searching() echo.HandlerFunc {
	return func(c echo.Context) error {
		quotes := c.QueryParam("q")
		log.Println(quotes)
		res, err := hc.srv.Searching(quotes)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "user not found"})
		}
		result := []Search{}
		for i := 0; i < len(res); i++ {
			result = append(result, SearchResponse(res[i]))
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    result,
			"message": "searching success",
		})
	}
}
