package userstates

import (
	"app/internal/adapters/database"
	"app/internal/adapters/database/repository"
	"app/internal/core/services"
	"app/internal/core/services/sessions"
	"app/internal/core/services/user-states"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UpdateJSON struct {
	ChatID string `json:"chat_id"`
}

func Update(c echo.Context) error {
	var data UpdateJSON

	if err := c.Bind(&data); err != nil {
		if errors.Is(err, echo.ErrUnsupportedMediaType) {
			return c.JSONPretty(http.StatusBadRequest, &services.APIMessage[any]{
				Error:  "content type should be application/json",
				Status: http.StatusBadRequest,
			}, "  ")
		}
		if errors.Is(err, &json.SyntaxError{}) || errors.Is(err, &json.UnsupportedTypeError{}) {
			return c.JSONPretty(http.StatusBadRequest, &services.APIMessage[any]{
				Error:  "invalid JSON syntax or unsupported data type",
				Status: http.StatusBadRequest,
			}, "  ")
		}

		return c.JSONPretty(http.StatusInternalServerError, &services.APIMessage[any]{
			Error:  "internal server error",
			Status: http.StatusInternalServerError,
		}, "  ")
	}

	db, err := database.GetDbConnection()
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

	WLSVC := sessions.NewWhosLoggedSvc(sessionRepo, c)

	WLMSG := WLSVC.WhosLogged(c.Request().Context())
	if !WLMSG.Succeed || WLMSG.Data.ID == 0 {
		return c.JSONPretty(http.StatusInternalServerError, &services.APIMessage[any]{
			Error:  "internal server error",
			Status: http.StatusInternalServerError,
		}, "  ")
	}

	userStateRepo, err := repository.NewUserStateRepository(db)
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, &services.APIMessage[any]{
			Error:  "internal server error",
			Status: http.StatusInternalServerError,
		}, "  ")
	}

	UCSVC := userstates.NewUpdateChatSVC(userStateRepo)

	UCMSG := UCSVC.UpdateChat(c.Request().Context(), WLMSG.Data.ID, data.ChatID)

	return c.JSONPretty(int(UCMSG.Status), &UCMSG, "  ")
}
