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

func (s *IsLoggedSvc) IsLogged(ctx context.Context, url string, ignoredPaths map[string]bool) services.APIMessage[any] {
	if _, ok := ignoredPaths[url]; ok {
		return services.NewAPIMessage[any]("none", nil, true, http.StatusOK, nil)
	}

	cookie, err := s.extractor.Cookie("_Session")
	if err != nil {
		return services.NewAPIMessage[any]("not logged in", nil, false, http.StatusUnauthorized, nil)
	}

	_, err = s.repo.SelectByKey(ctx, cookie.Value)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return services.NewAPIMessage[any]("not logged in", nil, false, http.StatusUnauthorized, nil)
		}

		return services.NewAPIMessage[any]("internal server error", nil, false, http.StatusInternalServerError, nil)
	}

	return services.NewAPIMessage[any]("logged in", nil, true, http.StatusOK, nil)
}
