package items

import (
	"net/http"
	"simple-user-inventory/db/models"
	"simple-user-inventory/operation/role"
	"simple-user-inventory/server/context"
	"simple-user-inventory/server/quick"
	"simple-user-inventory/server/session"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ItemRegistrationForm struct {
	Name        string `form:"name"`
	Description string `form:"description"`
	Price       uint64 `form:"price"`
}

func requireAdmin(c echo.Context, uuid string) (*models.User, error) {
	ctrl := c.(*context.Context).User()
	user, err := ctrl.ReadByUuid(uuid)
	if err == gorm.ErrRecordNotFound {
		c.Logger().Warn(err)
		return nil, quick.BadRequest()
	} else if err != nil {
		c.Logger().Error(err)
		return nil, quick.ServiceError()
	}
	if user.Role != role.Admin {
		return nil, quick.NotAllowed()
	}

	return user, nil
}

func Create(c echo.Context) error {
	sess, err := session.RequireSession(c)
	if err != nil {
		return err
	}
	_, err = requireAdmin(c, sess.Uuid)
	if err != nil {
		return err
	}

	formData := &ItemRegistrationForm{}
	if err := quick.ProcessFormData(c, formData); err != nil {
		c.Logger().Warn(err)
		return quick.BadRequest()
	}

	ctrl := c.(*context.Context).Item()
	if err = ctrl.Create(
		formData.Name,
		formData.Description,
		formData.Price,
	); err != nil {
		c.Logger().Error(err)
		return quick.ServiceError()
	}

	return c.NoContent(http.StatusOK)
}
