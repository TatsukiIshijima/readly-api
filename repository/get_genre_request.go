package repository

type GetGenreRequest struct {
	Name string
}

func NewGetGenreRequest(name string) GetGenreRequest {
	return GetGenreRequest{
		Name: name,
	}
}
