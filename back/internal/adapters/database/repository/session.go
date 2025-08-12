package repository

import (
	"app/internal/core/models/session"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
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

func (s *SessionRepository) SelectByUserID(ctx context.Context, userID uint32) (session.Session, error) {
	var sess session.Session

	return sess, s.db.WithContext(ctx).Where("user_id = ?", userID).First(&sess).Error
}

func (s *SessionRepository) SelectByKey(ctx context.Context, key string) (session.Session, error) {
	var sess session.Session

	return sess, s.db.WithContext(ctx).Where("key = ?", key).First(&sess).Error
}

func (s *SessionRepository) UpdateExpiresAtByKey(ctx context.Context, key string, expiresAt time.Time) error {
	return s.db.WithContext(ctx).Model(&session.Session{}).Where("key = ?", key).Update("expires_at", expiresAt).Error
}
