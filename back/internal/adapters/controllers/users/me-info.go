package users

import (
	"app/internal/adapters/database"
	"app/internal/adapters/database/repository"
	"app/internal/core/services"
	"app/internal/core/services/users"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Me(c echo.Context) error {
	db, err := database.GetDbConnection()
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, &services.APIMessage[any]{
			Error:  "internal server error",
			Status: http.StatusInternalServerError,
		}, "  ")
	}

	userRepo, err := repository.NewUserRepository(db)
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, &services.APIMessage[any]{
			Error:  "internal server error",
			Status: http.StatusInternalServerError,
		}, "  ")
	}

	sessionRepo, err := repository.NewSessionRepository(db)
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, &services.APIMessage[any]{
			Error:  "internal server error",
			Status: http.StatusInternalServerError,
		}, "  ")
	}

	meInfoSVC := users.NewMeInfoSVC(userRepo, sessionRepo, c)

	meInfoMSG := meInfoSVC.MeInfo(c.Request().Context())

	return c.JSONPretty(int(meInfoMSG.Status), &meInfoMSG, "  ")
}
