package usecase

import (
	"readly/feature/book/domain"
	bookRepo "readly/feature/book/repository"
	"readly/util"
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

func (r UpdateBookRequest) Validate() error {
	// Title validation
	if len(r.Title) == 0 {
		return newError(BadRequest, InvalidRequestError, "title is required")
	}
	if err := util.StringValidator(r.Title).ValidateLength(1, 255); err != nil {
		return newError(BadRequest, InvalidRequestError, "title must be between 1 and 255 characters")
	}

	// Description validation
	if r.Description != nil {
		if err := util.StringValidator(*r.Description).ValidateLength(0, 500); err != nil {
			return newError(BadRequest, InvalidRequestError, "description must be less than 500 characters")
		}
	}

	// CoverImageURL validation
	if r.CoverImageURL != nil {
		if err := util.StringValidator(*r.CoverImageURL).ValidateLength(0, 2048); err != nil {
			return newError(BadRequest, InvalidRequestError, "cover image URL must be less than 2048 characters")
		}
		if err := util.StringValidator(*r.CoverImageURL).ValidateURL(); err != nil {
			return newError(BadRequest, InvalidRequestError, "cover image URL has invalid format")
		}
	}

	// URL validation
	if r.URL != nil {
		if err := util.StringValidator(*r.URL).ValidateLength(0, 2048); err != nil {
			return newError(BadRequest, InvalidRequestError, "URL must be less than 2048 characters")
		}
		if err := util.StringValidator(*r.URL).ValidateURL(); err != nil {
			return newError(BadRequest, InvalidRequestError, "URL has invalid format")
		}
	}

	// Author validation
	if r.Author != nil {
		if err := util.StringValidator(*r.Author).ValidateLength(0, 255); err != nil {
			return newError(BadRequest, InvalidRequestError, "author name must be less than 255 characters")
		}
	}

	// Publisher validation
	if r.Publisher != nil {
		if err := util.StringValidator(*r.Publisher).ValidateLength(0, 255); err != nil {
			return newError(BadRequest, InvalidRequestError, "publisher name must be less than 255 characters")
		}
	}

	// ISBN validation
	if r.ISBN != nil {
		if err := util.StringValidator(*r.ISBN).ValidateISBN(); err != nil {
			return newError(BadRequest, InvalidRequestError, "ISBN must be 13 digits")
		}
	}

	// Date validation
	if r.StartDate != nil && r.EndDate != nil {
		if r.EndDate.Before(*r.StartDate) {
			return newError(BadRequest, InvalidRequestError, "end date must be after start date")
		}
	}

	return nil
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
