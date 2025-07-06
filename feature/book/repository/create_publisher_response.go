package repository

import sqlc "readly/db/sqlc"

type CreatePublisherResponse struct {
	Name string
}

func newCreatePublisherResponseFromSQLC(p sqlc.Publisher) *CreatePublisherResponse {
	return &CreatePublisherResponse{
		Name: p.Name,
	}
}
