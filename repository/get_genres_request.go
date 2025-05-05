package repository

type GetGenresByBookIDRequest struct {
	ID int64
}

func NewGetGenresByBookIDRequest(id int64) GetGenresByBookIDRequest {
	return GetGenresByBookIDRequest{
		ID: id,
	}
}
