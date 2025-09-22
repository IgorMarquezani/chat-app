package user

import "time"

type User struct {
	ID        uint32    `gorm:"primaryKey" json:"id"`
	Name      UserName  `gorm:"not null" json:"name"`
	Email     Email     `gorm:"unique" json:"email"`
	Password  Password  `gorm:"not null" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
}

type FullUserInfo struct {
	ID         uint32    `json:"id"`
	Name       UserName  `json:"name"`
	Email      Email     `json:"email"`
	CreatedAt  time.Time `json:"created_at"`
	LastChatID string    `json:"last_chat_id"`
}
