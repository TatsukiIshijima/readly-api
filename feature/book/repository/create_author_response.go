package repository

import sqlc "readly/db/sqlc"

type CreateAuthorResponse struct {
	Name string
}

func newCreateAuthorResponseFromSQLC(a sqlc.Author) *CreateAuthorResponse {
	return &CreateAuthorResponse{
		Name: a.Name,
	}
}
