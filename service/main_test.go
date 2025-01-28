package service

import (
	"github.com/gin-gonic/gin"
	"os"
	sqlc "readly/db/sqlc"
	"readly/env"
	"readly/repository"
	"readly/testdata"
	"readly/usecase"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func NewTestServer() (*Server, error) {
	fa := sqlc.FakeAdapter{}
	db, q := fa.Connect("", "")
	t := repository.New(db)
	bookRepo := repository.NewBookRepository(q)
	userRepo := repository.NewUserRepository(q)
	readingHistoryRepo := repository.NewReadingHistoryRepository(q)
	registerBookUseCase := usecase.NewRegisterBookUseCase(t, bookRepo, readingHistoryRepo, userRepo)
	deleteBookUseCase := usecase.NewDeleteBookUseCase(t, bookRepo, readingHistoryRepo, userRepo)
	signUpUseCase := usecase.NewSignUpUseCase(userRepo)
	signInUseCase := usecase.NewSignInUseCase(userRepo)
	bookService := NewBookService(registerBookUseCase, deleteBookUseCase)
	userService := NewUserService(signUpUseCase, signInUseCase)

	config := env.Config{
		TokenSymmetricKey:   testdata.RandomString(32),
		AccessTokenDuration: time.Minute,
	}
	return NewServer(config, bookService, userService)
}
