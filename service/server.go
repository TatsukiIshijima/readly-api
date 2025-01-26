package service

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	bookService BookService
	userService UserService
	router      *gin.Engine
}

func NewServer(bookService BookService, userService UserService) *Server {
	server := &Server{bookService: bookService, userService: userService}
	router := gin.Default()

	router.POST("/books", server.bookService.Register)
	router.DELETE("/books", server.bookService.Delete)
	router.POST("/users/signup", server.userService.SignUp)
	router.POST("/users/signin", server.userService.SignIn)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
