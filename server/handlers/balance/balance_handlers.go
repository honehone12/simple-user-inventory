package balance

import (
	"net/http"
	"simple-user-inventory/server/context"
	"simple-user-inventory/server/quick"
	"simple-user-inventory/server/session"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Coin(c echo.Context) error {
	sess, err := session.RequireSession(c)
	if err != nil {
		return err
	}

	ctrl := c.(*context.Context).Balance()
	coin, err := ctrl.Coin(sess.Id)
	if err == gorm.ErrRecordNotFound {
		c.Logger().Warn(err)
		return quick.BadRequest()
	} else if err != nil {
		c.Logger().Error(err)
		return quick.ServiceError()
	}

	return c.JSON(http.StatusOK, coin)
}

func Fund(c echo.Context) error {
	sess, err := session.RequireSession(c)
	if err != nil {
		return err
	}

	ctrl := c.(*context.Context).Balance()
	coin, err := ctrl.Fund(sess.Id, 500)
	if err == gorm.ErrRecordNotFound {
		c.Logger().Warn(err)
		return quick.BadRequest()
	} else if err != nil {
		c.Logger().Error(err)
		return quick.ServiceError()
	}

	return c.JSON(http.StatusOK, coin)
}
