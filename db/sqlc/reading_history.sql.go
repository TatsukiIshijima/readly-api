// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: reading_history.sql

package db

import (
	"context"
	"database/sql"
)

const getReadingHistoryByUserAndBook = `-- name: GetReadingHistoryByUserAndBook :one
SELECT user_id, book_id, status, start_date, end_date, created_at, updated_at
FROM reading_histories
WHERE user_id = $1
  AND book_id = $2
`

type GetReadingHistoryByUserAndBookParams struct {
	UserID int64 `json:"user_id"`
	BookID int64 `json:"book_id"`
}

func (q *Queries) GetReadingHistoryByUserAndBook(ctx context.Context, arg GetReadingHistoryByUserAndBookParams) (ReadingHistory, error) {
	row := q.db.QueryRowContext(ctx, getReadingHistoryByUserAndBook, arg.UserID, arg.BookID)
	var i ReadingHistory
	err := row.Scan(
		&i.UserID,
		&i.BookID,
		&i.Status,
		&i.StartDate,
		&i.EndDate,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getReadingHistoryByUserAndStatus = `-- name: GetReadingHistoryByUserAndStatus :many
SELECT user_id, book_id, status, start_date, end_date, created_at, updated_at
FROM reading_histories
WHERE user_id = $1
  AND status = $2
`

type GetReadingHistoryByUserAndStatusParams struct {
	UserID int64         `json:"user_id"`
	Status ReadingStatus `json:"status"`
}

func (q *Queries) GetReadingHistoryByUserAndStatus(ctx context.Context, arg GetReadingHistoryByUserAndStatusParams) ([]ReadingHistory, error) {
	rows, err := q.db.QueryContext(ctx, getReadingHistoryByUserAndStatus, arg.UserID, arg.Status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ReadingHistory
	for rows.Next() {
		var i ReadingHistory
		if err := rows.Scan(
			&i.UserID,
			&i.BookID,
			&i.Status,
			&i.StartDate,
			&i.EndDate,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getReadingHistoryByUserID = `-- name: GetReadingHistoryByUserID :many
SELECT user_id, book_id, status, start_date, end_date, created_at, updated_at
FROM reading_histories
WHERE user_id = $1
`

func (q *Queries) GetReadingHistoryByUserID(ctx context.Context, userID int64) ([]ReadingHistory, error) {
	rows, err := q.db.QueryContext(ctx, getReadingHistoryByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ReadingHistory
	for rows.Next() {
		var i ReadingHistory
		if err := rows.Scan(
			&i.UserID,
			&i.BookID,
			&i.Status,
			&i.StartDate,
			&i.EndDate,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertReadingHistory = `-- name: InsertReadingHistory :exec
INSERT INTO reading_histories (user_id, book_id, status, start_date, end_date)
VALUES ($1, $2, $3, $4, $5)
`

type InsertReadingHistoryParams struct {
	UserID    int64         `json:"user_id"`
	BookID    int64         `json:"book_id"`
	Status    ReadingStatus `json:"status"`
	StartDate sql.NullTime  `json:"start_date"`
	EndDate   sql.NullTime  `json:"end_date"`
}

func (q *Queries) InsertReadingHistory(ctx context.Context, arg InsertReadingHistoryParams) error {
	_, err := q.db.ExecContext(ctx, insertReadingHistory,
		arg.UserID,
		arg.BookID,
		arg.Status,
		arg.StartDate,
		arg.EndDate,
	)
	return err
}

const updateReadingDates = `-- name: UpdateReadingDates :exec
UPDATE reading_histories
SET start_date = $3,
    end_date   = $4,
    updated_at = now()
WHERE user_id = $1
  AND book_id = $2
`

type UpdateReadingDatesParams struct {
	UserID    int64        `json:"user_id"`
	BookID    int64        `json:"book_id"`
	StartDate sql.NullTime `json:"start_date"`
	EndDate   sql.NullTime `json:"end_date"`
}

func (q *Queries) UpdateReadingDates(ctx context.Context, arg UpdateReadingDatesParams) error {
	_, err := q.db.ExecContext(ctx, updateReadingDates,
		arg.UserID,
		arg.BookID,
		arg.StartDate,
		arg.EndDate,
	)
	return err
}

const updateReadingStatus = `-- name: UpdateReadingStatus :exec
UPDATE reading_histories
SET status     = $3,
    updated_at = now()
WHERE user_id = $1
  AND book_id = $2
`

type UpdateReadingStatusParams struct {
	UserID int64         `json:"user_id"`
	BookID int64         `json:"book_id"`
	Status ReadingStatus `json:"status"`
}

func (q *Queries) UpdateReadingStatus(ctx context.Context, arg UpdateReadingStatusParams) error {
	_, err := q.db.ExecContext(ctx, updateReadingStatus, arg.UserID, arg.BookID, arg.Status)
	return err
}
