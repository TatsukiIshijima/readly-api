package repository

import (
	"context"
	"github.com/stretchr/testify/require"
	db "readly/db/sqlc"
	"readly/test"
	"testing"
	"time"
)

func TestRegister(t *testing.T) {
	user, err := test.CreateRandomUser()
	require.NoError(t, err)

	n := 3
	results := make(chan BookResponse)
	errs := make(chan error)

	for i := 0; i < n; i++ {
		go func() {
			// TODO:チェネルでジャンルを増やす&共有
			genres := make([]string, i+1)
			for j := 0; j <= i; j++ {
				genres[j] = test.RandomString(6)
			}
			arg := RegisterRequest{
				UserID:        user.ID,
				Title:         test.RandomString(6),
				Genres:        genres,
				Description:   test.RandomString(12),
				CoverImageURL: "https://example.com",
				URL:           "https://example.com",
				AuthorName:    test.RandomString(6),
				PublisherName: test.RandomString(6),
				PublishDate:   time.Now(),
				ISBN:          test.RandomString(13),
			}
			result, err := repo.Register(context.Background(), arg)
			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results

		author, err := store.GetAuthorByName(context.Background(), result.AuthorName)
		require.NoError(t, err)
		require.NotEmpty(t, author)
		require.Equal(t, result.AuthorName, author.Name)

		publisher, err := store.GetPublisherByName(context.Background(), result.PublisherName)
		require.NoError(t, err)
		require.NotEmpty(t, publisher)
		require.Equal(t, result.PublisherName, publisher.Name)

		genres, err := store.GetGenresByBookID(context.Background(), result.ID)
		require.NoError(t, err)
		require.Equal(t, len(result.Genres), len(genres))
		for _, g := range genres {
			genre, err := store.GetGenreByName(context.Background(), g)
			require.NoError(t, err)
			require.NotEmpty(t, genre)
		}

		book, err := store.GetBookById(context.Background(), result.ID)
		require.NoError(t, err)
		require.NotEmpty(t, book)
		require.Equal(t, result.Title, book.Title.String)
		require.Equal(t, result.Description, book.Description.String)
		require.Equal(t, result.CoverImageURL, book.CoverImageUrl.String)
		require.Equal(t, result.URL, book.Url.String)
		require.Equal(t, result.AuthorName, book.AuthorName)
		require.Equal(t, result.PublisherName, book.PublisherName)
		require.WithinDuration(t, result.PublishDate, book.PublishedDate.Time.UTC(), time.Second)
		require.Equal(t, result.ISBN, book.Isbn.String)

		param := db.GetReadingHistoryByUserIDParams{
			UserID: user.ID,
			Limit:  10,
			Offset: 0,
		}
		histories, err := store.GetReadingHistoryByUserID(context.Background(), param)
		require.NoError(t, err)
		require.Equal(t, n, len(histories))
		for _, h := range histories {
			require.Equal(t, user.ID, h.UserID)
			require.Equal(t, db.ReadingStatusUnread, h.Status)
		}
	}
}

func TestList(t *testing.T) {
	user, err := test.CreateRandomUser()
	require.NoError(t, err)

	n := 3
	for i := 0; i < n; i++ {
		registerReq := RegisterRequest{
			UserID:        user.ID,
			Title:         test.RandomString(6),
			Genres:        []string{test.RandomString(6)},
			Description:   test.RandomString(12),
			CoverImageURL: "https://example.com",
			URL:           "https://example.com",
			AuthorName:    test.RandomString(6),
			PublisherName: test.RandomString(6),
			PublishDate:   time.Now(),
			ISBN:          test.RandomString(13),
		}
		book, err := repo.Register(context.Background(), registerReq)
		require.NoError(t, err)
		require.NotEmpty(t, book)
	}

	listReq := ListRequest{
		UserID: user.ID,
		Limit:  int32(n),
		Offset: 0,
	}
	books, err := repo.List(context.Background(), listReq)
	require.NoError(t, err)
	require.Equal(t, n, len(books))
}
