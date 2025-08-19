package privatemessages

import (
	privatemessage "app/internal/core/models/private-message"
	"app/internal/core/services"
	"app/internal/ports"
	"context"
	"net/http"
)

type CreateMessageReq struct {
	Data     string `json:"text"`
	ChatID   string `json:"chat_id"`
	SenderID uint32 `json:"-"`
}

type CreateMessageSvc struct {
	PMRepo ports.PrivateMessageRepository
}

func NewCreateMessageSvc(PMRepo ports.PrivateMessageRepository) *CreateMessageSvc {
	return &CreateMessageSvc{
		PMRepo: PMRepo,
	}
}

func (s *CreateMessageSvc) CreateMessage(ctx context.Context, req *CreateMessageReq) services.APIMessage[any] {
	pm := privatemessage.PrivateMessage{
		Data:     req.Data,
		ChatID:   req.ChatID,
		SenderID: req.SenderID,
	}

	if err := s.PMRepo.Insert(ctx, &pm); err != nil {
		return services.NewAPIMessage[any]("internal server error", nil, false, http.StatusInternalServerError, nil)
	}

	return services.NewAPIMessage[any]("message create", nil, true, http.StatusOK, nil)
}
