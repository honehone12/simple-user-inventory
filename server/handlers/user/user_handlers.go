package user

import (
	"net/http"
	"simple-user-inventory/db/models"
	"simple-user-inventory/server/context"
	"simple-user-inventory/server/quick"
	"simple-user-inventory/server/session"

	"github.com/labstack/echo/v4"
)

type UserItemsResponse map[uint]*models.ItemData

type PurchaseForm struct {
	Id uint `form:"id" validate:"required,min=1"`
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

	// consume coins

	ctrl := c.(*context.Context).User()
	if err = ctrl.Purchase(sess.Id, formData.Id); err != nil {
		c.Logger().Error(err)
		return quick.ServiceError()
	}

	return c.NoContent(http.StatusOK)
}

func Items(c echo.Context) error {
	sess, err := session.RequireSession(c)
	if err != nil {
		return err
	}

	ctrl := c.(*context.Context).User()
	it, err := ctrl.ReadOwnedItems(sess.Id)
	if err != nil {
		c.Logger().Error(err)
		return quick.ServiceError()
	}

	itemMap := make(UserItemsResponse)
	for _, item := range it {
		itemMap[item.ID] = item.ItemData
	}

	return c.JSON(http.StatusOK, itemMap)
}
