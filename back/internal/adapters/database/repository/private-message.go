package repository

import (
	privatemessage "app/internal/core/models/private-message"
	"context"
	"errors"

	"gorm.io/gorm"
)

type PrivateMessageRepository struct {
	db *gorm.DB
}

func NewPrivateMessageRepository(db *gorm.DB) (*PrivateMessageRepository, error) {
	if db == nil {
		return nil, errors.New("db cannot be nil")
	}
	return &PrivateMessageRepository{
		db: db,
	}, nil
}

func (r *PrivateMessageRepository) Insert(ctx context.Context, pm *privatemessage.PrivateMessage) error {
	return r.db.WithContext(ctx).Create(&pm).Error
}
