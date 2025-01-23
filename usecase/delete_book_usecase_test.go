package usecase

import (
	"context"
	"github.com/stretchr/testify/require"
	"readly/repository"
	"readly/testdata"
	"testing"
	"time"
)

func TestDeleteBook(t *testing.T) {
	createUserReq := repository.CreateUserRequest{
		Name:     testdata.RandomString(10),
		Email:    testdata.RandomEmail(),
		Password: testdata.RandomString(16),
	}
	user, err := userRepo.CreateUser(context.Background(), createUserReq)
	require.NoError(t, err)

	title := testdata.RandomString(10)
	genre := testdata.RandomString(6)
	desc := testdata.RandomString(100)
	coverImgURL := testdata.RandomString(255)
	url := testdata.RandomString(255)
	author := testdata.RandomString(10)
	publisher := testdata.RandomString(10)
	publishDate := testdata.TimeFrom("1970-01-01 00:00:00").UTC()
	ISBN := testdata.RandomString(13)
	startDate := time.Now().UTC()
	endDate := time.Now().Add(time.Duration(60*60*24) * time.Second).UTC()

	testCases := []struct {
		name        string
		registerReq *RegisterBookRequest
		deleteReq   DeleteBookRequest
		err         error
	}{
		{
			name: "Delete created book success",
			registerReq: &RegisterBookRequest{
				UserID:        user.ID,
				Title:         title,
				Genres:        []string{genre},
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
			},
			deleteReq: DeleteBookRequest{
				UserID: user.ID,
				BookID: 0,
			},
			err: nil,
		},
		{
			name:        "Delete not exist book failed",
			registerReq: nil,
			deleteReq: DeleteBookRequest{
				UserID: user.ID,
				BookID: 0,
			},
			err: repository.ErrNoRowsDeleted,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.registerReq != nil {
				b, err := registerBookUseCase.RegisterBook(context.Background(), *tc.registerReq)
				require.NoError(t, err)
				tc.deleteReq.BookID = b.ID
			}
			err = deleteBookUseCase.DeleteBook(context.Background(), tc.deleteReq)
			require.Equal(t, tc.err, err)
		})
	}
}

// TODO: Goroutineを使ったテストケースを追加する
