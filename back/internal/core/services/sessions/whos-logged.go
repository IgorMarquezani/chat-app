package sessions

import (
	"app/internal/core/models/user"
	"app/internal/core/services"
	"app/internal/ports"
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type WhosLoggedSvc struct {
	repo      ports.SessionRepository
	extractor ports.CookieExtractor
}

func NewWhosLoggedSvc(repo ports.SessionRepository, extractor ports.CookieExtractor) *WhosLoggedSvc {
	return &WhosLoggedSvc{
		repo:      repo,
		extractor: extractor,
	}
}

func (s *WhosLoggedSvc) WhosLogged(ctx context.Context) services.APIMessage[user.User] {
	cookie, err := s.extractor.Cookie("_Session")
	if err != nil {
		if errors.Is(err, echo.ErrCookieNotFound) {
			return services.NewAPIMessage("", nil, false, http.StatusUnauthorized, user.User{})
		}

		return services.NewAPIMessage("", nil, false, http.StatusInternalServerError, user.User{})
	}

	u, err := s.repo.SelectUserBySession(ctx, cookie.Value)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return services.NewAPIMessage("no users found", nil, false, http.StatusInternalServerError, u)
		}

		return services.NewAPIMessage("", nil, false, http.StatusInternalServerError, u)
	}

	return services.NewAPIMessage("", nil, true, http.StatusOK, u)
}
