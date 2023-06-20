package handlers

import (
	"errors"
	"net/http"
	"simple-user-inventory/server/context"
	"simple-user-inventory/server/quick"

	"github.com/labstack/echo/v4"
)

type RegistrationForm struct {
	Name     string `form:"name" validate:"required,alphanum,min=2,max=64"`
	Email    string `form:"email" validate:"required,email,max=64"`
	Password string `form:"password" validate:"required,alphanum,min=8,max=64"`
}

type LoginForm struct {
	Email    string `form:"email" validate:"required,email,max=64"`
	Password string `form:"password" validate:"required,alphanum,min=8,max=64"`
}

func Register(c echo.Context) error {
	formData := RegistrationForm{}
	if err := quick.ProcessFormData(c, &formData); err != nil {
		c.Logger().Warn(err)
		return quick.BadRequest()
	}

	ctrl := c.(*context.Context).User()
	if err := ctrl.Create(
		formData.Name,
		formData.Email,
		formData.Password,
	); err != nil {
		c.Logger().Error(err)
		return quick.ServiceError()
	}

	return c.NoContent(http.StatusOK)
}

func Login(c echo.Context) error {
	formData := LoginForm{}
	if err := quick.ProcessFormData(c, &formData); err != nil {
		c.Logger().Warn(err)
		return quick.BadRequest()
	}

	return errors.New("not implemented")
}
