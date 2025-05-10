package usecase

import (
	"context"
	"readly/entity"
	"readly/repository"
)

type GetBookListUseCase interface {
	GetBookList(ctx context.Context, req GetBookListRequest) (*GetBookListResponse, error)
}

type GetBookListUseCaseImpl struct {
	readingHistoryRepo repository.ReadingHistoryRepository
}

func NewGetBookListUseCase(
	readingHistoryRepo repository.ReadingHistoryRepository,
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
	books := make([]entity.Book, len(res))
	for i := 0; i < len(res); i++ {
		books[i] = entity.Book{
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
