package auth

import (
	"context"
	"readly/repository"
)

const savedTokenMaxLimit = 5

type TokenCleaner struct {
	sessionRepo repository.SessionRepository
}

func NewTokenCleaner(sessionRepo repository.SessionRepository) *TokenCleaner {
	return &TokenCleaner{
		sessionRepo: sessionRepo,
	}
}

func (tc *TokenCleaner) Clean(ctx context.Context, userID int64) (int64, error) {
	getReq := repository.GetSessionByUserIDRequest{
		UserID: userID,
	}
	sessions, err := tc.sessionRepo.GetSessionByUserID(ctx, getReq)
	if err != nil {
		return 0, err
	}
	sessionCount := len(sessions)
	if sessionCount <= savedTokenMaxLimit {
		return 0, nil
	}
	sessionToDeleteLimit := sessionCount - savedTokenMaxLimit + 1
	deleteReq := repository.DeleteSessionByUserIDRequest{
		UserID: userID,
		Limit:  int32(sessionToDeleteLimit),
	}
	return tc.sessionRepo.DeleteSessionByUserID(ctx, deleteReq)
}
