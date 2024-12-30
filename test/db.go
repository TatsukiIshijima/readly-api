package test

import (
	"database/sql"
	db "readly/db/sqlc"

	_ "github.com/lib/pq"
)

const (
	DBDriver = "postgres"
	DBSource = "postgresql://root:secret@localhost:5432/readly?sslmode=disable"
)

var DB *sql.DB
var Queries *db.Queries
