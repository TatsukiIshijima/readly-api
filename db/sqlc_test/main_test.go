package sqlc_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"math/rand"
	"os"
	sqlc "readly/db/sqlc"
	"readly/testdata"
	"testing"
	"time"
)

var querier sqlc.Querier

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func TestMain(m *testing.M) {
	h := &sqlc.DBAdapter{}
	_, q := h.Connect()
	querier = q
	os.Exit(m.Run())
}

func createRandomUser(t *testing.T) sqlc.User {
	arg := sqlc.CreateUserParams{
		Name:           testdata.RandomString(12),
		Email:          testdata.RandomString(6) + "@example.com",
		HashedPassword: testdata.RandomString(16),
	}
	user, err := querier.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	return user
}
