-- name: CreateReadingHistory :one
INSERT INTO reading_histories (user_id, book_id, status, start_date, end_date)
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetReadingHistoryByUserID :many
SELECT *
FROM reading_histories
WHERE user_id = $1
ORDER BY user_id LIMIT $2
OFFSET $3;

-- name: GetReadingHistoryByUserAndBook :one
SELECT *
FROM reading_histories
WHERE user_id = $1
  AND book_id = $2;

-- name: GetReadingHistoryByUserAndStatus :many
SELECT *
FROM reading_histories
WHERE user_id = $1
  AND status = $2
ORDER BY status LIMIT $3
OFFSET $4;

-- name: UpdateReadingHistory :one
UPDATE reading_histories
SET status     = $3,
    start_date = $4,
    end_date   = $5,
    updated_at = now()
WHERE user_id = $1
  AND book_id = $2 RETURNING *;

-- name: DeleteReadingHistory :exec
DELETE
FROM reading_histories
WHERE user_id = $1
  AND book_id = $2;