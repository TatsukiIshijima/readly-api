//go:build test

package server

import (
	"github.com/stretchr/testify/require"
	"readly/configs"
	sqlc "readly/db/sqlc"
	"readly/middleware/auth"
	"readly/repository"
	"readly/testdata"
	"readly/usecase"
	userRepo "readly/user/repository"
	"testing"
	"time"
)

func NewTestBookServer(t *testing.T) *BookServerImpl {
	config := configs.Config{
		TokenSymmetricKey:    testdata.RandomString(32),
		AccessTokenDuration:  time.Minute,
		RefreshTokenDuration: time.Hour,
	}
	fa := sqlc.FakeAdapter{}
	db, q := fa.Connect("", "")
	transaction := repository.New(db)

	userRepository := userRepo.NewUserRepository(q)
	bookRepo := repository.NewBookRepository(q)
	readingHistoryRepo := repository.NewReadingHistoryRepository(q)

	maker, err := auth.NewPasetoMaker(config.TokenSymmetricKey)
	require.NoError(t, err)

	registerBookUseCase := usecase.NewRegisterBookUseCase(transaction, bookRepo, readingHistoryRepo, userRepository)
	deleteBookUseCase := usecase.NewDeleteBookUseCase(transaction, bookRepo, readingHistoryRepo, userRepository)

	return NewBookServer(
		maker,
		registerBookUseCase,
		deleteBookUseCase,
	)
}
