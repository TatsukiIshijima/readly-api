package api

import (
	"os"
	sqlc "readly/db/sqlc"
	"readly/repository"
	"testing"
)

var server *Server

func TestMain(m *testing.M) {
	fa := sqlc.FakeAdapter{}
	db, q := fa.Connect("", "")
	bookRepo := repository.NewBookRepository(db, q)
	server = NewServer(bookRepo)
	os.Exit(m.Run())
}
