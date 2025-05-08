package usecase

import "readly/entity"

type GetBookListResponse struct {
	Books []entity.Book
}

func NewGetBookListResponse(books []entity.Book) *GetBookListResponse {
	return &GetBookListResponse{
		Books: books,
	}
}
