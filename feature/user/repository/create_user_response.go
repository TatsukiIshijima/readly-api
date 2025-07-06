package repository

import sqlc "readly/db/sqlc"

type CreateUserResponse struct {
	ID    int64
	Name  string
	Email string
}

func newCreateUserResponseFromSQLC(u sqlc.User) *CreateUserResponse {
	return &CreateUserResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}
