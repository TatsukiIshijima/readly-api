package repository

import (
	"context"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	sqlc "readly/db/sqlc"
	"readly/env"
	"testing"
	"time"
)

var querier sqlc.Querier
var bookRepo BookRepository
var userRepo UserRepository
var readingHistoryRepo ReadingHistoryRepository
var genres = []string{"ミステリー", "ファンタジー", "SF", "自己啓発", "ビジネス", "科学"}

func GetGenres() []string {
	return genres
}

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func TestMain(m *testing.M) {
	config, err := env.Load(filepath.Join(env.ProjectRoot(), "/env"))
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	a := &sqlc.Adapter{}
	_, q := a.Connect(config.DBDriver, config.DBSource)
	querier = q
	bookRepo = NewBookRepository(q)
	userRepo = NewUserRepository(q)
	readingHistoryRepo = NewReadingHistoryRepository(q)

	createGenresIfNeed()

	os.Exit(m.Run())
}

func createGenresIfNeed() {
	genres := GetGenres()
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
