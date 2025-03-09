package util

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStringValidator_ValidateUsername(t *testing.T) {
	testCases := []struct {
		username string
		check    func(err error)
	}{
		{
			username: "User123",
			check: func(err error) {
				require.NoError(t, err)
			},
		},
		{
			username: "User@123",
			check: func(err error) {
				require.Error(t, err)
			},
		},
		{
			username: "Sh0t",
			check: func(err error) {
				require.Error(t, err)
			},
		},
		{
			username: "1234567890ABCDEabcde12345678901",
			check: func(err error) {
				require.Error(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("username=%s", tc.username), func(t *testing.T) {
			err := StringValidator(tc.username).ValidateUsername()
			tc.check(err)
		})
	}
}

func TestStringValidator_ValidateEmail(t *testing.T) {
	testCases := []struct {
		email string
		check func(err error)
	}{
		{
			email: "test@example.com",
			check: func(err error) {
				require.NoError(t, err)
			},
		},
		{
			email: "invalid-email",
			check: func(err error) {
				require.Error(t, err)
			},
		},
		{
			email: "a@b.c",
			check: func(err error) {
				require.Error(t, err)
			},
		},
		{
			email: "verylongemailaddress@example",
			check: func(err error) {
				require.Error(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("email=%s", tc.email), func(t *testing.T) {
			err := StringValidator(tc.email).ValidateEmail()
			tc.check(err)
		})
	}
}

func TestStringValidator_ValidatePassword(t *testing.T) {
	testCases := []struct {
		password string
		check    func(err error)
	}{
		{
			password: "Abcd1-^$*.@",
			check: func(err error) {
				require.NoError(t, err)
			},
		},
		{
			password: "abc123!",
			check: func(err error) {
				require.Error(t, err)
			},
		},
		{
			password: "ABC123!",
			check: func(err error) {
				require.Error(t, err)
			},
		},
		{
			password: "Abcdefgh",
			check: func(err error) {
				require.Error(t, err)
			},
		},
		{
			password: "Abc1-",
			check: func(err error) {
				require.Error(t, err)
			},
		},
		{
			password: "AbcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*-=+:;,./",
			check: func(err error) {
				require.Error(t, err)
			},
		},
		{
			password: "Aa1!Aa1!Aa1!",
			check: func(err error) {
				require.Error(t, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("password=%s", tc.password), func(t *testing.T) {
			err := StringValidator(tc.password).ValidatePassword()
			tc.check(err)
		})
	}
}
