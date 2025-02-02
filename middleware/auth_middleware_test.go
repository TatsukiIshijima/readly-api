package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"readly/service/auth"
	"readly/testdata"
	"testing"
	"time"
)

func setAuthorizationHeader(
	t *testing.T,
	req *http.Request,
	maker auth.TokenMaker,
	authorizationType string,
	userID int64,
	duration time.Duration,
) {
	token, err := maker.Generate(userID, duration)
	require.NoError(t, err)

	authorizationHeader := fmt.Sprintf("%s %s", authorizationType, token)
	req.Header.Set(authorizationHeaderKey, authorizationHeader)
}

func TestAuthorize(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testCases := []struct {
		name  string
		setup func(t *testing.T, req *http.Request, maker auth.TokenMaker)
		check func(t *testing.T, rec *httptest.ResponseRecorder)
	}{
		{
			name: "authorize success",
			setup: func(t *testing.T, req *http.Request, maker auth.TokenMaker) {
				setAuthorizationHeader(t, req, maker, authorizationTypeBearer, 1, time.Minute)
			},
			check: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, rec.Code)
			},
		},
		{
			name: "unauthorized by no authorization header",
			setup: func(t *testing.T, req *http.Request, maker auth.TokenMaker) {

			},
			check: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, rec.Code)
			},
		},
		{
			name: "unauthorized by unsupported authorization type",
			setup: func(t *testing.T, req *http.Request, maker auth.TokenMaker) {
				setAuthorizationHeader(t, req, maker, "unsupportedType", 1, time.Minute)
			},
			check: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, rec.Code)
			},
		},
		{
			name: "unauthorized by invalid authorization header format",
			setup: func(t *testing.T, req *http.Request, maker auth.TokenMaker) {
				setAuthorizationHeader(t, req, maker, "", 1, time.Minute)
			},
			check: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, rec.Code)
			},
		},
		{
			name: "unauthorized by expired access token",
			setup: func(t *testing.T, req *http.Request, maker auth.TokenMaker) {
				setAuthorizationHeader(t, req, maker, authorizationTypeBearer, 1, -time.Minute)
			},
			check: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, rec.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			router := gin.Default()
			maker, err := auth.NewPasetoMaker(testdata.RandomString(32))
			require.NoError(t, err)

			router.GET("/test", Authorize(maker), func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{})
			})

			recorder := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodGet, "/test", nil)
			require.NoError(t, err)

			tc.setup(t, req, maker)
			router.ServeHTTP(recorder, req)
			tc.check(t, recorder)
		})
	}
}
