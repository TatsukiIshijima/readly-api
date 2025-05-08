package usecase

import "readly/repository"

type GetBookListRequest struct {
	UserID int64
	limit  int32
	offset int32
}

func NewGetBookListRequest(userID int64, limit, offset int32) GetBookListRequest {
	return GetBookListRequest{
		UserID: userID,
		limit:  limit,
		offset: offset,
	}
}

func (r GetBookListRequest) ToRepoRequest() repository.GetReadingHistoryByUserRequest {
	return repository.GetReadingHistoryByUserRequest{
		UserID: r.UserID,
		Limit:  r.limit,
		Offset: r.offset,
	}
}
