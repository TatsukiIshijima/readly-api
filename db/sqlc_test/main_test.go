package sqlc_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	sqlc "readly/db/sqlc"
	"readly/env"
	"readly/testdata"
	"testing"
	"time"
)

var querier sqlc.Querier

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func TestMain(m *testing.M) {
	config, err := env.Load(filepath.Join(env.ProjectRoot(), "/env"))
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	h := &sqlc.Adapter{}
	_, q := h.Connect(config.DBDriver, config.DBSource)
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
