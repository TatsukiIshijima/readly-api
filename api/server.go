package api

import (
	"github.com/gin-gonic/gin"
	"readly/repository"
)

type Server struct {
	bookRepo repository.BookRepository
	router   *gin.Engine
}

func NewServer(repo repository.BookRepository) *Server {
	server := &Server{bookRepo: repo}
	router := gin.Default()

	router.POST("/books", server.registerBook)
	router.GET("/books/:id", server.getBook)
	router.GET("/books", server.listBook)
	router.DELETE("/books", server.deleteBook)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
