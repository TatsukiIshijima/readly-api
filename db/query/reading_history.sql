-- name: InsertReadingHistory :exec
INSERT INTO reading_histories (user_id, book_id, status, start_date, end_date)
VALUES ($1, $2, $3, $4, $5);

-- name: GetReadingHistoryByUserID :many
SELECT *
FROM reading_histories
WHERE user_id = $1;

-- name: GetReadingHistoryByUserAndBook :one
SELECT *
FROM reading_histories
WHERE user_id = $1
  AND book_id = $2;

-- name: GetReadingHistoryByUserAndStatus :many
SELECT *
FROM reading_histories
WHERE user_id = $1
  AND status = $2;

-- name: UpdateReadingStatus :exec
UPDATE reading_histories
SET status     = $3,
    updated_at = now()
WHERE user_id = $1
  AND book_id = $2;

-- name: UpdateReadingDates :exec
UPDATE reading_histories
SET start_date = $3,
    end_date   = $4,
    updated_at = now()
WHERE user_id = $1
  AND book_id = $2;