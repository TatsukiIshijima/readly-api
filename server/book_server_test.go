//go:build test

package server

import (
	"github.com/stretchr/testify/require"
	sqlc "readly/db/sqlc"
	"readly/env"
	"readly/repository"
	"readly/service/auth"
	"readly/testdata"
	"readly/usecase"
	"testing"
	"time"
)

func NewTestBookServer(t *testing.T) *BookServerImpl {
	config := env.Config{
		TokenSymmetricKey:    testdata.RandomString(32),
		AccessTokenDuration:  time.Minute,
		RefreshTokenDuration: time.Hour,
	}
	fa := sqlc.FakeAdapter{}
	db, q := fa.Connect("", "")
	transaction := repository.New(db)

	userRepo := repository.NewUserRepository(q)
	bookRepo := repository.NewBookRepository(q)
	readingHistoryRepo := repository.NewReadingHistoryRepository(q)

	maker, err := auth.NewPasetoMaker(config.TokenSymmetricKey)
	require.NoError(t, err)

	registerBookUseCase := usecase.NewRegisterBookUseCase(transaction, bookRepo, readingHistoryRepo, userRepo)
	deleteBookUseCase := usecase.NewDeleteBookUseCase(transaction, bookRepo, readingHistoryRepo, userRepo)

	return NewBookServer(
		maker,
		registerBookUseCase,
		deleteBookUseCase,
	)
}
