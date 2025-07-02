//go:build test

package controller

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"os"
	"readly/configs"
	sqlc "readly/db/sqlc"
	"readly/repository"
	"readly/service/auth"
	"readly/testdata"
	"readly/usecase"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func setupControllers(t *testing.T) (BookController, UserController) {
	config := configs.Config{
		TokenSymmetricKey:    testdata.RandomString(32),
		AccessTokenDuration:  time.Minute,
		RefreshTokenDuration: time.Hour,
	}
	fa := sqlc.FakeAdapter{}
	db, q := fa.Connect("", "")
	transaction := repository.New(db)

	bookRepo := repository.NewBookRepository(q)
	userRepo := repository.NewUserRepository(q)
	readingHistoryRepo := repository.NewReadingHistoryRepository(q)
	sessionRepo := repository.NewSessionRepository(q)

	maker, err := auth.NewPasetoMaker(config.TokenSymmetricKey)
	require.NoError(t, err)

	registerBookUseCase := usecase.NewRegisterBookUseCase(transaction, bookRepo, readingHistoryRepo, userRepo)
	deleteBookUseCase := usecase.NewDeleteBookUseCase(transaction, bookRepo, readingHistoryRepo, userRepo)
	signUpUseCase := usecase.NewSignUpUseCase(config, maker, transaction, sessionRepo, userRepo)
	signInUseCase := usecase.NewSignInUseCase(config, maker, transaction, sessionRepo, userRepo)
	refreshTokenUseCase := usecase.NewRefreshAccessTokenUseCase(config, maker, sessionRepo)

	bookController := NewBookController(registerBookUseCase, deleteBookUseCase)
	userController := NewUserController(config, maker, signUpUseCase, signInUseCase, refreshTokenUseCase)

	return bookController, userController
}

func setupTestContext(method string, url string, body []byte) (*gin.Context, *httptest.ResponseRecorder) {
	r := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(r)

	req, _ := http.NewRequest(method, url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	return c, r
}
