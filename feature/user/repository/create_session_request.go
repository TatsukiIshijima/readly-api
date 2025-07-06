package repository

import (
	"database/sql"
	"github.com/google/uuid"
	sqlc "readly/db/sqlc"
	"time"
)

type CreateSessionRequest struct {
	ID           uuid.UUID
	UserID       int64
	RefreshToken string
	ExpiresAt    time.Time
	IPAddress    string
	UserAgent    string
}

func (r CreateSessionRequest) toSQLC() sqlc.CreateSessionParams {
	address := sql.NullString{String: "", Valid: false}
	if r.IPAddress != "" {
		address = sql.NullString{String: r.IPAddress, Valid: true}
	}
	ua := sql.NullString{String: "", Valid: false}
	if r.UserAgent != "" {
		ua = sql.NullString{String: r.UserAgent, Valid: true}
	}
	return sqlc.CreateSessionParams{
		ID:           r.ID,
		UserID:       r.UserID,
		RefreshToken: r.RefreshToken,
		ExpiresAt:    r.ExpiresAt,
		IpAddress:    address,
		UserAgent:    ua,
	}
}
