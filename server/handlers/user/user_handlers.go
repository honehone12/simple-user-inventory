package user

import (
	"net/http"
	"simple-user-inventory/db/controller"
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

	ctrl := c.(*context.Context).User()
	balance, err := ctrl.Purchase(sess.Id, formData.Id)
	if err == controller.ErrorBalanceNotEnough ||
		err == controller.ErrorAlreadyPurchased {

		c.Logger().Warn(err)
		return quick.BadRequest()
	}
	if err != nil {
		c.Logger().Error(err)
		return quick.ServiceError()
	}

	return c.JSON(http.StatusOK, balance)
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

	len := len(it)
	itemMap := make(UserItemsResponse, len)
	for i := 0; i < len; i++ {
		itemMap[it[i].ID] = &it[i].ItemData
	}

	return c.JSON(http.StatusOK, itemMap)
}
