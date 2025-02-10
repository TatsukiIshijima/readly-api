package usecase

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"readly/testdata"
	"testing"
)

func TestSignIn(t *testing.T) {
	signUpUseCase := newTestSignUpUseCase(t)
	signInUseCase := newTestSignInUseCase(t)

	testCases := []struct {
		name  string
		setup func(t *testing.T) SignInRequest
		check func(t *testing.T, req SignInRequest, res *SignInResponse, err error)
	}{
		{
			name: "Sign in success if correct email and password",
			setup: func(t *testing.T) SignInRequest {
				email := testdata.RandomEmail()
				password := testdata.RandomString(16)
				signUpReq := SignUpRequest{
					Name:      testdata.RandomString(16),
					Email:     email,
					Password:  password,
					IPAddress: "127.0.0.1",
					UserAgent: "Mozilla/5.0",
				}
				_, err := signUpUseCase.SignUp(context.Background(), signUpReq)
				require.NoError(t, err)
				return SignInRequest{
					Email:    email,
					Password: password,
				}
			},
			check: func(t *testing.T, req SignInRequest, res *SignInResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotZero(t, len(res.AccessToken))
				require.NotZero(t, len(res.RefreshToken))
				require.NotEmpty(t, res.UserID)
				require.NotEmpty(t, res.Name)
				require.Equal(t, req.Email, res.Email)
			},
		},
		{
			name: "Sign in failure if not found email",
			setup: func(t *testing.T) SignInRequest {
				password := testdata.RandomString(16)
				signUpReq := SignUpRequest{
					Name:      testdata.RandomString(16),
					Email:     testdata.RandomEmail(),
					Password:  password,
					IPAddress: "127.0.0.1",
					UserAgent: "Mozilla/5.0",
				}
				_, err := signUpUseCase.SignUp(context.Background(), signUpReq)
				require.NoError(t, err)
				return SignInRequest{
					Email:    "not-found-email",
					Password: password,
				}
			},
			check: func(t *testing.T, req SignInRequest, res *SignInResponse, err error) {
				require.Nil(t, res)
				require.Error(t, err)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, e.Code, NotFound)
				require.Equal(t, e.Message, sql.ErrNoRows.Error())
			},
		},
		{
			name: "Sign in failure if wrong password",
			setup: func(t *testing.T) SignInRequest {
				email := testdata.RandomEmail()
				signUpReq := SignUpRequest{
					Name:      testdata.RandomString(16),
					Email:     email,
					Password:  testdata.RandomString(16),
					IPAddress: "127.0.0.1",
					UserAgent: "Mozilla/5.0",
				}
				_, err := signUpUseCase.SignUp(context.Background(), signUpReq)
				require.NoError(t, err)
				return SignInRequest{
					Email:    email,
					Password: "wrong-password",
				}
			},
			check: func(t *testing.T, req SignInRequest, res *SignInResponse, err error) {
				require.Nil(t, res)
				require.Error(t, err)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, e.Code, UnAuthorized)
				require.Equal(t, e.Message, bcrypt.ErrMismatchedHashAndPassword.Error())
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := tc.setup(t)
			res, err := signInUseCase.SignIn(context.Background(), req)
			tc.check(t, req, res, err)
		})
	}
}
