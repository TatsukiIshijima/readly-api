//go:build test

package server

import (
	"context"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"readly/configs"
	sqlc "readly/db/sqlc"
	"readly/db/transaction"
	"readly/middleware/auth"
	"readly/pb/readly/v1"
	"readly/repository"
	"readly/testdata"
	userRepo "readly/user/repository"
	"readly/user/usecase"
	"testing"
	"time"
)

func NewTestUserServer(t *testing.T) *UserServerImpl {
	config := configs.Config{
		TokenSymmetricKey:    testdata.RandomString(32),
		AccessTokenDuration:  time.Minute,
		RefreshTokenDuration: time.Hour,
	}
	// TODO:本物のDBを使うよう変更
	fa := sqlc.FakeAdapter{}
	db, q := fa.Connect("", "")
	transactor := transaction.New(db)

	userRepository := userRepo.NewUserRepository(q)
	sessionRepo := repository.NewSessionRepository(q)

	maker, err := auth.NewPasetoMaker(config.TokenSymmetricKey)
	require.NoError(t, err)

	signUpUseCase := usecase.NewSignUpUseCase(config, maker, transactor, sessionRepo, userRepository)
	signInUseCase := usecase.NewSignInUseCase(config, maker, transactor, sessionRepo, userRepository)
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

func TestSignIn(t *testing.T) {
	us := NewTestUserServer(t)
	ctx := metadata.NewIncomingContext(
		context.Background(),
		NewTestMetadata(
			"grpc-node-js/1.11.0-postman.1",
			"",
			"127.0.0.1",
		),
	)

	wrongPasswordCaseEmail := testdata.RandomEmail()
	successCaseEmail := testdata.RandomEmail()

	testCases := []struct {
		name    string
		prepare func(t *testing.T)
		req     *pb.SignInRequest
		check   func(t *testing.T, res *pb.SignInResponse, err error)
	}{
		{
			name: "invalid request by missing required fields",
			prepare: func(t *testing.T) {
			},
			req: &pb.SignInRequest{},
			check: func(t *testing.T, res *pb.SignInResponse, err error) {
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
			req: &pb.SignInRequest{
				Email:    "invalid",
				Password: "1234abcD@",
			},
			check: func(t *testing.T, res *pb.SignInResponse, err error) {
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
			req: &pb.SignInRequest{
				Email:    testdata.RandomEmail(),
				Password: "short",
			},
			check: func(t *testing.T, res *pb.SignInResponse, err error) {
				require.Nil(t, res)
				require.Error(t, err)
				s, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.InvalidArgument, s.Code())
			},
		},
		{
			name: "failure: not found email",
			prepare: func(t *testing.T) {
				req := pb.SignUpRequest{
					Name:     "TestUser",
					Email:    testdata.RandomEmail(),
					Password: "1234abcD@",
				}
				_, err := us.SignUp(ctx, &req)
				require.NoError(t, err)
			},
			req: &pb.SignInRequest{
				Email:    testdata.RandomEmail(),
				Password: "1234abcD@",
			},
			check: func(t *testing.T, res *pb.SignInResponse, err error) {
				require.Nil(t, res)
				require.Error(t, err)
				s, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.InvalidArgument, s.Code())
			},
		},
		{
			name: "failure: wrong password",
			prepare: func(t *testing.T) {
				req := pb.SignUpRequest{
					Name:     "TestUser",
					Email:    wrongPasswordCaseEmail,
					Password: "1234abcD@",
				}
				_, err := us.SignUp(ctx, &req)
				require.NoError(t, err)
			},
			req: &pb.SignInRequest{
				Email:    wrongPasswordCaseEmail,
				Password: "1234abcD",
			},
			check: func(t *testing.T, res *pb.SignInResponse, err error) {
				require.Nil(t, res)
				require.Error(t, err)
				s, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.InvalidArgument, s.Code())
			},
		},
		{
			name: "success: valid request",
			prepare: func(t *testing.T) {
				req := pb.SignUpRequest{
					Name:     "TestUser",
					Email:    successCaseEmail,
					Password: "1234abcD@",
				}
				_, err := us.SignUp(ctx, &req)
				require.NoError(t, err)
			},
			req: &pb.SignInRequest{
				Email:    successCaseEmail,
				Password: "1234abcD@",
			},
			check: func(t *testing.T, res *pb.SignInResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotEmpty(t, res.GetAccessToken())
				require.NotEmpty(t, res.GetRefreshToken())
				require.NotEmpty(t, res.GetUserId())
				require.Equal(t, "TestUser", res.GetName())
				require.Equal(t, successCaseEmail, res.GetEmail())
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.prepare(t)
			res, err := us.SignIn(ctx, tc.req)
			tc.check(t, res, err)
		})
	}
}

func TestRefreshAccessToken(t *testing.T) {
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
		prepare func(t *testing.T) pb.RefreshTokenRequest
		check   func(t *testing.T, req *pb.RefreshTokenRequest, res *pb.RefreshTokenResponse, err error)
	}{
		{
			name: "Refresh token success if correct refresh token",
			prepare: func(t *testing.T) pb.RefreshTokenRequest {
				req := pb.SignUpRequest{
					Name:     "TestUser",
					Email:    testdata.RandomEmail(),
					Password: "1234abcD@",
				}
				res, err := us.SignUp(ctx, &req)
				require.NoError(t, err)

				return pb.RefreshTokenRequest{
					RefreshToken: res.GetRefreshToken(),
				}
			},
			check: func(t *testing.T, req *pb.RefreshTokenRequest, res *pb.RefreshTokenResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotEmpty(t, res.GetAccessToken())
				require.Equal(t, res.GetAccessToken(), res.GetAccessToken())
			},
		},
		{
			name: "Refresh token failure if request empty",
			prepare: func(t *testing.T) pb.RefreshTokenRequest {
				return pb.RefreshTokenRequest{}
			},
			check: func(t *testing.T, req *pb.RefreshTokenRequest, res *pb.RefreshTokenResponse, err error) {
				require.Nil(t, res)
				require.Error(t, err)
				s, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Unauthenticated, s.Code())
			},
		},
		{
			name: "Refresh token failure if invalid refresh token",
			prepare: func(t *testing.T) pb.RefreshTokenRequest {
				req := pb.SignUpRequest{
					Name:     "TestUser",
					Email:    testdata.RandomEmail(),
					Password: "1234abcD@",
				}
				res, err := us.SignUp(ctx, &req)
				require.NoError(t, err)

				return pb.RefreshTokenRequest{
					RefreshToken: res.GetRefreshToken() + "_invalid",
				}
			},
			check: func(t *testing.T, req *pb.RefreshTokenRequest, res *pb.RefreshTokenResponse, err error) {
				require.Nil(t, res)
				require.Error(t, err)
				s, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Unauthenticated, s.Code())
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := tc.prepare(t)
			res, err := us.RefreshToken(ctx, &req)
			tc.check(t, &req, res, err)
		})
	}
}
