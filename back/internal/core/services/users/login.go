package users

import (
	"app/internal/core/models/session"
	"app/internal/core/services"
	"app/internal/ports"
	"errors"
	"time"

	"context"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SignInReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	ClientIP string `json:"-"`
}

type SignInSvc struct {
	userRepo     ports.UserRepository
	sessionRepo  ports.SessionRepository
	hasher       ports.Hasher
	cookieSetter ports.CookieSetter
	header       ports.HeaderExtractor
}

func NewSignInSvc(
	userRepo ports.UserRepository,
	sessionRepo ports.SessionRepository,
	hasher ports.Hasher,
	cookie ports.CookieSetter,
	header ports.HeaderExtractor) *SignInSvc {
	return &SignInSvc{
		userRepo:     userRepo,
		sessionRepo:  sessionRepo,
		hasher:       hasher,
		cookieSetter: cookie,
		header:       header,
	}
}

func (s *SignInSvc) SignIn(ctx context.Context, req SignInReq) services.APIMessage[any] {
	u, err := s.userRepo.SelectByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return services.NewAPIMessage[any]("invalid credentials",
				[]string{"e-mail not registered or invalid password"},
				false, http.StatusBadRequest, nil)
		}

		return services.NewAPIMessage[any](
			"internal server error", nil,
			false, http.StatusInternalServerError, nil)
	}

	if !s.hasher.CompareHashAndPasswd(string(u.Password), req.Password) {
		return services.NewAPIMessage[any](
			"invalid credentials",
			[]string{"e-mail not registered or invalid password"},
			false, http.StatusBadRequest, nil)
	}

	sessions, err := s.sessionRepo.SelectByUserID(ctx, u.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return services.NewAPIMessage[any](
			"internal server error", nil,
			false, http.StatusInternalServerError, nil)
	}

	for _, sess := range sessions {
		t := time.Now().Add(time.Hour * 24)

		s.sessionRepo.UpdateExpiresAtByKey(ctx, sess.Key, t)

		s.cookieSetter.SetCookie(&http.Cookie{
			Name:     "_Session",
			Value:    sess.Key,
			Expires:  t,
			Path:     "/",
			HttpOnly: true,
		})

		return services.NewAPIMessage[any](
			"already logged in", nil, true, http.StatusAlreadyReported, nil,
		)
	}

	sess := session.Session{
		UserId:    u.ID,
		Key:       uuid.NewString(),
		UserAgent: s.header.Get("User-Agent"),
		ClientIp:  req.ClientIP,
		ExpiresAt: time.Now().Add(time.Hour * 24),
	}

	if len(sess.UserAgent) == 0 {
		return services.NewAPIMessage[any](
			"missing User-Agent header", nil,
			false, http.StatusBadRequest, nil,
		)
	}

	if err := s.sessionRepo.Insert(ctx, &sess); err != nil {
		return services.NewAPIMessage[any](
			"internal server error", nil,
			false, http.StatusInternalServerError, nil)
	}

	s.cookieSetter.SetCookie(&http.Cookie{
		Name:     "_Session",
		Value:    sess.Key,
		Expires:  sess.ExpiresAt,
		Path:     "/",
		HttpOnly: true,
	})

	return services.NewAPIMessage[any]("logged in", nil, true, http.StatusOK, nil)
}
