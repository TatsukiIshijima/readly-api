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

type GetBookRequest struct {
	UserID int64
	BookID int64
}

type GetBookResponse struct {
	Book entity.Book
}

func (g GetBookUseCaseImpl) GetBook(ctx context.Context, req GetBookRequest) (*GetBookResponse, error) {
	b, err := g.bookRepo.GetBookByID(ctx, req.BookID)
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

	return &GetBookResponse{
		Book: entity.Book{
			ID:            b.ID,
			Title:         b.Title,
			Genres:        b.Genres,
			Description:   b.Description,
			CoverImageURL: b.CoverImageURL,
			URL:           b.URL,
			AuthorName:    b.AuthorName,
			PublisherName: b.PublisherName,
			PublishDate:   b.PublishDate,
			ISBN:          b.ISBN,
			Status:        rh.Status,
			StartDate:     rh.StartDate,
			EndDate:       rh.EndDate,
		},
	}, nil
}
