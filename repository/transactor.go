package repository

import (
	"context"
	"database/sql"
	sqlc "readly/db/sqlc"
)

type Transactor interface {
	Exec(ctx context.Context, fn func() error) error
}

type TransactorImpl struct {
	db sqlc.DBTX
}

func (t TransactorImpl) Exec(ctx context.Context, fn func() error) error {
	switch t.db.(type) {
	case *sql.DB:
		return t.execTx(ctx, fn)
	default:
		return t.execFakeTx(fn)
	}
}

func (t TransactorImpl) execTx(ctx context.Context, fn func() error) error {
	db, ok := t.db.(*sql.DB)
	if !ok {
		panic("not implemented")
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	err = fn()
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}

func (t TransactorImpl) execFakeTx(fn func() error) error {
	return fn()
}
