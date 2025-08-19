package privatechats

import (
	"app/internal/adapters/database"
	"app/internal/adapters/database/repository"
	"app/internal/core/services/private-chats"
	"app/internal/core/services/sessions"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type createJSON struct {
	TargetID uint32 `json:"target_id"`
}

func Create(c echo.Context) error {
	var data createJSON

	ct := c.Request().Header.Get("Content-Type")
	if ct != "application/json" {
		return c.String(http.StatusBadRequest, "invalid content type")
	}

	if err := c.Bind(&data); err != nil {
		if errors.Is(err, &json.SyntaxError{}) || errors.Is(err, &json.UnsupportedTypeError{}) {
			return c.String(http.StatusBadRequest, "invalid JSON syntax or invalid data type")
		}

		return c.String(http.StatusInternalServerError, "internal server error")
	}

	db, err := database.GetDbConnection()
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	sessionRepo, err := repository.NewSessionRepository(db)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	WLSVC := sessions.NewWhosLoggedSvc(sessionRepo, c)

	WLMSG := WLSVC.WhosLogged(c.Request().Context())
	if !WLMSG.Succeed || WLMSG.Data.ID == 0 {
		return c.JSONPretty(int(WLMSG.Status), &WLMSG, "  ")
	}

	privateChatRepo, err := repository.NewPrivateChatRepository(db)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	userRepo, err := repository.NewUserRepository(db)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	CPCSVC := privatechats.NewCreatePrivateChatSVC(privateChatRepo, userRepo)

	CPCMSG := CPCSVC.Create(c.Request().Context(), WLMSG.Data.ID, data.TargetID)

	return c.JSONPretty(int(WLMSG.Status), &CPCMSG, "  ")
}
