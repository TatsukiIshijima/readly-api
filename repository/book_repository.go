package repository

import (
	"context"
	"database/sql"
	"errors"
	sqlc "readly/db/sqlc"
)

type BookRepository interface {
	CreateAuthor(ctx context.Context, req CreateAuthorRequest) (*CreateAuthorResponse, error)
	CreateBook(ctx context.Context, req CreateBookRequest) (*CreateBookResponse, error)
	CreateBookGenre(ctx context.Context, req CreateBookGenreRequest) (*CreateBookGenreResponse, error)
	CreateGenre(ctx context.Context, req CreateGenreRequest) (*CreateGenreResponse, error)
	CreatePublisher(ctx context.Context, req CreatePublisherRequest) (*CreatePublisherResponse, error)
	DeleteAuthor(ctx context.Context, req DeleteAuthorRequest) error
	DeleteBook(ctx context.Context, req DeleteBookRequest) error
	DeleteBookGenre(ctx context.Context, req DeleteBookGenreRequest) error
	DeleteGenre(ctx context.Context, req DeleteGenreRequest) error
	DeletePublisher(ctx context.Context, req DeletePublisherRequest) error
	GetBookByID(ctx context.Context, req GetBookRequest) (*GetBookResponse, error)
	GetGenresByBookID(ctx context.Context, request GetGenresByBookIDRequest) (*GetGenresByBookIDResponse, error)
}

type BookRepositoryImpl struct {
	querier sqlc.Querier
}

func NewBookRepository(q sqlc.Querier) BookRepository {
	return &BookRepositoryImpl{
		querier: q,
	}
}

func (r *BookRepositoryImpl) CreateAuthor(ctx context.Context, req CreateAuthorRequest) (*CreateAuthorResponse, error) {
	res, err := r.querier.CreateAuthor(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	return newCreateAuthorResponseFromSQLC(res), nil
}

func (r *BookRepositoryImpl) CreateBook(ctx context.Context, req CreateBookRequest) (*CreateBookResponse, error) {
	p := req.toSQLC()
	b, err := r.querier.CreateBook(ctx, p)
	if err != nil {
		return nil, err
	}
	return newCreateBookResponseFromSQLC(b), nil
}

type CreateBookGenreRequest struct {
	BookID    int64
	GenreName string
}

func (r CreateBookGenreRequest) toParams() sqlc.CreateBookGenreParams {
	return sqlc.CreateBookGenreParams{
		BookID:    r.BookID,
		GenreName: r.GenreName,
	}
}

type CreateBookGenreResponse struct {
	BookID    int64
	GenreName string
}

func newCreateBookGenreResponse(b sqlc.BookGenre) *CreateBookGenreResponse {
	return &CreateBookGenreResponse{
		BookID:    b.BookID,
		GenreName: b.GenreName,
	}
}

func (r *BookRepositoryImpl) CreateBookGenre(ctx context.Context, req CreateBookGenreRequest) (*CreateBookGenreResponse, error) {
	p := req.toParams()
	b, err := r.querier.CreateBookGenre(ctx, p)
	if err != nil {
		return nil, err
	}
	return newCreateBookGenreResponse(b), nil
}

func (r *BookRepositoryImpl) CreateGenre(ctx context.Context, req CreateGenreRequest) (*CreateGenreResponse, error) {
	g, err := r.querier.CreateGenre(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	return newCreateGenreResponseFromSQLC(g), nil
}

func (r *BookRepositoryImpl) CreatePublisher(ctx context.Context, req CreatePublisherRequest) (*CreatePublisherResponse, error) {
	p, err := r.querier.CreatePublisher(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	return newCreatePublisherResponseFromSQLC(p), nil
}

func (r *BookRepositoryImpl) DeleteAuthor(ctx context.Context, req DeleteAuthorRequest) error {
	return r.querier.DeleteAuthor(ctx, req.Name)
}

func (r *BookRepositoryImpl) DeleteBook(ctx context.Context, req DeleteBookRequest) error {
	rowsAffected, err := r.querier.DeleteBook(ctx, req.ID)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrNoRowsDeleted
	}
	return nil
}

type DeleteBookGenreRequest struct {
	BookID    int64
	GenreName string
}

func (r *BookRepositoryImpl) DeleteBookGenre(ctx context.Context, req DeleteBookGenreRequest) error {
	p := sqlc.DeleteBookGenreParams{
		BookID:    req.BookID,
		GenreName: req.GenreName,
	}
	rowsAffected, err := r.querier.DeleteBookGenre(ctx, p)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrNoRowsDeleted
	}
	return nil
}

func (r *BookRepositoryImpl) DeleteGenre(ctx context.Context, req DeleteGenreRequest) error {
	return r.querier.DeleteGenre(ctx, req.Name)
}

func (r *BookRepositoryImpl) DeletePublisher(ctx context.Context, req DeletePublisherRequest) error {
	return r.querier.DeletePublisher(ctx, req.Name)
}

func (r *BookRepositoryImpl) GetBookByID(ctx context.Context, req GetBookRequest) (*GetBookResponse, error) {
	b, err := r.querier.GetBooksByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRows
		}
		return nil, err
	}
	return newGetBookResponseFromSQLC(b), nil
}

func (r *BookRepositoryImpl) GetGenresByBookID(ctx context.Context, req GetGenresByBookIDRequest) (*GetGenresByBookIDResponse, error) {
	g, err := r.querier.GetGenresByBookID(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return newGetGenresByBookIDResponse(g), err
}
