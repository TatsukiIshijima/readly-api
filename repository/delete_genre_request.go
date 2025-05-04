package repository

type DeleteGenreRequest struct {
	Name string
}

func NewDeleteGenreRequest(name string) DeleteGenreRequest {
	return DeleteGenreRequest{
		Name: name,
	}
}
