package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"path/filepath"
	"readly/env"
)

type Connector interface {
	Connect() (DBTX, Querier)
}

type Adapter struct{}

type FakeAdapter struct{}

func (a *Adapter) Connect() (DBTX, Querier) {
	config, err := env.Load(filepath.Join(env.ProjectRoot(), "/env"))
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	q := New(db)
	return db, q
}

func (f *FakeAdapter) Connect() (DBTX, Querier) {
	db := FakeDB{}
	q := &FakeQuerier{}
	return db, q
}
