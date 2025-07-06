package repository

import (
	sqlc "readly/db/sqlc"
	"readly/feature/book/domain"
)

type CreateBookResponse struct {
	ID            int64
	Title         string
	Description   *string
	CoverImageURL *string
	URL           *string
	Author        *string
	Publisher     *string
	PublishDate   *domain.Date
	ISBN          *string
}

func newCreateBookResponseFromSQLC(b sqlc.Book) *CreateBookResponse {
	return &CreateBookResponse{
		ID:            b.ID,
		Title:         b.Title,
		Description:   nilString(b.Description),
		CoverImageURL: nilString(b.CoverImageUrl),
		URL:           nilString(b.Url),
		Author:        nilString(b.AuthorName),
		Publisher:     nilString(b.PublisherName),
		PublishDate:   domain.NewDateEntityFromNullTime(b.PublishedDate),
		ISBN:          nilString(b.Isbn),
	}
}
