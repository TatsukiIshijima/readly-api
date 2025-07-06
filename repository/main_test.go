package repository

import (
	"context"
	"github.com/stretchr/testify/require"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"readly/configs"
	sqlc "readly/db/sqlc"
	"readly/testdata"
	"testing"
	"time"
)

var querier sqlc.Querier
var bookRepo BookRepository
var readingHistoryRepo ReadingHistoryRepository

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func TestMain(m *testing.M) {
	config, err := configs.Load(filepath.Join(configs.ProjectRoot(), "/configs/env"))
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	a := &sqlc.Adapter{}
	_, q := a.Connect(config.DBDriver, config.DBSource)
	querier = q
	bookRepo = NewBookRepository(q)
	readingHistoryRepo = NewReadingHistoryRepository(q)

	createGenresIfNeed()

	os.Exit(m.Run())
}

func createGenresIfNeed() {
	genres := testdata.GetGenres()
	for _, genre := range genres {
		_, err := querier.GetGenreByName(context.Background(), genre)
		if err == nil {
			continue
		}
		_, err = querier.CreateGenre(context.Background(), genre)
		if err != nil {
			log.Fatalf("failed to create genre %s: %v", genre, err)
		}
	}
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
