package privatemessage

import (
	"app/internal/core/models/private-chat"
	"app/internal/core/models/user"

	"time"
)

type PrivateMessage struct {
	ID        uint32    `gorm:"primaryKey" json:"id"`
	Data      string    `gorm:"not null;" json:"text"`
	SenderID  uint32    `gorm:"not null" json:"sender_id"`
	ChatID    string    `gorm:"not null" json:"chat_id"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`

	Sender user.User               `gorm:"foreignKey:SenderID;references:ID" json:"-"`
	Chat   privatechat.PrivateChat `gorm:"foreignKey:ChatID;references:ID" json:"-"`
}
