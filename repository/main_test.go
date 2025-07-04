package repository

import (
	"context"
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
