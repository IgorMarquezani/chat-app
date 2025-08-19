package users

import (
	"app/internal/core/models/user"
	"app/internal/core/services"
	"app/internal/ports"
	"context"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type SearchByNameSvc struct {
	repo ports.UserRepository
}

func NewSearchSvc(repo ports.UserRepository) (*SearchByNameSvc, error) {
	if repo == nil {
		return nil, errors.New("repo cannot be nil")

	}

	return &SearchByNameSvc{
		repo: repo,
	}, nil
}

func (s *SearchByNameSvc) Search(ctx context.Context, userName string) services.APIMessage[[]user.User] {
	users, err := s.repo.SelectByName(ctx, userName)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return services.NewAPIMessage[[]user.User]("no users found", nil, true, http.StatusOK, nil)
		}

		return services.NewAPIMessage[[]user.User]("internal server error", nil, false, http.StatusInternalServerError, nil)
	}

	return services.NewAPIMessage("", nil, true, http.StatusOK, users)
}
