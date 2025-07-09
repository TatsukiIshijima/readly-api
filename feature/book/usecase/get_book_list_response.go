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
		pbBooks[i] = &pb.Book{
			Id:            book.ID,
			Title:         book.Title,
			Genres:        book.Genres,
			Description:   book.Description,
			CoverImageUrl: book.CoverImageURL,
			Url:           book.URL,
			AuthorName:    book.AuthorName,
			PublisherName: book.PublisherName,
			PublishDate:   book.PublishDate.ToProto(),
			Isbn:          book.ISBN,
			ReadingStatus: pb.ReadingStatus(book.Status),
			StartDate:     book.StartDate.ToProto(),
			EndDate:       book.EndDate.ToProto(),
		}
	}
	return &pb.GetBookListResponse{
		Books: pbBooks,
	}
}
