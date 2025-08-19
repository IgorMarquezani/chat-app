package users

import (
	"app/internal/core/models/user"
	userstate "app/internal/core/models/user-state"
	"app/internal/core/services"
	"app/internal/ports"
	"context"
	"errors"
	"log"
	"net/http"

	"gorm.io/gorm"
)

type NewAccountReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AccountSvc struct {
	userRepo      ports.UserRepository
	userStateRepo ports.UserStateRepository
	hasher        ports.Hasher
}

func NewAccountSvc(userRepo ports.UserRepository, userStateRepo ports.UserStateRepository, hasher ports.Hasher) (*AccountSvc, error) {
	if userRepo == nil || hasher == nil || userStateRepo == nil {
		return nil, errors.New("repositorys and hasher cannot be nil")
	}
	return &AccountSvc{
		userRepo:      userRepo,
		userStateRepo: userStateRepo,
		hasher:        hasher,
	}, nil
}

func (s *AccountSvc) NewAccount(ctx context.Context, account *NewAccountReq) services.APIMessage[any] {
	var u user.User

	name, err := user.NewUserName(account.Name)
	if err != nil {
		return services.APIMessage[any]{
			Error:   "invalid name",
			Details: []string{err.Error()},
			Status:  http.StatusBadRequest,
		}
	}
	u.Name = name

	email, err := user.NewEmail(account.Email)
	if err != nil {
		return services.APIMessage[any]{
			Error:   "invalid e-mail address",
			Details: []string{err.Error()},
			Status:  http.StatusBadRequest,
		}
	}
	u.Email = email

	passwd, errs := user.NewPassword(account.Password)
	if len(errs) > 0 {
		resp := services.APIMessage[any]{
			Error:   "invalid password",
			Details: make([]string, 0, len(errs)),
			Status:  http.StatusBadRequest,
		}
		for _, err := range errs {
			resp.Details = append(resp.Details, err.Error())
		}
		return resp
	}

	hash, err := s.hasher.Hash(string(passwd))
	if err != nil {
		return services.APIMessage[any]{
			Error:   "internal server error",
			Details: []string{"unexpected error ocurred"},
			Status:  http.StatusInternalServerError,
		}
	}
	u.Password = user.Password(hash)

	if err := s.userRepo.Insert(ctx, &u); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return services.APIMessage[any]{
				Error:  "e-mail already in use",
				Status: http.StatusAlreadyReported,
			}
		}

		return services.APIMessage[any]{
			Error:  "internal server error",
			Status: http.StatusInternalServerError,
		}
	}

	userState := userstate.UserState{
		UserID: u.ID,
	}

	if err := s.userStateRepo.Insert(ctx, &userState); err != nil {
		if err := s.userRepo.DeleteByID(ctx, u.ID); err != nil {
			log.Println(err)
		}

		return services.APIMessage[any]{
			Error:  "internal server error",
			Status: http.StatusInternalServerError,
		}
	}

	return services.APIMessage[any]{
		Error:   "ok",
		Succeed: true,
		Status:  http.StatusOK,
	}
}
