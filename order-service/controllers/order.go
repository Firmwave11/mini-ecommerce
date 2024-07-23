package controllers

import (
	"order-service/middleware"
	"order-service/models/request"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (c *controller) CreateOrderHandler(e echo.Context) error {
	var (
		req         = e.Request()
		ctx         = req.Context()
		reqProducts = []request.ProductRequest{}
		user        string
	)

	if userId, ok := ctx.Value(middleware.CtxUserId).(string); ok {
		user = userId
	}
	token := req.Header.Get("Authorization")
	if err := e.Bind(&reqProducts); err != nil {
		logrus.WithFields(logrus.Fields{
			"service": "InsertorUpdateShopHandler",
			"event":   "Bind",
		}).Error(err)
	}

	return c.orderService.CreateOrder(ctx, &reqProducts, user, token).Write(e)
}
