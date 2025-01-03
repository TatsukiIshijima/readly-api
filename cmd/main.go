package main

import (
	"database/sql"
	"log"
	"readly/api"
	db "readly/db/sqlc"
	"readly/env"
	"readly/repository"

	_ "github.com/lib/pq"
)

func main() {
	config, err := env.Load("./env")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	bookRepo := repository.NewBookRepository(conn, db.New(conn))
	server := api.NewServer(bookRepo)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
