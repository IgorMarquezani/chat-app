package groupmessage

import (
	groupparticipant "app/internal/core/models/group-participant"
	"time"
)

type GroupMessage struct {
	GroupID   string    `gorm:"not null"`
	Sender    uint32    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`

	GroupParticipant groupparticipant.GroupParticipant `gorm:"foreignKey:GroupID,Sender;references:GroupChatID,UserID"`
}
