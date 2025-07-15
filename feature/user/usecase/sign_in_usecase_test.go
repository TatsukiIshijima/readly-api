//go:build test

package usecase

import (
	"context"
	"github.com/stretchr/testify/require"
	"readly/testdata"
	"testing"
)

func TestSignIn(t *testing.T) {
	signUpUseCase := newTestSignUpUseCase(t)
	signInUseCase := newTestSignInUseCase(t)

	email := testdata.RandomEmail()
	password := testdata.RandomValidPassword()
	signUpReq := SignUpRequest{
		Name:      testdata.RandomString(16),
		Email:     email,
		Password:  password,
		IPAddress: "127.0.0.1",
		UserAgent: "Mozilla/5.0",
	}
	_, err := signUpUseCase.SignUp(context.Background(), signUpReq)
	require.NoError(t, err)

	testCases := []struct {
		name  string
		setup func(t *testing.T) SignInRequest
		check func(t *testing.T, req SignInRequest, res *SignInResponse, err error)
	}{
		{
			name: "Sign in success if correct email and password",
			setup: func(t *testing.T) SignInRequest {
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

				sessions, err := querier.GetSessionByUserID(context.Background(), res.UserID)
				require.NoError(t, err)
				require.Equal(t, len(sessions), 2)
			},
		},
		{
			name: "Sign in success from multi devices",
			setup: func(t *testing.T) SignInRequest {
				for i := 0; i < 6; i++ {
					_, err := signInUseCase.SignIn(context.Background(), SignInRequest{
						Email:     email,
						Password:  password,
						IPAddress: "127.0.0.1",
						UserAgent: "Mozilla/5.0",
					})
					require.NoError(t, err)
				}
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

				sessions, err := querier.GetSessionByUserID(context.Background(), res.UserID)
				require.NoError(t, err)
				require.Equal(t, 5, len(sessions))
			},
		},
		{
			name: "Sign in failure if not found email",
			setup: func(t *testing.T) SignInRequest {
				return SignInRequest{
					Email:    testdata.RandomEmail(),
					Password: password,
				}
			},
			check: func(t *testing.T, req SignInRequest, res *SignInResponse, err error) {
				require.Nil(t, res)
				require.Error(t, err)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, e.StatusCode, BadRequest)
				require.Equal(t, e.ErrorCode, NotFoundUserError)
			},
		},
		{
			name: "Sign in failure if wrong password",
			setup: func(t *testing.T) SignInRequest {
				return SignInRequest{
					Email:    email,
					Password: testdata.RandomValidPassword(),
				}
			},
			check: func(t *testing.T, req SignInRequest, res *SignInResponse, err error) {
				require.Nil(t, res)
				require.Error(t, err)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, e.StatusCode, BadRequest)
				require.Equal(t, e.ErrorCode, InvalidPasswordError)
			},
		},
		{
			name: "Sign in failed when email is empty",
			setup: func(t *testing.T) SignInRequest {
				return SignInRequest{
					Email:    "",
					Password: password,
				}
			},
			check: func(t *testing.T, req SignInRequest, res *SignInResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, BadRequest, e.StatusCode)
				require.Equal(t, InvalidRequestError, e.ErrorCode)
				require.Contains(t, e.Message, "email is required")
			},
		},
		{
			name: "Sign in failed when email has invalid format",
			setup: func(t *testing.T) SignInRequest {
				return SignInRequest{
					Email:    "invalid-email",
					Password: password,
				}
			},
			check: func(t *testing.T, req SignInRequest, res *SignInResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, BadRequest, e.StatusCode)
				require.Equal(t, InvalidRequestError, e.ErrorCode)
				require.Contains(t, e.Message, "email has invalid format")
			},
		},
		{
			name: "Sign in failed when email contains SQL injection",
			setup: func(t *testing.T) SignInRequest {
				return SignInRequest{
					Email:    "user.select@example.com",
					Password: password,
				}
			},
			check: func(t *testing.T, req SignInRequest, res *SignInResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, BadRequest, e.StatusCode)
				require.Equal(t, InvalidRequestError, e.ErrorCode)
				require.Contains(t, e.Message, "email contains potentially dangerous content")
			},
		},
		{
			name: "Sign in failed when password is empty",
			setup: func(t *testing.T) SignInRequest {
				return SignInRequest{
					Email:    email,
					Password: "",
				}
			},
			check: func(t *testing.T, req SignInRequest, res *SignInResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, BadRequest, e.StatusCode)
				require.Equal(t, InvalidRequestError, e.ErrorCode)
				require.Contains(t, e.Message, "password is required")
			},
		},
		{
			name: "Sign in failed when password is too short",
			setup: func(t *testing.T) SignInRequest {
				return SignInRequest{
					Email:    email,
					Password: "Pass1!",
				}
			},
			check: func(t *testing.T, req SignInRequest, res *SignInResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, BadRequest, e.StatusCode)
				require.Equal(t, InvalidRequestError, e.ErrorCode)
				require.Contains(t, e.Message, "password length must be between 8 and 48")
			},
		},
		{
			name: "Sign in failed when password is too long",
			setup: func(t *testing.T) SignInRequest {
				longPassword := testdata.RandomString(49) + "A1!"
				return SignInRequest{
					Email:    email,
					Password: longPassword,
				}
			},
			check: func(t *testing.T, req SignInRequest, res *SignInResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, BadRequest, e.StatusCode)
				require.Equal(t, InvalidRequestError, e.ErrorCode)
				require.Contains(t, e.Message, "password length must be between 8 and 48")
			},
		},
		{
			name: "Sign in failed when password has no uppercase letter",
			setup: func(t *testing.T) SignInRequest {
				return SignInRequest{
					Email:    email,
					Password: "password123!",
				}
			},
			check: func(t *testing.T, req SignInRequest, res *SignInResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, BadRequest, e.StatusCode)
				require.Equal(t, InvalidRequestError, e.ErrorCode)
				require.Contains(t, e.Message, "password must contain at least one uppercase letter")
			},
		},
		{
			name: "Sign in failed when password has no lowercase letter",
			setup: func(t *testing.T) SignInRequest {
				return SignInRequest{
					Email:    email,
					Password: "PASSWORD123!",
				}
			},
			check: func(t *testing.T, req SignInRequest, res *SignInResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, BadRequest, e.StatusCode)
				require.Equal(t, InvalidRequestError, e.ErrorCode)
				require.Contains(t, e.Message, "password must contain at least one lowercase letter")
			},
		},
		{
			name: "Sign in failed when password has no digit",
			setup: func(t *testing.T) SignInRequest {
				return SignInRequest{
					Email:    email,
					Password: "Password!",
				}
			},
			check: func(t *testing.T, req SignInRequest, res *SignInResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, BadRequest, e.StatusCode)
				require.Equal(t, InvalidRequestError, e.ErrorCode)
				require.Contains(t, e.Message, "password must contain at least one digit")
			},
		},
		{
			name: "Sign in failed when password has no symbol",
			setup: func(t *testing.T) SignInRequest {
				return SignInRequest{
					Email:    email,
					Password: "Password123",
				}
			},
			check: func(t *testing.T, req SignInRequest, res *SignInResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, BadRequest, e.StatusCode)
				require.Equal(t, InvalidRequestError, e.ErrorCode)
				require.Contains(t, e.Message, "password must contain at least one symbol")
			},
		},
		{
			name: "Sign in failed when password contains SQL injection",
			setup: func(t *testing.T) SignInRequest {
				return SignInRequest{
					Email:    email,
					Password: "Password123!'; DROP TABLE users; --",
				}
			},
			check: func(t *testing.T, req SignInRequest, res *SignInResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, BadRequest, e.StatusCode)
				require.Equal(t, InvalidRequestError, e.ErrorCode)
				require.Contains(t, e.Message, "password contains potentially dangerous content")
			},
		},
		{
			name: "Sign in failed when IP address has invalid format",
			setup: func(t *testing.T) SignInRequest {
				return SignInRequest{
					Email:     email,
					Password:  password,
					IPAddress: "invalid-ip",
				}
			},
			check: func(t *testing.T, req SignInRequest, res *SignInResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, BadRequest, e.StatusCode)
				require.Equal(t, InvalidRequestError, e.ErrorCode)
				require.Contains(t, e.Message, "IP address has invalid format")
			},
		},
		{
			name: "Sign in failed when user agent exceeds 2048 characters",
			setup: func(t *testing.T) SignInRequest {
				longUserAgent := testdata.RandomString(2049)
				return SignInRequest{
					Email:     email,
					Password:  password,
					UserAgent: longUserAgent,
				}
			},
			check: func(t *testing.T, req SignInRequest, res *SignInResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, BadRequest, e.StatusCode)
				require.Equal(t, InvalidRequestError, e.ErrorCode)
				require.Contains(t, e.Message, "user agent must be less than 2048 characters")
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
