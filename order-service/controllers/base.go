package controllers

import (
	orderService "order-service/services"

	"github.com/labstack/echo/v4"
)

type controller struct {
	orderService orderService.OrderService
}

type Controller interface {
	CreateOrderHandler(e echo.Context) error
}

func NewController(order orderService.OrderService) Controller {
	return &controller{
		orderService: order,
	}
}
