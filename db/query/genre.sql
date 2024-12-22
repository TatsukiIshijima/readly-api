-- name: GetAllGenres :many
SELECT *
FROM genres LIMIT $1
OFFSET $2;

-- name: GetGenreByName :one
SELECT *
FROM genres
WHERE name = $1;

-- name: CreateGenre :one
INSERT INTO genres (name)
VALUES ($1) RETURNING *;

-- name: DeleteGenre :exec
DELETE
FROM genres
WHERE name = $1;