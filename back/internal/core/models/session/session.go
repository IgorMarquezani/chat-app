package session

import (
	"app/internal/core/models/user"
	"time"
)

type Session struct {
	UserId    uint32    `gorm:"not null"`
	Key       string    `gorm:"primaryKey;type:UUID"`
	ClientIp  string    `gorm:"not null"`
	UserAgent string    `gorm:"not null"`
	ExpiresAt time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`

	User user.User
}
