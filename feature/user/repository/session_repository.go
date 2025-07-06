package repository

import (
	"context"
	sqlc "readly/db/sqlc"
)

type SessionRepository interface {
	CreateSession(ctx context.Context, req CreateSessionRequest) error
	GetSessionByID(ctx context.Context, req GetSessionByIDRequest) (*GetSessionResponse, error)
	GetSessionByUserID(ctx context.Context, req GetSessionByUserIDRequest) ([]GetSessionResponse, error)
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

func (r *SessionRepositoryImpl) GetSessionByID(ctx context.Context, req GetSessionByIDRequest) (*GetSessionResponse, error) {
	res, err := r.querier.GetSessionByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return newGetSessionResponseFromSQLC(res), nil
}

func (r *SessionRepositoryImpl) GetSessionByUserID(ctx context.Context, req GetSessionByUserIDRequest) ([]GetSessionResponse, error) {
	sessions, err := r.querier.GetSessionByUserID(ctx, req.UserID)
	if err != nil {
		return nil, err
	}
	var res []GetSessionResponse
	for _, session := range sessions {
		res = append(res, *newGetSessionResponseFromSQLC(session))
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
