package usecase

import (
	"context"
	"readly/book/domain"
	bookRepo "readly/book/repository"
)

type GetBookListUseCase interface {
	GetBookList(ctx context.Context, req GetBookListRequest) (*GetBookListResponse, error)
}

type GetBookListUseCaseImpl struct {
	readingHistoryRepo bookRepo.ReadingHistoryRepository
}

func NewGetBookListUseCase(
	readingHistoryRepo bookRepo.ReadingHistoryRepository,
) GetBookListUseCase {
	return &GetBookListUseCaseImpl{
		readingHistoryRepo: readingHistoryRepo,
	}
}

func (g GetBookListUseCaseImpl) GetBookList(ctx context.Context, req GetBookListRequest) (*GetBookListResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, handle(err)
	}
	res, err := g.readingHistoryRepo.GetByUser(ctx, req.ToRepoRequest())
	if err != nil {
		return nil, handle(err)
	}
	books := make([]domain.Book, len(res))
	for i := 0; i < len(res); i++ {
		books[i] = domain.Book{
			ID:            res[i].BookID,
			Title:         res[i].Title,
			Genres:        res[i].Genres,
			Description:   res[i].Description,
			CoverImageURL: res[i].CoverImageURL,
			URL:           res[i].URL,
			AuthorName:    res[i].AuthorName,
			PublisherName: res[i].PublisherName,
			PublishDate:   res[i].PublishDate,
			ISBN:          res[i].ISBN,
			Status:        res[i].Status,
			StartDate:     res[i].StartDate,
		}
	}

	return NewGetBookListResponse(books), nil
}
