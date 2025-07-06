package repository

import sqlc "readly/db/sqlc"

type GetUserResponse struct {
	ID       int64
	Name     string
	Password string
	Email    string
}

func newGetUserResponseFromSQLC(u sqlc.User) *GetUserResponse {
	return &GetUserResponse{
		ID:       u.ID,
		Name:     u.Name,
		Password: u.HashedPassword,
		Email:    u.Email,
	}
}
