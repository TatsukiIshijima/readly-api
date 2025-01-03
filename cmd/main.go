package main

import (
	"log"
	"path/filepath"
	"readly/api"
	sqlc "readly/db/sqlc"
	"readly/env"
	"readly/repository"

	_ "github.com/lib/pq"
)

func main() {
	config, err := env.Load(filepath.Join(env.ProjectRoot(), "/env"))
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	a := sqlc.Adapter{}
	db, q := a.Connect(config.DBDriver, config.DBSource)
	bookRepo := repository.NewBookRepository(db, q)
	server := api.NewServer(bookRepo)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
