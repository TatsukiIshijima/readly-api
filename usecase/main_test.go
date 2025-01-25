package usecase

import (
	"log"
	"math/rand"
	"os"
	"path/filepath"
	sqlc "readly/db/sqlc"
	"readly/env"
	"readly/repository"
	"testing"
	"time"
)

var userRepo repository.UserRepository
var registerBookUseCase RegisterBookUseCase
var deleteBookUseCase DeleteBookUseCase

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
	transactor := repository.New(db)
	bookRepo := repository.NewBookRepository(q)
	userRepo = repository.NewUserRepository(q)
	readingHistoryRepo := repository.NewReadingHistoryRepository(q)
	registerBookUseCase = NewRegisterBookUseCase(transactor, bookRepo, readingHistoryRepo, userRepo)
	deleteBookUseCase = NewDeleteBookUseCase(transactor, bookRepo, readingHistoryRepo, userRepo)
	os.Exit(m.Run())
}
