package usecase

import "readly/entity"

type GetBookResponse struct {
	Book entity.Book
}

func NewGetBookResponse(book entity.Book) *GetBookResponse {
	return &GetBookResponse{
		Book: book,
	}
}
