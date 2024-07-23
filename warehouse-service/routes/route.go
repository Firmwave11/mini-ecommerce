package routes

import (
	"warehouse-service/controllers"
	"warehouse-service/middleware"

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

	warehouse := baseRoute.Group("/warehouse", echo.WrapMiddleware(r.middleware.CheckToken))

	warehouse.GET("", r.controller.GetWarehouseHandler)
	warehouse.POST("", r.controller.InsertWarehouseHandler)

	warehouse.GET("/stock", r.controller.GetStockWarehouseHandler)
	warehouse.PUT("/stock", r.controller.UpdateStockWarehouseHandler)
	warehouse.POST("/stock/sync", r.controller.SyncStockWarehouseHandler)
	warehouse.POST("/stock", r.controller.InsertStockWarehouseHandler)

	return router
}
