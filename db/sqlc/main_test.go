package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"readly/env"
	"readly/testdata"
	"testing"
	"time"
)

var querier Querier

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func TestMain(m *testing.M) {
	config, err := env.Load(filepath.Join(env.ProjectRoot(), "/env"))
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	h := &Adapter{}
	_, q := h.Connect(config.DBDriver, config.DBSource)
	querier = q
	os.Exit(m.Run())
}

func createRandomUser(t *testing.T) User {
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
	return user
}

func EqualDate(t *testing.T, a, b time.Time) {
	require.Equal(t, a.Year(), b.Year())
	require.Equal(t, a.Month(), b.Month())
	require.Equal(t, a.Day(), b.Day())
}
