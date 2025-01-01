package main

import (
	"database/sql"
	"log"
	"readly/api"
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

	store := repository.NewStore(conn)
	bookRepo := repository.NewBookRepository(store)
	server := api.NewServer(bookRepo)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
