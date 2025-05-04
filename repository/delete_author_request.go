package repository

type DeleteAuthorRequest struct {
	Name string
}

func NewDeleteAuthorRequest(name string) DeleteAuthorRequest {
	return DeleteAuthorRequest{
		Name: name,
	}
}
