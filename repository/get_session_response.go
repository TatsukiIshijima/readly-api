package repository

import (
	sqlc "readly/db/sqlc"
	"time"
)

type GetSessionResponse struct {
	UserID       int64
	RefreshToken string
	ExpiredAt    time.Time
	IsRevoked    bool
}

func newGetSessionResponseFromSQLC(s sqlc.Session) *GetSessionResponse {
	return &GetSessionResponse{
		UserID:       s.UserID,
		RefreshToken: s.RefreshToken,
		ExpiredAt:    s.ExpiresAt,
		IsRevoked:    s.Revoked,
	}
}
