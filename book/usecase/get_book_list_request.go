package usecase

import bookRepo "readly/book/repository"

type GetBookListRequest struct {
	UserID int64
	Limit  int32
	Offset int32
}

func NewGetBookListRequest(userID int64, limit, offset int32) GetBookListRequest {
	return GetBookListRequest{
		UserID: userID,
		Limit:  limit,
		Offset: offset,
	}
}

func (r GetBookListRequest) Validate() error {
	if r.Limit < 0 {
		return newError(BadRequest, InvalidRequestError, "limit must be greater than 0")
	}
	if r.Offset < 0 {
		return newError(BadRequest, InvalidRequestError, "offset must be greater than 0")
	}
	return nil
}

func (r GetBookListRequest) ToRepoRequest() bookRepo.GetReadingHistoryByUserRequest {
	return bookRepo.GetReadingHistoryByUserRequest{
		UserID: r.UserID,
		Limit:  r.Limit,
		Offset: r.Offset,
	}
}
