package usecase

import bookRepo "readly/feature/book/repository"

type GetBookRequest struct {
	UserID int64
	BookID int64
}

func NewGetBookRequest(userID, bookID int64) GetBookRequest {
	return GetBookRequest{
		UserID: userID,
		BookID: bookID,
	}
}

func (r GetBookRequest) ToRepoRequest() bookRepo.GetBookRequest {
	return bookRepo.GetBookRequest{
		ID: r.BookID,
	}
}
