package usecase

import (
	"readly/feature/book/domain"
	pb "readly/pb/readly/v1"
)

type GetBookResponse struct {
	Book domain.Book
}

func NewGetBookResponse(book domain.Book) *GetBookResponse {
	return &GetBookResponse{
		Book: book,
	}
}

func (r *GetBookResponse) ToProto() *pb.GetBookResponse {
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
	return &pb.GetBookResponse{
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
