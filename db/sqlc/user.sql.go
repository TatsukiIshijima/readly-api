// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (name,
                   email,
                   hashed_password)
VALUES ($1,
        $2,
        $3) RETURNING id, name, email, hashed_password, created_at, updated_at
`

type CreateUserParams struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Name, arg.Email, arg.HashedPassword)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.HashedPassword,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE
FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getAllUsers = `-- name: GetAllUsers :many
SELECT id, name, email, hashed_password, created_at, updated_at
FROM users
ORDER BY id LIMIT $1
OFFSET $2
`

type GetAllUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetAllUsers(ctx context.Context, arg GetAllUsersParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getAllUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.HashedPassword,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, name, email, hashed_password, created_at, updated_at
FROM users
WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.HashedPassword,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, name, email, hashed_password, created_at, updated_at
FROM users
WHERE id = $1
`

func (q *Queries) GetUserById(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.HashedPassword,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET name            = $2,
    email           = $3,
    hashed_password = $4,
    updated_at      = now()
WHERE id = $1 RETURNING id, name, email, hashed_password, created_at, updated_at
`

type UpdateUserParams struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.HashedPassword,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.HashedPassword,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
