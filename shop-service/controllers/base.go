package controllers

import (
	shopService "shop-service/services"

	"github.com/labstack/echo/v4"
)

type controller struct {
	shopService shopService.ShopService
}

type Controller interface {
	InsertorUpdateShopHandler(e echo.Context) error
	GetWarehouseHandler(e echo.Context) error
}

func NewController(shop shopService.ShopService) Controller {
	return &controller{
		shopService: shop,
	}
}
