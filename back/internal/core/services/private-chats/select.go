package privatechats

import (
	privatechat "app/internal/core/models/private-chat"
	"app/internal/core/services"
	"app/internal/ports"
	"context"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type SelectSVC struct {
	repo ports.PrivateChatRepository
}

func NewSelectSVC(repo ports.PrivateChatRepository) *SelectSVC {
	return &SelectSVC{
		repo: repo,
	}
}

func (s *SelectSVC) Select(ctx context.Context, chatID string) services.APIMessage[privatechat.PrivateChat] {
	pc, err := s.repo.Select(ctx, chatID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return services.APIMessage[privatechat.PrivateChat]{
				Error:  "no such private chat with the given id",
				Status: http.StatusBadRequest,
			}
		}

		return services.APIMessage[privatechat.PrivateChat]{
			Error:  "internal server error",
			Status: http.StatusInternalServerError,
		}
	}

	return services.APIMessage[privatechat.PrivateChat]{
		Error:   "ok",
		Succeed: true,
		Status:  http.StatusOK,
		Data:    pc,
	}
}
