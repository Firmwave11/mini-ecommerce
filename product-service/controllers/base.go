package controllers

import (
	productService "product-service/services"

	"github.com/labstack/echo/v4"
)

type controller struct {
	productService productService.ProductService
}

type Controller interface {
	InsertorUpdateProductHandler(e echo.Context) error
	GetProductHandler(e echo.Context) error
}

func NewController(shop productService.ProductService) Controller {
	return &controller{
		productService: shop,
	}
}
