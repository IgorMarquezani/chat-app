package ports

import (
	"app/internal/core/models/user"
	"context"
)

type UserRepository interface {
	Insert(context.Context, *user.User) error
}
