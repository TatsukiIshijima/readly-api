package repository

import sqlc "readly/db/sqlc"

type UpdateBookResponse struct {
	BookID int64
}

func newUpdateBookResponseFromSQLC(b sqlc.Book) *UpdateBookResponse {
	return &UpdateBookResponse{
		BookID: b.ID,
	}
}
