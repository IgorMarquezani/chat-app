package controllers

import (
	"app/internal/adapters/controllers/users"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(app *echo.Echo) {
	app.GET("/api/status", func(c echo.Context) error {
		return c.String(http.StatusOK, "server available")
	})

	app.POST("/api/users/signup", users.SignUp)
	app.POST("/api/users/login", users.LogIn)
}
