package usecase

import (
	"context"
	"github.com/stretchr/testify/require"
	"readly/entity"
	"readly/testdata"
	"testing"
)

func TestDeleteBook(t *testing.T) {
	signUpUseCase := newTestSignUpUseCase(t)
	registerBookUseCase := newTestRegisterBookUseCase(t)
	deleteBookUseCase := newTestDeleteBookUseCase(t)

	signUpReq := SignUpRequest{
		Name:     testdata.RandomString(10),
		Email:    testdata.RandomEmail(),
		Password: testdata.RandomString(16),
	}
	signUpRes, err := signUpUseCase.SignUp(context.Background(), signUpReq)
	require.NoError(t, err)

	testCases := []struct {
		name  string
		setup func(t *testing.T) DeleteBookRequest
		check func(t *testing.T, err error)
	}{
		{
			name: "Delete created book success",
			setup: func(t *testing.T) DeleteBookRequest {
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

				registerReq := RegisterBookRequest{
					UserID:        signUpRes.UserID,
					Title:         testdata.RandomString(10),
					Genres:        []string{GetGenres()[0]},
					Description:   &desc,
					CoverImageURL: &coverImgURL,
					URL:           &url,
					AuthorName:    &author,
					PublisherName: &publisher,
					PublishDate:   &publishDate,
					ISBN:          &ISBN,
					Status:        0,
					StartDate:     &startDate,
					EndDate:       &endDate,
				}
				res, err := registerBookUseCase.RegisterBook(context.Background(), registerReq)
				require.NoError(t, err)

				return DeleteBookRequest{
					UserID: signUpRes.UserID,
					BookID: res.Book.ID,
				}
			},
			check: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		{
			name: "Delete not exist book failed",
			setup: func(t *testing.T) DeleteBookRequest {
				return DeleteBookRequest{
					UserID: signUpRes.UserID,
					BookID: 0,
				}
			},
			check: func(t *testing.T, err error) {
				require.Error(t, err)
				var e *Error
				require.ErrorAs(t, err, &e)
				require.Equal(t, e.StatusCode, BadRequest)
				require.Equal(t, e.ErrorCode, NotFoundBookError)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := tc.setup(t)
			err := deleteBookUseCase.DeleteBook(context.Background(), req)
			tc.check(t, err)
		})
	}
}

// TODO: Goroutineを使ったテストケースを追加する
