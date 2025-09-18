package usecase

import (
	"readly/feature/book/domain"
	pb "readly/pb/readly/v1"
)

type RegisterBookResponse struct {
	Book domain.Book
}

func NewRegisterBookResponse(book *domain.Book) *RegisterBookResponse {
	if book == nil {
		return nil
	}
	return &RegisterBookResponse{
		Book: *book,
	}
}

func (r *RegisterBookResponse) ToProto() *pb.RegisterBookResponse {
	var publishDate *pb.Date
	var startDate *pb.Date
	var endDate *pb.Date
	if r.Book.PublishDate != nil {
		publishDate = r.Book.PublishDate.ToProto()
	}

	if r.Book.StartDate != nil {
		startDate = r.Book.StartDate.ToProto()
	}
	if r.Book.EndDate != nil {
		endDate = r.Book.EndDate.ToProto()
	}
	return &pb.RegisterBookResponse{
		Book: &pb.Book{
			Id:            r.Book.ID,
			Title:         r.Book.Title,
			Genres:        r.Book.Genres,
			Description:   r.Book.Description,
			CoverImageUrl: r.Book.CoverImageURL,
			Url:           r.Book.URL,
			AuthorName:    r.Book.AuthorName,
			PublisherName: r.Book.PublisherName,
			PublishDate:   publishDate,
			Isbn:          r.Book.ISBN,
			ReadingStatus: pb.ReadingStatus(r.Book.Status),
			StartDate:     startDate,
			EndDate:       endDate,
		},
	}
}
