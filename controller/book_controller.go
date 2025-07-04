package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"readly/entity"
	"readly/middleware/auth"
	"readly/usecase"
)

type BookController interface {
	Register(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type BookControllerImpl struct {
	registerUseCase usecase.RegisterBookUseCase
	deleteUseCase   usecase.DeleteBookUseCase
}

func NewBookController(
	registerUseCase usecase.RegisterBookUseCase,
	deleteUseCase usecase.DeleteBookUseCase,
) BookController {
	return &BookControllerImpl{
		registerUseCase: registerUseCase,
		deleteUseCase:   deleteUseCase,
	}
}

type RegisterBookRequest struct {
	Title         string               `json:"title" binding:"required,min=1"`
	Genres        []string             `json:"genres" binding:"omitempty,max=5"`
	Description   *string              `json:"description" binding:"omitempty,max=500"`
	CoverImageURL *string              `json:"cover_image_url" binding:"omitempty,url,max=2048"`
	URL           *string              `json:"url" binding:"omitempty,url,max=2048"`
	AuthorName    *string              `json:"author_name" binding:"omitempty,max=255"`
	PublisherName *string              `json:"publisher_name" binding:"omitempty,max=255"`
	PublishDate   *entity.Date         `json:"publish_date" binding:"omitempty"`
	ISBN          *string              `json:"isbn" binding:"omitempty,isbn"`
	Status        entity.ReadingStatus `json:"status" binding:"required"`
	StartDate     *entity.Date         `json:"start_date" binding:"omitempty"`
	EndDate       *entity.Date         `json:"end_date" binding:"omitempty"`
}

func (bc *BookControllerImpl) Register(ctx *gin.Context) {
	var req RegisterBookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	claims := ctx.MustGet(auth.AuthorizationClaimKey).(*auth.Claims)

	args := usecase.RegisterBookRequest{
		UserID:        claims.UserID,
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
	book, err := bc.registerUseCase.RegisterBook(ctx, args)
	if err != nil {
		c, e := toHttpStatusCode(err)
		ctx.JSON(c, errorResponse(e))
		return
	}

	ctx.JSON(http.StatusOK, book)
}

type DeleteBookRequest struct {
	BookID int64 `json:"book_id" binding:"required"`
}

func (bc *BookControllerImpl) Delete(ctx *gin.Context) {
	var req DeleteBookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	claims := ctx.MustGet(auth.AuthorizationClaimKey).(*auth.Claims)

	args := usecase.DeleteBookRequest{
		UserID: claims.UserID,
		BookID: req.BookID,
	}
	err := bc.deleteUseCase.DeleteBook(ctx, args)
	if err != nil {
		c, e := toHttpStatusCode(err)
		ctx.JSON(c, errorResponse(e))
		return
	}

	ctx.Status(http.StatusOK)
}
