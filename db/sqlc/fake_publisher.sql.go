package db

import (
	"context"
	"database/sql"
	"github.com/lib/pq"
)

type PublisherTable struct {
	Columns []Publisher
}

var publisherTable = PublisherTable{}

func (q *FakeQuerier) CreatePublisher(_ context.Context, name string) (Publisher, error) {
	for _, publisher := range publisherTable.Columns {
		if publisher.Name == name {
			return Publisher{}, &pq.Error{Code: "23505", Message: "duplicate key value violates unique constraint"}
		}
	}
	p := Publisher{
		Name: name,
	}
	publisherTable.Columns = append(publisherTable.Columns, p)
	return p, nil
}

func (q *FakeQuerier) DeletePublisher(_ context.Context, name string) error {
	for i, publisher := range publisherTable.Columns {
		if publisher.Name == name {
			publisherTable.Columns = append(publisherTable.Columns[:i], publisherTable.Columns[i+1:]...)
			return nil
		}
	}
	return nil
}

func (q *FakeQuerier) GetAllPublishers(_ context.Context, arg GetAllPublishersParams) ([]Publisher, error) {
	// TODO: Implement
	return nil, nil
}

func (q *FakeQuerier) GetPublisherByName(_ context.Context, name string) (Publisher, error) {
	for _, publisher := range publisherTable.Columns {
		if publisher.Name == name {
			return publisher, nil
		}
	}
	return Publisher{}, sql.ErrNoRows
}
