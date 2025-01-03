package db

import (
	"context"
	"database/sql"
	"errors"
)

type Container struct {
	db      DBTX
	Querier Querier
}

func NewContainer(db DBTX, q Querier) Container {
	return Container{
		db:      db,
		Querier: q,
	}
}

type Transactor interface {
	Exec(ctx context.Context, fn func(queries Querier) error) error
}

func (c Container) Exec(ctx context.Context, fn func(queries Querier) error) error {
	switch c.db.(type) {
	case *sql.DB:
		return c.execTx(ctx, fn)
	default:
		return c.execFakeTx(fn)
	}
}

func (c Container) execTx(ctx context.Context, fn func(queries Querier) error) error {
	db, ok := c.db.(*sql.DB)
	if !ok {
		return errors.New("not implemented")
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	err = fn(c.Querier)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}

func (c Container) execFakeTx(fn func(queries Querier) error) error {
	return fn(c.Querier)
}
