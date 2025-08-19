package repository

import (
	userstate "app/internal/core/models/user-state"
	"context"
	"errors"

	"gorm.io/gorm"
)

type UserStateRepository struct {
	db *gorm.DB
}

func NewUserStateRepository(db *gorm.DB) (*UserStateRepository, error) {
	if db == nil {
		return nil, errors.New("db cannot be nil")
	}

	return &UserStateRepository{
		db: db,
	}, nil
}

func (r *UserStateRepository) Insert(ctx context.Context, us *userstate.UserState) error {
	return r.db.WithContext(ctx).Create(&us).Error
}
