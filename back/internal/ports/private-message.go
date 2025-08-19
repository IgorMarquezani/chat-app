package ports

import (
	"app/internal/core/models/private-message"
	"context"
)

type PrivateMessageRepository interface {
	Insert(ctx context.Context, pm *privatemessage.PrivateMessage) error
}
