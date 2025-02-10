package controller

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"readly/testdata"
	"testing"
)

func TestSignUp(t *testing.T) {
	_, uc := setupControllers(t)

	duplicateEmail := testdata.RandomEmail()

	testCases := []struct {
		name  string
		setup func(t *testing.T)
		req   SignUpRequest
		check func(t *testing.T, req SignUpRequest, rec *httptest.ResponseRecorder)
	}{
		{
			name:  "invalid request by missing required fields",
			setup: func(t *testing.T) {},
			req:   SignUpRequest{},
			check: func(t *testing.T, req SignUpRequest, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, rec.Code)
			},
		},
		{
			name:  "invalid request by empty name",
			setup: func(t *testing.T) {},
			req: SignUpRequest{
				Name:     "",
				Email:    testdata.RandomEmail(),
				Password: testdata.RandomString(16),
			},
			check: func(t *testing.T, req SignUpRequest, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, rec.Code)
			},
		},
		{
			name:  "invalid request by invalid email",
			setup: func(t *testing.T) {},
			req: SignUpRequest{
				Name:     testdata.RandomString(10),
				Email:    "invalid",
				Password: testdata.RandomString(16),
			},
			check: func(t *testing.T, req SignUpRequest, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, rec.Code)
			},
		},
		{
			name:  "invalid request by short password",
			setup: func(t *testing.T) {},
			req: SignUpRequest{
				Name:     testdata.RandomString(10),
				Email:    testdata.RandomEmail(),
				Password: "short",
			},
			check: func(t *testing.T, req SignUpRequest, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, rec.Code)
			},
		},
		{
			name: "invalid request by duplicate email",
			setup: func(t *testing.T) {
				req := SignUpRequest{
					Name:     testdata.RandomString(10),
					Email:    duplicateEmail,
					Password: testdata.RandomString(16),
				}
				body, err := json.Marshal(req)
				require.NoError(t, err)
				ctx, rec := setupTestContext(http.MethodPost, "/signup", body)
				uc.SignUp(ctx)
				require.Equal(t, http.StatusOK, rec.Code)
			},
			req: SignUpRequest{
				Name:     testdata.RandomString(10),
				Email:    duplicateEmail,
				Password: testdata.RandomString(16),
			},
			check: func(t *testing.T, req SignUpRequest, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusConflict, rec.Code)
			},
		},
		{
			name:  "success: valid request",
			setup: func(t *testing.T) {},
			req: SignUpRequest{
				Name:     testdata.RandomString(10),
				Email:    testdata.RandomEmail(),
				Password: testdata.RandomString(16),
			},
			check: func(t *testing.T, req SignUpRequest, rec *httptest.ResponseRecorder) {
				var res *SignUpResponse
				err := json.Unmarshal(rec.Body.Bytes(), &res)
				require.NoError(t, err)
				require.Equal(t, http.StatusOK, rec.Code)
				require.NotEmpty(t, res.UserID)
				require.Equal(t, req.Name, res.Name)
				require.Equal(t, req.Email, res.Email)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			url := "/signup"
			body, err := json.Marshal(tc.req)
			println("request: " + string(body))
			require.NoError(t, err)

			ctx, recorder := setupTestContext(http.MethodPost, url, body)
			tc.setup(t)
			uc.SignUp(ctx)
			tc.check(t, tc.req, recorder)
		})
	}
}

func TestSignIn(t *testing.T) {
	_, uc := setupControllers(t)

	wrongPasswordCaseEmail := testdata.RandomEmail()
	successCaseEmail := testdata.RandomEmail()
	successCasePassword := testdata.RandomString(16)

	signUp := func(t *testing.T, req SignUpRequest) {
		body, err := json.Marshal(req)
		require.NoError(t, err)
		ctx, rec := setupTestContext(http.MethodPost, "/signup", body)
		uc.SignUp(ctx)
		require.Equal(t, http.StatusOK, rec.Code)
	}

	testCases := []struct {
		name  string
		setup func(t *testing.T)
		req   SignInRequest
		check func(t *testing.T, req SignInRequest, rec *httptest.ResponseRecorder)
	}{
		{
			name:  "invalid request by missing required fields",
			setup: func(t *testing.T) {},
			req:   SignInRequest{},
			check: func(t *testing.T, req SignInRequest, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, rec.Code)
			},
		},
		{
			name: "invalid request by invalid email",
			setup: func(t *testing.T) {
				req := SignUpRequest{
					Name:     testdata.RandomString(10),
					Email:    testdata.RandomEmail(),
					Password: testdata.RandomString(16),
				}
				signUp(t, req)
			},
			req: SignInRequest{
				Email:    "invalid",
				Password: testdata.RandomString(16),
			},
			check: func(t *testing.T, req SignInRequest, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, rec.Code)
			},
		},
		{
			name: "invalid request by short password",
			setup: func(t *testing.T) {
				req := SignUpRequest{
					Name:     testdata.RandomString(10),
					Email:    testdata.RandomEmail(),
					Password: testdata.RandomString(16),
				}
				signUp(t, req)
			},
			req: SignInRequest{
				Email:    testdata.RandomEmail(),
				Password: "short",
			},
			check: func(t *testing.T, req SignInRequest, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, rec.Code)
			},
		},
		{
			name:  "failure: not found email",
			setup: func(t *testing.T) {},
			req: SignInRequest{
				Email:    testdata.RandomEmail(),
				Password: testdata.RandomString(16),
			},
			check: func(t *testing.T, req SignInRequest, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, rec.Code)
			},
		},
		{
			name: "failure: wrong password",
			setup: func(t *testing.T) {
				req := SignUpRequest{
					Name:     testdata.RandomString(10),
					Email:    wrongPasswordCaseEmail,
					Password: testdata.RandomString(16),
				}
				signUp(t, req)
			},
			req: SignInRequest{
				Email:    wrongPasswordCaseEmail,
				Password: testdata.RandomString(16),
			},
			check: func(t *testing.T, req SignInRequest, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, rec.Code)
			},
		},
		{
			name: "success: valid request",
			setup: func(t *testing.T) {
				req := SignUpRequest{
					Name:     testdata.RandomString(10),
					Email:    successCaseEmail,
					Password: successCasePassword,
				}
				signUp(t, req)
			},
			req: SignInRequest{
				Email:    successCaseEmail,
				Password: successCasePassword,
			},
			check: func(t *testing.T, req SignInRequest, rec *httptest.ResponseRecorder) {
				var res SignInResponse
				err := json.Unmarshal(rec.Body.Bytes(), &res)
				require.NoError(t, err)
				require.Equal(t, http.StatusOK, rec.Code)
				require.NotEmpty(t, res.AccessToken)
				require.NotEmpty(t, res.RefreshToken)
				require.NotEmpty(t, res.UserID)
				require.NotEmpty(t, res.Name)
				require.Equal(t, req.Email, res.Email)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			url := "/signin"
			body, err := json.Marshal(tc.req)
			println("request: " + string(body))
			require.NoError(t, err)

			ctx, recorder := setupTestContext(http.MethodPost, url, body)
			tc.setup(t)
			uc.SignIn(ctx)
			tc.check(t, tc.req, recorder)
		})
	}
}

func TestRefreshToken(t *testing.T) {
	_, uc := setupControllers(t)

	testCases := []struct {
		name  string
		setup func(t *testing.T) RefreshTokenRequest
		check func(t *testing.T, rec *httptest.ResponseRecorder)
	}{
		{
			name: "Refresh token success if correct refresh token",
			setup: func(t *testing.T) RefreshTokenRequest {
				signUpReq := SignUpRequest{
					Name:     testdata.RandomString(16),
					Email:    testdata.RandomEmail(),
					Password: testdata.RandomString(16),
				}
				url := "/signup"
				body, err := json.Marshal(signUpReq)
				require.NoError(t, err)

				ctx, recorder := setupTestContext(http.MethodPost, url, body)
				uc.SignUp(ctx)

				var res *SignUpResponse
				err = json.Unmarshal(recorder.Body.Bytes(), &res)
				require.NoError(t, err)

				return RefreshTokenRequest{
					RefreshToken: res.RefreshToken,
				}
			},
			check: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var res *RefreshTokenResponse
				err := json.Unmarshal(rec.Body.Bytes(), &res)
				require.NoError(t, err)
				require.Equal(t, http.StatusOK, rec.Code)
				require.NotZero(t, len(res.AccessToken))
			},
		},
		{
			name: "Refresh token failure if request empty",
			setup: func(t *testing.T) RefreshTokenRequest {
				return RefreshTokenRequest{
					RefreshToken: "",
				}
			},
			check: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, rec.Code)
			},
		},
		{
			name: "Refresh token failure if incorrect refresh token",
			setup: func(t *testing.T) RefreshTokenRequest {
				signUpReq := SignUpRequest{
					Name:     testdata.RandomString(16),
					Email:    testdata.RandomEmail(),
					Password: testdata.RandomString(16),
				}
				url := "/signup"
				body, err := json.Marshal(signUpReq)
				require.NoError(t, err)

				ctx, recorder := setupTestContext(http.MethodPost, url, body)
				uc.SignUp(ctx)

				var res *SignUpResponse
				err = json.Unmarshal(recorder.Body.Bytes(), &res)
				require.NoError(t, err)

				return RefreshTokenRequest{
					RefreshToken: "invalid_refresh_token",
				}
			},
			check: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, rec.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := tc.setup(t)

			url := "/refresh-token"
			body, err := json.Marshal(req)
			println("request: " + string(body))
			require.NoError(t, err)

			ctx, recorder := setupTestContext(http.MethodPost, url, body)
			uc.RefreshToken(ctx)
			tc.check(t, recorder)
		})
	}
}
