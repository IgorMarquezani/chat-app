package users

import (
	"app/internal/adapters/database"
	"app/internal/adapters/database/repository"
	"app/internal/core/services/hasher"
	"app/internal/core/services/users"

	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SignUp(c echo.Context) error {
	var JSON users.NewAccountReq

	if err := json.NewDecoder(c.Request().Body).Decode(&JSON); err != nil {
		if errors.Is(err, &json.SyntaxError{}) {
			return c.String(http.StatusBadRequest, fmt.Sprintf("JSON syntax error: %s", err.Error()))
		}

		return c.String(http.StatusInternalServerError, "internal server error")
	}

	db, err := database.GetDbConnection()
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	userRepo, err := repository.NewUserRepository(db)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	userStateRepo, err := repository.NewUserStateRepository(db)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	svc, err := users.NewAccountSvc(userRepo, userStateRepo, &hasher.Hasher{})
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	msg := svc.NewAccount(c.Request().Context(), &JSON)

	return c.JSONPretty(int(msg.Status), &msg, "  ")
}
