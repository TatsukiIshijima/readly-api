package repository

import sqlc "readly/db/sqlc"

type UpdateUserResponse struct {
	ID    int64
	Name  string
	Email string
}

func newUpdateUserResponseFromSQLC(u sqlc.User) *UpdateUserResponse {
	return &UpdateUserResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}
