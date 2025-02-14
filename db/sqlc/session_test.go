package db

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"readly/testdata"
	"testing"
	"time"
)

func createRandomSession(t *testing.T, user User) Session {
	id, err := uuid.NewRandom()
	require.NoError(t, err)
	refreshToken := testdata.RandomString(32)

	arg := CreateSessionParams{
		ID:           id,
		UserID:       user.ID,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().UTC(),
		IpAddress:    sql.NullString{String: "127.0.0.1", Valid: true},
		UserAgent:    sql.NullString{String: "Mozilla/5.0", Valid: true},
	}
	session, err := querier.CreateSession(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, session)

	require.Equal(t, arg.ID, session.ID)
	require.Equal(t, arg.UserID, session.UserID)
	require.Equal(t, arg.RefreshToken, session.RefreshToken)
	require.WithinDuration(t, arg.ExpiresAt, session.ExpiresAt, time.Second)
	require.Equal(t, arg.IpAddress, session.IpAddress)
	require.Equal(t, arg.UserAgent, session.UserAgent)

	return session
}

func TestCreateSession(t *testing.T) {
	user := createRandomUser(t)
	createRandomSession(t, user)
}

func TestDeleteSessionByUserID(t *testing.T) {
	user := createRandomUser(t)

	limit := 3

	for i := 0; i < limit; i++ {
		createRandomSession(t, user)
	}

	arg := DeleteSessionByUserIDParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}
	n, err := querier.DeleteSessionByUserID(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, int64(limit), n)
}

func TestGetSessionByID(t *testing.T) {
	user := createRandomUser(t)
	session1 := createRandomSession(t, user)

	session2, err := querier.GetSessionByID(context.Background(), session1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, session2)

	require.Equal(t, session1.ID, session2.ID)
	require.Equal(t, session1.UserID, session2.UserID)
	require.Equal(t, session1.RefreshToken, session2.RefreshToken)
	require.WithinDuration(t, session1.ExpiresAt, session2.ExpiresAt, time.Second)
	require.Equal(t, session1.IpAddress, session2.IpAddress)
	require.Equal(t, session1.UserAgent, session2.UserAgent)
}

func TestGetSessionByUserID(t *testing.T) {
	user1 := createRandomUser(t)
	user2 := createRandomUser(t)
	session1 := createRandomSession(t, user1)
	_ = createRandomSession(t, user2)

	sessions, err := querier.GetSessionByUserID(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, sessions)
	require.Len(t, sessions, 1)

	sessionByUser1 := sessions[0]

	require.Equal(t, session1.ID, sessionByUser1.ID)
	require.Equal(t, session1.UserID, sessionByUser1.UserID)
	require.Equal(t, session1.RefreshToken, sessionByUser1.RefreshToken)
	require.WithinDuration(t, session1.ExpiresAt, sessionByUser1.ExpiresAt, time.Second)
	require.Equal(t, session1.IpAddress, sessionByUser1.IpAddress)
	require.Equal(t, session1.UserAgent, sessionByUser1.UserAgent)
}
