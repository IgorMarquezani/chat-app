package repository

import (
	"context"
	"errors"

	"app/internal/core/models/user"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) (*UserRepository, error) {
	if db == nil {
		return nil, errors.New("db cannot be nil")
	}

	return &UserRepository{
		db: db,
	}, nil
}

func (r *UserRepository) Insert(ctx context.Context, u *user.User) error {
	return r.db.WithContext(ctx).Create(u).Error
}

func (r *UserRepository) SelectByEmail(ctx context.Context, email string) (user.User, error) {
	var u user.User

	return u, r.db.WithContext(ctx).Where("email = ?", email).First(&u).Error
}
