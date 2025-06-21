package usecase

type CreateGenresRequest struct {
	Names []string
}

func NewCreateGenresRequest(names []string) CreateGenresRequest {
	return CreateGenresRequest{
		Names: names,
	}
}
