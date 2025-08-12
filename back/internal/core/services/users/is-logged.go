package users

import (
	"app/internal/core/services"
	"app/internal/ports"
	"context"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type IsLoggedSvc struct {
	extractor ports.CookieExtractor
	repo      ports.SessionRepository
}

func NewIsLoggedSvc(extractor ports.CookieExtractor, repo ports.SessionRepository) *IsLoggedSvc {
	return &IsLoggedSvc{
		extractor: extractor,
		repo:      repo,
	}
}

func (s *IsLoggedSvc) IsLogged(ctx context.Context, url string, ignoredPaths map[string]bool) services.APIMessage {
	if _, ok := ignoredPaths[url]; ok {
		return services.NewAPIMessage("none", nil, true, http.StatusOK)
	}

	cookie, err := s.extractor.Cookie("_Session")
	if err != nil {
		return services.NewAPIMessage("not logged in", nil, false, http.StatusUnauthorized)
	}

	_, err = s.repo.SelectByKey(ctx, cookie.Value)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return services.NewAPIMessage("logged in", nil, true, http.StatusOK)
		}

		return services.NewAPIMessage("internal server error", nil, false, http.StatusInternalServerError)
	}

	return services.NewAPIMessage("not logged in", nil, false, http.StatusUnauthorized)
}
