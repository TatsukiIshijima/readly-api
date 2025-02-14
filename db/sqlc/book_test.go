package db

import (
	"context"
	"database/sql"
	"errors"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"readly/testdata"
	"sort"
	"strings"
	"testing"
	"time"
)

func createTestBook(t *testing.T, title string, author string, publisher string, isbn string) Book {
	desc := sql.NullString{
		String: testdata.RandomString(12),
		Valid:  true,
	}
	imgURL := sql.NullString{
		String: "https://example.com",
		Valid:  true,
	}
	URL := sql.NullString{
		String: "https://example.com",
		Valid:  true,
	}
	pubDate := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	ISBN := sql.NullString{
		String: isbn,
		Valid:  isbn != "",
	}

	auth, err := createAuthorIfNeed(t, author)
	require.NoError(t, err)
	pub, err := createPublisherIfNeed(t, publisher)
	require.NoError(t, err)

	arg := CreateBookParams{
		Title:         title,
		Description:   desc,
		CoverImageUrl: imgURL,
		Url:           URL,
		AuthorName:    auth,
		PublisherName: pub,
		PublishedDate: pubDate,
		Isbn:          ISBN,
	}

	book, err := querier.CreateBook(context.Background(), arg)
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

func createAuthorIfNeed(t *testing.T, author string) (sql.NullString, error) {
	auth := sql.NullString{
		String: author,
		Valid:  author != "",
	}
	if !auth.Valid {
		return auth, nil
	}
	_, err := querier.CreateAuthor(context.Background(), author)
	if err != nil {
		checkDuplicateKeyError(t, err)
		return auth, nil
	}
	ga, err := querier.GetAuthorByName(context.Background(), author)
	require.NoError(t, err)
	require.Equal(t, author, ga.Name)
	return sql.NullString{String: ga.Name, Valid: true}, nil
}

func createPublisherIfNeed(t *testing.T, publisher string) (sql.NullString, error) {
	pub := sql.NullString{
		String: publisher,
		Valid:  publisher != "",
	}
	if !pub.Valid {
		return pub, nil
	}
	_, err := querier.CreatePublisher(context.Background(), publisher)
	if err != nil {
		checkDuplicateKeyError(t, err)
		return pub, nil
	}
	dp, err := querier.GetPublisherByName(context.Background(), publisher)
	require.NoError(t, err)
	require.Equal(t, publisher, dp.Name)
	return sql.NullString{String: publisher, Valid: true}, nil
}

func checkDuplicateKeyError(t *testing.T, err error) {
	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		if pqErr.Code != "23505" {
			require.Fail(t, "unexpected error: %v", err)
		}
	} else {
		require.Fail(t, "unexpected error: %v", err)
	}
}

func TestCreateBook(t *testing.T) {
	createTestBook(
		t,
		testdata.RandomString(6),
		"",
		"",
		testdata.RandomString(13),
	)
}

func TestGetBookByID(t *testing.T) {
	bookWithEmptyGenres := createTestBook(
		t,
		testdata.RandomString(6),
		"",
		"",
		testdata.RandomString(13),
	)
	bookWithGenres := createTestBook(
		t,
		testdata.RandomString(6),
		testdata.RandomString(8),
		testdata.RandomString(10),
		testdata.RandomString(13),
	)
	genre1 := createRandomGenre(t)
	genre2 := createRandomGenre(t)
	genre3 := createRandomGenre(t)
	createRandomBookGenre(t, bookWithGenres, genre1)
	createRandomBookGenre(t, bookWithGenres, genre2)
	createRandomBookGenre(t, bookWithGenres, genre3)

	result, err := querier.GetBooksByID(context.Background(), bookWithEmptyGenres.ID)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, bookWithEmptyGenres.ID, result.ID)
	require.Equal(t, bookWithEmptyGenres.Title, result.Title)
	require.Empty(t, result.Genres)
	require.Equal(t, bookWithEmptyGenres.Description, result.Description)
	require.Equal(t, bookWithEmptyGenres.CoverImageUrl, result.CoverImageUrl)
	require.Equal(t, bookWithEmptyGenres.Url, result.Url)
	require.Equal(t, bookWithEmptyGenres.AuthorName, result.AuthorName)
	require.Equal(t, bookWithEmptyGenres.PublisherName, result.PublisherName)
	require.Equal(t, bookWithEmptyGenres.PublishedDate.Time, result.PublishedDate.Time)
	require.Equal(t, bookWithEmptyGenres.Isbn, result.Isbn)
	require.WithinDuration(t, bookWithEmptyGenres.CreatedAt, result.CreatedAt, time.Second)
	require.WithinDuration(t, bookWithEmptyGenres.UpdatedAt, result.UpdatedAt, time.Second)

	result, err = querier.GetBooksByID(context.Background(), bookWithGenres.ID)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, bookWithGenres.ID, result.ID)
	require.Equal(t, bookWithGenres.Title, result.Title)
	genres := []string{genre1.Name, genre2.Name, genre3.Name}
	sort.Strings(genres)
	require.Equal(t, genres, strings.Split(string(result.Genres), ", "))
	require.Equal(t, bookWithGenres.Description, result.Description)
	require.Equal(t, bookWithGenres.CoverImageUrl, result.CoverImageUrl)
	require.Equal(t, bookWithGenres.Url, result.Url)
	require.Equal(t, bookWithGenres.AuthorName, result.AuthorName)
	require.Equal(t, bookWithGenres.PublisherName, result.PublisherName)
	require.Equal(t, bookWithGenres.PublishedDate.Time, result.PublishedDate.Time)
	require.Equal(t, bookWithGenres.Isbn, result.Isbn)
	require.WithinDuration(t, bookWithGenres.CreatedAt, result.CreatedAt, time.Second)
	require.WithinDuration(t, bookWithGenres.UpdatedAt, result.UpdatedAt, time.Second)
}

func TestGetBooksByTitle(t *testing.T) {
	title := testdata.RandomString(8)
	bookWithEmptyGenres := createTestBook(
		t,
		title,
		"",
		"",
		testdata.RandomString(13),
	)
	bookWithGenres := createTestBook(
		t,
		title,
		testdata.RandomString(8),
		testdata.RandomString(10),
		testdata.RandomString(13),
	)
	genre1 := createRandomGenre(t)
	genre2 := createRandomGenre(t)
	genre3 := createRandomGenre(t)
	createRandomBookGenre(t, bookWithGenres, genre1)
	createRandomBookGenre(t, bookWithGenres, genre2)
	createRandomBookGenre(t, bookWithGenres, genre3)

	result, err := querier.GetBooksByTitle(context.Background(), title)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Len(t, result, 2)

	require.Equal(t, bookWithEmptyGenres.ID, result[0].ID)
	require.Equal(t, bookWithEmptyGenres.Title, result[0].Title)
	require.Empty(t, result[0].Genres)
	require.Equal(t, bookWithEmptyGenres.Description, result[0].Description)
	require.Equal(t, bookWithEmptyGenres.CoverImageUrl, result[0].CoverImageUrl)
	require.Equal(t, bookWithEmptyGenres.Url, result[0].Url)
	require.Equal(t, bookWithEmptyGenres.AuthorName, result[0].AuthorName)
	require.Equal(t, bookWithEmptyGenres.PublisherName, result[0].PublisherName)
	require.Equal(t, bookWithEmptyGenres.PublishedDate.Time, result[0].PublishedDate.Time)
	require.Equal(t, bookWithEmptyGenres.Isbn, result[0].Isbn)
	require.WithinDuration(t, bookWithEmptyGenres.CreatedAt, result[0].CreatedAt, time.Second)
	require.WithinDuration(t, bookWithEmptyGenres.UpdatedAt, result[0].UpdatedAt, time.Second)

	require.Equal(t, bookWithGenres.ID, result[1].ID)
	require.Equal(t, bookWithGenres.Title, result[1].Title)
	genres := []string{genre1.Name, genre2.Name, genre3.Name}
	sort.Strings(genres)
	require.Equal(t, genres, strings.Split(string(result[1].Genres), ", "))
	require.Equal(t, bookWithGenres.Description, result[1].Description)
	require.Equal(t, bookWithGenres.CoverImageUrl, result[1].CoverImageUrl)
	require.Equal(t, bookWithGenres.Url, result[1].Url)
	require.Equal(t, bookWithGenres.AuthorName, result[1].AuthorName)
	require.Equal(t, bookWithGenres.PublisherName, result[1].PublisherName)
	require.Equal(t, bookWithGenres.PublishedDate.Time, result[1].PublishedDate.Time)
	require.Equal(t, bookWithGenres.Isbn, result[1].Isbn)
	require.WithinDuration(t, bookWithGenres.CreatedAt, result[1].CreatedAt, time.Second)
	require.WithinDuration(t, bookWithGenres.UpdatedAt, result[1].UpdatedAt, time.Second)
}

func TestGetBooksByISBN(t *testing.T) {
	ISBN := testdata.RandomString(13)
	book := createTestBook(
		t,
		testdata.RandomString(6),
		"",
		"",
		ISBN,
	)
	genre := createRandomGenre(t)
	createRandomBookGenre(t, book, genre)

	result, err := querier.GetBooksByISBN(
		context.Background(),
		sql.NullString{
			String: ISBN,
			Valid:  true,
		},
	)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, book.ID, result[0].ID)
	require.Equal(t, book.Title, result[0].Title)
	require.Equal(
		t, []string{genre.Name},
		strings.Split(string(result[0].Genres), ", "),
	)
	require.Equal(t, book.Description, result[0].Description)
	require.Equal(t, book.CoverImageUrl, result[0].CoverImageUrl)
	require.Equal(t, book.Url, result[0].Url)
	require.Equal(t, book.AuthorName, result[0].AuthorName)
	require.Equal(t, book.PublisherName, result[0].PublisherName)
	require.Equal(t, book.PublishedDate.Time, result[0].PublishedDate.Time)
	require.Equal(t, book.Isbn, result[0].Isbn)
	require.WithinDuration(t, book.CreatedAt, result[0].CreatedAt, time.Second)
	require.WithinDuration(t, book.UpdatedAt, result[0].UpdatedAt, time.Second)
}

func TestGetBooksByAuthor(t *testing.T) {
	author := testdata.RandomString(8)
	bookWithEmptyGenres := createTestBook(
		t,
		testdata.RandomString(6),
		author,
		"",
		testdata.RandomString(13),
	)
	bookWithGenres := createTestBook(
		t,
		testdata.RandomString(6),
		author,
		testdata.RandomString(10),
		testdata.RandomString(13),
	)
	genre1 := createRandomGenre(t)
	genre2 := createRandomGenre(t)
	createRandomBookGenre(t, bookWithGenres, genre1)
	createRandomBookGenre(t, bookWithGenres, genre2)

	result, err := querier.GetBooksByAuthor(
		context.Background(),
		sql.NullString{
			String: author,
			Valid:  true,
		},
	)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Len(t, result, 2)

	require.Equal(t, bookWithEmptyGenres.ID, result[0].ID)
	require.Equal(t, bookWithEmptyGenres.Title, result[0].Title)
	require.Empty(t, result[0].Genres)
	require.Equal(t, bookWithEmptyGenres.Description, result[0].Description)
	require.Equal(t, bookWithEmptyGenres.CoverImageUrl, result[0].CoverImageUrl)
	require.Equal(t, bookWithEmptyGenres.Url, result[0].Url)
	require.Equal(t, bookWithEmptyGenres.AuthorName, result[0].AuthorName)
	require.Equal(t, bookWithEmptyGenres.PublisherName, result[0].PublisherName)
	require.Equal(t, bookWithEmptyGenres.PublishedDate.Time, result[0].PublishedDate.Time)
	require.Equal(t, bookWithEmptyGenres.Isbn, result[0].Isbn)
	require.WithinDuration(t, bookWithEmptyGenres.CreatedAt, result[0].CreatedAt, time.Second)
	require.WithinDuration(t, bookWithEmptyGenres.UpdatedAt, result[0].UpdatedAt, time.Second)

	require.Equal(t, bookWithGenres.ID, result[1].ID)
	require.Equal(t, bookWithGenres.Title, result[1].Title)
	genres := []string{genre1.Name, genre2.Name}
	sort.Strings(genres)
	require.Equal(t, genres, strings.Split(string(result[1].Genres), ", "))
	require.Equal(t, bookWithGenres.Description, result[1].Description)
	require.Equal(t, bookWithGenres.CoverImageUrl, result[1].CoverImageUrl)
	require.Equal(t, bookWithGenres.Url, result[1].Url)
	require.Equal(t, bookWithGenres.AuthorName, result[1].AuthorName)
	require.Equal(t, bookWithGenres.PublisherName, result[1].PublisherName)
	require.Equal(t, bookWithGenres.PublishedDate.Time, result[1].PublishedDate.Time)
	require.Equal(t, bookWithGenres.Isbn, result[1].Isbn)
	require.WithinDuration(t, bookWithGenres.CreatedAt, result[1].CreatedAt, time.Second)
	require.WithinDuration(t, bookWithGenres.UpdatedAt, result[1].UpdatedAt, time.Second)
}

func TestGetBooksByPublisher(t *testing.T) {
	publisher := testdata.RandomString(10)
	bookWithEmptyGenres := createTestBook(
		t,
		testdata.RandomString(6),
		"",
		publisher,
		testdata.RandomString(13),
	)
	bookWithGenres := createTestBook(
		t,
		testdata.RandomString(6),
		testdata.RandomString(8),
		publisher,
		testdata.RandomString(13),
	)
	genre1 := createRandomGenre(t)
	genre2 := createRandomGenre(t)
	createRandomBookGenre(t, bookWithGenres, genre1)
	createRandomBookGenre(t, bookWithGenres, genre2)

	result, err := querier.GetBooksByPublisher(
		context.Background(),
		sql.NullString{
			String: publisher,
			Valid:  true,
		},
	)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Len(t, result, 2)

	require.Equal(t, bookWithEmptyGenres.ID, result[0].ID)
	require.Equal(t, bookWithEmptyGenres.Title, result[0].Title)
	require.Empty(t, result[0].Genres)
	require.Equal(t, bookWithEmptyGenres.Description, result[0].Description)
	require.Equal(t, bookWithEmptyGenres.CoverImageUrl, result[0].CoverImageUrl)
	require.Equal(t, bookWithEmptyGenres.Url, result[0].Url)
	require.Equal(t, bookWithEmptyGenres.AuthorName, result[0].AuthorName)
	require.Equal(t, bookWithEmptyGenres.PublisherName, result[0].PublisherName)
	require.Equal(t, bookWithEmptyGenres.PublishedDate.Time, result[0].PublishedDate.Time)
	require.Equal(t, bookWithEmptyGenres.Isbn, result[0].Isbn)
	require.WithinDuration(t, bookWithEmptyGenres.CreatedAt, result[0].CreatedAt, time.Second)
	require.WithinDuration(t, bookWithEmptyGenres.UpdatedAt, result[0].UpdatedAt, time.Second)

	require.Equal(t, bookWithGenres.ID, result[1].ID)
	require.Equal(t, bookWithGenres.Title, result[1].Title)
	genres := []string{genre1.Name, genre2.Name}
	sort.Strings(genres)
	require.Equal(t, genres, strings.Split(string(result[1].Genres), ", "))
	require.Equal(t, bookWithGenres.Description, result[1].Description)
	require.Equal(t, bookWithGenres.CoverImageUrl, result[1].CoverImageUrl)
	require.Equal(t, bookWithGenres.Url, result[1].Url)
	require.Equal(t, bookWithGenres.AuthorName, result[1].AuthorName)
	require.Equal(t, bookWithGenres.PublisherName, result[1].PublisherName)
	require.Equal(t, bookWithGenres.PublishedDate.Time, result[1].PublishedDate.Time)
	require.Equal(t, bookWithGenres.Isbn, result[1].Isbn)
	require.WithinDuration(t, bookWithGenres.CreatedAt, result[1].CreatedAt, time.Second)
	require.WithinDuration(t, bookWithGenres.UpdatedAt, result[1].UpdatedAt, time.Second)
}

func TestUpdateBook(t *testing.T) {
	book1 := createTestBook(
		t,
		testdata.RandomString(6),
		"",
		"",
		testdata.RandomString(13),
	)

	arg := UpdateBookParams{
		ID:            book1.ID,
		Title:         book1.Title,
		Description:   sql.NullString{String: testdata.RandomString(12), Valid: true},
		CoverImageUrl: book1.CoverImageUrl,
		Url:           book1.Url,
		AuthorName:    book1.AuthorName,
		PublisherName: book1.PublisherName,
		PublishedDate: book1.PublishedDate,
		Isbn:          book1.Isbn,
	}

	book2, err := querier.UpdateBook(context.Background(), arg)
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
	book1 := createTestBook(
		t,
		testdata.RandomString(6),
		"",
		"",
		testdata.RandomString(13),
	)
	rowsAffected, err := querier.DeleteBook(context.Background(), book1.ID)
	require.NoError(t, err)
	require.Equal(t, int64(1), rowsAffected)

	book2, err := querier.GetBooksByID(context.Background(), book1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, book2)
}
