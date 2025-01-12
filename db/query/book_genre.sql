-- name: CreateBookGenre :one
INSERT INTO book_genres (book_id, genre_name)
VALUES ($1, $2) RETURNING *;

-- name: GetGenresByBookID :many
SELECT genre_name
FROM book_genres
WHERE book_id = $1;

-- name: DeleteBookGenre :exec
DELETE
FROM book_genres
WHERE book_id = $1
  AND genre_name = $2;