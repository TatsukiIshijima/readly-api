package repository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	sqlc "readly/db/sqlc"
	"time"
)

type SessionRepository interface {
	CreateSession(ctx context.Context, req CreateSessionRequest) error
	GetSessionByID(ctx context.Context, req GetSessionByIDRequest) (*SessionResponse, error)
	GetSessionByUserID(ctx context.Context, req GetSessionByUserIDRequest) ([]SessionResponse, error)
	DeleteSessionByUserID(ctx context.Context, req DeleteSessionByUserIDRequest) (int64, error)
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
	ID           uuid.UUID
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
		ID:           req.ID,
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

type GetSessionByIDRequest struct {
	ID uuid.UUID
}

type SessionResponse struct {
	UserID       int64
	RefreshToken string
	ExpiredAt    time.Time
	IsRevoked    bool
}

func (r *SessionRepositoryImpl) GetSessionByID(ctx context.Context, req GetSessionByIDRequest) (*SessionResponse, error) {
	session, err := r.querier.GetSessionByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return &SessionResponse{
		UserID:       session.UserID,
		RefreshToken: session.RefreshToken,
		ExpiredAt:    session.ExpiresAt,
		IsRevoked:    session.Revoked,
	}, nil
}

type GetSessionByUserIDRequest struct {
	UserID int64
}

func (r *SessionRepositoryImpl) GetSessionByUserID(ctx context.Context, req GetSessionByUserIDRequest) ([]SessionResponse, error) {
	sessions, err := r.querier.GetSessionByUserID(ctx, req.UserID)
	if err != nil {
		return nil, err
	}
	var res []SessionResponse
	for _, session := range sessions {
		res = append(res, SessionResponse{
			UserID:       session.UserID,
			RefreshToken: session.RefreshToken,
			ExpiredAt:    session.ExpiresAt,
			IsRevoked:    session.Revoked,
		})
	}
	return res, nil
}

type DeleteSessionByUserIDRequest struct {
	UserID int64
	Limit  int32
}

func (r *SessionRepositoryImpl) DeleteSessionByUserID(ctx context.Context, req DeleteSessionByUserIDRequest) (int64, error) {
	args := sqlc.DeleteSessionByUserIDParams{
		UserID: req.UserID,
		Limit:  req.Limit,
	}
	count, err := r.querier.DeleteSessionByUserID(ctx, args)
	if err != nil {
		return 0, err
	}
	return count, nil
}
