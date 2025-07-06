package repository

import sqlc "readly/db/sqlc"

type UpdateUserRequest struct {
	ID       int64
	Name     string
	Email    string
	Password string
}

func (r UpdateUserRequest) toSQLC() sqlc.UpdateUserParams {
	return sqlc.UpdateUserParams{
		ID:             r.ID,
		Name:           r.Name,
		Email:          r.Email,
		HashedPassword: r.Password,
	}
}
