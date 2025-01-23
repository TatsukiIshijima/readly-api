package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"readly/usecase"
	"time"
)

type BookService interface {
	Register(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type BookServiceImpl struct {
	registerUseCase usecase.RegisterBookUseCase
	deleteUseCase   usecase.DeleteBookUseCase
}

func NewBookService(registerUseCase usecase.RegisterBookUseCase, deleteUseCase usecase.DeleteBookUseCase) BookServiceImpl {
	return BookServiceImpl{
		registerUseCase: registerUseCase,
		deleteUseCase:   deleteUseCase,
	}
}

type RegisterBookRequest struct {
	UserID        int64
	Title         string
	Genres        []string
	Description   *string
	CoverImageURL *string
	URL           *string
	AuthorName    *string
	PublisherName *string
	PublishDate   *time.Time
	ISBN          *string
	Status        int
	StartDate     *time.Time
	EndDate       *time.Time
}

func (s BookServiceImpl) Register(ctx *gin.Context) {
	var req RegisterBookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := usecase.RegisterBookRequest{
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
		Status:        req.Status,
		StartDate:     req.StartDate,
		EndDate:       req.EndDate,
	}
	book, err := s.registerUseCase.RegisterBook(ctx, args)
	if err != nil {
		c, e := toHttpStatusCode(err)
		ctx.JSON(c, errorResponse(e))
		return
	}

	ctx.JSON(http.StatusOK, book)
}

type DeleteBookRequest struct {
	UserID int64
	BookID int64
}

func (s BookServiceImpl) Delete(ctx *gin.Context) {
	var req DeleteBookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := usecase.DeleteBookRequest{
		UserID: req.UserID,
		BookID: req.BookID,
	}
	err := s.deleteUseCase.DeleteBook(ctx, args)
	if err != nil {
		c, e := toHttpStatusCode(err)
		ctx.JSON(c, errorResponse(e))
		return
	}

	ctx.Status(http.StatusOK)
}
