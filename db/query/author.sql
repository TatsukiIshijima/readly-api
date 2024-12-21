-- name: GetAllAuthors :many
SELECT *
FROM authors;

-- name: ListAuthors :many
SELECT *
FROM authors LIMIT $1
OFFSET $2;

-- name: GetAuthorByName :one
SELECT *
FROM authors
WHERE name = $1;

-- name: InsertAuthor :one
INSERT INTO authors (name)
VALUES ($1) RETURNING *;

-- name: DeleteAuthor :exec
DELETE
FROM authors
WHERE name = $1;