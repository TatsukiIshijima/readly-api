package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"readly/domain"
	"readly/repository"
	"time"
)

type RegisterRequest struct {
	// TODO:削除（ログイン情報から取得するため）
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

type RegisterResponse struct {
	domain.Book
}

func (server *Server) registerBook(ctx *gin.Context) {
	var req RegisterRequest
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

	ctx.JSON(http.StatusOK, RegisterResponse{Book: book})
}

type GetBookRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

type GetBookResponse struct {
	domain.Book
}

func (server *Server) getBook(ctx *gin.Context) {
	var req GetBookRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	book, err := server.bookRepo.Get(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, GetBookResponse{Book: *book})
}

type ListBooksRequest struct {
	// limitとoffsetのままだと0が始点になる。requiredタグを使用した場合
	// int32のゼロ値は無効な値として扱われるため、ページング処理を入れる
	Page int32 `form:"page" binding:"required,min=1"`
	Size int32 `form:"size" binding:"required,min=10,max=50"`
}

// TODO:削除（ログイン情報から取得するため）
type ListBooksTempRequest struct {
	UserID int64 `json:"user_id" binding:"required"`
}

func (server *Server) listBook(ctx *gin.Context) {
	var req ListBooksRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var tempReq ListBooksTempRequest
	if err := ctx.ShouldBindJSON(&tempReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	books, err := server.bookRepo.List(ctx, repository.ListRequest{
		UserID: tempReq.UserID,
		Limit:  req.Size,
		Offset: (req.Page - 1) * req.Size,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, books)
}

type DeleteBookRequest struct {
	// TODO:削除（ログイン情報から取得するため）
	UserID int64 `json:"user_id" binding:"required"`
	BookID int64 `json:"book_id" binding:"required"`
}

func (server *Server) deleteBook(ctx *gin.Context) {
	var req DeleteBookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := repository.DeleteRequest{
		UserID: req.UserID,
		BookID: req.BookID,
	}
	err := server.bookRepo.Delete(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.Status(http.StatusOK)
}
