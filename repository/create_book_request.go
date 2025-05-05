package repository

import (
	"database/sql"
	sqlc "readly/db/sqlc"
	"readly/entity"
	"time"
)

type CreateBookRequest struct {
	Title         string
	Description   *string
	CoverImageURL *string
	URL           *string
	Author        *string
	Publisher     *string
	PublishDate   *entity.Date
	ISBN          *string
}

func (r CreateBookRequest) toSQLC() sqlc.CreateBookParams {
	desc := sql.NullString{String: "", Valid: false}
	coverImgURL := sql.NullString{String: "", Valid: false}
	URL := sql.NullString{String: "", Valid: false}
	a := sql.NullString{String: "", Valid: false}
	p := sql.NullString{String: "", Valid: false}
	pd := sql.NullTime{Time: time.Time{}, Valid: false}
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
		a = sql.NullString{String: *r.Author, Valid: true}
	}
	if r.Publisher != nil {
		p = sql.NullString{String: *r.Publisher, Valid: true}
	}
	if r.PublishDate != nil {
		t := r.PublishDate.ToTime()
		pd = sql.NullTime{Time: *t, Valid: true}
	}
	if r.ISBN != nil {
		ISBN = sql.NullString{String: *r.ISBN, Valid: true}
	}
	return sqlc.CreateBookParams{
		Title:         r.Title,
		Description:   desc,
		CoverImageUrl: coverImgURL,
		Url:           URL,
		AuthorName:    a,
		PublisherName: p,
		PublishedDate: pd,
		Isbn:          ISBN,
	}
}
