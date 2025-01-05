package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
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

func handle(err error) (int, error) {
	var e *repository.Error
	if !errors.As(err, &e) {
		return http.StatusInternalServerError, err
	}
	var sc int
	switch e.Code {
	case repository.BadRequest:
		sc = http.StatusBadRequest
	case repository.Forbidden:
		sc = http.StatusForbidden
	case repository.NotFound:
		sc = http.StatusNotFound
	case repository.Conflict:
		sc = http.StatusConflict
	case repository.Internal:
		sc = http.StatusInternalServerError
	default:
		sc = http.StatusInternalServerError
	}
	return sc, e
}
