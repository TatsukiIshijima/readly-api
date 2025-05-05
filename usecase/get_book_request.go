package usecase

import "readly/repository"

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

func (r GetBookRequest) ToRepoRequest() repository.GetBookRequest {
	return repository.GetBookRequest{
		ID: r.BookID,
	}
}
