package usecase

import (
	"context"
	"github.com/stretchr/testify/require"
	"readly/entity"
	"readly/testdata"
	"testing"
)

func TestUpdateBook(t *testing.T) {
	signUpUseCase := newTestSignUpUseCase(t)
	registerBookUseCase := newTestRegisterBookUseCase(t)
	getBookUseCase := newTestGetBookUseCase(t)
	updateBookUseCase := newTestUpdateBookUseCase(t)

	signUpReq := SignUpRequest{
		Name:     testdata.RandomString(10),
		Email:    testdata.RandomEmail(),
		Password: testdata.RandomString(16),
	}
	signUpRes, err := signUpUseCase.SignUp(context.Background(), signUpReq)
	require.NoError(t, err)
	registerReq := RegisterBookRequest{
		UserID: signUpRes.UserID,
		Title:  testdata.RandomString(10),
		Genres: []string{testdata.GetGenres()[0]},
		Status: entity.Unread,
	}
	registerBookRes, err := registerBookUseCase.RegisterBook(context.Background(), registerReq)
	require.NoError(t, err)

	testCases := []struct {
		name  string
		setup func(t *testing.T) UpdateBookRequest
		check func(t *testing.T, res *UpdateBookResponse, err error)
	}{
		{
			name: "Update book success",
			setup: func(t *testing.T) UpdateBookRequest {
				return UpdateBookRequest{
					UserID:        signUpRes.UserID,
					BookID:        registerBookRes.Book.ID,
					Title:         "SampleTitle",
					Genres:        registerBookRes.Book.Genres,
					Description:   registerBookRes.Book.Description,
					CoverImageURL: registerBookRes.Book.CoverImageURL,
					URL:           registerBookRes.Book.URL,
					Author:        registerBookRes.Book.AuthorName,
					Publisher:     registerBookRes.Book.PublisherName,
					PublishedDate: registerBookRes.Book.PublishDate,
					ISBN:          registerBookRes.Book.ISBN,
					Status:        entity.Reading,
					StartDate:     &entity.Date{Year: 2025, Month: 1, Day: 1},
					EndDate:       registerBookRes.Book.EndDate,
				}
			},
			check: func(t *testing.T, res *UpdateBookResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				getBookReq := GetBookRequest{
					UserID: signUpRes.UserID,
					BookID: res.BookID,
				}
				getBookRes, e := getBookUseCase.GetBook(context.Background(), getBookReq)
				require.NoError(t, e)
				require.Equal(t, "SampleTitle", getBookRes.Book.Title)
				require.Equal(t, registerBookRes.Book.Genres, getBookRes.Book.Genres)
				require.Equal(t, registerBookRes.Book.Description, getBookRes.Book.Description)
				require.Equal(t, registerBookRes.Book.CoverImageURL, getBookRes.Book.CoverImageURL)
				require.Equal(t, registerBookRes.Book.URL, getBookRes.Book.URL)
				require.Equal(t, registerBookRes.Book.AuthorName, getBookRes.Book.AuthorName)
				require.Equal(t, registerBookRes.Book.PublisherName, getBookRes.Book.PublisherName)
				require.Equal(t, registerBookRes.Book.PublishDate, getBookRes.Book.PublishDate)
				require.Equal(t, registerBookRes.Book.ISBN, getBookRes.Book.ISBN)
				require.Equal(t, entity.Reading, getBookRes.Book.Status)
				require.Equal(t, int32(2025), getBookRes.Book.StartDate.Year)
				require.Equal(t, int32(1), getBookRes.Book.StartDate.Month)
				require.Equal(t, int32(1), getBookRes.Book.StartDate.Day)
				require.Equal(t, registerBookRes.Book.EndDate, getBookRes.Book.EndDate)
			},
		},
		{
			name: "Update book failure when not registered book",
			setup: func(t *testing.T) UpdateBookRequest {
				return UpdateBookRequest{
					UserID:        signUpRes.UserID,
					BookID:        99999999,
					Title:         "SampleTitle",
					Genres:        registerBookRes.Book.Genres,
					Description:   registerBookRes.Book.Description,
					CoverImageURL: registerBookRes.Book.CoverImageURL,
					URL:           registerBookRes.Book.URL,
					Author:        registerBookRes.Book.AuthorName,
					Publisher:     registerBookRes.Book.PublisherName,
					PublishedDate: registerBookRes.Book.PublishDate,
					ISBN:          registerBookRes.Book.ISBN,
					Status:        entity.Reading,
					StartDate:     &entity.Date{Year: 2025, Month: 1, Day: 1},
					EndDate:       registerBookRes.Book.EndDate,
				}
			},
			check: func(t *testing.T, res *UpdateBookResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, BadRequest, e.StatusCode)
				require.Equal(t, InvalidRequestError, e.ErrorCode)
			},
		},
		{
			name: "Update book failure when another user registered book",
			setup: func(t *testing.T) UpdateBookRequest {
				return UpdateBookRequest{
					UserID:        99999999,
					BookID:        registerBookRes.Book.ID,
					Title:         "SampleTitle",
					Genres:        registerBookRes.Book.Genres,
					Description:   registerBookRes.Book.Description,
					CoverImageURL: registerBookRes.Book.CoverImageURL,
					URL:           registerBookRes.Book.URL,
					Author:        registerBookRes.Book.AuthorName,
					Publisher:     registerBookRes.Book.PublisherName,
					PublishedDate: registerBookRes.Book.PublishDate,
					ISBN:          registerBookRes.Book.ISBN,
					Status:        entity.Reading,
					StartDate:     &entity.Date{Year: 2025, Month: 1, Day: 1},
					EndDate:       registerBookRes.Book.EndDate,
				}
			},
			check: func(t *testing.T, res *UpdateBookResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, BadRequest, e.StatusCode)
				require.Equal(t, InvalidRequestError, e.ErrorCode)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := tc.setup(t)
			res, err := updateBookUseCase.UpdateBook(context.Background(), req)
			tc.check(t, res, err)
		})
	}
}
