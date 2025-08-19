package privatechat

import (
	"app/internal/core/models/user"
	"time"
)

type PrivateChat struct {
	ID           string    `gorm:"type:uuid;unique;not null" json:"id"`
	User1ID      uint32    `gorm:"primaryKey" json:"-"`
	User2ID      uint32    `gorm:"primaryKey" json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	LastActivity time.Time `json:"last_activity"`

	User1 user.User `gorm:"foreignKey:User1ID;references:ID" json:"-"`
	User2 user.User `gorm:"foreignKey:User2ID;references:ID" json:"-"`
}

type UserPrivateChat struct {
	ChatID     string `json:"chat_id"`
	FriendName string `json:"friend_name"`
}
