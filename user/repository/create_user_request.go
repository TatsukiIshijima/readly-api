package repository

import sqlc "readly/db/sqlc"

type CreateUserRequest struct {
	Name     string
	Email    string
	Password string
}

func (r CreateUserRequest) toSQLC() sqlc.CreateUserParams {
	return sqlc.CreateUserParams{
		Name:           r.Name,
		Email:          r.Email,
		HashedPassword: r.Password,
	}
}
