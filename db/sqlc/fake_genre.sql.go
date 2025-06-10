//go:build test

package db

import (
	"context"
	"database/sql"
	"github.com/lib/pq"
)

type GenreTable struct {
	Columns []Genre
}

var genreTable = GenreTable{}

func (q *FakeQuerier) CreateGenre(_ context.Context, name string) (Genre, error) {
	for _, genre := range genreTable.Columns {
		if genre.Name == name {
			return Genre{}, &pq.Error{Code: "23505", Message: "duplicate key value violates unique constraint"}
		}
	}
	g := Genre{
		Name: name,
	}
	genreTable.Columns = append(genreTable.Columns, g)
	return g, nil
}

func (q *FakeQuerier) DeleteGenre(_ context.Context, name string) error {
	for i, genre := range genreTable.Columns {
		if genre.Name == name {
			genreTable.Columns = append(genreTable.Columns[:i], genreTable.Columns[i+1:]...)
			return nil
		}
	}
	return nil
}

func (q *FakeQuerier) GetAllGenres(_ context.Context) ([]string, error) {
	genres := make([]string, len(genreTable.Columns))
	for i, genre := range genreTable.Columns {
		genres[i] = genre.Name
	}
	return genres, nil
}

func (q *FakeQuerier) GetGenreByName(_ context.Context, name string) (Genre, error) {
	for _, genre := range genreTable.Columns {
		if genre.Name == name {
			return genre, nil
		}
	}
	return Genre{}, sql.ErrNoRows
}
