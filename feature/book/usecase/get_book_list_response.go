package usecase

import (
	"readly/feature/book/domain"
	pb "readly/pb/readly/v1"
)

type GetBookListResponse struct {
	Books []domain.Book
}

func NewGetBookListResponse(books []domain.Book) *GetBookListResponse {
	return &GetBookListResponse{
		Books: books,
	}
}

func (r *GetBookListResponse) ToProto() *pb.GetBookListResponse {
	pbBooks := make([]*pb.Book, len(r.Books))
	for i, book := range r.Books {
		var publishDate *pb.Date
		var startDate *pb.Date
		var endDate *pb.Date
		if book.PublishDate != nil {
			publishDate = book.PublishDate.ToProto()
		}
		if book.StartDate != nil {
			startDate = book.StartDate.ToProto()
		}
		if book.EndDate != nil {
			endDate = book.EndDate.ToProto()
		}
		pbBooks[i] = &pb.Book{
			Id:            book.ID,
			Title:         book.Title,
			Genres:        book.Genres,
			Description:   book.Description,
			CoverImageUrl: book.CoverImageURL,
			Url:           book.URL,
			AuthorName:    book.AuthorName,
			PublisherName: book.PublisherName,
			PublishDate:   publishDate,
			Isbn:          book.ISBN,
			ReadingStatus: pb.ReadingStatus(book.Status),
			StartDate:     startDate,
			EndDate:       endDate,
		}
	}
	return &pb.GetBookListResponse{
		Books: pbBooks,
	}
}
