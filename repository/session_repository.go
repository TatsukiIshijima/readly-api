package repository

import (
	"context"
	"github.com/google/uuid"
	sqlc "readly/db/sqlc"
	"time"
)

type SessionRepository interface {
	CreateSession(ctx context.Context, req CreateSessionRequest) (*CreateSessionResponse, error)
}

type SessionRepositoryImpl struct {
	querier sqlc.Querier
}

func NewSessionRepository(q sqlc.Querier) SessionRepositoryImpl {
	return SessionRepositoryImpl{
		querier: q,
	}
}

type CreateSessionRequest struct {
	ID           uuid.UUID
	UserID       int64
	RefreshToken string
	ExpiresAt    time.Time
}

type CreateSessionResponse struct {
	RefreshToken string
	ExpiresAt    time.Time
}

func (r *SessionRepositoryImpl) CreateSession(ctx context.Context, req CreateSessionRequest) (*CreateSessionResponse, error) {
	args := sqlc.CreateSessionParams{
		UserID:       req.UserID,
		RefreshToken: req.RefreshToken,
		ExpiresAt:    req.ExpiresAt,
	}
	res, err := r.querier.CreateSession(ctx, args)
	if err != nil {
		return nil, err
	}
	return &CreateSessionResponse{
		RefreshToken: res.RefreshToken,
		ExpiresAt:    res.ExpiresAt,
	}, nil
}
