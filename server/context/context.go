package context

import (
	"simple-user-inventory/db"

	"github.com/labstack/echo/v4"
)

type Context struct {
	echo.Context
	db.Connection
}
