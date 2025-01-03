package repository

import (
	"context"
	"math/rand"
	"os"
	sqlc "readly/db/sqlc"
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
	a := &sqlc.Adapter{}
	db, q := a.Connect()
	querier = q
	repo = NewBookRepository(db, q)
	os.Exit(m.Run())
}

func createRandomUser() (sqlc.User, error) {
	arg := sqlc.CreateUserParams{
		Name:           testdata.RandomString(12),
		Email:          testdata.RandomString(6) + "@example.com",
		HashedPassword: testdata.RandomString(16),
	}
	return querier.CreateUser(context.Background(), arg)
}
