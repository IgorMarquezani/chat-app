package ports

import (
	"app/internal/core/models/session"

	"context"
	"time"
)

type SessionRepository interface {
	Insert(context.Context, *session.Session) error
	SelectByUserID(context.Context, uint32) (session.Session, error)
	SelectByKey(context.Context, string) (session.Session, error)
	UpdateExpiresAtByKey(context.Context, string, time.Time) error
}
