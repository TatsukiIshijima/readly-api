-- name: GetAllGenres :many
SELECT * FROM genres;

-- name: GetGenreByID :one
SELECT * FROM genres WHERE id = $1;

-- name: InsertGenre :exec
INSERT INTO genres (name) VALUES ($1);

-- name: UpdateGenre :exec
UPDATE genres SET name = $1 WHERE id = $2;

-- name: DeleteGenre :exec
DELETE FROM genres WHERE id = $1;