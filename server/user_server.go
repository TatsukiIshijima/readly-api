package server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"readly/env"
	"readly/pb/readly/v1"
	"readly/service/auth"
	"readly/usecase"
	"readly/util"
)

type UserServerImpl struct {
	pb.UnimplementedUserServiceServer
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

	validateFunc := func() error {
		err := util.StringValidator(req.GetEmail()).ValidateEmail()
		if err != nil {
			return err
		}
		err = util.StringValidator(req.GetPassword()).ValidatePassword()
		if err != nil {
			return err
		}
		return nil
	}

	err := validateFunc()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	meta := newMetadataFrom(ctx)
	args := usecase.NewSignInRequest(req, meta.IPAddress, meta.UserAgent)
	result, err := s.signInUseCase.SignIn(ctx, args)
	if err != nil {
		return nil, gRPCStatusError(err)
	}
	return result.ToProto(), nil
}

func (s *UserServerImpl) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {

	validateFunc := func() error {
		err := util.StringValidator(req.GetName()).ValidateUsername()
		if err != nil {
			return err
		}
		err = util.StringValidator(req.GetEmail()).ValidateEmail()
		if err != nil {
			return err
		}
		err = util.StringValidator(req.GetPassword()).ValidatePassword()
		if err != nil {
			return err
		}
		return nil
	}

	err := validateFunc()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	meta := newMetadataFrom(ctx)
	args := usecase.NewSignUpRequest(req, meta.IPAddress, meta.UserAgent)
	result, err := s.signUpUseCase.SignUp(ctx, args)
	if err != nil {
		return nil, gRPCStatusError(err)
	}
	return result.ToProto(), nil
}

func (s *UserServerImpl) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	args := usecase.NewRefreshAccessTokenRequest(req)

	result, err := s.refreshTokenUseCase.Refresh(ctx, args)
	if err != nil {
		return nil, gRPCStatusError(err)
	}

	return result.ToProto(), nil
}
