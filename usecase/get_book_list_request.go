package usecase

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

func (r GetBookListRequest) ToRepoRequest() GetBookListRequest {
	return GetBookListRequest{
		UserID: r.UserID,
		limit:  r.limit,
		offset: r.offset,
	}
}
