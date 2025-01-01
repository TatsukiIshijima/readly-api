// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: author.sql

package db

import (
	"context"
)

const createAuthor = `-- name: CreateAuthor :one
INSERT INTO authors (name)
VALUES ($1) RETURNING name, created_at
`

func (q *Queries) CreateAuthor(ctx context.Context, name string) (Author, error) {
	row := q.db.QueryRowContext(ctx, createAuthor, name)
	var i Author
	err := row.Scan(&i.Name, &i.CreatedAt)
	return i, err
}

const deleteAuthor = `-- name: DeleteAuthor :exec
DELETE
FROM authors
WHERE name = $1
`

func (q *Queries) DeleteAuthor(ctx context.Context, name string) error {
	_, err := q.db.ExecContext(ctx, deleteAuthor, name)
	return err
}

const getAllAuthors = `-- name: GetAllAuthors :many
SELECT name, created_at
FROM authors LIMIT $1
OFFSET $2
`

type GetAllAuthorsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetAllAuthors(ctx context.Context, arg GetAllAuthorsParams) ([]Author, error) {
	rows, err := q.db.QueryContext(ctx, getAllAuthors, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Author{}
	for rows.Next() {
		var i Author
		if err := rows.Scan(&i.Name, &i.CreatedAt); err != nil {
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

const getAuthorByName = `-- name: GetAuthorByName :one
SELECT name, created_at
FROM authors
WHERE name = $1
`

func (q *Queries) GetAuthorByName(ctx context.Context, name string) (Author, error) {
	row := q.db.QueryRowContext(ctx, getAuthorByName, name)
	var i Author
	err := row.Scan(&i.Name, &i.CreatedAt)
	return i, err
}
