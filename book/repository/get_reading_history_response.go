package repository

import (
	"readly/book/domain"
	sqlc "readly/db/sqlc"
)

type GetReadingHistoryByUserResponse struct {
	BookID        int64
	Title         string
	Genres        []string
	Description   *string
	CoverImageURL *string
	URL           *string
	AuthorName    *string
	PublisherName *string
	PublishDate   *domain.Date
	ISBN          *string
	Status        domain.ReadingStatus
	StartDate     *domain.Date
	EndDate       *domain.Date
}

func newGetReadingHistoryByUserResponseFromSQLC(r sqlc.GetReadingHistoryByUserRow) GetReadingHistoryByUserResponse {
	id := nilInt64(r.ID)
	t := nilString(r.Title)
	g := newGenres(r.Genres)
	desc := nilString(r.Description)
	coverImgURL := nilString(r.CoverImageUrl)
	URL := nilString(r.Url)
	a := nilString(r.AuthorName)
	p := nilString(r.PublisherName)
	ISBN := nilString(r.Isbn)
	return GetReadingHistoryByUserResponse{
		BookID:        *id,
		Title:         *t,
		Genres:        g,
		Description:   desc,
		CoverImageURL: coverImgURL,
		URL:           URL,
		AuthorName:    a,
		PublisherName: p,
		PublishDate:   domain.NewDateEntityFromNullTime(r.PublishedDate),
		ISBN:          ISBN,
		Status:        domain.NewReadingStatusFromSQLC(r.Status),
		StartDate:     domain.NewDateEntityFromNullTime(r.StartDate),
		EndDate:       domain.NewDateEntityFromNullTime(r.EndDate),
	}
}

type GetReadingHistoryByUserAndBookResponse struct {
	BookID        int64
	Title         string
	Genres        []string
	Description   *string
	CoverImageURL *string
	URL           *string
	AuthorName    *string
	PublisherName *string
	PublishDate   *domain.Date
	ISBN          *string
	Status        domain.ReadingStatus
	StartDate     *domain.Date
	EndDate       *domain.Date
}

func newGetReadingHistoryByUserAndBookResponseFromSQLC(r sqlc.GetReadingHistoryByUserAndBookRow) *GetReadingHistoryByUserAndBookResponse {
	id := nilInt64(r.ID)
	t := nilString(r.Title)
	g := newGenres(r.Genres)
	desc := nilString(r.Description)
	coverImgURL := nilString(r.CoverImageUrl)
	URL := nilString(r.Url)
	a := nilString(r.AuthorName)
	p := nilString(r.PublisherName)
	ISBN := nilString(r.Isbn)
	return &GetReadingHistoryByUserAndBookResponse{
		BookID:        *id,
		Title:         *t,
		Genres:        g,
		Description:   desc,
		CoverImageURL: coverImgURL,
		URL:           URL,
		AuthorName:    a,
		PublisherName: p,
		PublishDate:   domain.NewDateEntityFromNullTime(r.PublishedDate),
		ISBN:          ISBN,
		Status:        domain.NewReadingStatusFromSQLC(r.Status),
		StartDate:     domain.NewDateEntityFromNullTime(r.StartDate),
		EndDate:       domain.NewDateEntityFromNullTime(r.EndDate),
	}
}

type GetReadingHistoryByUserAndStatusResponse struct {
	BookID        int64
	Title         string
	Genres        []string
	Description   *string
	CoverImageURL *string
	URL           *string
	AuthorName    *string
	PublisherName *string
	PublishDate   *domain.Date
	ISBN          *string
	Status        domain.ReadingStatus
	StartDate     *domain.Date
	EndDate       *domain.Date
}

func newGetReadingHistoryByUserAndStatusResponseFromSQLC(r sqlc.GetReadingHistoryByUserAndStatusRow) GetReadingHistoryByUserAndStatusResponse {
	id := nilInt64(r.ID)
	t := nilString(r.Title)
	g := newGenres(r.Genres)
	desc := nilString(r.Description)
	coverImgURL := nilString(r.CoverImageUrl)
	URL := nilString(r.Url)
	a := nilString(r.AuthorName)
	p := nilString(r.PublisherName)
	ISBN := nilString(r.Isbn)
	return GetReadingHistoryByUserAndStatusResponse{
		BookID:        *id,
		Title:         *t,
		Genres:        g,
		Description:   desc,
		CoverImageURL: coverImgURL,
		URL:           URL,
		AuthorName:    a,
		PublisherName: p,
		PublishDate:   domain.NewDateEntityFromNullTime(r.PublishedDate),
		ISBN:          ISBN,
		Status:        domain.NewReadingStatusFromSQLC(r.Status),
		StartDate:     domain.NewDateEntityFromNullTime(r.StartDate),
		EndDate:       domain.NewDateEntityFromNullTime(r.EndDate),
	}
}
