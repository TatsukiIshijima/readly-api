package repository

import (
	sqlc "readly/db/sqlc"
	"readly/feature/book/domain"
	"strings"
)

type GetBookResponse struct {
	ID            int64
	Title         string
	Genres        []string
	Description   *string
	CoverImageURL *string
	URL           *string
	AuthorName    *string
	PublisherName *string
	PublishDate   *domain.Date
	ISBN          *string
}

func newGetBookResponseFromSQLC(row sqlc.GetBooksByIDRow) *GetBookResponse {
	return &GetBookResponse{
		ID:            row.ID,
		Title:         row.Title,
		Genres:        strings.Split(string(row.Genres), ", "),
		Description:   nilString(row.Description),
		CoverImageURL: nilString(row.CoverImageUrl),
		URL:           nilString(row.Url),
		AuthorName:    nilString(row.AuthorName),
		PublisherName: nilString(row.PublisherName),
		PublishDate:   domain.NewDateEntityFromNullTime(row.PublishedDate),
		ISBN:          nilString(row.Isbn),
	}
}
