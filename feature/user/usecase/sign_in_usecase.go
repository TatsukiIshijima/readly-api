package usecase

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"readly/configs"
	"readly/db/transaction"
	sessionRepo "readly/feature/user/repository"
	userRepo "readly/feature/user/repository"
	"readly/middleware/auth"
)

const maxSaveToken = 5

type SignInUseCase interface {
	SignIn(ctx context.Context, req SignInRequest) (*SignInResponse, error)
}

type SignInUseCaseImpl struct {
	config      configs.Config
	maker       auth.TokenMaker
	transactor  transaction.Transactor
	sessionRepo sessionRepo.SessionRepository
	userRepo    userRepo.UserRepository
}

func NewSignInUseCase(
	config configs.Config,
	maker auth.TokenMaker,
	transactor transaction.Transactor,
	sessionRepo sessionRepo.SessionRepository,
	userRepo userRepo.UserRepository,
) SignInUseCase {
	return &SignInUseCaseImpl{
		config:      config,
		maker:       maker,
		transactor:  transactor,
		sessionRepo: sessionRepo,
		userRepo:    userRepo,
	}
}

func (u *SignInUseCaseImpl) SignIn(ctx context.Context, req SignInRequest) (*SignInResponse, error) {
	var res *SignInResponse
	err := u.transactor.Exec(ctx, func() error {
		user, err := u.userRepo.GetUserByEmail(ctx, userRepo.NewGetUserByEmailRequest(req.Email))
		if err != nil {
			return newError(BadRequest, NotFoundUserError, "user not found")
		}

		err = u.checkPasswordHash(req.Password, user.Password)
		if err != nil {
			return newError(BadRequest, InvalidPasswordError, "invalid password")
		}

		accessTokenPayload, err := u.maker.Generate(user.ID, u.config.AccessTokenDuration)
		if err != nil {
			return err
		}

		refreshTokenPayload, err := u.maker.Generate(user.ID, u.config.RefreshTokenDuration)
		if err != nil {
			return err
		}

		err = u.cleanSessions(ctx, user.ID)
		if err != nil {
			return err
		}

		sessionReq := sessionRepo.CreateSessionRequest{
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
		res = NewSignInResponse(accessTokenPayload.Token, refreshTokenPayload.Token, user)
		return nil
	})
	return res, handle(err)
}

func (u *SignInUseCaseImpl) checkPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func (u *SignInUseCaseImpl) cleanSessions(ctx context.Context, userID int64) error {
	getReq := sessionRepo.GetSessionByUserIDRequest{
		UserID: userID,
	}
	sessions, err := u.sessionRepo.GetSessionByUserID(ctx, getReq)
	if err != nil {
		return err
	}
	sessionsCount := len(sessions)
	if sessionsCount < maxSaveToken {
		return nil
	}
	sessionToDeleteLimit := sessionsCount - maxSaveToken + 1
	deleteReq := sessionRepo.DeleteSessionByUserIDRequest{
		UserID: userID,
		Limit:  int32(sessionToDeleteLimit),
	}
	_, err = u.sessionRepo.DeleteSessionByUserID(ctx, deleteReq)
	return err
}
