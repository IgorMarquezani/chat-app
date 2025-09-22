package controllers

import (
	privatechats "app/internal/adapters/controllers/private-chats"
	privatemessages "app/internal/adapters/controllers/private-messages"
	userstates "app/internal/adapters/controllers/user-states"
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
	app.GET("/api/users/me", users.Me)
	app.GET("/api/users/search/:name", users.Search)
	app.PUT("/api/users/state", userstates.Update)

	app.POST("/api/private/chat/create", privatechats.Create)
	app.GET("/api/private/chat/list", privatechats.List)

	app.GET("/api/ws/private/:id/connect", privatemessages.Chat)
}
