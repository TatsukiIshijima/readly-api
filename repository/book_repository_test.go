package repository

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"readly/entity"
	"readly/testdata"
	"testing"
)

func createRandomAuthor(t *testing.T) *CreateAuthorResponse {
	name := testdata.RandomString(8)
	res, err := bookRepo.CreateAuthor(context.Background(), NewCreateAuthorRequest(name))
	require.NoError(t, err)
	return res
}

func TestCreateAuthor(t *testing.T) {
	a := createRandomAuthor(t)
	ga, err := querier.GetAuthorByName(context.Background(), a.Name)
	require.NoError(t, err)
	require.Equal(t, a.Name, ga.Name)
}

func createRandomBook(t *testing.T) *CreateBookResponse {
	title := testdata.RandomString(8)
	description := testdata.RandomString(10)
	coverImageUrl := "https://example.com"
	url := "https://example.com"
	publishDate := entity.Now()
	isbn := testdata.RandomString(13)

	req := CreateBookRequest{
		Title:         title,
		Description:   &description,
		CoverImageURL: &coverImageUrl,
		URL:           &url,
		Author:        nil,
		Publisher:     nil,
		PublishDate:   &publishDate,
		ISBN:          &isbn,
	}

	b, err := bookRepo.CreateBook(context.Background(), req)
	require.NoError(t, err)
	return b
}

func TestCreateBook(t *testing.T) {
	b := createRandomBook(t)
	gb, err := querier.GetBooksByID(context.Background(), b.ID)
	require.NoError(t, err)

	require.Equal(t, b.ID, gb.ID)
	require.Equal(t, b.Title, gb.Title)
	require.Equal(t, *b.Description, gb.Description.String)
	require.Equal(t, *b.CoverImageURL, gb.CoverImageUrl.String)
	require.Equal(t, *b.URL, gb.Url.String)
	require.Equal(t, sql.NullString{}, gb.AuthorName)
	require.Equal(t, sql.NullString{}, gb.PublisherName)
	require.Equal(t, *b.PublishDate, *entity.NewDateEntityFromNullTime(gb.PublishedDate))
	require.Equal(t, *b.ISBN, gb.Isbn.String)
}

func TestCreateBookGenre(t *testing.T) {
	b := createRandomBook(t)
	g := createRandomGenre(t)

	req := CreateBookGenreRequest{
		BookID:    b.ID,
		GenreName: g.Name,
	}
	res, err := bookRepo.CreateBookGenre(context.Background(), req)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	gg, err := querier.GetGenresByBookID(context.Background(), b.ID)
	require.NoError(t, err)
	require.Equal(t, g.Name, gg[0])
}

func createRandomGenre(t *testing.T) *CreateGenreResponse {
	name := testdata.RandomString(8)
	res, err := bookRepo.CreateGenre(context.Background(), NewCreateGenreRequest(name))
	require.NoError(t, err)
	return res
}

func TestCreateGenre(t *testing.T) {
	g := createRandomGenre(t)
	gg, err := querier.GetGenreByName(context.Background(), g.Name)
	require.NoError(t, err)
	require.Equal(t, g.Name, gg.Name)
}

func createRandomPublisher(t *testing.T) *CreatePublisherResponse {
	name := testdata.RandomString(8)
	p, err := bookRepo.CreatePublisher(context.Background(), NewCreatePublisherRequest(name))
	require.NoError(t, err)
	return p
}

func TestCreatePublisher(t *testing.T) {
	p := createRandomPublisher(t)
	gp, err := querier.GetPublisherByName(context.Background(), p.Name)
	require.NoError(t, err)
	require.Equal(t, p.Name, gp.Name)
}

func TestDeleteAuthor(t *testing.T) {
	a := createRandomAuthor(t)
	err := bookRepo.DeleteAuthor(context.Background(), NewDeleteAuthorRequest(a.Name))
	require.NoError(t, err)

	_, err = querier.GetAuthorByName(context.Background(), a.Name)
	require.ErrorIs(t, err, sql.ErrNoRows)
}

func TestDeleteBook(t *testing.T) {
	b := createRandomBook(t)
	err := bookRepo.DeleteBook(context.Background(), NewDeleteBookRequest(b.ID))
	require.NoError(t, err)

	_, err = querier.GetBooksByID(context.Background(), b.ID)
	require.ErrorIs(t, err, sql.ErrNoRows)
}

func TestDeleteBookGenre(t *testing.T) {
	b := createRandomBook(t)
	g := createRandomGenre(t)

	req := CreateBookGenreRequest{
		BookID:    b.ID,
		GenreName: g.Name,
	}
	_, err := bookRepo.CreateBookGenre(context.Background(), req)
	require.NoError(t, err)

	err = bookRepo.DeleteBookGenre(context.Background(), DeleteBookGenreRequest{
		BookID:    b.ID,
		GenreName: g.Name,
	})
	require.NoError(t, err)

	gg, err := querier.GetGenresByBookID(context.Background(), b.ID)
	require.NoError(t, err)
	require.Empty(t, gg)
}

func TestDeleteGenre(t *testing.T) {
	g := createRandomGenre(t)
	err := bookRepo.DeleteGenre(context.Background(), NewDeleteGenreRequest(g.Name))
	require.NoError(t, err)

	_, err = querier.GetGenreByName(context.Background(), g.Name)
	require.ErrorIs(t, err, sql.ErrNoRows)
}

func TestUpdateBook(t *testing.T) {
	b := createRandomBook(t)
	newTitle := testdata.RandomString(8)
	req := UpdateBookRequest{
		BookID: b.ID,
		Title:  newTitle,
	}
	res, err := bookRepo.UpdateBook(context.Background(), req)
	require.NoError(t, err)
	updatedBook, err := querier.GetBooksByID(context.Background(), res.BookID)
	require.NoError(t, err)
	require.Equal(t, newTitle, updatedBook.Title)
}

//func TestRegister(t *testing.T) {
//	user, err := repository.createRandomUser()
//	require.NoError(t, err)
//
//	n := 3
//	results := make(chan entity.Book)
//	errs := make(chan error)
//
//	for i := 0; i < n; i++ {
//		go func() {
//			// TODO:チェネルでジャンルを増やす&共有
//			genres := make([]string, i+1)
//			for j := 0; j <= i; j++ {
//				genres[j] = testdata.RandomString(6)
//			}
//			arg := RegisterRequest{
//				UserID:        user.ID,
//				Title:         testdata.RandomString(6),
//				Genres:        genres,
//				Description:   testdata.RandomString(12),
//				CoverImageURL: "https://example.com",
//				URL:           "https://example.com",
//				AuthorName:    testdata.RandomString(6),
//				PublisherName: testdata.RandomString(6),
//				PublishDate:   time.Now(),
//				ISBN:          testdata.RandomString(13),
//			}
//			result, err := repository.repo.Register(context.Background(), arg)
//			errs <- err
//			results <- result
//		}()
//	}
//
//	for i := 0; i < n; i++ {
//		err := <-errs
//		require.NoError(t, err)
//
//		result := <-results
//
//		author, err := repository.querier.GetAuthorByName(context.Background(), result.AuthorName)
//		require.NoError(t, err)
//		require.NotEmpty(t, author)
//		require.Equal(t, result.AuthorName, author.Name)
//
//		publisher, err := repository.querier.GetPublisherByName(context.Background(), result.PublisherName)
//		require.NoError(t, err)
//		require.NotEmpty(t, publisher)
//		require.Equal(t, result.PublisherName, publisher.Name)
//
//		genres, err := repository.querier.GetGenresByBookID(context.Background(), result.ID)
//		require.NoError(t, err)
//		require.Equal(t, len(result.Genres), len(genres))
//		for _, g := range genres {
//			genre, err := repository.querier.GetGenreByName(context.Background(), g)
//			require.NoError(t, err)
//			require.NotEmpty(t, genre)
//		}
//
//		book, err := repository.querier.GetBookById(context.Background(), result.ID)
//		require.NoError(t, err)
//		require.NotEmpty(t, book)
//		require.Equal(t, result.Title, book.Title.String)
//		require.Equal(t, result.Description, book.Description.String)
//		require.Equal(t, result.CoverImageURL, book.CoverImageUrl.String)
//		require.Equal(t, result.URL, book.Url.String)
//		require.Equal(t, result.AuthorName, book.AuthorName)
//		require.Equal(t, result.PublisherName, book.PublisherName)
//		require.WithinDuration(t, result.PublishDate, book.PublishedDate.Time.UTC(), time.Second)
//		require.Equal(t, result.ISBN, book.Isbn.String)
//	}
//
//	param := db.GetReadingHistoryByUserIDParams{
//		UserID: user.ID,
//		Limit:  10,
//		Offset: 0,
//	}
//	histories, err := repository.querier.GetReadingHistoryByUserID(context.Background(), param)
//	require.NoError(t, err)
//	require.Equal(t, n, len(histories))
//	for _, h := range histories {
//		require.Equal(t, user.ID, h.UserID)
//		require.Equal(t, db.ReadingStatusUnread, h.Status)
//	}
//}
//
//func TestGet(t *testing.T) {
//	user, err := repository.createRandomUser()
//	require.NoError(t, err)
//
//	registerReq := RegisterRequest{
//		UserID:        user.ID,
//		Title:         testdata.RandomString(6),
//		Genres:        []string{testdata.RandomString(6)},
//		Description:   testdata.RandomString(12),
//		CoverImageURL: "https://example.com",
//		URL:           "https://example.com",
//		AuthorName:    testdata.RandomString(6),
//		PublisherName: testdata.RandomString(6),
//		PublishDate:   time.Now(),
//		ISBN:          testdata.RandomString(13),
//	}
//	registeredBook, err := repository.repo.Register(context.Background(), registerReq)
//	require.NoError(t, err)
//
//	book, err := repository.repo.Get(context.Background(), registeredBook.ID)
//	require.NoError(t, err)
//	require.Equal(t, registeredBook.ID, book.ID)
//	require.Equal(t, registeredBook.Title, book.Title)
//	require.Equal(t, registeredBook.Genres[0], book.Genres[0])
//	require.Equal(t, registeredBook.Description, book.Description)
//	require.Equal(t, registeredBook.CoverImageURL, book.CoverImageURL)
//	require.Equal(t, registeredBook.URL, book.URL)
//	require.Equal(t, registeredBook.AuthorName, book.AuthorName)
//	require.Equal(t, registeredBook.PublisherName, book.PublisherName)
//	require.WithinDuration(t, registeredBook.PublishDate, book.PublishDate, time.Second)
//	require.Equal(t, registeredBook.ISBN, book.ISBN)
//}
//
//func TestList(t *testing.T) {
//	user, err := repository.createRandomUser()
//	require.NoError(t, err)
//
//	n := 3
//	requests := make([]RegisterRequest, 0, n)
//	for i := 0; i < n; i++ {
//		registerReq := RegisterRequest{
//			UserID:        user.ID,
//			Title:         testdata.RandomString(6),
//			Genres:        []string{testdata.RandomString(6)},
//			Description:   testdata.RandomString(12),
//			CoverImageURL: "https://example.com",
//			URL:           "https://example.com",
//			AuthorName:    testdata.RandomString(6),
//			PublisherName: testdata.RandomString(6),
//			PublishDate:   time.Now(),
//			ISBN:          testdata.RandomString(13),
//		}
//		requests = append(requests, registerReq)
//		book, err := repository.repo.Register(context.Background(), registerReq)
//		require.NoError(t, err)
//		require.NotEmpty(t, book)
//	}
//
//	listReq := ListRequest{
//		UserID: user.ID,
//		Limit:  int32(n),
//		Offset: 0,
//	}
//	books, err := repository.repo.List(context.Background(), listReq)
//	require.NoError(t, err)
//	require.Equal(t, n, len(books))
//
//	for i, book := range books {
//		require.Equal(t, requests[i].Title, book.Title)
//		require.Equal(t, requests[i].Genres[0], book.Genres[0])
//		require.Equal(t, requests[i].Description, book.Description)
//		require.Equal(t, requests[i].CoverImageURL, book.CoverImageURL)
//		require.Equal(t, requests[i].URL, book.URL)
//		require.Equal(t, requests[i].AuthorName, book.AuthorName)
//		require.Equal(t, requests[i].PublisherName, book.PublisherName)
//		require.WithinDuration(t, requests[i].PublishDate, book.PublishDate, time.Second)
//		require.Equal(t, requests[i].ISBN, book.ISBN)
//	}
//}
//
//func TestDelete(t *testing.T) {
//	user, err := repository.createRandomUser()
//	require.NoError(t, err)
//
//	registerReq := RegisterRequest{
//		UserID:        user.ID,
//		Title:         testdata.RandomString(6),
//		Genres:        []string{testdata.RandomString(6)},
//		Description:   testdata.RandomString(12),
//		CoverImageURL: "https://example.com",
//		URL:           "https://example.com",
//		AuthorName:    testdata.RandomString(6),
//		PublisherName: testdata.RandomString(6),
//		PublishDate:   time.Now(),
//		ISBN:          testdata.RandomString(13),
//	}
//	registeredBook, err := repository.repo.Register(context.Background(), registerReq)
//	require.NoError(t, err)
//
//	err = repository.repo.Delete(context.Background(), DeleteRequest{
//		UserID: user.ID,
//		BookID: registeredBook.ID,
//	})
//	require.NoError(t, err)
//
//	historyParam := db.GetReadingHistoryByUserIDParams{
//		UserID: user.ID,
//		Limit:  10,
//		Offset: 0,
//	}
//	histories, err := repository.querier.GetReadingHistoryByUserID(context.Background(), historyParam)
//	require.NoError(t, err)
//	require.Empty(t, histories)
//
//	genres, err := repository.querier.GetGenresByBookID(context.Background(), registeredBook.ID)
//	require.NoError(t, err)
//	require.Empty(t, genres)
//
//	deletedBook, err := repository.querier.GetBookById(context.Background(), registeredBook.ID)
//	require.Zero(t, deletedBook)
//	require.ErrorIs(t, err, sql.ErrNoRows)
//}
