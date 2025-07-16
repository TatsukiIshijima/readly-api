package usecase

import (
	"context"
	"github.com/stretchr/testify/require"
	"readly/feature/book/domain"
	"readly/testdata"
	"testing"
)

func TestUpdateBook(t *testing.T) {
	registerBookUseCase := newTestRegisterBookUseCase(t)
	getBookUseCase := newTestGetBookUseCase(t)
	updateBookUseCase := newTestUpdateBookUseCase(t)

	user := createRandomUser(t)

	registerReq := RegisterBookRequest{
		UserID: user.ID,
		Title:  testdata.RandomString(10),
		Genres: []string{testdata.GetGenres()[0]},
		Status: domain.Unread,
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
					UserID:        user.ID,
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
					Status:        domain.Reading,
					StartDate:     &domain.Date{Year: 2025, Month: 1, Day: 1},
					EndDate:       registerBookRes.Book.EndDate,
				}
			},
			check: func(t *testing.T, res *UpdateBookResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				getBookReq := GetBookRequest{
					UserID: user.ID,
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
				require.Equal(t, domain.Reading, getBookRes.Book.Status)
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
					UserID:        user.ID,
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
					Status:        domain.Reading,
					StartDate:     &domain.Date{Year: 2025, Month: 1, Day: 1},
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
					Status:        domain.Reading,
					StartDate:     &domain.Date{Year: 2025, Month: 1, Day: 1},
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
			name: "Update book failed when title is empty",
			setup: func(t *testing.T) UpdateBookRequest {
				return UpdateBookRequest{
					UserID: user.ID,
					BookID: registerBookRes.Book.ID,
					Title:  "",
					Genres: registerBookRes.Book.Genres,
					Status: domain.Reading,
				}
			},
			check: func(t *testing.T, res *UpdateBookResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, BadRequest, e.StatusCode)
				require.Equal(t, InvalidRequestError, e.ErrorCode)
				require.Contains(t, e.Message, "title is required")
			},
		},
		{
			name: "Update book failed when title exceeds 255 characters",
			setup: func(t *testing.T) UpdateBookRequest {
				longTitle := testdata.RandomString(256)
				return UpdateBookRequest{
					UserID: user.ID,
					BookID: registerBookRes.Book.ID,
					Title:  longTitle,
					Genres: registerBookRes.Book.Genres,
					Status: domain.Reading,
				}
			},
			check: func(t *testing.T, res *UpdateBookResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, BadRequest, e.StatusCode)
				require.Equal(t, InvalidRequestError, e.ErrorCode)
				require.Contains(t, e.Message, "title must be between 1 and 255 characters")
			},
		},
		{
			name: "Update book failed when description exceeds 500 characters",
			setup: func(t *testing.T) UpdateBookRequest {
				longDesc := testdata.RandomString(501)
				return UpdateBookRequest{
					UserID:      user.ID,
					BookID:      registerBookRes.Book.ID,
					Title:       testdata.RandomString(10),
					Description: &longDesc,
					Genres:      registerBookRes.Book.Genres,
					Status:      domain.Reading,
				}
			},
			check: func(t *testing.T, res *UpdateBookResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, BadRequest, e.StatusCode)
				require.Equal(t, InvalidRequestError, e.ErrorCode)
				require.Contains(t, e.Message, "description must be less than 500 characters")
			},
		},
		{
			name: "Update book failed when URL has invalid format",
			setup: func(t *testing.T) UpdateBookRequest {
				invalidURL := "invalid-url"
				return UpdateBookRequest{
					UserID: user.ID,
					BookID: registerBookRes.Book.ID,
					Title:  testdata.RandomString(10),
					URL:    &invalidURL,
					Genres: registerBookRes.Book.Genres,
					Status: domain.Reading,
				}
			},
			check: func(t *testing.T, res *UpdateBookResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, BadRequest, e.StatusCode)
				require.Equal(t, InvalidRequestError, e.ErrorCode)
				require.Contains(t, e.Message, "URL has invalid format")
			},
		},
		{
			name: "Update book failed when URL exceeds 2048 characters",
			setup: func(t *testing.T) UpdateBookRequest {
				longURL := "https://example.com/" + testdata.RandomString(2030)
				return UpdateBookRequest{
					UserID: user.ID,
					BookID: registerBookRes.Book.ID,
					Title:  testdata.RandomString(10),
					URL:    &longURL,
					Genres: registerBookRes.Book.Genres,
					Status: domain.Reading,
				}
			},
			check: func(t *testing.T, res *UpdateBookResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, BadRequest, e.StatusCode)
				require.Equal(t, InvalidRequestError, e.ErrorCode)
				require.Contains(t, e.Message, "URL must be less than 2048 characters")
			},
		},
		{
			name: "Update book failed when cover image URL has invalid format",
			setup: func(t *testing.T) UpdateBookRequest {
				invalidURL := "not-a-url"
				return UpdateBookRequest{
					UserID:        user.ID,
					BookID:        registerBookRes.Book.ID,
					Title:         testdata.RandomString(10),
					CoverImageURL: &invalidURL,
					Genres:        registerBookRes.Book.Genres,
					Status:        domain.Reading,
				}
			},
			check: func(t *testing.T, res *UpdateBookResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, BadRequest, e.StatusCode)
				require.Equal(t, InvalidRequestError, e.ErrorCode)
				require.Contains(t, e.Message, "cover image URL has invalid format")
			},
		},
		{
			name: "Update book failed when author name exceeds 255 characters",
			setup: func(t *testing.T) UpdateBookRequest {
				longAuthor := testdata.RandomString(256)
				return UpdateBookRequest{
					UserID: user.ID,
					BookID: registerBookRes.Book.ID,
					Title:  testdata.RandomString(10),
					Author: &longAuthor,
					Genres: registerBookRes.Book.Genres,
					Status: domain.Reading,
				}
			},
			check: func(t *testing.T, res *UpdateBookResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, BadRequest, e.StatusCode)
				require.Equal(t, InvalidRequestError, e.ErrorCode)
				require.Contains(t, e.Message, "author name must be less than 255 characters")
			},
		},
		{
			name: "Update book failed when publisher name exceeds 255 characters",
			setup: func(t *testing.T) UpdateBookRequest {
				longPublisher := testdata.RandomString(256)
				return UpdateBookRequest{
					UserID:    user.ID,
					BookID:    registerBookRes.Book.ID,
					Title:     testdata.RandomString(10),
					Publisher: &longPublisher,
					Genres:    registerBookRes.Book.Genres,
					Status:    domain.Reading,
				}
			},
			check: func(t *testing.T, res *UpdateBookResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, BadRequest, e.StatusCode)
				require.Equal(t, InvalidRequestError, e.ErrorCode)
				require.Contains(t, e.Message, "publisher name must be less than 255 characters")
			},
		},
		{
			name: "Update book failed when ISBN has invalid format",
			setup: func(t *testing.T) UpdateBookRequest {
				invalidISBN := "123abc"
				return UpdateBookRequest{
					UserID: user.ID,
					BookID: registerBookRes.Book.ID,
					Title:  testdata.RandomString(10),
					ISBN:   &invalidISBN,
					Genres: registerBookRes.Book.Genres,
					Status: domain.Reading,
				}
			},
			check: func(t *testing.T, res *UpdateBookResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, BadRequest, e.StatusCode)
				require.Equal(t, InvalidRequestError, e.ErrorCode)
				require.Contains(t, e.Message, "ISBN must be 13 digits")
			},
		},
		{
			name: "Update book failed when end date is before start date",
			setup: func(t *testing.T) UpdateBookRequest {
				startDate := domain.Date{Year: 2023, Month: 12, Day: 31}
				endDate := domain.Date{Year: 2023, Month: 1, Day: 1}
				return UpdateBookRequest{
					UserID:    user.ID,
					BookID:    registerBookRes.Book.ID,
					Title:     testdata.RandomString(10),
					Genres:    registerBookRes.Book.Genres,
					Status:    domain.Done,
					StartDate: &startDate,
					EndDate:   &endDate,
				}
			},
			check: func(t *testing.T, res *UpdateBookResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, BadRequest, e.StatusCode)
				require.Equal(t, InvalidRequestError, e.ErrorCode)
				require.Contains(t, e.Message, "end date must be after start date")
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
