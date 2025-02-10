package usecase

import (
	"context"
	"fmt"
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
		// FIXME: エラー定義
		err := fmt.Errorf("refresh token is revoked")
		return nil, err
	}
	if session.UserID != payload.UserID {
		// FIXME: エラー定義
		err := fmt.Errorf("incorrect user")
		return nil, err
	}
	if session.RefreshToken != req.RefreshToken {
		// FIXME: エラー定義
		err := fmt.Errorf("mismatched refresh token")
		return nil, err
	}
	if time.Now().After(session.ExpiredAt) {
		// FIXME: エラー定義
		err := fmt.Errorf("refresh token is expired")
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
