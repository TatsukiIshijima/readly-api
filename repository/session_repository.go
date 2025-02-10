package repository

import (
	"context"
	"database/sql"
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
	IPAddress    string
	UserAgent    string
}

func (r *SessionRepositoryImpl) CreateSession(ctx context.Context, req CreateSessionRequest) error {
	ipAddress := sql.NullString{String: "", Valid: false}
	if req.IPAddress != "" {
		ipAddress = sql.NullString{String: req.IPAddress, Valid: true}
	}
	userAgent := sql.NullString{String: "", Valid: false}
	if req.UserAgent != "" {
		userAgent = sql.NullString{String: req.UserAgent, Valid: true}
	}

	args := sqlc.CreateSessionParams{
		UserID:       req.UserID,
		RefreshToken: req.RefreshToken,
		ExpiresAt:    req.ExpiresAt,
		IpAddress:    ipAddress,
		UserAgent:    userAgent,
	}
	_, err := r.querier.CreateSession(ctx, args)
	if err != nil {
		return err
	}
	return nil
}
