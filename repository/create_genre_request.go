package repository

type CreateGenreRequest struct {
	Name string
}

func NewCreateGenreRequest(name string) CreateGenreRequest {
	return CreateGenreRequest{
		Name: name,
	}
}
