package usecase

import (
	"context"
	"github.com/stretchr/testify/require"
	"readly/entity"
	"readly/testdata"
	"testing"
)

func TestGetBook(t *testing.T) {
	signUpUseCase := newTestSignUpUseCase(t)
	registerBookUseCase := newTestRegisterBookUseCase(t)
	getBookUseCase := newTestGetBookUseCase(t)

	signUpReq := SignUpRequest{
		Name:     testdata.RandomString(10),
		Email:    testdata.RandomEmail(),
		Password: testdata.RandomString(16),
	}
	signUpRes, err := signUpUseCase.SignUp(context.Background(), signUpReq)
	require.NoError(t, err)
	registerBookReq := RegisterBookRequest{
		UserID: signUpRes.UserID,
		Title:  testdata.RandomString(10),
		Genres: []string{GetGenres()[0]},
		Status: entity.Unread,
	}
	registerBookRes, err := registerBookUseCase.RegisterBook(context.Background(), registerBookReq)
	require.NoError(t, err)

	testCases := []struct {
		name  string
		req   GetBookRequest
		check func(t *testing.T, res *GetBookResponse, err error)
	}{
		{
			name: "Get book by ID success when book is exists.",
			req: GetBookRequest{
				UserID: signUpRes.UserID,
				BookID: registerBookRes.Book.ID,
			},
			check: func(t *testing.T, res *GetBookResponse, err error) {
				require.NoError(t, err)
				require.Equal(t, registerBookRes.Book.ID, res.Book.ID)
				require.Equal(t, registerBookRes.Book.Title, res.Book.Title)
				require.Equal(t, registerBookRes.Book.Genres, res.Book.Genres)
				require.Equal(t, registerBookRes.Book.Status, res.Book.Status)
			},
		},
		{
			name: "Get book by ID failure when book is not exists.",
			req: GetBookRequest{
				UserID: signUpRes.UserID,
				BookID: 0,
			},
			check: func(t *testing.T, res *GetBookResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, e.StatusCode, BadRequest)
				require.Equal(t, e.ErrorCode, NotFoundBookError)
			},
		},
		{
			name: "Get book by ID failure when user is not exists.",
			req: GetBookRequest{
				UserID: 0,
				BookID: registerBookRes.Book.ID,
			},
			check: func(t *testing.T, res *GetBookResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, e.StatusCode, BadRequest)
				require.Equal(t, e.ErrorCode, NotFoundBookError)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := getBookUseCase.GetBook(context.Background(), tc.req)
			tc.check(t, res, err)
		})
	}
}
