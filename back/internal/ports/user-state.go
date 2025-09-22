package ports

import (
	userstate "app/internal/core/models/user-state"
	"context"
)

type UserStateRepository interface {
	Insert(ctx context.Context, us *userstate.UserState) error
	UpdateLastChatID(ctx context.Context, userID uint32, chatID string) error
}
