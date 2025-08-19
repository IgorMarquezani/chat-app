package userstate

import (
	privatechat "app/internal/core/models/private-chat"
	"app/internal/core/models/user"
)

type UserState struct {
	UserID     uint32  `json:"-" gorm:"primaryKey"`
	LastChatID *string `json:"last_chat_id" gorm:"type:uuid"`

	User user.User               `gorm:"foreignKey:UserID;references:ID"`
	Chat privatechat.PrivateChat `gorm:"foreignKey:LastChatID;references:ID"`
}
