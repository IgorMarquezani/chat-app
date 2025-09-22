package userstates

import (
	"app/internal/core/services"
	"app/internal/ports"
	"context"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type UpdateChatSVC struct {
	repo ports.UserStateRepository
}

func NewUpdateChatSVC(repo ports.UserStateRepository) *UpdateChatSVC {
	return &UpdateChatSVC{
		repo: repo,
	}
}

func (s *UpdateChatSVC) UpdateChat(ctx context.Context, userID uint32, chatID string) services.APIMessage[any] {
	if err := s.repo.UpdateLastChatID(ctx, userID, chatID); err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return services.APIMessage[any]{
				Error:  "no such chat or user with the given id",
				Status: http.StatusBadRequest,
			}
		}

		return services.APIMessage[any]{
			Error:  "internal server error",
			Status: http.StatusInternalServerError,
		}
	}

	return services.APIMessage[any]{
		Error:   "ok",
		Succeed: true,
		Status:  http.StatusOK,
	}
}
