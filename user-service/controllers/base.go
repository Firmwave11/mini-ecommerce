package controllers

import (
	"user-service/services/token"
	"user-service/services/user"

	"github.com/labstack/echo"
)

type controller struct {
	userService  user.UserService
	tokenService token.TokenService
}

type Controller interface {
	LoginHandler(e echo.Context) error
	RegisterHandler(e echo.Context) error
	ValidateTokenHandler(e echo.Context) error
	GetUserHandler(e echo.Context) error
}

func NewController(user user.UserService, token token.TokenService) Controller {
	return &controller{
		userService:  user,
		tokenService: token,
	}
}
