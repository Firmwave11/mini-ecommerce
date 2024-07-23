package controllers

import (
	model "user-service/models/token"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func (c *controller) LoginHandler(e echo.Context) error {
	var (
		req      = e.Request()
		ctx      = req.Context()
		reqLogin = new(model.Login)
	)

	if err := e.Bind(&reqLogin); err != nil {
		logrus.WithFields(logrus.Fields{
			"service": "LoginHandler",
			"event":   "Bind",
		}).Error(err)
	}

	return c.tokenService.LoginAction(ctx, *reqLogin).Write(e)
}

func (c *controller) ValidateTokenHandler(e echo.Context) error {
	var (
		req      = e.Request()
		ctx      = req.Context()
		reqToken = new(model.Token)
	)

	if err := e.Bind(&reqToken); err != nil {
		logrus.WithFields(logrus.Fields{
			"service": "LoginHandler",
			"event":   "Bind",
		}).Error(err)
	}

	return c.tokenService.ValidateToken(ctx, reqToken.Token).Write(e)
}

func (c *controller) GetUserHandler(e echo.Context) error {
	var (
		req      = e.Request()
		ctx      = req.Context()
		reqToken = e.QueryParam("userId")
	)

	return c.tokenService.User(ctx, reqToken).Write(e)
}
