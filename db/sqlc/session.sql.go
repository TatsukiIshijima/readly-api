// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: session.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createSession = `-- name: CreateSession :one
INSERT INTO sessions (id,
                      user_id,
                      refresh_token,
                      expires_at,
                      ip_address,
                      user_agent)
VALUES ($1,
        $2,
        $3,
        $4,
        $5,
        $6) RETURNING id, user_id, refresh_token, expires_at, created_at, ip_address, user_agent, revoked, revoked_at
`

type CreateSessionParams struct {
	ID           uuid.UUID      `json:"id"`
	UserID       int64          `json:"user_id"`
	RefreshToken string         `json:"refresh_token"`
	ExpiresAt    time.Time      `json:"expires_at"`
	IpAddress    sql.NullString `json:"ip_address"`
	UserAgent    sql.NullString `json:"user_agent"`
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error) {
	row := q.db.QueryRowContext(ctx, createSession,
		arg.ID,
		arg.UserID,
		arg.RefreshToken,
		arg.ExpiresAt,
		arg.IpAddress,
		arg.UserAgent,
	)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.RefreshToken,
		&i.ExpiresAt,
		&i.CreatedAt,
		&i.IpAddress,
		&i.UserAgent,
		&i.Revoked,
		&i.RevokedAt,
	)
	return i, err
}

const getSessionByID = `-- name: GetSessionByID :one
SELECT id, user_id, refresh_token, expires_at, created_at, ip_address, user_agent, revoked, revoked_at
FROM sessions
WHERE id = $1
`

func (q *Queries) GetSessionByID(ctx context.Context, id uuid.UUID) (Session, error) {
	row := q.db.QueryRowContext(ctx, getSessionByID, id)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.RefreshToken,
		&i.ExpiresAt,
		&i.CreatedAt,
		&i.IpAddress,
		&i.UserAgent,
		&i.Revoked,
		&i.RevokedAt,
	)
	return i, err
}

const updateSession = `-- name: UpdateSession :one
UPDATE sessions
SET refresh_token = $2,
    expires_at    = $3,
    ip_address    = $4,
    user_agent    = $5,
    revoked       = $6,
    revoked_at    = $7
WHERE id = $1 RETURNING id, user_id, refresh_token, expires_at, created_at, ip_address, user_agent, revoked, revoked_at
`

type UpdateSessionParams struct {
	ID           uuid.UUID      `json:"id"`
	RefreshToken string         `json:"refresh_token"`
	ExpiresAt    time.Time      `json:"expires_at"`
	IpAddress    sql.NullString `json:"ip_address"`
	UserAgent    sql.NullString `json:"user_agent"`
	Revoked      bool           `json:"revoked"`
	RevokedAt    sql.NullTime   `json:"revoked_at"`
}

func (q *Queries) UpdateSession(ctx context.Context, arg UpdateSessionParams) (Session, error) {
	row := q.db.QueryRowContext(ctx, updateSession,
		arg.ID,
		arg.RefreshToken,
		arg.ExpiresAt,
		arg.IpAddress,
		arg.UserAgent,
		arg.Revoked,
		arg.RevokedAt,
	)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.RefreshToken,
		&i.ExpiresAt,
		&i.CreatedAt,
		&i.IpAddress,
		&i.UserAgent,
		&i.Revoked,
		&i.RevokedAt,
	)
	return i, err
}
