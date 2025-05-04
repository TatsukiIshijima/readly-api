package repository

type CreateAuthorRequest struct {
	Name string
}

func NewCreateAuthorRequest(name string) CreateAuthorRequest {
	return CreateAuthorRequest{
		Name: name,
	}
}
