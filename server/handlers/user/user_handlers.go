package user

import (
	"net/http"
	"simple-user-inventory/server/context"
	"simple-user-inventory/server/quick"
	"simple-user-inventory/server/session"

	"github.com/labstack/echo/v4"
)

type PurchaseForm struct {
	Id uint `form:"id"`
}

func Purchase(c echo.Context) error {
	sess, err := session.RequireSession(c)
	if err != nil {
		return err
	}

	formData := &PurchaseForm{}
	if err = quick.ProcessFormData(c, formData); err != nil {
		c.Logger().Warn(err)
		return quick.BadRequest()
	}
	if formData.Id == 0 {
		return quick.BadRequest()
	}

	ctrl := c.(*context.Context).User()
	if err = ctrl.Purchase(sess.Id, formData.Id); err != nil {
		c.Logger().Error(err)
		return quick.ServiceError()
	}

	return c.NoContent(http.StatusOK)
}
