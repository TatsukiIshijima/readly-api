package usecase

import (
	"context"
	"github.com/stretchr/testify/require"
	"readly/feature/book/domain"
	"readly/testdata"
	"testing"
)

func TestRegisterBook(t *testing.T) {
	registerBookUseCase := newTestRegisterBookUseCase(t)

	user := createRandomUser(t)

	testCases := []struct {
		name  string
		setup func(t *testing.T) RegisterBookRequest
		check func(t *testing.T, req RegisterBookRequest, res *RegisterBookResponse, err error)
	}{
		{
			name: "New unread book with required fields register success",
			setup: func(t *testing.T) RegisterBookRequest {
				return RegisterBookRequest{
					UserID: user.ID,
					Title:  testdata.RandomString(10),
					Genres: []string{testdata.GetGenres()[0]},
					Status: 0,
				}
			},
			check: func(t *testing.T, req RegisterBookRequest, res *RegisterBookResponse, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, res.Book.ID)
				require.Equal(t, req.Title, res.Book.Title)
				require.Equal(t, req.Genres, res.Book.Genres)
				require.Equal(t, req.Description, res.Book.Description)
				require.Equal(t, req.CoverImageURL, res.Book.CoverImageURL)
				require.Equal(t, req.URL, res.Book.URL)
				require.Equal(t, req.AuthorName, res.Book.AuthorName)
				require.Equal(t, req.PublisherName, res.Book.PublisherName)
				require.Equal(t, req.PublishDate, res.Book.PublishDate)
				require.Equal(t, req.ISBN, res.Book.ISBN)
				require.Equal(t, req.Status, res.Book.Status)
				require.True(t, isSameDate(req.StartDate, res.Book.StartDate))
				require.True(t, isSameDate(req.EndDate, res.Book.EndDate))
			},
		},
		{
			name: "New read book with all fields register success",
			setup: func(t *testing.T) RegisterBookRequest {
				desc := testdata.RandomString(100)
				coverImgURL := testdata.RandomURL()
				url := testdata.RandomURL()
				author := testdata.RandomString(10)
				publisher := testdata.RandomString(10)
				publishDate := domain.Date{Year: 2018, Month: 12, Day: 31}
				ISBN := testdata.RandomISBN()
				startDate := domain.Date{Year: 2018, Month: 12, Day: 31}
				endDate := domain.Date{Year: 2019, Month: 1, Day: 30}

				return RegisterBookRequest{
					UserID:        user.ID,
					Title:         testdata.RandomString(10),
					Genres:        []string{testdata.GetGenres()[1]},
					Description:   &desc,
					CoverImageURL: &coverImgURL,
					URL:           &url,
					AuthorName:    &author,
					PublisherName: &publisher,
					PublishDate:   &publishDate,
					ISBN:          &ISBN,
					Status:        2,
					StartDate:     &startDate,
					EndDate:       &endDate,
				}
			},
			check: func(t *testing.T, req RegisterBookRequest, res *RegisterBookResponse, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, res.Book.ID)
				require.Equal(t, req.Title, res.Book.Title)
				require.Equal(t, req.Genres, res.Book.Genres)
				require.Equal(t, req.Description, res.Book.Description)
				require.Equal(t, req.CoverImageURL, res.Book.CoverImageURL)
				require.Equal(t, req.URL, res.Book.URL)
				require.Equal(t, req.AuthorName, res.Book.AuthorName)
				require.Equal(t, req.PublisherName, res.Book.PublisherName)
				require.Equal(t, req.PublishDate, res.Book.PublishDate)
				require.Equal(t, req.ISBN, res.Book.ISBN)
				require.Equal(t, req.Status, res.Book.Status)
				require.True(t, isSameDate(req.StartDate, res.Book.StartDate))
				require.True(t, isSameDate(req.EndDate, res.Book.EndDate))
			},
		},
		{
			name: "New reading book with author & publisher register success when author & publisher are already exist.",
			setup: func(t *testing.T) RegisterBookRequest {
				author := testdata.RandomString(10)
				publisher := testdata.RandomString(10)
				startDate := domain.Now()

				req := RegisterBookRequest{
					UserID:        user.ID,
					Title:         testdata.RandomString(10),
					Genres:        []string{testdata.GetGenres()[1]},
					AuthorName:    &author,
					PublisherName: &publisher,
					Status:        0,
				}
				_, err := registerBookUseCase.RegisterBook(context.Background(), req)
				require.NoError(t, err)

				return RegisterBookRequest{
					UserID:        user.ID,
					Title:         testdata.RandomString(10),
					Genres:        []string{testdata.GetGenres()[2]},
					AuthorName:    &author,
					PublisherName: &publisher,
					Status:        1,
					StartDate:     &startDate,
				}
			},
			check: func(t *testing.T, req RegisterBookRequest, res *RegisterBookResponse, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, res.Book.ID)
				require.Equal(t, req.Title, res.Book.Title)
				require.Equal(t, req.Genres, res.Book.Genres)
				require.Equal(t, req.Description, res.Book.Description)
				require.Equal(t, req.CoverImageURL, res.Book.CoverImageURL)
				require.Equal(t, req.URL, res.Book.URL)
				require.Equal(t, req.AuthorName, res.Book.AuthorName)
				require.Equal(t, req.PublisherName, res.Book.PublisherName)
				require.Equal(t, req.PublishDate, res.Book.PublishDate)
				require.Equal(t, req.ISBN, res.Book.ISBN)
				require.Equal(t, req.Status, res.Book.Status)
				require.True(t, isSameDate(req.StartDate, res.Book.StartDate))
				require.True(t, isSameDate(req.EndDate, res.Book.EndDate))
			},
		},
		{
			name: "New unread book register failed when genres are not exist.",
			setup: func(t *testing.T) RegisterBookRequest {
				return RegisterBookRequest{
					UserID: user.ID,
					Title:  testdata.RandomString(10),
					Genres: []string{testdata.RandomString(6)},
					Status: 0,
				}
			},
			check: func(t *testing.T, req RegisterBookRequest, res *RegisterBookResponse, err error) {
				require.Error(t, err)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, e.StatusCode, BadRequest)
				require.Equal(t, e.ErrorCode, InvalidGenreError)
			},
		},
		{
			name: "Register book failed when title is empty",
			setup: func(t *testing.T) RegisterBookRequest {
				return RegisterBookRequest{
					UserID: user.ID,
					Title:  "",
					Genres: []string{testdata.GetGenres()[0]},
					Status: 0,
				}
			},
			check: func(t *testing.T, req RegisterBookRequest, res *RegisterBookResponse, err error) {
				require.Error(t, err)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, e.StatusCode, BadRequest)
				require.Equal(t, e.ErrorCode, InvalidRequestError)
				require.Contains(t, e.Message, "title is required")
			},
		},
		{
			name: "Register book failed when title exceeds 255 characters",
			setup: func(t *testing.T) RegisterBookRequest {
				longTitle := testdata.RandomString(256)
				return RegisterBookRequest{
					UserID: user.ID,
					Title:  longTitle,
					Genres: []string{testdata.GetGenres()[0]},
					Status: 0,
				}
			},
			check: func(t *testing.T, req RegisterBookRequest, res *RegisterBookResponse, err error) {
				require.Error(t, err)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, e.StatusCode, BadRequest)
				require.Equal(t, e.ErrorCode, InvalidRequestError)
				require.Contains(t, e.Message, "title must be between 1 and 255 characters")
			},
		},
		{
			name: "Register book failed when description exceeds 500 characters",
			setup: func(t *testing.T) RegisterBookRequest {
				longDesc := testdata.RandomString(501)
				return RegisterBookRequest{
					UserID:      user.ID,
					Title:       testdata.RandomString(10),
					Description: &longDesc,
					Genres:      []string{testdata.GetGenres()[0]},
					Status:      0,
				}
			},
			check: func(t *testing.T, req RegisterBookRequest, res *RegisterBookResponse, err error) {
				require.Error(t, err)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, e.StatusCode, BadRequest)
				require.Equal(t, e.ErrorCode, InvalidRequestError)
				require.Contains(t, e.Message, "description must be less than 500 characters")
			},
		},
		{
			name: "Register book failed when URL has invalid format",
			setup: func(t *testing.T) RegisterBookRequest {
				invalidURL := "invalid-url"
				return RegisterBookRequest{
					UserID: user.ID,
					Title:  testdata.RandomString(10),
					URL:    &invalidURL,
					Genres: []string{testdata.GetGenres()[0]},
					Status: 0,
				}
			},
			check: func(t *testing.T, req RegisterBookRequest, res *RegisterBookResponse, err error) {
				require.Error(t, err)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, e.StatusCode, BadRequest)
				require.Equal(t, e.ErrorCode, InvalidRequestError)
				require.Contains(t, e.Message, "URL has invalid format")
			},
		},
		{
			name: "Register book failed when URL exceeds 2048 characters",
			setup: func(t *testing.T) RegisterBookRequest {
				longURL := "https://example.com/" + testdata.RandomString(2030)
				return RegisterBookRequest{
					UserID: user.ID,
					Title:  testdata.RandomString(10),
					URL:    &longURL,
					Genres: []string{testdata.GetGenres()[0]},
					Status: 0,
				}
			},
			check: func(t *testing.T, req RegisterBookRequest, res *RegisterBookResponse, err error) {
				require.Error(t, err)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, e.StatusCode, BadRequest)
				require.Equal(t, e.ErrorCode, InvalidRequestError)
				require.Contains(t, e.Message, "URL must be less than 2048 characters")
			},
		},
		{
			name: "Register book failed when cover image URL has invalid format",
			setup: func(t *testing.T) RegisterBookRequest {
				invalidURL := "not-a-url"
				return RegisterBookRequest{
					UserID:        user.ID,
					Title:         testdata.RandomString(10),
					CoverImageURL: &invalidURL,
					Genres:        []string{testdata.GetGenres()[0]},
					Status:        0,
				}
			},
			check: func(t *testing.T, req RegisterBookRequest, res *RegisterBookResponse, err error) {
				require.Error(t, err)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, e.StatusCode, BadRequest)
				require.Equal(t, e.ErrorCode, InvalidRequestError)
				require.Contains(t, e.Message, "cover image URL has invalid format")
			},
		},
		{
			name: "Register book failed when author name exceeds 255 characters",
			setup: func(t *testing.T) RegisterBookRequest {
				longAuthor := testdata.RandomString(256)
				return RegisterBookRequest{
					UserID:     user.ID,
					Title:      testdata.RandomString(10),
					AuthorName: &longAuthor,
					Genres:     []string{testdata.GetGenres()[0]},
					Status:     0,
				}
			},
			check: func(t *testing.T, req RegisterBookRequest, res *RegisterBookResponse, err error) {
				require.Error(t, err)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, e.StatusCode, BadRequest)
				require.Equal(t, e.ErrorCode, InvalidRequestError)
				require.Contains(t, e.Message, "author name must be less than 255 characters")
			},
		},
		{
			name: "Register book failed when publisher name exceeds 255 characters",
			setup: func(t *testing.T) RegisterBookRequest {
				longPublisher := testdata.RandomString(256)
				return RegisterBookRequest{
					UserID:        user.ID,
					Title:         testdata.RandomString(10),
					PublisherName: &longPublisher,
					Genres:        []string{testdata.GetGenres()[0]},
					Status:        0,
				}
			},
			check: func(t *testing.T, req RegisterBookRequest, res *RegisterBookResponse, err error) {
				require.Error(t, err)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, e.StatusCode, BadRequest)
				require.Equal(t, e.ErrorCode, InvalidRequestError)
				require.Contains(t, e.Message, "publisher name must be less than 255 characters")
			},
		},
		{
			name: "Register book failed when ISBN has invalid format",
			setup: func(t *testing.T) RegisterBookRequest {
				invalidISBN := "123abc"
				return RegisterBookRequest{
					UserID: user.ID,
					Title:  testdata.RandomString(10),
					ISBN:   &invalidISBN,
					Genres: []string{testdata.GetGenres()[0]},
					Status: 0,
				}
			},
			check: func(t *testing.T, req RegisterBookRequest, res *RegisterBookResponse, err error) {
				require.Error(t, err)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, e.StatusCode, BadRequest)
				require.Equal(t, e.ErrorCode, InvalidRequestError)
				require.Contains(t, e.Message, "ISBN must be 13 digits")
			},
		},
		{
			name: "Register book failed when end date is before start date",
			setup: func(t *testing.T) RegisterBookRequest {
				startDate := domain.Date{Year: 2023, Month: 12, Day: 31}
				endDate := domain.Date{Year: 2023, Month: 1, Day: 1}
				return RegisterBookRequest{
					UserID:    user.ID,
					Title:     testdata.RandomString(10),
					Genres:    []string{testdata.GetGenres()[0]},
					Status:    2,
					StartDate: &startDate,
					EndDate:   &endDate,
				}
			},
			check: func(t *testing.T, req RegisterBookRequest, res *RegisterBookResponse, err error) {
				require.Error(t, err)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, e.StatusCode, BadRequest)
				require.Equal(t, e.ErrorCode, InvalidRequestError)
				require.Contains(t, e.Message, "end date must be after start date")
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := tc.setup(t)
			res, err := registerBookUseCase.RegisterBook(context.Background(), req)
			tc.check(t, req, res, err)
		})
	}
}

func isSameDate(t1, t2 *domain.Date) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	return t1.Year == t2.Year && t1.Month == t2.Month && t1.Day == t2.Day
}

// TODO: Goroutineを使ったテストケースを追加する
