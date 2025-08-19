package privatechats

import (
	"app/internal/adapters/database"
	"app/internal/adapters/database/repository"
	"app/internal/core/services/private-chats"
	"app/internal/core/services/sessions"
	"net/http"

	"github.com/labstack/echo/v4"
)

func List(c echo.Context) error {
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

	u := WLMSG.Data

	privateChatRepo, err := repository.NewPrivateChatRepository(db)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	LPCSVC := privatechats.NewListPrivateChatSVC(privateChatRepo)

	LPCMSG := LPCSVC.List(c.Request().Context(), u.ID)

	return c.JSONPretty(int(LPCMSG.Status), &LPCMSG, "  ")
}
