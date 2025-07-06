package usecase

import (
	"context"
	"errors"
	"readly/feature/book/domain"
	bookRepo "readly/feature/book/repository"
)

type GetBookUseCase interface {
	GetBook(ctx context.Context, req GetBookRequest) (*GetBookResponse, error)
}

type GetBookUseCaseImpl struct {
	bookRepo           bookRepo.BookRepository
	readingHistoryRepo bookRepo.ReadingHistoryRepository
}

func NewGetBookUseCase(
	bookRepo bookRepo.BookRepository,
	readingHistoryRepo bookRepo.ReadingHistoryRepository,
) GetBookUseCase {
	return &GetBookUseCaseImpl{
		bookRepo:           bookRepo,
		readingHistoryRepo: readingHistoryRepo,
	}
}

func (g GetBookUseCaseImpl) GetBook(ctx context.Context, req GetBookRequest) (*GetBookResponse, error) {
	res, err := g.bookRepo.GetBookByID(ctx, req.ToRepoRequest())
	if err != nil {
		if errors.Is(err, bookRepo.ErrNoRows) {
			return nil, newError(BadRequest, NotFoundBookError, "book not found")
		}
		return nil, handle(err)
	}
	rh, err := g.readingHistoryRepo.GetByUserAndBook(ctx, bookRepo.GetReadingHistoryByUserAndBookRequest{
		UserID: req.UserID,
		BookID: req.BookID,
	})
	if err != nil {
		if errors.Is(err, bookRepo.ErrNoRows) {
			return nil, newError(BadRequest, NotFoundBookError, "book not found")
		}
		return nil, handle(err)
	}

	book := domain.Book{
		ID:            res.ID,
		Title:         res.Title,
		Genres:        res.Genres,
		Description:   res.Description,
		CoverImageURL: res.CoverImageURL,
		URL:           res.URL,
		AuthorName:    res.AuthorName,
		PublisherName: res.PublisherName,
		PublishDate:   res.PublishDate,
		ISBN:          res.ISBN,
		Status:        rh.Status,
		StartDate:     rh.StartDate,
		EndDate:       rh.EndDate,
	}

	return NewGetBookResponse(book), nil
}
