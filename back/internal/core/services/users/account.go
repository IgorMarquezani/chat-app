package users

import (
	"app/internal/core/models/user"
	"app/internal/core/services"
	"app/internal/ports"
	"context"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type NewAccountReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AccountSvc struct {
	repo   ports.UserRepository
	hasher ports.Hasher
}

func NewAccountSvc(repo ports.UserRepository, hasher ports.Hasher) (*AccountSvc, error) {
	if repo == nil || hasher == nil {
		return nil, errors.New("repo and hasher cannot be nil")
	}
	return &AccountSvc{

		repo:   repo,
		hasher: hasher,
	}, nil
}

func (s *AccountSvc) NewAccount(ctx context.Context, account *NewAccountReq) services.APIMessage {
	var (
		u    user.User
		resp = services.APIMessage{
			Error:   "none",
			Details: make([]string, 0, 1),
			Status:  http.StatusOK,
		}
	)

	name, err := user.NewUserName(account.Name)
	if err != nil {
		resp.Error = err.Error()
		resp.Status = http.StatusBadRequest
		return resp
	}
	u.Name = name

	email, err := user.NewEmail(account.Email)
	if err != nil {
		resp.Error = "invalid e-mail address"
		resp.Details = append(resp.Details, err.Error())
		resp.Status = http.StatusBadRequest
		return resp
	}
	u.Email = email

	passwd, errs := user.NewPassword(account.Password)
	if len(errs) > 0 {
		resp.Error = "invalid password"
		resp.Status = http.StatusBadRequest
		for _, err := range errs {
			resp.Details = append(resp.Details, err.Error())
		}
		return resp
	}

	hash, err := s.hasher.Hash(string(passwd))
	if err != nil {
		resp.Error = "internal server error"
		resp.Details = append(resp.Details, "unexpected error ocurred")
		resp.Status = http.StatusInternalServerError
		return resp
	}
	u.Password = user.Password(hash)

	if err := s.repo.Insert(ctx, &u); err != nil {
		resp.Error = "internal server error"
		resp.Status = http.StatusInternalServerError

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			resp.Error = "e-mail already in use"
			resp.Status = http.StatusAlreadyReported
		}

		return resp
	}

	resp.Succeed = true

	return resp
}
