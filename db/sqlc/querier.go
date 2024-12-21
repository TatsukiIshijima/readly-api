// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	DeleteGenre(ctx context.Context, id int32) error
	GetAllGenres(ctx context.Context) ([]Genre, error)
	GetGenreByID(ctx context.Context, id int32) (Genre, error)
	InsertGenre(ctx context.Context, name string) error
	UpdateGenre(ctx context.Context, arg UpdateGenreParams) error
}

var _ Querier = (*Queries)(nil)
