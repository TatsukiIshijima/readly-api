package usecase

import (
	"context"
	"github.com/google/uuid"
	"readly/env"
	"readly/repository"
	"readly/service/auth"
	"time"
)

type RefreshAccessTokenUseCase interface {
	Refresh(ctx context.Context, req RefreshAccessTokenRequest) (*RefreshAccessTokenResponse, error)
}

type RefreshAccessTokenUseCaseImpl struct {
	config      env.Config
	marker      auth.TokenMaker
	sessionRepo repository.SessionRepository
}

func NewRefreshAccessTokenUseCase(
	config env.Config,
	maker auth.TokenMaker,
	sessionRepo repository.SessionRepository,
) RefreshAccessTokenUseCase {
	return &RefreshAccessTokenUseCaseImpl{
		config:      config,
		marker:      maker,
		sessionRepo: sessionRepo,
	}
}

type RefreshAccessTokenRequest struct {
	RefreshToken string
}

type RefreshAccessTokenResponse struct {
	AccessToken string
}

func (u *RefreshAccessTokenUseCaseImpl) Refresh(ctx context.Context, req RefreshAccessTokenRequest) (res *RefreshAccessTokenResponse, err error) {
	defer func() {
		if err != nil {
			err = handle(err)
		}
	}()

	payload, err := u.marker.Verify(req.RefreshToken)
	if err != nil {
		return nil, err
	}
	id, err := uuid.Parse(payload.ID)
	if err != nil {
		return nil, err
	}

	getSessionReq := repository.GetSessionByIDRequest{
		ID: id,
	}
	session, err := u.sessionRepo.GetSessionByID(ctx, getSessionReq)
	if err != nil {
		return nil, err
	}
	if session.IsRevoked {
		err := newError(UnAuthorized, InvalidTokenError, "refresh token is revoked")
		return nil, err
	}
	if session.UserID != payload.UserID {
		err := newError(UnAuthorized, InvalidTokenError, "incorrect user")
		return nil, err
	}
	if session.RefreshToken != req.RefreshToken {
		err := newError(UnAuthorized, InvalidTokenError, "mismatched refresh token")
		return nil, err
	}
	if time.Now().After(session.ExpiredAt) {
		err := newError(UnAuthorized, InvalidTokenError, "refresh token is expired")
		return nil, err
	}

	accessTokenPayload, err := u.marker.Generate(payload.UserID, u.config.AccessTokenDuration)
	if err != nil {
		return nil, err
	}

	return &RefreshAccessTokenResponse{
		AccessToken: accessTokenPayload.Token,
	}, nil
}
