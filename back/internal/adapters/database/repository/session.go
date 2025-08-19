package repository

import (
	"app/internal/core/models/session"
	"app/internal/core/models/user"
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	ErrInvalidUUIDFormat = errors.New("invalid UUID format")
)

type SessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) (*SessionRepository, error) {
	if db == nil {
		return nil, errors.New("db cannot be nil")
	}

	return &SessionRepository{
		db: db,
	}, nil
}

func (s *SessionRepository) Insert(ctx context.Context, sess *session.Session) error {
	return s.db.WithContext(ctx).Create(sess).Error
}

func (s *SessionRepository) SelectByUserID(ctx context.Context, userID uint32) ([]session.Session, error) {
	sess := make([]session.Session, 0)

	return sess, s.db.WithContext(ctx).Model(&session.Session{}).Where("user_id = ?", userID).Scan(&sess).Error
}

func (s *SessionRepository) SelectByKey(ctx context.Context, key string) (session.Session, error) {
	var sess session.Session

	return sess, s.db.WithContext(ctx).Where("key = ?", key).First(&sess).Error
}

func (s *SessionRepository) UpdateExpiresAtByKey(ctx context.Context, key string, expiresAt time.Time) error {
	return s.db.WithContext(ctx).Model(&session.Session{}).Where("key = ?", key).Update("expires_at", expiresAt).Error
}

func (s *SessionRepository) SelectUserBySession(ctx context.Context, key string) (user.User, error) {
	var u user.User

	err := uuid.Validate(key)
	if err != nil {
		return u, ErrInvalidUUIDFormat
	}

	err = s.db.WithContext(ctx).
		Raw(`
        SELECT users.id, users.name, users.email
        FROM users
        JOIN sessions ON users.id = sessions.user_id
        WHERE sessions.key = ?
    `, key).
		Scan(&u).Error

	return u, err
}
