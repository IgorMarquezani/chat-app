package users

import (
	"app/internal/core/services"
	"app/internal/ports"
	"context"
)

type SignInReq struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type SignInSvc struct {
	repo ports.UserRepository
}

func (s *SignInSvc) SignIn(ctx context.Context, account SignInReq) services.APIMessage {
	var (
		msg = services.APIMessage{}
	)

	return msg
}
