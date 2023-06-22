package handlers

import (
	"net/http"
	"simple-user-inventory/db/controller"
	"simple-user-inventory/server/context"
	"simple-user-inventory/server/quick"
	"simple-user-inventory/server/session"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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

type RootResponse struct {
	Name    string
	Version string
}

func Root(c echo.Context) error {
	ctx := c.(*context.Context)
	return c.JSON(http.StatusOK, RootResponse{
		Name:    ctx.Name,
		Version: ctx.Version,
	})
}

func Register(c echo.Context) error {
	formData := &RegistrationForm{}
	if err := quick.ProcessFormData(c, formData); err != nil {
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
	formData := &LoginForm{}
	if err := quick.ProcessFormData(c, formData); err != nil {
		c.Logger().Warn(err)
		return quick.BadRequest()
	}

	ctrl := c.(*context.Context).User()
	uuid, err := ctrl.VerifyPassword(formData.Email, formData.Password)
	if err == gorm.ErrRecordNotFound || err == controller.ErrorInvalidPassword {
		c.Logger().Error(err)
		return quick.BadRequest()
	} else if err != nil {
		c.Logger().Error(err)
		return quick.ServiceError()
	}

	err = session.Set(c, uuid)
	if err != nil {
		c.Logger().Error(err)
		return quick.ServiceError()
	}
	return c.NoContent(http.StatusOK)
}
