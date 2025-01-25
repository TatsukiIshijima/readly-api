package db

import (
	"context"
	"database/sql"
	"readly/testdata"
)

func (q *FakeQuerier) GetBooksByID(_ context.Context, id int64) (GetBooksByIDRow, error) {
	// FIXME:インメモリ管理
	if id != 1 {
		return GetBooksByIDRow{}, sql.ErrNoRows
	}
	return GetBooksByIDRow{
		ID:    1,
		Title: "Title",
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
		AuthorName: sql.NullString{
			String: "Author",
			Valid:  true,
		},
		PublisherName: sql.NullString{
			String: "Publisher",
			Valid:  true,
		},
		PublishedDate: sql.NullTime{
			Time:  testdata.TimeFrom("1970-01-01 00:00:00"),
			Valid: true,
		},
		Isbn: sql.NullString{
			String: "1234567890123",
			Valid:  true,
		},
		CreatedAt: testdata.TimeFrom("2025-01-01 00:00:00"),
		UpdatedAt: testdata.TimeFrom("2025-01-01 00:00:00"),
	}, nil
}

func (q *FakeQuerier) GetGenresByBookID(_ context.Context, bookID int64) ([]string, error) {
	// FIXME:インメモリ管理
	if bookID == 1 {
		return []string{"Genre1", "Genre2"}, nil
	} else {
		return []string{}, nil
	}
}
