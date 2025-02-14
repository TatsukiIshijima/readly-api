//go:build test

package db

import (
	"bytes"
	"context"
	"database/sql"
	"github.com/lib/pq"
	"time"
)

type ReadingHistoryTable struct {
	Columns []ReadingHistory
}

var readingHistoryTable = ReadingHistoryTable{}

func (q *FakeQuerier) CreateReadingHistory(_ context.Context, arg CreateReadingHistoryParams) (ReadingHistory, error) {
	for _, h := range readingHistoryTable.Columns {
		if h.UserID == arg.UserID && h.BookID == arg.BookID {
			return ReadingHistory{}, &pq.Error{Code: "23505", Message: "duplicate key value violates unique constraint"}
		}
	}
	h := ReadingHistory{
		UserID:    arg.UserID,
		BookID:    arg.BookID,
		Status:    arg.Status,
		StartDate: arg.StartDate,
		EndDate:   arg.EndDate,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	readingHistoryTable.Columns = append(readingHistoryTable.Columns, h)
	return h, nil
}

func (q *FakeQuerier) DeleteReadingHistory(_ context.Context, arg DeleteReadingHistoryParams) (int64, error) {
	for i, h := range readingHistoryTable.Columns {
		if h.UserID == arg.UserID && h.BookID == arg.BookID {
			readingHistoryTable.Columns = append(readingHistoryTable.Columns[:i], readingHistoryTable.Columns[i+1:]...)
			return 1, nil
		}
	}
	return 0, nil
}

func (q *FakeQuerier) GetReadingHistoryByUser(ctx context.Context, arg GetReadingHistoryByUserParams) ([]GetReadingHistoryByUserRow, error) {
	// TODO:ページング対応
	var rows []GetReadingHistoryByUserRow
	for _, r := range readingHistoryTable.Columns {
		if r.UserID == arg.UserID {
			b, err := q.GetBooksByID(ctx, r.BookID)
			if err != nil {
				return rows, err
			}
			g, err := q.GetGenresByBookID(ctx, r.BookID)
			if err != nil {
				return rows, err
			}
			genres := chars(g).toByte()
			rows = append(rows, GetReadingHistoryByUserRow{
				ID:            sql.NullInt64{Int64: b.ID, Valid: true},
				Title:         sql.NullString{String: b.Title, Valid: true},
				Genres:        genres,
				Description:   b.Description,
				CoverImageUrl: b.CoverImageUrl,
				Url:           b.Url,
				AuthorName:    b.AuthorName,
				PublisherName: b.PublisherName,
				PublishedDate: b.PublishedDate,
				Isbn:          b.Isbn,
				Status:        r.Status,
				StartDate:     r.StartDate,
				EndDate:       r.EndDate,
			})
		}
	}
	return rows, nil
}

func (q *FakeQuerier) GetReadingHistoryByUserAndBook(ctx context.Context, arg GetReadingHistoryByUserAndBookParams) (GetReadingHistoryByUserAndBookRow, error) {
	for _, r := range readingHistoryTable.Columns {
		if r.UserID == arg.UserID && r.BookID == arg.BookID {
			b, err := q.GetBooksByID(ctx, r.BookID)
			if err != nil {
				return GetReadingHistoryByUserAndBookRow{}, err
			}
			g, err := q.GetGenresByBookID(ctx, r.BookID)
			if err != nil {
				return GetReadingHistoryByUserAndBookRow{}, err
			}
			genres := chars(g).toByte()
			return GetReadingHistoryByUserAndBookRow{
				ID:            sql.NullInt64{Int64: b.ID, Valid: true},
				Title:         sql.NullString{String: b.Title, Valid: true},
				Genres:        genres,
				Description:   b.Description,
				CoverImageUrl: b.CoverImageUrl,
				Url:           b.Url,
				AuthorName:    b.AuthorName,
				PublisherName: b.PublisherName,
				PublishedDate: b.PublishedDate,
				Isbn:          b.Isbn,
				Status:        r.Status,
				StartDate:     r.StartDate,
				EndDate:       r.EndDate,
			}, nil
		}
	}
	return GetReadingHistoryByUserAndBookRow{}, sql.ErrNoRows
}

func (q *FakeQuerier) GetReadingHistoryByUserAndStatus(ctx context.Context, arg GetReadingHistoryByUserAndStatusParams) ([]GetReadingHistoryByUserAndStatusRow, error) {
	// TODO:ページング対応
	var rows []GetReadingHistoryByUserAndStatusRow
	for _, r := range readingHistoryTable.Columns {
		if r.UserID == arg.UserID && r.Status == arg.Status {
			b, err := q.GetBooksByID(ctx, r.BookID)
			if err != nil {
				return rows, err
			}
			g, err := q.GetGenresByBookID(ctx, r.BookID)
			if err != nil {
				return rows, err
			}
			genres := chars(g).toByte()
			rows = append(rows, GetReadingHistoryByUserAndStatusRow{
				ID:            sql.NullInt64{Int64: b.ID, Valid: true},
				Title:         sql.NullString{String: b.Title, Valid: true},
				Genres:        genres,
				Description:   b.Description,
				CoverImageUrl: b.CoverImageUrl,
				Url:           b.Url,
				AuthorName:    b.AuthorName,
				PublisherName: b.PublisherName,
				PublishedDate: b.PublishedDate,
				Isbn:          b.Isbn,
				Status:        r.Status,
				StartDate:     r.StartDate,
				EndDate:       r.EndDate,
			})
		}
	}
	return rows, nil
}

func (q *FakeQuerier) UpdateReadingHistory(ctx context.Context, arg UpdateReadingHistoryParams) (ReadingHistory, error) {
	for i, r := range readingHistoryTable.Columns {
		if r.UserID == arg.UserID && r.BookID == arg.BookID {
			readingHistoryTable.Columns[i].Status = arg.Status
			readingHistoryTable.Columns[i].StartDate = arg.StartDate
			readingHistoryTable.Columns[i].EndDate = arg.EndDate
			readingHistoryTable.Columns[i].UpdatedAt = time.Now().UTC()
			return readingHistoryTable.Columns[i], nil
		}
	}
	return ReadingHistory{}, sql.ErrNoRows
}

type chars []string

func (c chars) toByte() []byte {
	var buffer bytes.Buffer

	for _, str := range c {
		buffer.WriteString(str)
	}

	byteSlice := buffer.Bytes()
	return byteSlice
}
