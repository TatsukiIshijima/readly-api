package repository

import (
	"context"
	"database/sql"
	sqlc "readly/db/sqlc"
	"strings"
	"time"
)

type BookRepository interface {
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

type BookRepositoryImpl struct {
	querier sqlc.Querier
}

func NewBookRepository(q sqlc.Querier) BookRepository {
	return BookRepositoryImpl{
		querier: q,
	}
}

func (r BookRepositoryImpl) CreateAuthor(ctx context.Context, name string) (*string, error) {
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
	desc := sql.NullString{String: "", Valid: false}
	coverImgURL := sql.NullString{String: "", Valid: false}
	URL := sql.NullString{String: "", Valid: false}
	a := sql.NullString{String: "", Valid: false}
	p := sql.NullString{String: "", Valid: false}
	pd := sql.NullTime{Time: time.Time{}, Valid: false}
	ISBN := sql.NullString{String: "", Valid: false}
	if r.Description != nil {
		desc = sql.NullString{String: *r.Description, Valid: true}
	}
	if r.CoverImageURL != nil {
		coverImgURL = sql.NullString{String: *r.CoverImageURL, Valid: true}
	}
	if r.URL != nil {
		URL = sql.NullString{String: *r.URL, Valid: true}
	}
	if r.Author != nil {
		a = sql.NullString{String: *r.Author, Valid: true}
	}
	if r.Publisher != nil {
		p = sql.NullString{String: *r.Publisher, Valid: true}
	}
	if r.PublishDate != nil {
		pd = sql.NullTime{Time: *r.PublishDate, Valid: true}
	}
	if r.ISBN != nil {
		ISBN = sql.NullString{String: *r.ISBN, Valid: true}
	}
	return sqlc.CreateBookParams{
		Title:         r.Title,
		Description:   desc,
		CoverImageUrl: coverImgURL,
		Url:           URL,
		AuthorName:    a,
		PublisherName: p,
		PublishedDate: pd,
		Isbn:          ISBN,
	}
}

type CreateBookResponse struct {
	ID            int64
	Title         string
	Description   *string
	CoverImageURL *string
	URL           *string
	Author        *string
	Publisher     *string
	PublishDate   *time.Time
	ISBN          *string
}

func newCreateResponse(b sqlc.Book) *CreateBookResponse {
	return &CreateBookResponse{
		ID:            b.ID,
		Title:         b.Title,
		Description:   nilString(b.Description),
		CoverImageURL: nilString(b.CoverImageUrl),
		URL:           nilString(b.Url),
		Author:        nilString(b.AuthorName),
		Publisher:     nilString(b.PublisherName),
		PublishDate:   nilTime(b.PublishedDate),
		ISBN:          nilString(b.Isbn),
	}
}

func (r BookRepositoryImpl) CreateBook(ctx context.Context, req CreateBookRequest) (*CreateBookResponse, error) {
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

func (r BookRepositoryImpl) CreateBookGenre(ctx context.Context, req CreateBookGenreRequest) (*CreateBookGenreResponse, error) {
	p := req.toParams()
	b, err := r.querier.CreateBookGenre(ctx, p)
	if err != nil {
		return nil, err
	}
	return newCreateBookGenreResponse(b), nil
}

func (r BookRepositoryImpl) CreateGenre(ctx context.Context, name string) (*string, error) {
	g, err := r.querier.CreateGenre(ctx, name)
	if err != nil {
		return nil, err
	}
	return &g.Name, nil
}

func (r BookRepositoryImpl) CreatePublisher(ctx context.Context, name string) (*string, error) {
	p, err := r.querier.CreatePublisher(ctx, name)
	if err != nil {
		return nil, err
	}
	return &p.Name, nil
}

func (r BookRepositoryImpl) DeleteAuthor(ctx context.Context, name string) error {
	return r.querier.DeleteAuthor(ctx, name)
}

func (r BookRepositoryImpl) DeleteBook(ctx context.Context, id int64) error {
	return r.querier.DeleteBook(ctx, id)
}

type DeleteBookGenreRequest struct {
	BookID    int64
	GenreName string
}

func (r BookRepositoryImpl) DeleteBookGenre(ctx context.Context, req DeleteBookGenreRequest) error {
	p := sqlc.DeleteBookGenreParams{
		BookID:    req.BookID,
		GenreName: req.GenreName,
	}
	return r.querier.DeleteBookGenre(ctx, p)
}

func (r BookRepositoryImpl) DeleteGenre(ctx context.Context, name string) error {
	return r.querier.DeleteGenre(ctx, name)
}

func (r BookRepositoryImpl) DeletePublisher(ctx context.Context, name string) error {
	return r.querier.DeletePublisher(ctx, name)
}

type GetBookResponse struct {
	ID            int64
	Title         string
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
		Title:         b.Title,
		Genres:        strings.Split(string(b.Genres), ", "),
		Description:   nilString(b.Description),
		CoverImageURL: nilString(b.CoverImageUrl),
		URL:           nilString(b.Url),
		AuthorName:    nilString(b.AuthorName),
		PublisherName: nilString(b.PublisherName),
		PublishDate:   nilTime(b.PublishedDate),
		ISBN:          nilString(b.Isbn),
	}
}

func (r BookRepositoryImpl) GetBookByID(ctx context.Context, id int64) (*GetBookResponse, error) {
	b, err := r.querier.GetBooksByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return newGetBookResponse(b), nil
}
