-- name: GetAllGenres :many
SELECT *
FROM genres;

-- name: GetGenreByName :one
SELECT *
FROM genres
WHERE name = $1;

-- name: InsertGenre :one
INSERT INTO genres (name)
VALUES ($1) RETURNING *;

-- name: DeleteGenre :exec
DELETE
FROM genres
WHERE name = $1;