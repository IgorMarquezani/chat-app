package groupparticipant

import (
	groupchat "app/internal/core/models/group-chat"
	"app/internal/core/models/user"
)

type GroupParticipant struct {
	GroupChatID string `gorm:"primaryKey"`
	UserID      uint32 `gorm:"primaryKey"`

	GroupChat groupchat.GroupChat `gorm:"foreignKey:GroupChatID;references:ID"`
	User      user.User           `gorm:"foreignKey:UserID;references:ID"`
}
