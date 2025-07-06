//go:build test

package db

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"sort"
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
	// Find all sessions for the user
	var userSessions []Session
	for _, s := range sessionTable.Columns {
		if s.UserID == arg.UserID {
			userSessions = append(userSessions, s)
		}
	}

	// If no sessions found, return 0
	if len(userSessions) == 0 {
		return 0, nil
	}

	// Sort sessions by creation time (ascending)
	sort.Slice(userSessions, func(i, j int) bool {
		return userSessions[i].CreatedAt.Before(userSessions[j].CreatedAt)
	})

	// Determine how many sessions to delete (limited by arg.Limit)
	deleteCount := int(arg.Limit)
	if deleteCount > len(userSessions) {
		deleteCount = len(userSessions)
	}

	// Get the IDs of sessions to delete
	sessionsToDelete := make(map[uuid.UUID]bool)
	for i := 0; i < deleteCount; i++ {
		sessionsToDelete[userSessions[i].ID] = true
	}

	// Create a new slice without the deleted sessions
	var newColumns []Session
	for _, s := range sessionTable.Columns {
		if !sessionsToDelete[s.ID] {
			newColumns = append(newColumns, s)
		}
	}

	// Calculate how many were actually deleted
	deletedCount := int64(len(sessionTable.Columns) - len(newColumns))

	// Update the table
	sessionTable.Columns = newColumns

	return deletedCount, nil
}
