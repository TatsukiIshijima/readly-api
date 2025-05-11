package usecase

import (
	"readly/entity"
	"readly/repository"
)

type UpdateBookRequest struct {
	UserID        int64
	BookID        int64
	Title         string
	Genres        []string
	Description   *string
	CoverImageURL *string
	URL           *string
	Author        *string
	Publisher     *string
	PublishedDate *entity.Date
	ISBN          *string
	Status        entity.ReadingStatus
	StartDate     *entity.Date
	EndDate       *entity.Date
}

func (r UpdateBookRequest) isValid() bool {
	// TODO:バリデーション増やす
	return len(r.Title) > 0
}

func (r UpdateBookRequest) toBookRepoRequest() repository.UpdateBookRequest {
	return repository.UpdateBookRequest{
		BookID:        r.BookID,
		Title:         r.Title,
		Description:   r.Description,
		CoverImageURL: r.CoverImageURL,
		URL:           r.URL,
		Author:        r.Author,
		Publisher:     r.Publisher,
		PublishedDate: r.PublishedDate,
		ISBN:          r.ISBN,
	}
}

func (r UpdateBookRequest) toReadingHistoryRepoRequest() repository.UpdateReadingHistoryRequest {
	return repository.UpdateReadingHistoryRequest{
		UserID:    r.UserID,
		BookID:    r.BookID,
		Status:    r.Status,
		StartDate: r.StartDate,
		EndDate:   r.EndDate,
	}
}
