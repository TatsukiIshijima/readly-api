package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"path/filepath"
	"readly/env"
)

type DBConnector interface {
	Connect() (DBTX, Querier)
}

type DBAdapter struct{}

type FakeDBAdapter struct{}

func (a *DBAdapter) Connect() (DBTX, Querier) {
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

func (f *FakeDBAdapter) Connect() (DBTX, Querier) {
	db := FakeDB{}
	q := &FakeQuerier{}
	return db, q
}
