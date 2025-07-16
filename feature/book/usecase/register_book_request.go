package usecase

import (
	"readly/feature/book/domain"
	pb "readly/pb/readly/v1"
	"readly/util"
)

type RegisterBookRequest struct {
	UserID        int64
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

func NewRegisterBookRequest(userID int64, proto *pb.RegisterBookRequest) RegisterBookRequest {
	return RegisterBookRequest{
		UserID:        userID,
		Title:         proto.GetTitle(),
		Genres:        proto.GetGenres(),
		Description:   util.ToStringOrNil(proto.GetDescription()),
		CoverImageURL: util.ToStringOrNil(proto.GetCoverImageUrl()),
		URL:           util.ToStringOrNil(proto.GetUrl()),
		AuthorName:    util.ToStringOrNil(proto.GetAuthorName()),
		PublisherName: util.ToStringOrNil(proto.GetPublisherName()),
		PublishDate:   domain.NewDateEntityFromProto(proto.GetPublishDate()),
		ISBN:          util.ToStringOrNil(proto.GetIsbn()),
		Status:        domain.NewReadingStatusFromProto(proto.GetReadingStatus()),
		StartDate:     domain.NewDateEntityFromProto(proto.GetStartDate()),
		EndDate:       domain.NewDateEntityFromProto(proto.GetEndDate()),
	}
}

func (r RegisterBookRequest) Validate() error {
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

	// AuthorName validation
	if r.AuthorName != nil {
		if err := util.StringValidator(*r.AuthorName).ValidateLength(0, 255); err != nil {
			return newError(BadRequest, InvalidRequestError, "author name must be less than 255 characters")
		}
	}

	// PublisherName validation
	if r.PublisherName != nil {
		if err := util.StringValidator(*r.PublisherName).ValidateLength(0, 255); err != nil {
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
