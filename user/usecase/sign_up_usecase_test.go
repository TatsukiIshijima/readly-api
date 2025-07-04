package usecase

import (
	"context"
	"github.com/stretchr/testify/require"
	"readly/testdata"
	"testing"
)

func TestSignUp(t *testing.T) {
	signUpUseCase := newTestSignUpUseCase(t)

	testCases := []struct {
		name  string
		setup func(t *testing.T) SignUpRequest
		check func(t *testing.T, req SignUpRequest, res *SignUpResponse, err error)
	}{
		{
			name: "Sign up success",
			setup: func(t *testing.T) SignUpRequest {
				return SignUpRequest{
					Name:      testdata.RandomString(10),
					Email:     testdata.RandomEmail(),
					Password:  testdata.RandomString(16),
					IPAddress: "127.0.0.1",
					UserAgent: "Mozilla/5.0",
				}
			},
			check: func(t *testing.T, req SignUpRequest, res *SignUpResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotZero(t, len(res.AccessToken))
				require.NotZero(t, len(res.RefreshToken))
				require.NotEmpty(t, res.UserID)
				require.Equal(t, req.Name, res.Name)
				require.Equal(t, req.Email, res.Email)
			},
		},
		{
			name: "Sign up failure if same email already exists",
			setup: func(t *testing.T) SignUpRequest {
				email := testdata.RandomEmail()

				req := SignUpRequest{
					Name:      testdata.RandomString(10),
					Email:     email,
					Password:  testdata.RandomString(16),
					IPAddress: "127.0.0.1",
				}
				_, err := signUpUseCase.SignUp(context.Background(), req)
				require.NoError(t, err)

				return SignUpRequest{
					Name:      testdata.RandomString(10),
					Email:     email,
					Password:  testdata.RandomString(16),
					IPAddress: "127.0.0.1",
				}
			},
			check: func(t *testing.T, req SignUpRequest, res *SignUpResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, e.StatusCode, Conflict)
				require.Equal(t, e.ErrorCode, EmailAlreadyRegisteredError)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := tc.setup(t)
			res, err := signUpUseCase.SignUp(context.Background(), req)
			tc.check(t, req, res, err)
		})
	}
}
