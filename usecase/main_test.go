package usecase

import (
	"github.com/stretchr/testify/require"
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
	os.Exit(m.Run())
}

func setupMain(t *testing.T) (env.Config, sqlc.Querier, repository.Transactor, auth.TokenMaker) {
	config, err := env.Load(filepath.Join(env.ProjectRoot(), "/env"))
	require.NoError(t, err)

	a := &sqlc.Adapter{}
	db, q := a.Connect(config.DBDriver, config.DBSource)

	tx := repository.New(db)

	maker, err := auth.NewPasetoMaker(config.TokenSymmetricKey)
	require.NoError(t, err)

	return config, q, tx, maker
}

func newTestSignInUseCase(t *testing.T) SignInUseCase {
	config, querier, _, maker := setupMain(t)
	userRepo := repository.NewUserRepository(querier)
	sessionRepo := repository.NewSessionRepository(querier)
	return NewSignInUseCase(config, maker, sessionRepo, userRepo)
}

func newTestSignUpUseCase(t *testing.T) SignUpUseCase {
	config, querier, tx, maker := setupMain(t)
	userRepo := repository.NewUserRepository(querier)
	sessionRepo := repository.NewSessionRepository(querier)
	return NewSignUpUseCase(config, maker, tx, sessionRepo, userRepo)
}

func newTestRegisterBookUseCase(t *testing.T) RegisterBookUseCase {
	_, querier, tx, _ := setupMain(t)
	userRepo := repository.NewUserRepository(querier)
	bookRepo := repository.NewBookRepository(querier)
	readingHistoryRepo := repository.NewReadingHistoryRepository(querier)
	return NewRegisterBookUseCase(tx, bookRepo, readingHistoryRepo, userRepo)
}

func newTestDeleteBookUseCase(t *testing.T) DeleteBookUseCase {
	_, querier, tx, _ := setupMain(t)
	userRepo := repository.NewUserRepository(querier)
	bookRepo := repository.NewBookRepository(querier)
	readingHistoryRepo := repository.NewReadingHistoryRepository(querier)
	return NewDeleteBookUseCase(tx, bookRepo, readingHistoryRepo, userRepo)
}
