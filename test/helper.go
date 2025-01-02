package test

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"path/filepath"
	sqlc "readly/db/sqlc"
	"readly/env"
)

// TODO:ファイル名変更(helper→db_adapter)&dbパッケージへ移動

type DBConnector interface {
	Connect() (sqlc.DBTX, sqlc.Querier)
}

type DBAdapter struct{}

type FakeDBAdapter struct{}

func (a *DBAdapter) Connect() (sqlc.DBTX, sqlc.Querier) {
	config, err := env.Load(filepath.Join(env.ProjectRoot(), "/env"))
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	q := sqlc.New(db)
	return db, q
}

func (f *FakeDBAdapter) Connect() (sqlc.DBTX, sqlc.Querier) {
	db := sqlc.FakeDB{}
	q := &sqlc.FakeQuerier{}
	return db, q
}
