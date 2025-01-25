package repository

import (
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
	os.Exit(m.Run())
}
