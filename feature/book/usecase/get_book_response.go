package usecase

import "readly/feature/book/domain"

type GetBookResponse struct {
	Book domain.Book
}

func NewGetBookResponse(book domain.Book) *GetBookResponse {
	return &GetBookResponse{
		Book: book,
	}
}
