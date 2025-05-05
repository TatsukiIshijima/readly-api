package usecase

import (
	"context"
	"github.com/stretchr/testify/require"
	"readly/entity"
	"readly/testdata"
	"testing"
)

func TestRegisterBook(t *testing.T) {
	signUpUseCase := newTestSignUpUseCase(t)
	registerBookUseCase := newTestRegisterBookUseCase(t)

	signUpReq := SignUpRequest{
		Name:     testdata.RandomString(10),
		Email:    testdata.RandomEmail(),
		Password: testdata.RandomString(16),
	}
	signUpRes, err := signUpUseCase.SignUp(context.Background(), signUpReq)
	require.NoError(t, err)

	testCases := []struct {
		name  string
		setup func(t *testing.T) RegisterBookRequest
		check func(t *testing.T, req RegisterBookRequest, res *RegisterBookResponse, err error)
	}{
		{
			name: "New unread book with required fields register success",
			setup: func(t *testing.T) RegisterBookRequest {
				return RegisterBookRequest{
					UserID: signUpRes.UserID,
					Title:  testdata.RandomString(10),
					Genres: []string{testdata.RandomString(6)},
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
				coverImgURL := testdata.RandomString(255)
				url := testdata.RandomString(255)
				author := testdata.RandomString(10)
				publisher := testdata.RandomString(10)
				publishDate := entity.Date{Year: 2018, Month: 12, Day: 31}
				require.NoError(t, err)
				ISBN := testdata.RandomString(13)
				startDate := entity.Date{Year: 2018, Month: 12, Day: 31}
				endDate := entity.Date{Year: 2019, Month: 1, Day: 30}

				return RegisterBookRequest{
					UserID:        signUpRes.UserID,
					Title:         testdata.RandomString(10),
					Genres:        []string{testdata.RandomString(6)},
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
			name: "New unread book register success when genres are already exist.",
			setup: func(t *testing.T) RegisterBookRequest {
				genre := testdata.RandomString(6)

				req := RegisterBookRequest{
					UserID: signUpRes.UserID,
					Title:  testdata.RandomString(10),
					Genres: []string{genre},
					Status: 0,
				}
				_, err := registerBookUseCase.RegisterBook(context.Background(), req)
				require.NoError(t, err)

				return RegisterBookRequest{
					UserID: signUpRes.UserID,
					Title:  testdata.RandomString(10),
					Genres: []string{genre},
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
			name: "New reading book with author & publisher register success when author & publisher are already exist.",
			setup: func(t *testing.T) RegisterBookRequest {
				author := testdata.RandomString(10)
				publisher := testdata.RandomString(10)
				startDate := entity.Now()

				req := RegisterBookRequest{
					UserID:        signUpRes.UserID,
					Title:         testdata.RandomString(10),
					Genres:        []string{testdata.RandomString(6)},
					AuthorName:    &author,
					PublisherName: &publisher,
					Status:        0,
				}
				_, err := registerBookUseCase.RegisterBook(context.Background(), req)
				require.NoError(t, err)

				return RegisterBookRequest{
					UserID:        signUpRes.UserID,
					Title:         testdata.RandomString(10),
					Genres:        []string{testdata.RandomString(6)},
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
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := tc.setup(t)
			res, err := registerBookUseCase.RegisterBook(context.Background(), req)
			tc.check(t, req, res, err)
		})
	}
}

func isSameDate(t1, t2 *entity.Date) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	return t1.Year == t2.Year && t1.Month == t2.Month && t1.Day == t2.Day
}

// TODO: Goroutineを使ったテストケースを追加する
