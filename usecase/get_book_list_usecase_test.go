package usecase

import (
	"context"
	"github.com/stretchr/testify/require"
	"readly/entity"
	"readly/testdata"
	"testing"
)

func TestGetBookList(t *testing.T) {
	signUpUseCase := newTestSignUpUseCase(t)
	registerBookUseCase := newTestRegisterBookUseCase(t)
	getBookListUseCase := newTestGetBookListUseCase(t)

	signUpReq := SignUpRequest{
		Name:     testdata.RandomString(10),
		Email:    testdata.RandomEmail(),
		Password: testdata.RandomString(16),
	}
	signUpRes, err := signUpUseCase.SignUp(context.Background(), signUpReq)
	require.NoError(t, err)
	for i := 0; i < 5; i++ {
		registerBookReq := RegisterBookRequest{
			UserID: signUpRes.UserID,
			Title:  testdata.RandomString(10),
			Genres: []string{testdata.RandomString(6)},
			Status: entity.Unread,
		}
		_, err := registerBookUseCase.RegisterBook(context.Background(), registerBookReq)
		require.NoError(t, err)
	}

	testCases := []struct {
		name  string
		req   GetBookListRequest
		check func(t *testing.T, res *GetBookListResponse, err error)
	}{
		{
			name: "Get book list success when user has books.",
			req: GetBookListRequest{
				UserID: signUpRes.UserID,
				Limit:  10,
				Offset: 0,
			},
			check: func(t *testing.T, res *GetBookListResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.Equal(t, 5, len(res.Books))
			},
		},
		{
			name: "Get book list success when user has no books.",
			req: GetBookListRequest{
				UserID: 0,
				Limit:  10,
				Offset: 0,
			},
			check: func(t *testing.T, res *GetBookListResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.Equal(t, 0, len(res.Books))
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := getBookListUseCase.GetBookList(context.Background(), testCase.req)
			testCase.check(t, res, err)
		})
	}
}
