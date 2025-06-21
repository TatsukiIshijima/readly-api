package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"log"
	"readly/testdata"
	"testing"
)

func createTestBookGenre(t *testing.T, book Book, genre Genre) {
	arg := CreateBookGenreParams{
		BookID:    book.ID,
		GenreName: genre.Name,
	}
	bookGenre, err := querier.CreateBookGenre(context.Background(), arg)
	log.Printf("create book genre: %+v", bookGenre)
	require.NoError(t, err)
	require.NotEmpty(t, bookGenre)
	require.NotZero(t, bookGenre.BookID)
	require.NotEmpty(t, bookGenre.GenreName)
	require.Equal(t, book.ID, bookGenre.BookID)
	require.Equal(t, genre.Name, bookGenre.GenreName)
}

func TestCreateBookGenre(t *testing.T) {
	book := createTestBook(
		t,
		testdata.RandomString(6),
		"",
		"",
		testdata.RandomString(13),
	)
	genre := createGenreIfNeed(t)
	createTestBookGenre(t, book, genre)
}

func TestDeleteBookGenre(t *testing.T) {
	book := createTestBook(
		t,
		testdata.RandomString(6),
		"",
		"",
		testdata.RandomString(13),
	)
	createRandomGenre := func() Genre {
		genre, err := querier.CreateGenre(context.Background(), testdata.RandomString(6))
		require.NoError(t, err)
		require.NotEmpty(t, genre)
		return genre
	}
	deleteGenre := func(genre Genre) {
		err := querier.DeleteGenre(context.Background(), genre.Name)
		require.NoError(t, err)
	}

	genre := createRandomGenre()
	createTestBookGenre(t, book, genre)

	genres, err := querier.GetGenresByBookID(context.Background(), book.ID)
	require.NoError(t, err)
	require.Len(t, genres, 1)

	deleteArgs := DeleteBookGenreParams{
		BookID:    book.ID,
		GenreName: genre.Name,
	}
	rowsAffected, err := querier.DeleteBookGenre(context.Background(), deleteArgs)
	require.NoError(t, err)
	require.Equal(t, int64(1), rowsAffected)

	genres, err = querier.GetGenresByBookID(context.Background(), book.ID)
	require.NoError(t, err)
	require.Len(t, genres, 0)

	deleteGenre(genre)
}
