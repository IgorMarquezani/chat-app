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

type ListPrivateChatSVC struct {
	repo ports.PrivateChatRepository
}

func NewListPrivateChatSVC(repo ports.PrivateChatRepository) *ListPrivateChatSVC {
	return &ListPrivateChatSVC{
		repo: repo,
	}
}

func (s *ListPrivateChatSVC) List(ctx context.Context, userID uint32) services.APIMessage[[]privatechat.UserPrivateChat] {
	data, err := s.repo.SelectByUserID(ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return services.APIMessage[[]privatechat.UserPrivateChat]{
				Error:   "no chats found",
				Details: nil,
				Succeed: true,
				Status:  http.StatusOK,
			}
		}

		return services.APIMessage[[]privatechat.UserPrivateChat]{
			Error:   "internal server error",
			Details: nil,
			Succeed: false,
			Status:  http.StatusInternalServerError,
		}
	}

	return services.APIMessage[[]privatechat.UserPrivateChat]{
		Error:   "ok",
		Details: nil,
		Succeed: true,
		Status:  http.StatusOK,
		Data:    data,
	}
}
