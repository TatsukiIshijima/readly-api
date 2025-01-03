package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Connector interface {
	Connect(dbDriver string, dbSource string) (DBTX, Querier)
}

type Adapter struct{}

type FakeAdapter struct{}

func (a *Adapter) Connect(dbDriver string, dbSource string) (DBTX, Querier) {
	db, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	q := New(db)
	return db, q
}

func (f *FakeAdapter) Connect(string, string) (DBTX, Querier) {
	db := FakeDB{}
	q := &FakeQuerier{}
	return db, q
}
