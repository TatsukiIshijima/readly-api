package repository

import (
	"context"
	"database/sql"
	sqlc "readly/db/sqlc"
	"strings"
	"time"
)

type Repository interface {
	CreateAuthor(ctx context.Context, name string) (*string, error)
	CreateBook(ctx context.Context, req CreateBookRequest) (*CreateBookResponse, error)
	CreateBookGenre(ctx context.Context, req CreateBookGenreRequest) (*CreateBookGenreResponse, error)
	CreateGenre(ctx context.Context, name string) (*string, error)
	CreatePublisher(ctx context.Context, name string) (*string, error)
	DeleteAuthor(ctx context.Context, name string) error
	DeleteBook(ctx context.Context, id int64) error
	DeleteBookGenre(ctx context.Context, req DeleteBookGenreRequest) error
	DeleteGenre(ctx context.Context, name string) error
	DeletePublisher(ctx context.Context, name string) error
	GetBookByID(ctx context.Context, id int64) (*GetBookResponse, error)
}

type RepositoryImpl struct {
	querier sqlc.Querier
}

func NewBookRepository(q sqlc.Querier) Repository {
	return RepositoryImpl{
		querier: q,
	}
}

func (r RepositoryImpl) CreateAuthor(ctx context.Context, name string) (*string, error) {
	author, err := r.querier.CreateAuthor(ctx, name)
	if err != nil {
		return nil, err
	}
	return &author.Name, nil
}

type CreateBookRequest struct {
	Title         string
	Description   *string
	CoverImageURL *string
	URL           *string
	Author        *string
	Publisher     *string
	PublishDate   *time.Time
	ISBN          *string
}

func (r CreateBookRequest) toParams() sqlc.CreateBookParams {
	return sqlc.CreateBookParams{
		Title:         sql.NullString{String: r.Title, Valid: true},
		Description:   sql.NullString{String: *r.Description, Valid: r.Description != nil},
		CoverImageUrl: sql.NullString{String: *r.CoverImageURL, Valid: r.CoverImageURL != nil},
		Url:           sql.NullString{String: *r.URL, Valid: r.URL != nil},
		AuthorName:    sql.NullString{String: *r.Author, Valid: r.Author != nil},
		PublisherName: sql.NullString{String: *r.Publisher, Valid: r.Publisher != nil},
		PublishedDate: sql.NullTime{Time: *r.PublishDate, Valid: r.PublishDate != nil},
		Isbn:          sql.NullString{String: *r.ISBN, Valid: r.ISBN != nil},
	}
}

type CreateBookResponse struct {
	ID            int64
	Title         *string
	Description   *string
	CoverImageURL *string
	URL           *string
	Author        *string
	Publisher     *string
	PublishDate   *time.Time
	ISBN          *string
}

// FIXME:移動
func newString(ns sql.NullString) *string {
	if !ns.Valid {
		return nil
	}
	return &ns.String
}

// FIXME:移動
func newTime(nt sql.NullTime) *time.Time {
	if !nt.Valid {
		return nil
	}
	return &nt.Time
}

func newCreateResponse(b sqlc.Book) *CreateBookResponse {
	return &CreateBookResponse{
		ID:            b.ID,
		Title:         newString(b.Title),
		Description:   newString(b.Description),
		CoverImageURL: newString(b.CoverImageUrl),
		URL:           newString(b.Url),
		Author:        newString(b.AuthorName),
		Publisher:     newString(b.PublisherName),
		PublishDate:   newTime(b.PublishedDate),
		ISBN:          newString(b.Isbn),
	}
}

func (r RepositoryImpl) CreateBook(ctx context.Context, req CreateBookRequest) (*CreateBookResponse, error) {
	p := req.toParams()
	b, err := r.querier.CreateBook(ctx, p)
	if err != nil {
		return nil, err
	}
	return newCreateResponse(b), nil
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

func (r RepositoryImpl) CreateBookGenre(ctx context.Context, req CreateBookGenreRequest) (*CreateBookGenreResponse, error) {
	p := req.toParams()
	b, err := r.querier.CreateBookGenre(ctx, p)
	if err != nil {
		return nil, err
	}
	return newCreateBookGenreResponse(b), nil
}

func (r RepositoryImpl) CreateGenre(ctx context.Context, name string) (*string, error) {
	g, err := r.querier.CreateGenre(ctx, name)
	if err != nil {
		return nil, err
	}
	return &g.Name, nil
}

func (r RepositoryImpl) CreatePublisher(ctx context.Context, name string) (*string, error) {
	p, err := r.querier.CreatePublisher(ctx, name)
	if err != nil {
		return nil, err
	}
	return &p.Name, nil
}

func (r RepositoryImpl) DeleteAuthor(ctx context.Context, name string) error {
	return r.querier.DeleteAuthor(ctx, name)
}

func (r RepositoryImpl) DeleteBook(ctx context.Context, id int64) error {
	return r.querier.DeleteBook(ctx, id)
}

type DeleteBookGenreRequest struct {
	BookID    int64
	GenreName string
}

func (r RepositoryImpl) DeleteBookGenre(ctx context.Context, req DeleteBookGenreRequest) error {
	p := sqlc.DeleteBookGenreParams{
		BookID:    req.BookID,
		GenreName: req.GenreName,
	}
	return r.querier.DeleteBookGenre(ctx, p)
}

func (r RepositoryImpl) DeleteGenre(ctx context.Context, name string) error {
	return r.querier.DeleteGenre(ctx, name)
}

func (r RepositoryImpl) DeletePublisher(ctx context.Context, name string) error {
	return r.querier.DeletePublisher(ctx, name)
}

type GetBookResponse struct {
	ID            int64
	Title         *string
	Genres        []string
	Description   *string
	CoverImageURL *string
	URL           *string
	AuthorName    *string
	PublisherName *string
	PublishDate   *time.Time
	ISBN          *string
}

func newGetBookResponse(b sqlc.GetBooksByIDRow) *GetBookResponse {
	return &GetBookResponse{
		ID:            b.ID,
		Title:         newString(b.Title),
		Genres:        strings.Split(string(b.Genres), ", "),
		Description:   newString(b.Description),
		CoverImageURL: newString(b.CoverImageUrl),
		URL:           newString(b.Url),
		AuthorName:    newString(b.AuthorName),
		PublisherName: newString(b.PublisherName),
		PublishDate:   newTime(b.PublishedDate),
		ISBN:          newString(b.Isbn),
	}
}

func (r RepositoryImpl) GetBookByID(ctx context.Context, id int64) (*GetBookResponse, error) {
	b, err := r.querier.GetBooksByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return newGetBookResponse(b), nil
}
