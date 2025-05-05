package usecase

import (
	"readly/entity"
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
	PublishDate   *entity.Date
	ISBN          *string
	Status        entity.ReadingStatus
	StartDate     *entity.Date
	EndDate       *entity.Date
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
		PublishDate:   entity.NewDateEntityFromProto(proto.GetPublishDate()),
		ISBN:          util.ToStringOrNil(proto.GetIsbn()),
		Status:        entity.NewReadingStatusFromProto(proto.GetReadingStatus()),
		StartDate:     entity.NewDateEntityFromProto(proto.GetStartDate()),
		EndDate:       entity.NewDateEntityFromProto(proto.GetEndDate()),
	}
}
