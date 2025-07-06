//go:build test

package usecase

import (
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"readly/configs"
	sqlc "readly/db/sqlc"
	"readly/db/transaction"
	userRepo "readly/feature/user/repository"
	"readly/middleware/auth"
	"readly/repository"
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
var userRepository userRepo.UserRepository
var sessionRepo repository.SessionRepository

func setupMain() {
	c, err := configs.Load(filepath.Join(configs.ProjectRoot(), "/configs/env"))
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}
	config = c

	fa := sqlc.FakeAdapter{}
	db, q := fa.Connect(c.DBDriver, c.DBSource)
	querier = q

	tx = transaction.New(db)

	maker, err = auth.NewPasetoMaker(c.TokenSymmetricKey)
	if err != nil {
		log.Fatalf("cannot create token maker: %v", err)
	}

	userRepository = userRepo.NewUserRepository(querier)
	sessionRepo = repository.NewSessionRepository(querier)
}

func newTestSignInUseCase(t *testing.T) SignInUseCase {
	return NewSignInUseCase(config, maker, tx, sessionRepo, userRepository)
}

func newTestSignUpUseCase(t *testing.T) SignUpUseCase {
	return NewSignUpUseCase(config, maker, tx, sessionRepo, userRepository)
}

func newTestRefreshAccessTokenUseCase(t *testing.T) RefreshAccessTokenUseCase {
	return NewRefreshAccessTokenUseCase(config, maker, sessionRepo)
}
