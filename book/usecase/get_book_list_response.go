package usecase

import "readly/book/domain"

type GetBookListResponse struct {
	Books []domain.Book
}

func NewGetBookListResponse(books []domain.Book) *GetBookListResponse {
	return &GetBookListResponse{
		Books: books,
	}
}
