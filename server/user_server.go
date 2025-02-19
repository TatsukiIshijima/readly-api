package server

import (
	"context"
	"readly/pb"
)

type UserServerImpl struct {
	pb.UnimplementedUserServer
}

func NewUserServer() pb.UserServer {
	return &UserServerImpl{}
}

func (s *UserServerImpl) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	// TODO: Implement
}

func (s *UserServerImpl) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	// TODO: Implement

}

func (s *UserServerImpl) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	// TODO: Implement
}
