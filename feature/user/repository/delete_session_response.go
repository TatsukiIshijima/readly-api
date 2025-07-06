package repository

type DeleteSessionByUserIDResponse struct {
	Count int64
}

func newDeleteSessionByUserIDResponse(count int64) *DeleteSessionByUserIDResponse {
	return &DeleteSessionByUserIDResponse{
		Count: count,
	}
}
