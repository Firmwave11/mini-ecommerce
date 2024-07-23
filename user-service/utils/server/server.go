package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

// Wrapper wraping the routing server with cors handler
func Wrapper(route *echo.Echo) error {
	// cors condition
	origins := viper.GetStringSlice("cors.whitelist")
	route.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     origins,
		AllowCredentials: true,
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	route.HideBanner = true

	port := fmt.Sprintf(":%d", viper.GetInt("application.port"))
	return route.Start(port)
}
