-- name: ListUsers :many
SELECT *
FROM users LIMIT $1
OFFSET $2;

-- name: GetUserById :one
SELECT *
FROM users
WHERE id = $1;

-- name: GetUserByName :one
SELECT *
FROM users
WHERE name = $1;

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

-- name: UpdateUserName :one
UPDATE users
SET name       = $2,
    updated_at = $3
WHERE id = $1 RETURNING *;

-- name: UpdateUserEmail :one
UPDATE users
SET email      = $2,
    updated_at = $3
WHERE id = $1 RETURNING *;

-- name: UpdateUserPassword :one
UPDATE users
SET hashed_password = $2,
    updated_at      = $3
WHERE id = $1 RETURNING *;