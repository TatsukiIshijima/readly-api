//go:build test

package db

import (
	"context"
	"github.com/lib/pq"
)

type BookGenreTable struct {
	Columns []BookGenre
}

var bookGenreTable = BookGenreTable{}

func (q *FakeQuerier) CreateBookGenre(ctx context.Context, arg CreateBookGenreParams) (BookGenre, error) {
	for _, bookGenre := range bookGenreTable.Columns {
		if bookGenre.BookID == arg.BookID && bookGenre.GenreName == arg.GenreName {
			return BookGenre{}, &pq.Error{Code: "23505", Message: "duplicate key value violates unique constraint"}
		}
	}
	b := BookGenre{
		BookID:    arg.BookID,
		GenreName: arg.GenreName,
	}
	bookGenreTable.Columns = append(bookGenreTable.Columns, b)
	return b, nil
}

func (q *FakeQuerier) DeleteBookGenre(ctx context.Context, arg DeleteBookGenreParams) (int64, error) {
	for i, bookGenre := range bookGenreTable.Columns {
		if bookGenre.BookID == arg.BookID && bookGenre.GenreName == arg.GenreName {
			bookGenreTable.Columns = append(bookGenreTable.Columns[:i], bookGenreTable.Columns[i+1:]...)
			return 1, nil
		}
	}
	return 0, nil
}

func (q *FakeQuerier) GetGenresByBookID(ctx context.Context, bookID int64) ([]string, error) {
	var genres []string
	for _, bookGenre := range bookGenreTable.Columns {
		if bookGenre.BookID == bookID {
			genres = append(genres, bookGenre.GenreName)
		}
	}
	return genres, nil
}
