package users

import (
	"app/internal/adapters/database"
	"app/internal/adapters/database/repository"
	"app/internal/core/services/users"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Search(c echo.Context) error {
	db, err := database.GetDbConnection()
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	repo, err := repository.NewUserRepository(db)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	svc, err := users.NewSearchSvc(repo)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	msg := svc.Search(c.Request().Context(), c.Param("name"))
	if !msg.Succeed {
		return c.JSONPretty(int(msg.Status), &msg, "  ")
	}

	return c.JSONPretty(int(msg.Status), &msg, "  ")
}
