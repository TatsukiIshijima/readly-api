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

var store Store
var repo BookRepository

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func TestMain(m *testing.M) {
	a := &sqlc.DBAdapter{}
	db, _ := a.Connect()
	store = NewStore(db)
	repo = NewBookRepository(store)
	os.Exit(m.Run())
}

func createRandomUser() (sqlc.User, error) {
	arg := sqlc.CreateUserParams{
		Name:           testdata.RandomString(12),
		Email:          testdata.RandomString(6) + "@example.com",
		HashedPassword: testdata.RandomString(16),
	}
	return store.CreateUser(context.Background(), arg)
}
