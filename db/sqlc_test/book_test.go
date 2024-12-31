package sqlc_test

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"readly/db/sqlc"
	"readly/test"
	"testing"
	"time"
)

func createRandomBook(t *testing.T) db.Book {
	author := createRandomAuthor(t)
	publisher := createRandomPublisher(t)
	arg := db.CreateBookParams{
		Title: sql.NullString{
			String: test.RandomString(6),
			Valid:  true,
		},
		Description: sql.NullString{
			String: test.RandomString(12),
			Valid:  true,
		},
		CoverImageUrl: sql.NullString{
			String: "https://example.com",
			Valid:  true,
		},
		Url: sql.NullString{
			String: "https://example.com",
			Valid:  true,
		},
		AuthorName:    author.Name,
		PublisherName: publisher.Name,
		PublishedDate: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		Isbn: sql.NullString{
			String: test.RandomString(13),
			Valid:  true,
		},
	}

	book, err := test.Queries.CreateBook(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, book)

	require.NotZero(t, book.ID)
	require.Equal(t, arg.Title, book.Title)
	require.Equal(t, arg.Description, book.Description)
	require.Equal(t, arg.CoverImageUrl, book.CoverImageUrl)
	require.Equal(t, arg.Url, book.Url)
	require.Equal(t, arg.AuthorName, book.AuthorName)
	require.Equal(t, arg.PublisherName, book.PublisherName)
	require.WithinDuration(t, arg.PublishedDate.Time, book.PublishedDate.Time, time.Second)
	require.Equal(t, arg.Isbn, book.Isbn)

	return book
}

func checkSameBook(t *testing.T, book1 db.Book, book2 db.Book) {
	require.Equal(t, book1.ID, book2.ID)
	require.Equal(t, book1.Title, book2.Title)
	require.Equal(t, book1.Description, book2.Description)
	require.Equal(t, book1.CoverImageUrl, book2.CoverImageUrl)
	require.Equal(t, book1.Url, book2.Url)
	require.Equal(t, book1.AuthorName, book2.AuthorName)
	require.Equal(t, book1.PublisherName, book2.PublisherName)
	require.WithinDuration(t, book1.PublishedDate.Time, book2.PublishedDate.Time, time.Second)
	require.Equal(t, book1.Isbn, book2.Isbn)
}

func TestCreateBook(t *testing.T) {
	createRandomBook(t)
}

func TestGetBookById(t *testing.T) {
	book1 := createRandomBook(t)
	book2, err := test.Queries.GetBookById(context.Background(), book1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, book2)

	checkSameBook(t, book1, book2)
}

func TestGetBooksByAuthorName(t *testing.T) {
	author := createRandomAuthor(t)
	for i := 0; i < 2; i++ {
		createRandomBook(t)
	}

	books, err := test.Queries.GetBooksByAuthorName(context.Background(), author.Name)
	require.NoError(t, err)

	for _, book := range books {
		require.NotEmpty(t, book)
		require.Equal(t, author.Name, book.AuthorName)
	}
}

func TestGetBooksByIsbn(t *testing.T) {
	book1 := createRandomBook(t)
	books, err := test.Queries.GetBooksByIsbn(context.Background(), book1.Isbn)
	require.NoError(t, err)
	require.Equal(t, len(books), 1)
	book2 := books[0]

	checkSameBook(t, book1, book2)
}

func TestGetBooksByTitle(t *testing.T) {
	book1 := createRandomBook(t)
	books, err := test.Queries.GetBooksByTitle(context.Background(), book1.Title)
	require.NoError(t, err)
	require.Equal(t, len(books), 1)
	book2 := books[0]

	checkSameBook(t, book1, book2)
}

func TestUpdateBook(t *testing.T) {
	book1 := createRandomBook(t)

	arg := db.UpdateBookParams{
		ID:            book1.ID,
		Title:         book1.Title,
		Description:   sql.NullString{String: test.RandomString(12), Valid: true},
		CoverImageUrl: book1.CoverImageUrl,
		Url:           book1.Url,
		AuthorName:    book1.AuthorName,
		PublisherName: book1.PublisherName,
		PublishedDate: book1.PublishedDate,
		Isbn:          book1.Isbn,
	}

	book2, err := test.Queries.UpdateBook(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, book2)

	require.Equal(t, book1.ID, book2.ID)
	require.Equal(t, book1.Title, book2.Title)
	require.Equal(t, arg.Description, book2.Description)
	require.Equal(t, book1.CoverImageUrl, book2.CoverImageUrl)
	require.Equal(t, book1.Url, book2.Url)
	require.Equal(t, book1.AuthorName, book2.AuthorName)
	require.Equal(t, book1.PublisherName, book2.PublisherName)
	require.Equal(t, book1.PublishedDate.Time, book2.PublishedDate.Time)
	require.Equal(t, book1.Isbn, book2.Isbn)
}

func TestDeleteBook(t *testing.T) {
	book1 := createRandomBook(t)
	err := test.Queries.DeleteBook(context.Background(), book1.ID)
	require.NoError(t, err)

	book2, err := test.Queries.GetBookById(context.Background(), book1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, book2)
}
