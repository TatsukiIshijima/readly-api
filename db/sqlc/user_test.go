package db

import (
	"context"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"readly/testdata"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func checkSameUser(t *testing.T, user1 User, user2 User) {
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Name, user2.Name)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	require.WithinDuration(t, user1.UpdatedAt, user2.UpdatedAt, time.Second)
}

func checkPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func TestCreateUser(t *testing.T) {
	password := testdata.RandomString(16)
	hashedPassword, err := testdata.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	arg := CreateUserParams{
		Name:           testdata.RandomString(12),
		Email:          testdata.RandomEmail(),
		HashedPassword: hashedPassword,
	}
	user, err := querier.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.NotZero(t, user.ID)
	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Email, user.Email)
	require.NoError(t, checkPassword(password, user.HashedPassword))
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.UpdatedAt)
}

func TestGetUserById(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := querier.GetUserByID(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	checkSameUser(t, user1, user2)
}

func TestGetUserByEmail(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := querier.GetUserByEmail(context.Background(), user1.Email)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	checkSameUser(t, user1, user2)
}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)

	arg := UpdateUserParams{
		ID:             user1.ID,
		Name:           testdata.RandomString(12),
		Email:          user1.Email,
		HashedPassword: user1.HashedPassword,
	}

	user2, err := querier.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, arg.Name, user2.Name)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	require.WithinDuration(t, user1.UpdatedAt, user2.UpdatedAt, time.Second)
}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)
	err := querier.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := querier.GetUserByID(context.Background(), user1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}

func TestGetAllUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	arg1 := GetAllUsersParams{
		Limit:  5,
		Offset: 0,
	}

	users1, err := querier.GetAllUsers(context.Background(), arg1)
	require.NoError(t, err)
	require.Len(t, users1, 5)

	for _, user := range users1 {
		require.NotEmpty(t, user)
	}

	arg2 := GetAllUsersParams{
		Limit:  5,
		Offset: 5,
	}

	users2, err := querier.GetAllUsers(context.Background(), arg2)
	require.NoError(t, err)
	require.Len(t, users2, 5)

	for _, user := range users2 {
		require.NotEmpty(t, user)
	}
}
