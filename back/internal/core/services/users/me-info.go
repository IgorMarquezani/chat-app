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

type MeInfoSVC struct {
	userRepo    ports.UserRepository
	sessionRepo ports.SessionRepository
	extractor   ports.CookieExtractor
}

func NewMeInfoSVC(
	userRepo ports.UserRepository,
	sessionRepo ports.SessionRepository,
	extractor ports.CookieExtractor,
) *MeInfoSVC {
	return &MeInfoSVC{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
		extractor:   extractor,
	}
}

func (s *MeInfoSVC) MeInfo(ctx context.Context) services.APIMessage[user.FullUserInfo] {
	cookie, err := s.extractor.Cookie("_Session")
	if err != nil {
		return services.APIMessage[user.FullUserInfo]{
			Error:  "not loged in",
			Status: http.StatusUnauthorized,
		}
	}

	sess, err := s.sessionRepo.SelectByKey(ctx, cookie.Value)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return services.APIMessage[user.FullUserInfo]{
				Error:  "not loged in",
				Status: http.StatusUnauthorized,
			}
		}

		return services.APIMessage[user.FullUserInfo]{
			Error:  "internal server error",
			Status: http.StatusInternalServerError,
		}
	}

	u, err := s.userRepo.SelectJoinUserState(ctx, sess.UserId)
	if err != nil {
		return services.APIMessage[user.FullUserInfo]{
			Error:  "internal server error",
			Status: http.StatusInternalServerError,
		}
	}

	return services.APIMessage[user.FullUserInfo]{
		Error:   "ok",
		Succeed: true,
		Status:  http.StatusOK,
		Data:    u,
	}
}
