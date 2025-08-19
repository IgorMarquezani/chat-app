package ports

import (
	"app/internal/core/models/user"
	"context"
)

type UserRepository interface {
	Insert(context.Context, *user.User) error
	SelectByEmail(context.Context, string) (user.User, error)
	SelectByName(context.Context, string) ([]user.User, error)
	SelectByID(ctx context.Context, id uint32) (user.User, error)
	DeleteByID(ctx context.Context, id uint32) error
}
