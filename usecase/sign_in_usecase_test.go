package usecase

import (
	"context"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"readly/entity"
	"readly/testdata"
	"testing"
)

func TestSignIn(t *testing.T) {
	name := testdata.RandomString(10)
	email := testdata.RandomEmail()
	password := testdata.RandomString(16)

	// Sign up
	signUpReq := SignUpRequest{
		Name:     name,
		Email:    email,
		Password: password,
	}
	user, err := signUpUseCase.SignUp(context.Background(), signUpReq)
	require.NoError(t, err)

	testCases := []struct {
		name string
		req  SignInRequest
		exp  *entity.User
		err  error
	}{
		{
			name: "Sign in success",
			req: SignInRequest{
				Email:    email,
				Password: password,
			},
			exp: user,
			err: nil,
		},
		{
			name: "Sign in failure by wrong password",
			req: SignInRequest{
				Email:    email,
				Password: testdata.RandomString(16),
			},
			exp: nil,
			err: newError(bcrypt.ErrMismatchedHashAndPassword.Error(), UnAuthorized),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := signInUseCase.SignIn(context.Background(), tc.req)
			if tc.err == nil {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.Equal(t, tc.exp, res)
			} else {
				require.Nil(t, res)
				var tcErr *Error
				var ucErr *Error
				require.ErrorAs(t, err, &ucErr)
				require.ErrorAs(t, tc.err, &tcErr)
				require.Equal(t, tcErr.Code, ucErr.Code)
			}
		})
	}
}
