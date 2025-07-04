// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: genre.sql

package db

import (
	"context"
)

const createGenre = `-- name: CreateGenre :one
INSERT INTO genres (name)
VALUES ($1) RETURNING name, created_at
`

func (q *Queries) CreateGenre(ctx context.Context, name string) (Genre, error) {
	row := q.db.QueryRowContext(ctx, createGenre, name)
	var i Genre
	err := row.Scan(&i.Name, &i.CreatedAt)
	return i, err
}

const deleteGenre = `-- name: DeleteGenre :exec
DELETE
FROM genres
WHERE name = $1
`

func (q *Queries) DeleteGenre(ctx context.Context, name string) error {
	_, err := q.db.ExecContext(ctx, deleteGenre, name)
	return err
}

const getAllGenres = `-- name: GetAllGenres :many
SELECT name
FROM genres
`

func (q *Queries) GetAllGenres(ctx context.Context) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getAllGenres)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []string{}
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		items = append(items, name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getGenreByName = `-- name: GetGenreByName :one
SELECT name, created_at
FROM genres
WHERE name = $1
`

func (q *Queries) GetGenreByName(ctx context.Context, name string) (Genre, error) {
	row := q.db.QueryRowContext(ctx, getGenreByName, name)
	var i Genre
	err := row.Scan(&i.Name, &i.CreatedAt)
	return i, err
}
