package controllers

import (
	models "shop-service/models/entity"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (c *controller) GetWarehouseHandler(e echo.Context) error {
	var (
		req = e.Request()
		ctx = req.Context()
	)

	return c.shopService.GetShop(ctx).Write(e)
}

func (c *controller) InsertorUpdateShopHandler(e echo.Context) error {
	var (
		req     = e.Request()
		ctx     = req.Context()
		reqShop = &models.Shop{}
	)

	if err := e.Bind(&reqShop); err != nil {
		logrus.WithFields(logrus.Fields{
			"service": "InsertorUpdateShopHandler",
			"event":   "Bind",
		}).Error(err)
	}

	return c.shopService.InsertorUpdateShop(ctx, *reqShop).Write(e)
}
