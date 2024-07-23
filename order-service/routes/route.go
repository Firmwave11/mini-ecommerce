package routes

import (
	"order-service/controllers"
	"order-service/middleware"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
)

type route struct {
	middleware middleware.Clients
	controller controllers.Controller
}

type Route interface {
	Route() *echo.Echo
}

func NewRouter(mdl middleware.Clients, controller controllers.Controller) Route {
	return &route{
		middleware: mdl,
		controller: controller,
	}
}

func (r *route) Route() *echo.Echo {
	router := echo.New()
	router.Use(echoprometheus.NewMiddleware("myapp"))
	router.GET("/metrics", echoprometheus.NewHandler())

	baseRoute := router.Group("", echo.WrapMiddleware(r.middleware.Logger))

	order := baseRoute.Group("/order", echo.WrapMiddleware(r.middleware.CheckToken))

	order.POST("", r.controller.CreateOrderHandler)

	return router
}
