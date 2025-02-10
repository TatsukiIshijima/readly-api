package usecase

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"readly/env"
	"readly/repository"
	"readly/service/auth"
)

type SignUpUseCase interface {
	SignUp(ctx context.Context, req SignUpRequest) (*SignUpResponse, error)
}

type SignUpUseCaseImpl struct {
	config      env.Config
	maker       auth.TokenMaker
	transactor  repository.Transactor
	sessionRepo repository.SessionRepository
	userRepo    repository.UserRepository
}

func NewSignUpUseCase(
	config env.Config,
	maker auth.TokenMaker,
	transactor repository.Transactor,
	sessionRepo repository.SessionRepository,
	userRepo repository.UserRepository,
) SignUpUseCase {
	return &SignUpUseCaseImpl{
		config:      config,
		maker:       maker,
		transactor:  transactor,
		sessionRepo: sessionRepo,
		userRepo:    userRepo,
	}
}

type SignUpRequest struct {
	Name      string
	Email     string
	Password  string
	IPAddress string
	UserAgent string
}

type SignUpResponse struct {
	AccessToken  string
	RefreshToken string
	UserID       int64
	Name         string
	Email        string
}

func (u *SignUpUseCaseImpl) SignUp(ctx context.Context, req SignUpRequest) (*SignUpResponse, error) {
	var res *SignUpResponse
	err := u.transactor.Exec(ctx, func() error {
		hashedPassword, err := generateHashedPassword(req.Password)
		if err != nil {
			return err
		}

		user, err := u.userRepo.CreateUser(ctx, repository.CreateUserRequest{
			Name:     req.Name,
			Email:    req.Email,
			Password: hashedPassword,
		})
		if err != nil {
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

		res = &SignUpResponse{
			AccessToken:  accessTokenPayload.Token,
			RefreshToken: refreshTokenPayload.Token,
			UserID:       user.ID,
			Name:         user.Name,
			Email:        user.Email,
		}
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
