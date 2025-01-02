package sqlc_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"math/rand"
	"os"
	sqlc "readly/db/sqlc"
	"readly/test"
	"testing"
	"time"
)

var querier sqlc.Querier

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func TestMain(m *testing.M) {
	h := &test.DBAdapter{}
	_, q := h.Connect()
	querier = q
	os.Exit(m.Run())
}

func createRandomUser(t *testing.T) sqlc.User {
	arg := sqlc.CreateUserParams{
		Name:           test.RandomString(12),
		Email:          test.RandomString(6) + "@example.com",
		HashedPassword: test.RandomString(16),
	}
	user, err := querier.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	return user
}
