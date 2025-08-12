package middleware

import (
	"app/internal/adapters/database"
	"app/internal/adapters/database/repository"
	"app/internal/core/services/users"
	"net/http"

	"github.com/labstack/echo/v4"
)

var ignoredPaths = map[string]bool{
	"/api/status":       true,
	"/api/users/login":  true,
	"/api/users/signup": true,
}

func Authenticate() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			db, err := database.GetDbConnection()
			if err != nil {
				return c.String(http.StatusInternalServerError, "internal server error")
			}

			repo, err := repository.NewSessionRepository(db)
			if err != nil {
				return c.String(http.StatusInternalServerError, "internal server error")
			}

			svc := users.NewIsLoggedSvc(c, repo)

			println(c.Request().URL.Path)

			msg := svc.IsLogged(c.Request().Context(), c.Request().URL.Path, ignoredPaths)

			if msg.Status == http.StatusOK {
				return next(c)
			}

			return c.String(http.StatusUnauthorized, "not logged in")
		}
	}
}
