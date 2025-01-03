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
	book := createRandomBook(t)
	genre := createRandomGenre(t)
	createRandomBookGenre(t, book, genre)
}

func TestGetBooksByGenreName(t *testing.T) {
	book1 := createRandomBook(t)
	book2 := createRandomBook(t)
	genre := createRandomGenre(t)
	createRandomBookGenre(t, book1, genre)
	createRandomBookGenre(t, book2, genre)

	args := db.GetBooksByGenreNameParams{
		GenreName: genre.Name,
		Limit:     5,
		Offset:    0,
	}

	books, err := querier.GetBooksByGenreName(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, books)
	require.Len(t, books, 2)
}

func TestGetGenresByBookID(t *testing.T) {
	book := createRandomBook(t)
	genre1 := createRandomGenre(t)
	genre2 := createRandomGenre(t)
	createRandomBookGenre(t, book, genre1)
	createRandomBookGenre(t, book, genre2)

	genres, err := querier.GetGenresByBookID(context.Background(), book.ID)
	require.NoError(t, err)
	require.NotEmpty(t, genres)
	require.Len(t, genres, 2)
}

func TestDeleteBookGenre(t *testing.T) {
	book := createRandomBook(t)
	genre := createRandomGenre(t)
	createRandomBookGenre(t, book, genre)

	arg1 := db.DeleteGenreForBookParams{
		BookID:    book.ID,
		GenreName: genre.Name,
	}
	err := querier.DeleteGenreForBook(context.Background(), arg1)
	require.NoError(t, err)

	arg2 := db.GetBooksByGenreNameParams{
		GenreName: genre.Name,
		Limit:     5,
		Offset:    0,
	}

	books, err := querier.GetBooksByGenreName(context.Background(), arg2)
	require.NoError(t, err)
	require.Len(t, books, 0)

	genres, err := querier.GetGenresByBookID(context.Background(), book.ID)
	require.NoError(t, err)
	require.Len(t, genres, 0)
}
