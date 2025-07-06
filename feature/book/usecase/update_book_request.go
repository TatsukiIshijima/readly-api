package usecase

import (
	"readly/feature/book/domain"
	bookRepo "readly/feature/book/repository"
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
	PublishedDate *domain.Date
	ISBN          *string
	Status        domain.ReadingStatus
	StartDate     *domain.Date
	EndDate       *domain.Date
}

func (r UpdateBookRequest) isValid() bool {
	// TODO:バリデーション増やす
	return len(r.Title) > 0
}

func (r UpdateBookRequest) toBookRepoRequest() bookRepo.UpdateBookRequest {
	return bookRepo.UpdateBookRequest{
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

func (r UpdateBookRequest) toReadingHistoryRepoRequest() bookRepo.UpdateReadingHistoryRequest {
	return bookRepo.UpdateReadingHistoryRequest{
		UserID:    r.UserID,
		BookID:    r.BookID,
		Status:    r.Status,
		StartDate: r.StartDate,
		EndDate:   r.EndDate,
	}
}
