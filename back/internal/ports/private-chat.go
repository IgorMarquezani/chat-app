package ports

import (
	privatechat "app/internal/core/models/private-chat"
	"context"
)

type PrivateChatRepository interface {
	Insert(ctx context.Context, chat *privatechat.PrivateChat) error
	SelectByUserID(ctx context.Context, userID uint32) ([]privatechat.UserPrivateChat, error)
}
