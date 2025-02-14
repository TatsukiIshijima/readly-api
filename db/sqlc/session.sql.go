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

const deleteSessionByUserID = `-- name: DeleteSessionByUserID :execrows
DELETE
FROM sessions
WHERE id IN (SELECT s.id
             FROM sessions AS s
             WHERE s.user_id = $1
             ORDER BY s.created_at ASC
    LIMIT $2
    )
`

type DeleteSessionByUserIDParams struct {
	UserID int64 `json:"user_id"`
	Limit  int32 `json:"limit"`
}

func (q *Queries) DeleteSessionByUserID(ctx context.Context, arg DeleteSessionByUserIDParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, deleteSessionByUserID, arg.UserID, arg.Limit)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
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

const getSessionByUserID = `-- name: GetSessionByUserID :many
SELECT id, user_id, refresh_token, expires_at, created_at, ip_address, user_agent, revoked, revoked_at
FROM sessions
WHERE user_id = $1
`

func (q *Queries) GetSessionByUserID(ctx context.Context, userID int64) ([]Session, error) {
	rows, err := q.db.QueryContext(ctx, getSessionByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Session{}
	for rows.Next() {
		var i Session
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.RefreshToken,
			&i.ExpiresAt,
			&i.CreatedAt,
			&i.IpAddress,
			&i.UserAgent,
			&i.Revoked,
			&i.RevokedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
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
