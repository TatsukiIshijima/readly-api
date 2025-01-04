package db

import (
	"context"
	"database/sql"
	"time"
)

func (q *FakeQuerier) GetBookById(ctx context.Context, id int64) (Book, error) {
	// FIXME:インメモリ管理
	return Book{
		ID: 1,
		Title: sql.NullString{
			String: "Title",
			Valid:  true,
		},
		Description: sql.NullString{
			String: "Description",
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
		AuthorName:    "Author",
		PublisherName: "Publisher",
		PublishedDate: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		Isbn: sql.NullString{
			String: "1234567890123",
			Valid:  true,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (q *FakeQuerier) GetGenresByBookID(ctx context.Context, bookID int64) ([]string, error) {
	// FIXME:インメモリ管理
	if bookID == 1 {
		return []string{"genre1", "genre2"}, nil
	} else {
		return []string{}, nil
	}
}
