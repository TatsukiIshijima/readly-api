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

// TODO:ここinterface返すではなくてOK？
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
	// TODO: ShouldBindJSONの代わりの実装
	args := usecase.SignInRequest{
		Email:    req.Email,
		Password: req.Password,
		// TODO:値の取得
		IPAddress: "",
		UserAgent: "",
	}
	result, err := s.signInUseCase.SignIn(ctx, args)
	if err != nil {
		// TODO:エラーをどう返すか
		return nil, err
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
	// TODO: ShouldBindJSONの代わりの実装
	args := usecase.SignUpRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		// TODO:値の取得
		IPAddress: "",
		UserAgent: "",
	}
	result, err := s.signUpUseCase.SignUp(ctx, args)
	if err != nil {
		// TODO:エラーをどう返すか
		return nil, err
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
	// TODO: ShouldBindJSONの代わりの実装
	args := usecase.RefreshAccessTokenRequest{
		RefreshToken: req.RefreshToken,
	}

	result, err := s.refreshTokenUseCase.Refresh(ctx, args)
	if err != nil {
		// TODO:エラーをどう返すか
		return nil, err
	}

	return &pb.RefreshTokenResponse{
		AccessToken: result.AccessToken,
	}, nil
}
