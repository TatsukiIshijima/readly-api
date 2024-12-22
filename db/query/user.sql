-- name: ListUsers :many
SELECT *
FROM users LIMIT $1
OFFSET $2;

-- name: GetUserById :one
SELECT *
FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;

-- name: InsertUser :one
INSERT INTO users (name,
                   email,
                   hashed_password)
VALUES ($1,
        $2,
        $3) RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET name            = $2,
    email           = $3,
    hashed_password = $4,
    updated_at      = now()
WHERE id = $1 RETURNING *;