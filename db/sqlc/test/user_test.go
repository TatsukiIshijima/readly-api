package db

import (
	"context"
	db "readly/db/sqlc"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInsertUser(t *testing.T) {
	arg := db.CreateUserParams{
		Name:           randomString(12),
		Email:          randomString(6) + "@example.com",
		HashedPassword: randomString(16),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.NotZero(t, user.ID)
	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.UpdatedAt)
}
