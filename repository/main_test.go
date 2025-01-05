package repository

import (
	"context"
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
var repo BookRepository

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func TestMain(m *testing.M) {
	config, err := env.Load(filepath.Join(env.ProjectRoot(), "/env"))
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	a := &sqlc.Adapter{}
	db, q := a.Connect(config.DBDriver, config.DBSource)
	querier = q
	repo = NewBookRepository(db, q)
	os.Exit(m.Run())
}

func createRandomUser() (sqlc.User, error) {
	arg := sqlc.CreateUserParams{
		Name:           testdata.RandomString(12),
		Email:          testdata.RandomEmail(),
		HashedPassword: testdata.RandomString(16),
	}
	return querier.CreateUser(context.Background(), arg)
}
