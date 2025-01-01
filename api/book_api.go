package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"readly/domain"
	"readly/repository"
	"time"
)

type registerRequest struct {
	UserID        int64     `json:"user_id" binding:"required"`
	Title         string    `json:"title" binding:"required"`
	Genres        []string  `json:"genres"`
	Description   string    `json:"description"`
	CoverImageURL string    `json:"cover_image_url"`
	URL           string    `json:"url"`
	AuthorName    string    `json:"author_name"`
	PublisherName string    `json:"publisher_name"`
	PublishDate   time.Time `json:"publish_date"`
	ISBN          string    `json:"isbn"`
}

type registerResponse struct {
	domain.Book
}

func (server *Server) registerBook(ctx *gin.Context) {
	var req registerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := repository.RegisterRequest{
		UserID:        req.UserID,
		Title:         req.Title,
		Genres:        req.Genres,
		Description:   req.Description,
		CoverImageURL: req.CoverImageURL,
		URL:           req.URL,
		AuthorName:    req.AuthorName,
		PublisherName: req.PublisherName,
		PublishDate:   req.PublishDate,
		ISBN:          req.ISBN,
	}
	book, err := server.bookRepo.Register(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := registerResponse{Book: book}
	ctx.JSON(http.StatusOK, res)
}

type getBookRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

func (server *Server) getBook(ctx *gin.Context) {
	var req getBookRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	book, err := server.bookRepo.Get(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, book)
}
