package userstates

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Update(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}
