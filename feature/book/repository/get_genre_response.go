package repository

import sqlc "readly/db/sqlc"

type GetGenreResponse struct {
	Name string
}

func newGetGenreResponseFromSQLC(g sqlc.Genre) *GetGenreResponse {
	return &GetGenreResponse{
		Name: g.Name,
	}
}
