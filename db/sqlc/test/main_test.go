package db

import (
	"database/sql"
	"log"
	"math/rand"
	"os"
	db "readly/db/sqlc"
	"strings"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

const (
	dbDriver  = "postgres"
	dbSource  = "postgresql://root:secret@localhost:5432/readly?sslmode=disable"
	alplhabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var testQueries *db.Queries

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func randomInt(min int64, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func randomString(n int) string {
	var sb strings.Builder
	k := len(alplhabet)

	for i := 0; i < n; i++ {
		c := alplhabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = db.New(conn)

	os.Exit(m.Run())
}
