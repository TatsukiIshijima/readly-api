package usecase

import (
	"context"
	"github.com/stretchr/testify/require"
	"readly/entity"
	"readly/testdata"
	"testing"
)

func TestSignUp(t *testing.T) {
	name := testdata.RandomString(10)
	email := testdata.RandomEmail()

	testCases := []struct {
		name string
		req  SignUpRequest
		exp  *entity.User
		err  error
	}{
		{
			name: "Sign up success",
			req: SignUpRequest{
				Name:     name,
				Email:    email,
				Password: testdata.RandomString(16),
			},
			exp: &entity.User{
				// 出力されるIDは自動採番のためIDは比較対象としないとし、適当な値を入れている
				ID:    0,
				Name:  name,
				Email: email,
			},
			err: nil,
		},
		{
			name: "Sign up failure by same email",
			req: SignUpRequest{
				Name:     name,
				Email:    email,
				Password: testdata.RandomString(16),
			},
			exp: nil,
			err: newError("duplicate key value violates unique constraint \"users_email_key\"", Conflict),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := signUpUseCase.SignUp(context.Background(), tc.req)
			if tc.err == nil {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.Equal(t, tc.exp.Name, res.Name)
				require.Equal(t, tc.exp.Email, res.Email)
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
