//go:build test

package server

import (
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
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

	testCases := []struct {
		name  string
		req   *pb.SignUpRequest
		check func(t *testing.T, res *pb.SignUpResponse, err error)
	}{
		{
			name: "invalid request by missing required fields",
			req:  &pb.SignUpRequest{},
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
			req: &pb.SignUpRequest{
				Name:     "",
				Email:    testdata.RandomEmail(),
				Password: testdata.RandomString(16),
			},
			check: func(t *testing.T, res *pb.SignUpResponse, err error) {
				require.Nil(t, res)
				require.Error(t, err)
				s, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.InvalidArgument, s.Code())
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := us.SignUp(nil, tc.req)
			tc.check(t, res, err)
		})
	}
}
