package test

import (
	"database/sql"
	"log"
	"math/rand"
	"readly/db/sqlc"
	"strings"

	_ "github.com/lib/pq"
)

const (
	dbDriver  = "postgres"
	dbSource  = "postgresql://root:secret@localhost:5432/readly?sslmode=disable"
	alplhabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var DB *sql.DB
var Queries *db.Queries

func RandomInt(min int64, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alplhabet)

	for i := 0; i < n; i++ {
		c := alplhabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func Connect() {
	var err error
	DB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	Queries = db.New(DB)
}
