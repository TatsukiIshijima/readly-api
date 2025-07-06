package usecase

import (
	"context"
	"github.com/stretchr/testify/require"
	"readly/entity"
	"readly/testdata"
	"testing"
)

func TestGetBookList(t *testing.T) {
	registerBookUseCase := newTestRegisterBookUseCase(t)
	getBookListUseCase := newTestGetBookListUseCase(t)

	user := createRandomUser(t)

	for i := 0; i < 5; i++ {
		registerBookReq := RegisterBookRequest{
			UserID: user.ID,
			Title:  testdata.RandomString(10),
			Genres: []string{testdata.GetGenres()[i]},
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
				UserID: user.ID,
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
		{
			name: "Get book list failure when invalid request.",
			req: GetBookListRequest{
				UserID: 0,
				Limit:  -1,
				Offset: -1,
			},
			check: func(t *testing.T, res *GetBookListResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, e.StatusCode, BadRequest)
				require.Equal(t, e.ErrorCode, InvalidRequestError)
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
