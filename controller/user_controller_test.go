package controller

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"net/http"
	"readly/entity"
	"readly/testdata"
	"testing"
)

func TestSignUp(t *testing.T) {
	_, uc := setupControllers(t)

	testCases := []struct {
		name string
		req  SignUpRequest
		code int
		res  entity.User
	}{
		{
			name: "invalid request by missing required fields",
			req:  SignUpRequest{},
			code: http.StatusBadRequest,
			res:  entity.User{},
		},
		{
			name: "invalid request by empty name",
			req: SignUpRequest{
				Name:     "",
				Email:    testdata.RandomEmail(),
				Password: testdata.RandomString(16),
			},
			code: http.StatusBadRequest,
			res:  entity.User{},
		},
		{
			name: "invalid request by invalid email",
			req: SignUpRequest{
				Name:     testdata.RandomString(10),
				Email:    "invalid",
				Password: testdata.RandomString(16),
			},
			code: http.StatusBadRequest,
			res:  entity.User{},
		},
		{
			name: "invalid request by short password",
			req: SignUpRequest{
				Name:     testdata.RandomString(10),
				Email:    testdata.RandomEmail(),
				Password: "short",
			},
			code: http.StatusBadRequest,
			res:  entity.User{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			url := "/signup"
			body, err := json.Marshal(tc.req)
			println("request: " + string(body))
			require.NoError(t, err)

			ctx, recorder := setupTestContext("POST", url, body)
			uc.SignUp(ctx)

			if recorder.Code != tc.code {
				t.Fail()
			} else {
				switch recorder.Code {
				case http.StatusOK:
					var res entity.User
					err := json.Unmarshal(recorder.Body.Bytes(), &res)
					require.NoError(t, err)
				}
			}
		})
	}
}

func TestSignIn(t *testing.T) {
	_, uc := setupControllers(t)

	testCases := []struct {
		name string
		req  SignInRequest
		code int
		res  entity.User
	}{
		{
			name: "invalid request by missing required fields",
			req:  SignInRequest{},
			code: http.StatusBadRequest,
			res:  entity.User{},
		},
		{
			name: "invalid request by invalid email",
			req: SignInRequest{
				Email:    "invalid",
				Password: testdata.RandomString(16),
			},
			code: http.StatusBadRequest,
			res:  entity.User{},
		},
		{
			name: "invalid request by short password",
			req: SignInRequest{
				Email:    testdata.RandomEmail(),
				Password: "short",
			},
			code: http.StatusBadRequest,
			res:  entity.User{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			url := "/signin"
			body, err := json.Marshal(tc.req)
			println("request: " + string(body))
			require.NoError(t, err)

			ctx, recorder := setupTestContext("POST", url, body)
			uc.SignIn(ctx)

			if recorder.Code != tc.code {
				t.Fail()
			} else {
				switch recorder.Code {
				case http.StatusOK:
					var res entity.User
					err := json.Unmarshal(recorder.Body.Bytes(), &res)
					require.NoError(t, err)
				}
			}
		})
	}
}
