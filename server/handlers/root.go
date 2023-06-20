package handlers

import (
	"errors"
	"net/http"
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
	if err := quick.QuickFormData(c, &formData); err != nil {
		return quick.QuickErrorResponse(err)
	}

	return c.JSON(http.StatusOK, formData)
}

func Login(c echo.Context) error {
	formData := LoginForm{}
	if err := quick.QuickFormData(c, &formData); err != nil {
		return quick.QuickErrorResponse(err)
	}

	return errors.New("not implemented")
}
