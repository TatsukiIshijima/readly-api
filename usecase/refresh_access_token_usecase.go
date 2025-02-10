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

func (u *RefreshAccessTokenUseCaseImpl) Refresh(ctx context.Context, req RefreshAccessTokenRequest) (*RefreshAccessTokenResponse, error) {
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
		// FIXME:エラー定義&handleで返す
		err := newError("refresh token is revoked", UnAuthorized)
		return nil, err
	}
	if session.UserID != payload.UserID {
		// FIXME:エラー定義&handleで返す
		err := newError("incorrect user", UnAuthorized)
		return nil, err
	}
	if session.RefreshToken != req.RefreshToken {
		// FIXME:エラー定義&handleで返す
		err := newError("mismatched refresh token", UnAuthorized)
		return nil, err
	}
	if time.Now().After(session.ExpiredAt) {
		// FIXME:エラー定義&handleで返す
		err := newError("refresh token is expired", UnAuthorized)
		return nil, err
	}

	accessTokenPayload, err := u.marker.Generate(payload.UserID, u.config.AccessTokenDuration)
	if err != nil {
		return nil, err
	}

	res := &RefreshAccessTokenResponse{
		AccessToken: accessTokenPayload.Token,
	}
	return res, nil
}
