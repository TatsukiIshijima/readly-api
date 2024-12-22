package db

import (
	"context"
	db "readly/db/sqlc"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInsertUser(t *testing.T) {
	arg := db.InsertUserParams{
		Name:           "test1",
		Email:          "test1@example.com",
		HashedPassword: "hashed_password",
	}

	user, err := testQueries.InsertUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.NotZero(t, user.ID)
	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.UpdatedAt)
}
