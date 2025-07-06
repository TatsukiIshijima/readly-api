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
