package groupchat

import (
	"app/internal/core/models/user"
)

type GroupChat struct {
	ID      string `gorm:"primaryKey;type:UUID"`
	Name    string `gorm:"not null"`
	OwnerID uint32 `gorm:"not null"`

	Owner user.User `gorm:"foreignKey:OwnerID;references:ID"`
}
