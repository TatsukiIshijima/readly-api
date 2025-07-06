//go:build test

package repository

import (
	"context"
	"github.com/stretchr/testify/require"
	"readly/testdata"
	"testing"
)

func createRandomUser(t *testing.T) *CreateUserResponse {
	password := testdata.RandomString(16)
	hashedPassword, err := testdata.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	name := testdata.RandomString(10)
	email := testdata.RandomEmail()

	req := CreateUserRequest{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}
	u, err := userRepo.CreateUser(context.Background(), req)
	require.NoError(t, err)
	return u
}

func TestCreateUser(t *testing.T) {
	createUser := createRandomUser(t)
	getUser, err := userRepo.GetUserByID(context.Background(), GetUserByIDRequest{createUser.ID})
	require.NoError(t, err)

	require.Equal(t, createUser.ID, getUser.ID)
	require.Equal(t, createUser.Name, getUser.Name)
	require.Equal(t, createUser.Email, getUser.Email)
}
