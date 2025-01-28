package service

import (
	"github.com/gin-gonic/gin"
	"readly/auth"
	"readly/env"
)

type Server struct {
	config      env.Config
	bookService BookService
	userService UserService
	maker       auth.Maker
	router      *gin.Engine
}

func NewServer(config env.Config, bookService BookService, userService UserService) (*Server, error) {
	maker, err := auth.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, err
	}
	server := &Server{
		config:      config,
		bookService: bookService,
		userService: userService,
		maker:       maker,
	}
	router := gin.Default()

	router.POST("/books", server.bookService.Register)
	router.DELETE("/books", server.bookService.Delete)
	router.POST("/signup", server.userService.SignUp)
	router.POST("/signin", server.userService.SignIn)

	server.router = router
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
