package repository

import (
	"context"
	"github.com/google/uuid"
	sqlc "readly/db/sqlc"
	"time"
)

type SessionRepository interface {
	CreateSession(ctx context.Context, req CreateSessionRequest) error
	GetSessionByID(ctx context.Context, req GetSessionByIDRequest) (*SessionResponse, error)
	GetSessionByUserID(ctx context.Context, req GetSessionByUserIDRequest) ([]SessionResponse, error)
	DeleteSessionByUserID(ctx context.Context, req DeleteSessionByUserIDRequest) (*DeleteSessionByUserIDResponse, error)
}

type SessionRepositoryImpl struct {
	querier sqlc.Querier
}

func NewSessionRepository(q sqlc.Querier) SessionRepository {
	return &SessionRepositoryImpl{
		querier: q,
	}
}

func (r *SessionRepositoryImpl) CreateSession(ctx context.Context, req CreateSessionRequest) error {
	_, err := r.querier.CreateSession(ctx, req.toSQLC())
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

func (r *SessionRepositoryImpl) DeleteSessionByUserID(ctx context.Context, req DeleteSessionByUserIDRequest) (*DeleteSessionByUserIDResponse, error) {
	count, err := r.querier.DeleteSessionByUserID(ctx, req.toSQLC())
	if err != nil {
		return nil, err
	}
	return newDeleteSessionByUserIDResponse(count), nil
}
