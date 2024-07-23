package routes

import (
	"user-service/controllers"
	"user-service/middleware"

	"github.com/labstack/echo"
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

	baseRoute := router.Group("", echo.WrapMiddleware(r.middleware.Logger))

	baseRoute.POST("/login", r.controller.LoginHandler)
	baseRoute.POST("/register", r.controller.RegisterHandler)
	baseRoute.POST("/validate-token", r.controller.ValidateTokenHandler)
	baseRoute.GET("/user", r.controller.GetUserHandler)

	return router
}
