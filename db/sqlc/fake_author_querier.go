package db

import (
	"context"
	"database/sql"
	"github.com/lib/pq"
)

type AuthorTable struct {
	Columns []Author
}

var authorTable = AuthorTable{}

func (q *FakeQuerier) CreateAuthor(_ context.Context, name string) (Author, error) {
	for _, author := range authorTable.Columns {
		if author.Name == name {
			return Author{}, &pq.Error{Code: "23505", Message: "duplicate key value violates unique constraint"}
		}
	}
	a := Author{
		Name: name,
	}
	authorTable.Columns = append(authorTable.Columns, a)
	return a, nil
}

func (q *FakeQuerier) DeleteAuthor(_ context.Context, name string) error {
	for i, author := range authorTable.Columns {
		if author.Name == name {
			authorTable.Columns = append(authorTable.Columns[:i], authorTable.Columns[i+1:]...)
			return nil
		}
	}
	return nil
}

func (q *FakeQuerier) GetAllAuthors(_ context.Context, arg GetAllAuthorsParams) ([]Author, error) {
	// TODO: Implement
	return nil, nil
}

func (q *FakeQuerier) GetAuthorByName(_ context.Context, name string) (Author, error) {
	for _, author := range authorTable.Columns {
		if author.Name == name {
			return author, nil
		}
	}
	return Author{}, sql.ErrNoRows
}
