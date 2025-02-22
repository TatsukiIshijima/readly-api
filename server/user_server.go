package server

import (
	"context"
	"readly/env"
	"readly/pb"
	"readly/service/auth"
	"readly/usecase"
)

type UserServerImpl struct {
	pb.UnimplementedUserServer
	config              env.Config
	maker               auth.TokenMaker
	signUpUseCase       usecase.SignUpUseCase
	signInUseCase       usecase.SignInUseCase
	refreshTokenUseCase usecase.RefreshAccessTokenUseCase
}

func NewUserServer(
	config env.Config,
	maker auth.TokenMaker,
	signUpUseCase usecase.SignUpUseCase,
	signInUseCase usecase.SignInUseCase,
	refreshTokenUseCase usecase.RefreshAccessTokenUseCase,
) *UserServerImpl {
	return &UserServerImpl{
		config:              config,
		maker:               maker,
		signUpUseCase:       signUpUseCase,
		signInUseCase:       signInUseCase,
		refreshTokenUseCase: refreshTokenUseCase,
	}
}

func (s *UserServerImpl) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	// TODO: バリデーション

	meta := newMetadataFrom(ctx)
	args := usecase.SignInRequest{
		Email:     req.GetEmail(),
		Password:  req.GetPassword(),
		IPAddress: meta.IPAddress,
		UserAgent: meta.UserAgent,
	}
	result, err := s.signInUseCase.SignIn(ctx, args)
	if err != nil {
		return nil, gRPCStatusError(err)
	}
	return &pb.SignInResponse{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
		UserId:       result.UserID,
		Name:         result.Name,
		Email:        result.Email,
	}, nil
}

func (s *UserServerImpl) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	// TODO:メアドやパスワード等のバリデーション
	meta := newMetadataFrom(ctx)
	args := usecase.SignUpRequest{
		Name:      req.GetName(),
		Email:     req.GetEmail(),
		Password:  req.GetPassword(),
		IPAddress: meta.IPAddress,
		UserAgent: meta.UserAgent,
	}
	result, err := s.signUpUseCase.SignUp(ctx, args)
	if err != nil {
		return nil, gRPCStatusError(err)
	}
	return &pb.SignUpResponse{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
		UserId:       result.UserID,
		Name:         result.Name,
		Email:        result.Email,
	}, nil
}

func (s *UserServerImpl) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	args := usecase.RefreshAccessTokenRequest{
		RefreshToken: req.RefreshToken,
	}

	result, err := s.refreshTokenUseCase.Refresh(ctx, args)
	if err != nil {
		return nil, gRPCStatusError(err)
	}

	return &pb.RefreshTokenResponse{
		AccessToken: result.AccessToken,
	}, nil
}
