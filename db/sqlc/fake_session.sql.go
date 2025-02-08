package db

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type SessionTable struct {
	Columns []Session
}

var sessionTable = SessionTable{}

func (q *FakeQuerier) CreateSession(_ context.Context, arg CreateSessionParams) (Session, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return Session{}, err
	}
	s := Session{
		ID:           id,
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
