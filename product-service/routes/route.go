package routes

import (
	"product-service/controllers"
	"product-service/middleware"

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

	shop := baseRoute.Group("/product", echo.WrapMiddleware(r.middleware.CheckToken))

	shop.GET("", r.controller.GetProductHandler)
	shop.POST("", r.controller.InsertorUpdateProductHandler)

	return router
}
