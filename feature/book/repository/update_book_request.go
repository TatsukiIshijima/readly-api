package repository

import (
	"database/sql"
	sqlc "readly/db/sqlc"
	"readly/feature/book/domain"
	"time"
)

type UpdateBookRequest struct {
	BookID        int64
	Title         string
	Description   *string
	CoverImageURL *string
	URL           *string
	Author        *string
	Publisher     *string
	PublishedDate *domain.Date
	ISBN          *string
}

func (r UpdateBookRequest) toSQLC() sqlc.UpdateBookParams {
	desc := sql.NullString{String: "", Valid: false}
	coverImgURL := sql.NullString{String: "", Valid: false}
	URL := sql.NullString{String: "", Valid: false}
	author := sql.NullString{String: "", Valid: false}
	publisher := sql.NullString{String: "", Valid: false}
	publishedDate := sql.NullTime{Time: time.Time{}, Valid: false}
	ISBN := sql.NullString{String: "", Valid: false}
	if r.Description != nil {
		desc = sql.NullString{String: *r.Description, Valid: true}
	}
	if r.CoverImageURL != nil {
		coverImgURL = sql.NullString{String: *r.CoverImageURL, Valid: true}
	}
	if r.URL != nil {
		URL = sql.NullString{String: *r.URL, Valid: true}
	}
	if r.Author != nil {
		author = sql.NullString{String: *r.Author, Valid: true}
	}
	if r.Publisher != nil {
		publisher = sql.NullString{String: *r.Publisher, Valid: true}
	}
	if r.PublishedDate != nil {
		t := r.PublishedDate.ToTime()
		publishedDate = sql.NullTime{Time: *t, Valid: true}
	}
	if r.ISBN != nil {
		ISBN = sql.NullString{String: *r.ISBN, Valid: true}
	}
	return sqlc.UpdateBookParams{
		ID:            r.BookID,
		Title:         r.Title,
		Description:   desc,
		CoverImageUrl: coverImgURL,
		Url:           URL,
		AuthorName:    author,
		PublisherName: publisher,
		PublishedDate: publishedDate,
		Isbn:          ISBN,
	}
}
