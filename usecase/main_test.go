package usecase

import (
	"log"
	"math/rand"
	"os"
	"path/filepath"
	sqlc "readly/db/sqlc"
	"readly/env"
	"readly/repository"
	"readly/service/auth"
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

var config env.Config
var querier sqlc.Querier
var tx repository.Transactor
var maker auth.TokenMaker

func setupMain() {
	c, err := env.Load(filepath.Join(env.ProjectRoot(), "/env"))
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}
	config = c

	a := &sqlc.Adapter{}
	db, q := a.Connect(c.DBDriver, c.DBSource)
	querier = q

	tx = repository.New(db)

	maker, err = auth.NewPasetoMaker(c.TokenSymmetricKey)
	if err != nil {
		log.Fatalf("cannot create token maker: %v", err)
	}
}

func newTestSignInUseCase(t *testing.T) SignInUseCase {
	userRepo := repository.NewUserRepository(querier)
	sessionRepo := repository.NewSessionRepository(querier)
	return NewSignInUseCase(config, maker, tx, sessionRepo, userRepo)
}

func newTestSignUpUseCase(t *testing.T) SignUpUseCase {
	userRepo := repository.NewUserRepository(querier)
	sessionRepo := repository.NewSessionRepository(querier)
	return NewSignUpUseCase(config, maker, tx, sessionRepo, userRepo)
}

func newTestRegisterBookUseCase(t *testing.T) RegisterBookUseCase {
	userRepo := repository.NewUserRepository(querier)
	bookRepo := repository.NewBookRepository(querier)
	readingHistoryRepo := repository.NewReadingHistoryRepository(querier)
	return NewRegisterBookUseCase(tx, bookRepo, readingHistoryRepo, userRepo)
}

func newTestDeleteBookUseCase(t *testing.T) DeleteBookUseCase {
	userRepo := repository.NewUserRepository(querier)
	bookRepo := repository.NewBookRepository(querier)
	readingHistoryRepo := repository.NewReadingHistoryRepository(querier)
	return NewDeleteBookUseCase(tx, bookRepo, readingHistoryRepo, userRepo)
}

func newTestGetBookUseCase(t *testing.T) GetBookUseCase {
	bookRepo := repository.NewBookRepository(querier)
	readingHistoryRepo := repository.NewReadingHistoryRepository(querier)
	return NewGetBookUseCase(bookRepo, readingHistoryRepo)
}

func newTestRefreshAccessTokenUseCase(t *testing.T) RefreshAccessTokenUseCase {
	sessionRepo := repository.NewSessionRepository(querier)
	return NewRefreshAccessTokenUseCase(config, maker, sessionRepo)
}
