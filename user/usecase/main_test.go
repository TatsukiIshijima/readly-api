package usecase

import (
	"context"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"readly/configs"
	sqlc "readly/db/sqlc"
	"readly/middleware/auth"
	"readly/repository"
	"readly/testdata"
	userRepo "readly/user/repository"
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
var tx repository.Transactor
var maker auth.TokenMaker
var userRepository userRepo.UserRepository
var sessionRepo repository.SessionRepository

func setupMain() {
	c, err := configs.Load(filepath.Join(configs.ProjectRoot(), "/configs/env"))
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}
	config = c

	a := &sqlc.Adapter{}
	db, q := a.Connect(c.DBDriver, c.DBSource)
	querier = q

	tx = repository.New(db)

	maker, err = auth.NewPasetoMaker(c.TokenSymmetricKey)
	if err != nil {
		log.Fatalf("cannot create token maker: %v", err)
	}

	userRepository = userRepo.NewUserRepository(querier)
	sessionRepo = repository.NewSessionRepository(querier)

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

func newTestSignInUseCase(t *testing.T) SignInUseCase {
	return NewSignInUseCase(config, maker, tx, sessionRepo, userRepository)
}

func newTestSignUpUseCase(t *testing.T) SignUpUseCase {
	return NewSignUpUseCase(config, maker, tx, sessionRepo, userRepository)
}

func newTestRefreshAccessTokenUseCase(t *testing.T) RefreshAccessTokenUseCase {
	return NewRefreshAccessTokenUseCase(config, maker, sessionRepo)
}
