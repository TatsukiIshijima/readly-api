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
			PublishDate:   r.Book.PublishDate.ToProto(),
			Isbn:          r.Book.ISBN,
			ReadingStatus: pb.ReadingStatus(r.Book.Status),
			StartDate:     r.Book.StartDate.ToProto(),
			EndDate:       r.Book.EndDate.ToProto(),
		},
	}
}
