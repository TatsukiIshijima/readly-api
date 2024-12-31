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

	//router.POST("/books", server.createBook)
	//router.GET("/books", server.listBooks)
	//router.GET("/books/:id", server.getBook)
	//router.PUT("/books/:id", server.updateBook)
	//router.DELETE("/books/:id", server.deleteBook)

	server.router = router
	return server
}
