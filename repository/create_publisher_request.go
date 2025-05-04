package repository

type CreatePublisherRequest struct {
	Name string
}

func NewCreatePublisherRequest(name string) CreatePublisherRequest {
	return CreatePublisherRequest{
		Name: name,
	}
}
