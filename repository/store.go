package repository

import (
	"context"
	"database/sql"
	"fmt"
	sqlc "readly/db/sqlc"
)

type Store interface {
	sqlc.Querier
	execTx(ctx context.Context, fn func(queries sqlc.Querier) error) error
}

type DefaultStore struct {
	*sqlc.Queries
	db *sql.DB
}

type FakeStore struct {
	*sqlc.FakeQuerier
	db sqlc.FakeDB
}

func NewStore(dbtx sqlc.DBTX) Store {
	switch db := dbtx.(type) {
	case *sql.DB:
		return &DefaultStore{
			Queries: sqlc.New(db),
			db:      db,
		}
	default:
		return &FakeStore{
			FakeQuerier: &sqlc.FakeQuerier{},
			db:          sqlc.FakeDB{},
		}
	}
}

func (store *DefaultStore) execTx(ctx context.Context, fn func(queries sqlc.Querier) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := sqlc.New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

func (store *FakeStore) execTx(_ context.Context, fn func(queries sqlc.Querier) error) error {
	err := fn(store.FakeQuerier)
	if err != nil {
		return err
	}
	return nil
}
