package repository

import (
	privatechat "app/internal/core/models/private-chat"
	"context"
	"errors"

	"gorm.io/gorm"
)

type PrivateChatRepository struct {
	db *gorm.DB
}

func NewPrivateChatRepository(db *gorm.DB) (*PrivateChatRepository, error) {
	if db == nil {
		return nil, errors.New("db cannot be nil")
	}

	return &PrivateChatRepository{
		db: db,
	}, nil
}

func (r *PrivateChatRepository) Insert(ctx context.Context, chat *privatechat.PrivateChat) error {
	return r.db.WithContext(ctx).Create(&chat).Error
}

func (r *PrivateChatRepository) SelectByUserID(ctx context.Context, userID uint32) ([]privatechat.UserPrivateChat, error) {
	arr := make([]privatechat.UserPrivateChat, 0)

	return arr, r.db.WithContext(ctx).Raw(`
    SELECT pc.id as chat_id, u.name AS friend_name
    FROM private_chats pc
    JOIN users u 
    ON u.id = CASE 
      WHEN pc.user1_id = ? THEN pc.user2_id
      ELSE pc.user1_id
      END
    WHERE pc.user1_id = ? OR pc.user2_id = ? order by last_activity, pc.created_at;`, userID, userID, userID).Scan(&arr).Error
}
