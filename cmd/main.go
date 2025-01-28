package main

import (
	_ "github.com/lib/pq"
	"log"
	"path/filepath"
	sqlc "readly/db/sqlc"
	"readly/env"
	"readly/repository"
	"readly/service"
	"readly/usecase"
)

func main() {
	config, err := env.Load(filepath.Join(env.ProjectRoot(), "/env"))
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	a := sqlc.Adapter{}
	db, q := a.Connect(config.DBDriver, config.DBSource)
	t := repository.New(db)
	bookRepo := repository.NewBookRepository(q)
	userRepo := repository.NewUserRepository(q)
	readingHistoryRepo := repository.NewReadingHistoryRepository(q)
	registerBookUseCase := usecase.NewRegisterBookUseCase(t, bookRepo, readingHistoryRepo, userRepo)
	deleteBookUseCase := usecase.NewDeleteBookUseCase(t, bookRepo, readingHistoryRepo, userRepo)
	signUpUseCase := usecase.NewSignUpUseCase(userRepo)
	signInUseCase := usecase.NewSignInUseCase(userRepo)
	bookService := service.NewBookService(registerBookUseCase, deleteBookUseCase)
	userService := service.NewUserService(signUpUseCase, signInUseCase)
	server, err := service.NewServer(config, bookService, userService)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
