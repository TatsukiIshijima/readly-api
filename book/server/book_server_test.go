//go:build test

package server

import (
	"github.com/stretchr/testify/require"
	"readly/book/repository"
	"readly/book/usecase"
	"readly/configs"
	sqlc "readly/db/sqlc"
	"readly/middleware/auth"
	"readly/testdata"
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

	bookRepo := repository.NewBookRepository(q)
	readingHistoryRepo := repository.NewReadingHistoryRepository(q)

	maker, err := auth.NewPasetoMaker(config.TokenSymmetricKey)
	require.NoError(t, err)

	registerBookUseCase := usecase.NewRegisterBookUseCase(transaction, bookRepo, readingHistoryRepo)
	deleteBookUseCase := usecase.NewDeleteBookUseCase(transaction, bookRepo, readingHistoryRepo)

	return NewBookServer(
		maker,
		registerBookUseCase,
		deleteBookUseCase,
	)
}
