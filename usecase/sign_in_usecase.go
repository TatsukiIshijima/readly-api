package usecase

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"readly/env"
	"readly/repository"
	"readly/service/auth"
)

type SignInUseCase interface {
	SignIn(ctx context.Context, req SignInRequest) (*SignInResponse, error)
}

type SignInUseCaseImpl struct {
	config      env.Config
	maker       auth.TokenMaker
	sessionRepo repository.SessionRepository
	userRepo    repository.UserRepository
}

func NewSignInUseCase(
	config env.Config,
	maker auth.TokenMaker,
	sessionRepo repository.SessionRepository,
	userRepo repository.UserRepository,
) SignInUseCase {
	return &SignInUseCaseImpl{
		config:      config,
		maker:       maker,
		sessionRepo: sessionRepo,
		userRepo:    userRepo,
	}
}

type SignInRequest struct {
	Email     string
	Password  string
	IPAddress string
	UserAgent string
}

type SignInResponse struct {
	AccessToken  string
	RefreshToken string
	UserID       int64
	Name         string
	Email        string
}

func (u *SignInUseCaseImpl) SignIn(ctx context.Context, req SignInRequest) (res *SignInResponse, err error) {
	defer func() {
		if err != nil {
			err = handle(err)
		}
	}()

	user, err := u.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, newError(BadRequest, NotFoundUserError, "user not found")
	}

	err = checkPasswordHash(req.Password, user.Password)
	if err != nil {
		return nil, newError(BadRequest, InvalidPasswordError, "invalid password")
	}

	accessTokenPayload, err := u.maker.Generate(user.ID, u.config.AccessTokenDuration)
	if err != nil {
		return nil, err
	}

	refreshTokenPayload, err := u.maker.Generate(user.ID, u.config.RefreshTokenDuration)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return &SignInResponse{
		AccessToken:  accessTokenPayload.Token,
		RefreshToken: refreshTokenPayload.Token,
		UserID:       user.ID,
		Name:         user.Name,
		Email:        user.Email,
	}, nil
}

func checkPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
