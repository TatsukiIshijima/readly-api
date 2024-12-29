package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	DB DBTX
}

func NewStore(DB DBTX) *Store {
	return &Store{
		DB: DB,
	}
}

func (store *Store) ExecTx(ctx context.Context, fn func(*Queries) error) error {
	db, ok := store.DB.(*sql.DB)
	if !ok {
		return fmt.Errorf("store.DB is not a sql.DB")
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
