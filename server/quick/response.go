package quick

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func QuickErrorResponse(err error) error {
	return echo.NewHTTPError(
		http.StatusBadRequest,
		err.Error(),
	)
}
