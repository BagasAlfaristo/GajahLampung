package routers

import (
	_productData "gajahlampung/features/product/data"
	_productHandler "gajahlampung/features/product/handler"
	_productService "gajahlampung/features/product/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoutuer(e *echo.Echo, db *gorm.DB) {
	productData := _productData.New(db)
	productService := _productService.New(productData)
	productHandlerAPI := _productHandler.New(productService)

	e.POST("/products", productHandlerAPI.AddProduct)
	e.GET("/products", productHandlerAPI.GetProduct)
	e.GET("/products/name", productHandlerAPI.GetProductByName)
}
