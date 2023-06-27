package jewel

import (
	"net/http"
	"simple-user-inventory/db/models"
	"simple-user-inventory/operation/jewel"
	"simple-user-inventory/server/context"
	"simple-user-inventory/server/quick"
	"simple-user-inventory/server/session"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Jewels(c echo.Context) error {
	sess, err := session.RequireSession(c)
	if err != nil {
		return err
	}

	ctrl := c.(*context.Context).Jewel()
	j, err := ctrl.Jewels(sess.Id)
	if err == gorm.ErrRecordNotFound {
		c.Logger().Warn(err)
		return quick.BadRequest()
	} else if err != nil {
		c.Logger().Error(err)
		return quick.ServiceError()
	}

	return c.JSON(http.StatusOK, j)
}

func Gain(c echo.Context) error {
	sess, err := session.RequireSession(c)
	if err != nil {
		return err
	}

	add := &models.JewelData{
		Red:    jewel.RandomRed(),
		Blue:   jewel.RandomBlue(),
		Green:  jewel.RandomGreen(),
		Yellow: jewel.RandomYellow(),
		Black:  jewel.RandomBlack(),
	}
	ctrl := c.(*context.Context).Jewel()
	j, err := ctrl.Gain(sess.Id, add)
	if err == gorm.ErrRecordNotFound {
		c.Logger().Warn(err)
		return quick.BadRequest()
	} else if err != nil {
		c.Logger().Error(err)
		return quick.ServiceError()
	}

	return c.JSON(http.StatusOK, j)
}
