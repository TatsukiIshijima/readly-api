package usecase

import (
	"context"
	"readly/entity"
	"readly/repository"
)

type GetBookUseCase interface {
	GetBook(ctx context.Context, req GetBookRequest) (*GetBookResponse, error)
}

type GetBookUseCaseImpl struct {
	bookRepo           repository.BookRepository
	readingHistoryRepo repository.ReadingHistoryRepository
}

func NewGetBookUseCase(
	bookRepo repository.BookRepository,
	readingHistoryRepo repository.ReadingHistoryRepository,
) GetBookUseCase {
	return &GetBookUseCaseImpl{
		bookRepo:           bookRepo,
		readingHistoryRepo: readingHistoryRepo,
	}
}

func (g GetBookUseCaseImpl) GetBook(ctx context.Context, req GetBookRequest) (*GetBookResponse, error) {
	res, err := g.bookRepo.GetBookByID(ctx, req.ToRepoRequest())
	if err != nil {
		return nil, handle(err)
	}
	rh, err := g.readingHistoryRepo.GetByUserAndBook(ctx, repository.GetReadingHistoryByUserAndBookRequest{
		UserID: req.UserID,
		BookID: req.BookID,
	})
	if err != nil {
		return nil, handle(err)
	}

	book := entity.Book{
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
