package usecase

import (
	"context"
	"github.com/stretchr/testify/require"
	"readly/entity"
	"readly/repository"
	"readly/testdata"
	"testing"
	"time"
)

func TestRegisterBook(t *testing.T) {
	createUserReq := repository.CreateUserRequest{
		Name:     testdata.RandomString(10),
		Email:    testdata.RandomEmail(),
		Password: testdata.RandomString(16),
	}
	user, err := userRepo.CreateUser(context.Background(), createUserReq)
	require.NoError(t, err)

	title := testdata.RandomString(10)
	genre1 := testdata.RandomString(6)
	genre2 := testdata.RandomString(6)
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
		name string
		req  RegisterBookRequest
		exp  entity.Book
	}{
		{
			name: "New unread book with required fields register success",
			req: RegisterBookRequest{
				UserID: user.ID,
				Title:  title,
				Genres: []string{genre1},
				Status: 0,
			},
			exp: entity.Book{
				Title:  title,
				Genres: []string{genre1},
				Status: 0,
			},
		},
		{
			name: "New unread book with all fields register success",
			req: RegisterBookRequest{
				UserID:        user.ID,
				Title:         title,
				Genres:        []string{genre2},
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
			},
			exp: entity.Book{
				Title:         title,
				Genres:        []string{genre2},
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
			},
		},
		{
			name: "New unread book register success when genres are already exist.",
			req: RegisterBookRequest{
				UserID: user.ID,
				Title:  title,
				Genres: []string{genre1, genre2},
				Status: 0,
			},
			exp: entity.Book{
				Title:  title,
				Genres: []string{genre1, genre2},
				Status: 0,
			},
		},
		{
			name: "New read book with author & publisher register success",
			req: RegisterBookRequest{
				UserID:        user.ID,
				Title:         title,
				AuthorName:    &author,
				PublisherName: &publisher,
				Status:        1,
				StartDate:     &startDate,
			},
			exp: entity.Book{
				Title:         title,
				AuthorName:    &author,
				PublisherName: &publisher,
				Status:        1,
				StartDate:     &startDate,
			},
		},
		{
			name: "New read book with author & publisher register success when author & publisher are already exist.",
			req: RegisterBookRequest{
				UserID:        user.ID,
				Title:         title,
				AuthorName:    &author,
				PublisherName: &publisher,
				Status:        1,
				StartDate:     &startDate,
			},
			exp: entity.Book{
				Title:         title,
				AuthorName:    &author,
				PublisherName: &publisher,
				Status:        1,
				StartDate:     &startDate,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := registerBookUseCase.RegisterBook(context.Background(), tc.req)
			require.NoError(t, err)
			// IDは自動採番なので比較対象から外す
			require.Equal(t, tc.exp.Title, res.Title)
			require.Equal(t, tc.exp.Genres, res.Genres)
			require.Equal(t, tc.exp.Description, res.Description)
			require.Equal(t, tc.exp.CoverImageURL, res.CoverImageURL)
			require.Equal(t, tc.exp.URL, res.URL)
			require.Equal(t, tc.exp.AuthorName, res.AuthorName)
			require.Equal(t, tc.exp.PublisherName, res.PublisherName)
			require.Equal(t, tc.exp.PublishDate, res.PublishDate)
			require.Equal(t, tc.exp.ISBN, res.ISBN)
			require.Equal(t, tc.exp.Status, res.Status)
			require.True(t, isSameDate(tc.exp.StartDate, res.StartDate))
			require.True(t, isSameDate(tc.exp.EndDate, res.EndDate))
		})
	}
}

func isSameDate(t1, t2 *time.Time) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	return t1.Year() == t2.Year() && t1.Month() == t2.Month() && t1.Day() == t2.Day()
}
