package service

import (
	"github.com/gin-gonic/gin"
	"os"
	sqlc "readly/db/sqlc"
	"readly/repository"
	"readly/usecase"
	"testing"
)

var server *Server

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	fa := sqlc.FakeAdapter{}
	db, q := fa.Connect("", "")
	t := repository.New(db)
	bookRepo := repository.NewBookRepository(q)
	userRepo := repository.NewUserRepository(q)
	readingHistoryRepo := repository.NewReadingHistoryRepository(q)
	registerBookUseCase := usecase.NewRegisterBookUseCase(t, bookRepo, readingHistoryRepo, userRepo)
	deleteBookUseCase := usecase.NewDeleteBookUseCase(t, bookRepo, readingHistoryRepo, userRepo)
	bookService := NewBookService(registerBookUseCase, deleteBookUseCase)
	server = NewServer(bookService)

	os.Exit(m.Run())
}
