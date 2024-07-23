package controllers

import (
	model "user-service/models/user"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func (c *controller) RegisterHandler(e echo.Context) error {
	var (
		req         = e.Request()
		ctx         = req.Context()
		reqRegister = new(model.Register)
	)

	if err := e.Bind(&reqRegister); err != nil {
		logrus.WithFields(logrus.Fields{
			"service": "CreateRegistrationHandler",
			"event":   "Bind",
		}).Error(err)
	}

	return c.userService.RegisterAction(ctx, *reqRegister).Write(e)
}
