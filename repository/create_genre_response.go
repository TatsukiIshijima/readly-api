package repository

import sqlc "readly/db/sqlc"

type CreateGenreResponse struct {
	Name string
}

func newCreateGenreResponseFromSQLC(g sqlc.Genre) *CreateGenreResponse {
	return &CreateGenreResponse{
		Name: g.Name,
	}
}
