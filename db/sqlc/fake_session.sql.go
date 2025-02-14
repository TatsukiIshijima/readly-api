//go:build test

package db

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"time"
)

type SessionTable struct {
	Columns []Session
}

var sessionTable = SessionTable{}

func (q *FakeQuerier) CreateSession(_ context.Context, arg CreateSessionParams) (Session, error) {
	for _, s := range sessionTable.Columns {
		if s.ID == arg.ID {
			return Session{}, &pq.Error{Code: "23505", Message: "duplicate key value violates unique constraint"}
		}
	}
	s := Session{
		ID:           arg.ID,
		UserID:       arg.UserID,
		RefreshToken: arg.RefreshToken,
		ExpiresAt:    arg.ExpiresAt,
		CreatedAt:    time.Now().UTC(),
		IpAddress:    arg.IpAddress,
		UserAgent:    arg.UserAgent,
		Revoked:      false,
		RevokedAt:    sql.NullTime{Valid: false},
	}
	sessionTable.Columns = append(sessionTable.Columns, s)
	return s, nil
}

func (q *FakeQuerier) GetSessionByID(ctx context.Context, id uuid.UUID) (Session, error) {
	for _, s := range sessionTable.Columns {
		if s.ID == id {
			return s, nil
		}
	}
	return Session{}, sql.ErrNoRows
}

func (q *FakeQuerier) GetSessionByUserID(ctx context.Context, userID int64) ([]Session, error) {
	var sessions []Session
	for _, s := range sessionTable.Columns {
		if s.UserID == userID {
			sessions = append(sessions, s)
		}
	}
	return sessions, nil
}

func (q *FakeQuerier) DeleteSessionByUserID(ctx context.Context, arg DeleteSessionByUserIDParams) (int64, error) {
	var count int64
	for i, s := range sessionTable.Columns {
		if s.UserID == arg.UserID {
			sessionTable.Columns = append(sessionTable.Columns[:i], sessionTable.Columns[i+1:]...)
			count++
		}
	}
	return count, nil
}
