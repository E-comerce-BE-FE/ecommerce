package main

import (
	"ecommerce/config"
	prdData "ecommerce/features/product/data"
	prdHdl "ecommerce/features/product/handler"
	prdSrv "ecommerce/features/product/services"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)

	// gunakan migrate
	config.Migrate(db)

	pData := prdData.New(db)
	pSrv := prdSrv.New(pData)
	pHDL := prdHdl.New((pSrv))
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	// Product
	e.GET("/product", pHDL.AllProduct())
	// ========== Run Program ===========
	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
