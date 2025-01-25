package service

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	bookService BookService
	router      *gin.Engine
}

func NewServer(bookService BookService) *Server {
	server := &Server{bookService: bookService}
	router := gin.Default()

	router.POST("/books", server.bookService.Register)
	router.DELETE("/books", server.bookService.Delete)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
