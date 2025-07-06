package repository

type DeleteBookRequest struct {
	ID int64
}

func NewDeleteBookRequest(id int64) DeleteBookRequest {
	return DeleteBookRequest{
		ID: id,
	}
}
