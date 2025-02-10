package usecase

import (
	"context"
	"github.com/stretchr/testify/require"
	"readly/testdata"
	"testing"
)

func TestRefreshToken(t *testing.T) {
	signUpUseCase := newTestSignUpUseCase(t)
	refreshAccessTokenUseCase := newTestRefreshAccessTokenUseCase(t)

	testCases := []struct {
		name  string
		setup func(t *testing.T) RefreshAccessTokenRequest
		check func(t *testing.T, req RefreshAccessTokenRequest, res *RefreshAccessTokenResponse, err error)
	}{
		{
			name: "Refresh token success if correct refresh token",
			setup: func(t *testing.T) RefreshAccessTokenRequest {
				email := testdata.RandomEmail()
				password := testdata.RandomString(16)
				signUpReq := SignUpRequest{
					Name:     testdata.RandomString(16),
					Email:    email,
					Password: password,
				}
				res, err := signUpUseCase.SignUp(context.Background(), signUpReq)
				require.NoError(t, err)

				return RefreshAccessTokenRequest{
					RefreshToken: res.RefreshToken,
				}
			},
			check: func(t *testing.T, req RefreshAccessTokenRequest, res *RefreshAccessTokenResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotZero(t, len(res.AccessToken))
			},
		},
		{
			name: "Refresh token failure if incorrect refresh token",
			setup: func(t *testing.T) RefreshAccessTokenRequest {
				return RefreshAccessTokenRequest{
					RefreshToken: "invalid_refresh_token",
				}
			},
			check: func(t *testing.T, req RefreshAccessTokenRequest, res *RefreshAccessTokenResponse, err error) {
				require.Nil(t, res)
				require.Error(t, err)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, e.StatusCode, UnAuthorized)
				require.Equal(t, e.ErrorCode, InvalidTokenError)
				require.Equal(t, e.Message, "invalid refresh token")
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := tc.setup(t)
			res, err := refreshAccessTokenUseCase.Refresh(context.Background(), req)
			tc.check(t, req, res, err)
		})
	}
}
