package repository

import (
	"context"
	sqlc "readly/db/sqlc"
	"time"
)

type SessionRepository interface {
	CreateSession(ctx context.Context, req CreateSessionRequest) error
}

type SessionRepositoryImpl struct {
	querier sqlc.Querier
}

func NewSessionRepository(q sqlc.Querier) SessionRepository {
	return &SessionRepositoryImpl{
		querier: q,
	}
}

type CreateSessionRequest struct {
	UserID       int64
	RefreshToken string
	ExpiresAt    time.Time
}

func (r *SessionRepositoryImpl) CreateSession(ctx context.Context, req CreateSessionRequest) error {
	args := sqlc.CreateSessionParams{
		UserID:       req.UserID,
		RefreshToken: req.RefreshToken,
		ExpiresAt:    req.ExpiresAt,
	}
	_, err := r.querier.CreateSession(ctx, args)
	if err != nil {
		return err
	}
	return nil
}
