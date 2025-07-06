package usecase

import (
	"context"
	"errors"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"readly/configs"
	"readly/db/transaction"
	userRepo "readly/feature/user/repository"
	"readly/middleware/auth"
	"readly/repository"
)

type SignUpUseCase interface {
	SignUp(ctx context.Context, req SignUpRequest) (*SignUpResponse, error)
}

type SignUpUseCaseImpl struct {
	config      configs.Config
	maker       auth.TokenMaker
	transactor  transaction.Transactor
	sessionRepo repository.SessionRepository
	userRepo    userRepo.UserRepository
}

func NewSignUpUseCase(
	config configs.Config,
	maker auth.TokenMaker,
	transactor transaction.Transactor,
	sessionRepo repository.SessionRepository,
	userRepo userRepo.UserRepository,
) SignUpUseCase {
	return &SignUpUseCaseImpl{
		config:      config,
		maker:       maker,
		transactor:  transactor,
		sessionRepo: sessionRepo,
		userRepo:    userRepo,
	}
}

func (u *SignUpUseCaseImpl) SignUp(ctx context.Context, req SignUpRequest) (*SignUpResponse, error) {
	var res *SignUpResponse
	err := u.transactor.Exec(ctx, func() error {
		hashedPassword, err := generateHashedPassword(req.Password)
		if err != nil {
			return err
		}

		user, err := u.userRepo.CreateUser(ctx, userRepo.CreateUserRequest{
			Name:     req.Name,
			Email:    req.Email,
			Password: hashedPassword,
		})
		if err != nil {
			var pqErr *pq.Error
			if errors.As(err, &pqErr) {
				if pqErr.Code == "23505" {
					return newError(Conflict, EmailAlreadyRegisteredError, "email already exists")
				}
			}
			return err
		}

		accessTokenPayload, err := u.maker.Generate(user.ID, u.config.AccessTokenDuration)
		if err != nil {
			return err
		}

		refreshTokenPayload, err := u.maker.Generate(user.ID, u.config.RefreshTokenDuration)
		if err != nil {
			return err
		}

		sessionReq := repository.CreateSessionRequest{
			ID:           refreshTokenPayload.ID,
			UserID:       user.ID,
			RefreshToken: refreshTokenPayload.Token,
			ExpiresAt:    refreshTokenPayload.ExpiredAt,
			IPAddress:    req.IPAddress,
			UserAgent:    req.UserAgent,
		}
		err = u.sessionRepo.CreateSession(ctx, sessionReq)
		if err != nil {
			return err
		}

		res = NewSignUpResponse(accessTokenPayload.Token, refreshTokenPayload.Token, user)
		return nil
	})
	return res, handle(err)
}

func generateHashedPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
