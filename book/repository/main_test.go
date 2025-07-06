//go:build test

package repository

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

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func TestMain(m *testing.M) {
	setupMain()
	os.Exit(m.Run())
}

var bookRepo BookRepository
var readingHistoryRepo ReadingHistoryRepository
var querier sqlc.Querier

func setupMain() {
	fa := sqlc.FakeAdapter{}
	_, q := fa.Connect("", "")
	querier = q
	bookRepo = NewBookRepository(q)
	readingHistoryRepo = NewReadingHistoryRepository(q)
}

func createRandomUser(t *testing.T) sqlc.User {
	password := testdata.RandomString(16)
	hashedPassword, err := testdata.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	arg := sqlc.CreateUserParams{
		Name:           testdata.RandomString(12),
		Email:          testdata.RandomEmail(),
		HashedPassword: hashedPassword,
	}
	user, err := querier.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	return user
}

func createRandomBook(t *testing.T) sqlc.Book {
	arg := sqlc.CreateBookParams{
		Title: testdata.RandomString(10),
	}
	book, err := querier.CreateBook(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, book)
	return book
}
