package test

import (
	"context"
	"database/sql"
	"log"
	"math/rand"
	"path/filepath"
	"readly/db/sqlc"
	"readly/env"
	"strings"

	_ "github.com/lib/pq"
)

const (
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
	config, err := env.Load(filepath.Join(env.ProjectRoot(), "/env"))
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	DB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	Queries = db.New(DB)
}

func CreateRandomUser() (db.User, error) {
	arg := db.CreateUserParams{
		Name:           RandomString(12),
		Email:          RandomString(6) + "@example.com",
		HashedPassword: RandomString(16),
	}
	return Queries.CreateUser(context.Background(), arg)
}
