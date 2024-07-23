package controllers

import (
	models "product-service/models/entity"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (c *controller) GetProductHandler(e echo.Context) error {
	var (
		req = e.Request()
		ctx = req.Context()
	)

	token := req.Header.Get("Authorization")
	return c.productService.GetProduct(ctx, token).Write(e)
}

func (c *controller) InsertorUpdateProductHandler(e echo.Context) error {
	var (
		req        = e.Request()
		ctx        = req.Context()
		reqProduct = &models.Product{}
	)

	if err := e.Bind(&reqProduct); err != nil {
		logrus.WithFields(logrus.Fields{
			"service": "InsertorUpdateShopHandler",
			"event":   "Bind",
		}).Error(err)
	}

	return c.productService.InsertOrUpdateProduct(ctx, *reqProduct).Write(e)
}
