package main

import (
	_ "github.com/lib/pq"
	"log"
	"path/filepath"
	"readly/controller"
	sqlc "readly/db/sqlc"
	"readly/env"
	"readly/middleware"
	"readly/repository"
	"readly/router"
	"readly/service/auth"
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
	sessionRepo := repository.NewSessionRepository(q)

	maker, err := auth.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

	registerBookUseCase := usecase.NewRegisterBookUseCase(t, bookRepo, readingHistoryRepo, userRepo)
	deleteBookUseCase := usecase.NewDeleteBookUseCase(t, bookRepo, readingHistoryRepo, userRepo)
	signUpUseCase := usecase.NewSignUpUseCase(config, maker, t, sessionRepo, userRepo)
	signInUseCase := usecase.NewSignInUseCase(config, maker, sessionRepo, userRepo)

	bookController := controller.NewBookController(registerBookUseCase, deleteBookUseCase)
	userController := controller.NewUserController(config, maker, signUpUseCase, signInUseCase)

	r := router.Setup(middleware.Authorize(maker), bookController, userController)
	err = r.Run(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
