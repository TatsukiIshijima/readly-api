package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"readly/testdata"
	"testing"
)

func createRandomBookGenre(t *testing.T, book Book, genre Genre) {
	arg := CreateBookGenreParams{
		BookID:    book.ID,
		GenreName: genre.Name,
	}
	bookGenre, err := querier.CreateBookGenre(context.Background(), arg)
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
	genre := createRandomGenre(t)
	createRandomBookGenre(t, book, genre)
}

func TestDeleteBookGenre(t *testing.T) {
	book := createTestBook(
		t,
		testdata.RandomString(6),
		"",
		"",
		testdata.RandomString(13),
	)
	genre1 := createRandomGenre(t)
	genre2 := createRandomGenre(t)
	genre3 := createRandomGenre(t)
	createRandomBookGenre(t, book, genre1)
	createRandomBookGenre(t, book, genre2)
	createRandomBookGenre(t, book, genre3)

	genres, err := querier.GetGenresByBookID(context.Background(), book.ID)
	require.NoError(t, err)
	require.Len(t, genres, 3)

	deleteArgs := DeleteBookGenreParams{
		BookID:    book.ID,
		GenreName: genre1.Name,
	}
	rowsAffected, err := querier.DeleteBookGenre(context.Background(), deleteArgs)
	require.NoError(t, err)
	require.Equal(t, int64(1), rowsAffected)

	genres, err = querier.GetGenresByBookID(context.Background(), book.ID)
	require.NoError(t, err)
	require.Len(t, genres, 2)
}
