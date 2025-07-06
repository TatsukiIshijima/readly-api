package usecase

import (
	"context"
	"github.com/stretchr/testify/require"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"readly/configs"
	sqlc "readly/db/sqlc"
	"readly/db/transaction"
	"readly/feature/book/repository"
	"readly/middleware/auth"
	"readly/testdata"
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

var config configs.Config
var querier sqlc.Querier
var tx transaction.Transactor
var maker auth.TokenMaker

func setupMain() {
	c, err := configs.Load(filepath.Join(configs.ProjectRoot(), "/configs/env"))
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}
	config = c

	a := &sqlc.Adapter{}
	db, q := a.Connect(c.DBDriver, c.DBSource)
	querier = q

	tx = transaction.New(db)

	maker, err = auth.NewPasetoMaker(c.TokenSymmetricKey)
	if err != nil {
		log.Fatalf("cannot create token maker: %v", err)
	}

	createGenresIfNeed()
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

func newTestRegisterBookUseCase(t *testing.T) RegisterBookUseCase {
	bookRepo := repository.NewBookRepository(querier)
	readingHistoryRepo := repository.NewReadingHistoryRepository(querier)
	return NewRegisterBookUseCase(tx, bookRepo, readingHistoryRepo)
}

func newTestDeleteBookUseCase(t *testing.T) DeleteBookUseCase {
	bookRepo := repository.NewBookRepository(querier)
	readingHistoryRepo := repository.NewReadingHistoryRepository(querier)
	return NewDeleteBookUseCase(tx, bookRepo, readingHistoryRepo)
}

func newTestGetBookUseCase(t *testing.T) GetBookUseCase {
	bookRepo := repository.NewBookRepository(querier)
	readingHistoryRepo := repository.NewReadingHistoryRepository(querier)
	return NewGetBookUseCase(bookRepo, readingHistoryRepo)
}

func newTestGetBookListUseCase(t *testing.T) GetBookListUseCase {
	readingHistoryRepo := repository.NewReadingHistoryRepository(querier)
	return NewGetBookListUseCase(readingHistoryRepo)
}

func newTestUpdateBookUseCase(t *testing.T) UpdateBookUseCase {
	bookRepo := repository.NewBookRepository(querier)
	readingHistoryRepo := repository.NewReadingHistoryRepository(querier)
	return NewUpdateBookUseCase(tx, bookRepo, readingHistoryRepo)
}

func createRandomUser(t *testing.T) sqlc.User {
	password := testdata.RandomString(16)
	hashedPassword, err := testdata.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	arg := sqlc.CreateUserParams{
		Name:           testdata.RandomString(12),
		Email:          testdata.RandomEmail(),
		HashedPassword: hashedPassword,
	}
	user, err := querier.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	return user
}
