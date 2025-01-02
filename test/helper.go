package test

import (
	"database/sql"
	"log"
	"math/rand"
	"path/filepath"
	sqlc "readly/db/sqlc"
	"readly/env"
	"strings"

	_ "github.com/lib/pq"
)

// TODO:ファイル名変更(helper→db_adapter)&dbパッケージへ移動

// FIXME:移動
const (
	alplhabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

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

// RandomInt FIXME:移動
func RandomInt(min int64, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString FIXME:移動
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alplhabet)

	for i := 0; i < n; i++ {
		c := alplhabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}
