package users

import (
	"app/internal/adapters/database"
	"app/internal/adapters/database/repository"
	"app/internal/core/services/hasher"
	"app/internal/core/services/users"
	"encoding/json"

	"net/http"

	"github.com/labstack/echo/v4"
)

func LogIn(c echo.Context) error {
	db, err := database.GetDbConnection()
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	userRepo, err := repository.NewUserRepository(db)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	sessionRepo, err := repository.NewSessionRepository(db)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	hasher := hasher.Hasher{}

	svc := users.NewSignInSvc(userRepo, sessionRepo, &hasher, c, c.Request().Header)

	JSON := users.SignInReq{
		ClientIP: c.RealIP(),
	}

	json.NewDecoder(c.Request().Body).Decode(&JSON)

	msg := svc.SignIn(c.Request().Context(), JSON)

	return c.JSONPretty(int(msg.Status), &msg, "  ")
}
