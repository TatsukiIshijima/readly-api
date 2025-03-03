//go:build test

package server

import (
	"context"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	sqlc "readly/db/sqlc"
	"readly/env"
	"readly/pb"
	"readly/repository"
	"readly/service/auth"
	"readly/testdata"
	"readly/usecase"
	"testing"
	"time"
)

func NewTestUserServer(t *testing.T) *UserServerImpl {
	config := env.Config{
		TokenSymmetricKey:    testdata.RandomString(32),
		AccessTokenDuration:  time.Minute,
		RefreshTokenDuration: time.Hour,
	}
	fa := sqlc.FakeAdapter{}
	db, q := fa.Connect("", "")
	transaction := repository.New(db)

	userRepo := repository.NewUserRepository(q)
	sessionRepo := repository.NewSessionRepository(q)

	maker, err := auth.NewPasetoMaker(config.TokenSymmetricKey)
	require.NoError(t, err)

	signUpUseCase := usecase.NewSignUpUseCase(config, maker, transaction, sessionRepo, userRepo)
	signInUseCase := usecase.NewSignInUseCase(config, maker, transaction, sessionRepo, userRepo)
	refreshTokenUseCase := usecase.NewRefreshAccessTokenUseCase(config, maker, sessionRepo)

	return NewUserServer(
		config,
		maker,
		signUpUseCase,
		signInUseCase,
		refreshTokenUseCase,
	)
}

func TestSignUp(t *testing.T) {
	us := NewTestUserServer(t)
	ctx := metadata.NewIncomingContext(
		context.Background(),
		NewTestMetadata(
			"grpc-node-js/1.11.0-postman.1",
			"",
			"127.0.0.1",
		),
	)

	testCases := []struct {
		name    string
		prepare func(t *testing.T)
		req     *pb.SignUpRequest
		check   func(t *testing.T, res *pb.SignUpResponse, err error)
	}{
		{
			name:    "invalid request by missing required fields",
			prepare: func(t *testing.T) {},
			req:     &pb.SignUpRequest{},
			check: func(t *testing.T, res *pb.SignUpResponse, err error) {
				require.Nil(t, res)
				require.Error(t, err)
				s, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.InvalidArgument, s.Code())
			},
		},
		{
			name: "invalid request by empty name",
			prepare: func(t *testing.T) {

			},
			req: &pb.SignUpRequest{
				Name:     "",
				Email:    testdata.RandomEmail(),
				Password: "1234abcD@",
			},
			check: func(t *testing.T, res *pb.SignUpResponse, err error) {
				require.Nil(t, res)
				require.Error(t, err)
				s, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.InvalidArgument, s.Code())
			},
		},
		{
			name: "invalid request by invalid email",
			prepare: func(t *testing.T) {

			},
			req: &pb.SignUpRequest{
				Name:     "TestUser",
				Email:    "invalid",
				Password: "1234abcD@",
			},
			check: func(t *testing.T, res *pb.SignUpResponse, err error) {
				require.Nil(t, res)
				require.Error(t, err)
				s, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.InvalidArgument, s.Code())
			},
		},
		{
			name: "invalid request by short password",
			prepare: func(t *testing.T) {

			},
			req: &pb.SignUpRequest{
				Name:     testdata.RandomString(10),
				Email:    testdata.RandomEmail(),
				Password: "short",
			},
			check: func(t *testing.T, res *pb.SignUpResponse, err error) {
				require.Nil(t, res)
				require.Error(t, err)
				s, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.InvalidArgument, s.Code())
			},
		},
		{
			name: "invalid request by duplicate email",
			prepare: func(t *testing.T) {
				req := &pb.SignUpRequest{
					Name:     "TestUser",
					Email:    "duplicate@example.com",
					Password: "1234abcD@",
				}
				_, err := us.SignUp(ctx, req)
				require.NoError(t, err)
			},
			req: &pb.SignUpRequest{
				Name:     "TestUser",
				Email:    "duplicate@example.com",
				Password: "1234abcD@",
			},
			check: func(t *testing.T, res *pb.SignUpResponse, err error) {
				require.Nil(t, res)
				require.Error(t, err)
				s, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.AlreadyExists, s.Code())
			},
		},
		{
			name: "success: valid request",
			prepare: func(t *testing.T) {

			},
			req: &pb.SignUpRequest{
				Name:     "TestUser",
				Email:    testdata.RandomEmail(),
				Password: "1234abcD@",
			},
			check: func(t *testing.T, res *pb.SignUpResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotEmpty(t, res.GetAccessToken())
				require.NotEmpty(t, res.GetRefreshToken())
				require.NotEmpty(t, res.GetUserId())
				require.Equal(t, "TestUser", res.GetName())
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.prepare(t)
			res, err := us.SignUp(ctx, tc.req)
			tc.check(t, res, err)
		})
	}
}
