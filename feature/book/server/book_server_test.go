//go:build test

package server

import (
	"github.com/stretchr/testify/require"
	"readly/configs"
	sqlc "readly/db/sqlc"
	"readly/db/transaction"
	"readly/feature/book/repository"
	"readly/feature/book/usecase"
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
	transactor := transaction.New(db)

	bookRepo := repository.NewBookRepository(q)
	readingHistoryRepo := repository.NewReadingHistoryRepository(q)

	maker, err := auth.NewPasetoMaker(config.TokenSymmetricKey)
	require.NoError(t, err)

	registerBookUseCase := usecase.NewRegisterBookUseCase(transactor, bookRepo, readingHistoryRepo)
	deleteBookUseCase := usecase.NewDeleteBookUseCase(transactor, bookRepo, readingHistoryRepo)
	getBookUseCase := usecase.NewGetBookUseCase(bookRepo, readingHistoryRepo)
	getBookListUseCase := usecase.NewGetBookListUseCase(readingHistoryRepo)

	return NewBookServer(
		maker,
		registerBookUseCase,
		deleteBookUseCase,
		getBookUseCase,
		getBookListUseCase,
	)
}
