package privatechats

import (
	privatechat "app/internal/core/models/private-chat"
	"app/internal/core/services"
	"app/internal/ports"
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreatePrivateChatSVC struct {
	PMRepo   ports.PrivateChatRepository
	userRepo ports.UserRepository
}

func NewCreatePrivateChatSVC(PMRepo ports.PrivateChatRepository, userRepo ports.UserRepository) *CreatePrivateChatSVC {
	return &CreatePrivateChatSVC{
		PMRepo:   PMRepo,
		userRepo: userRepo,
	}
}

func (s *CreatePrivateChatSVC) Create(ctx context.Context, senderID uint32, targetID uint32) services.APIMessage[privatechat.UserPrivateChat] {
	chat := privatechat.PrivateChat{ID: uuid.NewString()}
	if senderID > targetID {
		chat.User1ID = senderID
		chat.User2ID = targetID
	} else {
		chat.User1ID = targetID
		chat.User2ID = senderID
	}

	if err := s.PMRepo.Insert(ctx, &chat); err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return services.APIMessage[privatechat.UserPrivateChat]{
				Error:   "no such user with the given id",
				Details: nil,
				Succeed: false,
				Status:  http.StatusBadRequest,
			}
		}

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return services.APIMessage[privatechat.UserPrivateChat]{
				Error:   "chat already exists",
				Details: nil,
				Succeed: false,
				Status:  http.StatusAlreadyReported,
			}
		}

		return services.APIMessage[privatechat.UserPrivateChat]{
			Error:   "internal server error",
			Details: nil,
			Succeed: false,
			Status:  http.StatusInternalServerError,
		}
	}

	u, err := s.userRepo.SelectByID(ctx, targetID)
	if err != nil {
		return services.APIMessage[privatechat.UserPrivateChat]{
			Error:   "internal server error",
			Details: nil,
			Succeed: false,
			Status:  http.StatusInternalServerError,
		}
	}

	userChat := privatechat.UserPrivateChat{
		ChatID:     chat.ID,
		FriendName: string(u.Name),
	}

	return services.APIMessage[privatechat.UserPrivateChat]{
		Error:   "ok",
		Details: nil,
		Succeed: true,
		Status:  http.StatusOK,
		Data:    userChat,
	}
}
