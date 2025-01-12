package sqlc_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"readly/db/sqlc"
	"testing"
)

func createRandomBookGenre(t *testing.T, book db.Book, genre db.Genre) {
	arg := db.CreateBookGenreParams{
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
	book := createBook(t, "", "", "", "")
	genre := createRandomGenre(t)
	createRandomBookGenre(t, book, genre)
}

func TestDeleteBookGenre(t *testing.T) {
	book := createBook(t, "", "", "", "")
	genre1 := createRandomGenre(t)
	genre2 := createRandomGenre(t)
	genre3 := createRandomGenre(t)
	createRandomBookGenre(t, book, genre1)
	createRandomBookGenre(t, book, genre2)
	createRandomBookGenre(t, book, genre3)

	genres, err := querier.GetGenresByBookID(context.Background(), book.ID)
	require.NoError(t, err)
	require.Len(t, genres, 3)

	deleteArgs := db.DeleteBookGenreParams{
		BookID:    book.ID,
		GenreName: genre1.Name,
	}
	err = querier.DeleteBookGenre(context.Background(), deleteArgs)
	require.NoError(t, err)

	genres, err = querier.GetGenresByBookID(context.Background(), book.ID)
	require.NoError(t, err)
	require.Len(t, genres, 2)
}
