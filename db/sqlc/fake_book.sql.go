//go:build test

package db

import (
	"context"
	"database/sql"
	"time"
)

type BookTable struct {
	// 自動インクリメンタル用
	NextID  int64
	Columns []Book
}

var bookTable = BookTable{NextID: 1}

func (q *FakeQuerier) CreateBook(ctx context.Context, arg CreateBookParams) (Book, error) {
	now := time.Now().UTC()
	b := Book{
		ID:            bookTable.NextID,
		Title:         arg.Title,
		Description:   arg.Description,
		CoverImageUrl: arg.CoverImageUrl,
		Url:           arg.Url,
		AuthorName:    arg.AuthorName,
		PublisherName: arg.PublisherName,
		PublishedDate: arg.PublishedDate,
		Isbn:          arg.Isbn,
		CreatedAt:     now,
		UpdatedAt:     now,
	}
	bookTable.Columns = append(bookTable.Columns, b)
	bookTable.NextID++
	return b, nil
}

func (q *FakeQuerier) DeleteBook(ctx context.Context, id int64) (int64, error) {
	for i, b := range bookTable.Columns {
		if b.ID == id {
			bookTable.Columns = append(bookTable.Columns[:i], bookTable.Columns[i+1:]...)
			return 1, nil
		}
	}
	return 0, nil
}

func (q *FakeQuerier) GetBooksByAuthor(ctx context.Context, authorName sql.NullString) ([]GetBooksByAuthorRow, error) {
	var rows []GetBooksByAuthorRow
	for _, b := range bookTable.Columns {
		if b.AuthorName == authorName {
			rows = append(rows, GetBooksByAuthorRow{
				ID:            b.ID,
				Title:         b.Title,
				Description:   b.Description,
				CoverImageUrl: b.CoverImageUrl,
				Url:           b.Url,
				AuthorName:    b.AuthorName,
				PublisherName: b.PublisherName,
				PublishedDate: b.PublishedDate,
				Isbn:          b.Isbn,
				CreatedAt:     b.CreatedAt,
				UpdatedAt:     b.UpdatedAt,
			})
		}
	}
	return rows, nil
}

func (q *FakeQuerier) GetBooksByID(_ context.Context, id int64) (GetBooksByIDRow, error) {
	for _, b := range bookTable.Columns {
		if b.ID == id {
			return GetBooksByIDRow{
				ID:            b.ID,
				Title:         b.Title,
				Description:   b.Description,
				CoverImageUrl: b.CoverImageUrl,
				Url:           b.Url,
				AuthorName:    b.AuthorName,
				PublisherName: b.PublisherName,
				PublishedDate: b.PublishedDate,
				Isbn:          b.Isbn,
				CreatedAt:     b.CreatedAt,
				UpdatedAt:     b.UpdatedAt,
			}, nil
		}
	}
	return GetBooksByIDRow{}, nil
}

func (q *FakeQuerier) GetBooksByISBN(ctx context.Context, isbn sql.NullString) ([]GetBooksByISBNRow, error) {
	var rows []GetBooksByISBNRow
	for _, b := range bookTable.Columns {
		if b.Isbn == isbn {
			rows = append(rows, GetBooksByISBNRow{
				ID:            b.ID,
				Title:         b.Title,
				Description:   b.Description,
				CoverImageUrl: b.CoverImageUrl,
				Url:           b.Url,
				AuthorName:    b.AuthorName,
				PublisherName: b.PublisherName,
				PublishedDate: b.PublishedDate,
				Isbn:          b.Isbn,
				CreatedAt:     b.CreatedAt,
				UpdatedAt:     b.UpdatedAt,
			})
		}
	}
	return rows, nil
}

func (q *FakeQuerier) GetBooksByPublisher(ctx context.Context, publisherName sql.NullString) ([]GetBooksByPublisherRow, error) {
	var rows []GetBooksByPublisherRow
	for _, b := range bookTable.Columns {
		if b.PublisherName == publisherName {
			rows = append(rows, GetBooksByPublisherRow{
				ID:            b.ID,
				Title:         b.Title,
				Description:   b.Description,
				CoverImageUrl: b.CoverImageUrl,
				Url:           b.Url,
				AuthorName:    b.AuthorName,
				PublisherName: b.PublisherName,
				PublishedDate: b.PublishedDate,
				Isbn:          b.Isbn,
				CreatedAt:     b.CreatedAt,
				UpdatedAt:     b.UpdatedAt,
			})
		}
	}
	return rows, nil
}

func (q *FakeQuerier) GetBooksByTitle(ctx context.Context, title string) ([]GetBooksByTitleRow, error) {
	var rows []GetBooksByTitleRow
	for _, b := range bookTable.Columns {
		if b.Title == title {
			rows = append(rows, GetBooksByTitleRow{
				ID:            b.ID,
				Title:         b.Title,
				Description:   b.Description,
				CoverImageUrl: b.CoverImageUrl,
				Url:           b.Url,
				AuthorName:    b.AuthorName,
				PublisherName: b.PublisherName,
				PublishedDate: b.PublishedDate,
				Isbn:          b.Isbn,
				CreatedAt:     b.CreatedAt,
				UpdatedAt:     b.UpdatedAt,
			})
		}
	}
	return rows, nil
}

func (q *FakeQuerier) UpdateBook(ctx context.Context, arg UpdateBookParams) (Book, error) {
	for i, b := range bookTable.Columns {
		if b.ID == arg.ID {
			bookTable.Columns[i].Title = arg.Title
			bookTable.Columns[i].Description = arg.Description
			bookTable.Columns[i].CoverImageUrl = arg.CoverImageUrl
			bookTable.Columns[i].Url = arg.Url
			bookTable.Columns[i].AuthorName = arg.AuthorName
			bookTable.Columns[i].PublisherName = arg.PublisherName
			bookTable.Columns[i].PublishedDate = arg.PublishedDate
			bookTable.Columns[i].Isbn = arg.Isbn
			bookTable.Columns[i].UpdatedAt = time.Now().UTC()
			return bookTable.Columns[i], nil
		}
	}
	return Book{}, sql.ErrNoRows
}
