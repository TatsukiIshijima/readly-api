-- name: CreateReadingHistory :one
INSERT INTO reading_histories (user_id, book_id, status, start_date, end_date)
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetReadingHistoryByUser :many
WITH genre_aggregation AS (SELECT bg.book_id,
                                  STRING_AGG(g.name, ', ') AS genres
                           FROM book_genres bg
                                    LEFT JOIN genres g ON bg.genre_name = g.name
                           GROUP BY bg.book_id)

SELECT b.id,
       b.title,
       ga.genres,
       b.description,
       b.cover_image_url,
       b.url,
       b.author_name,
       b.publisher_name,
       b.published_date,
       b.isbn,
       rh.status,
       rh.start_date,
       rh.end_date
FROM reading_histories rh
         LEFT JOIN books b ON b.id = rh.book_id
         LEFT JOIN genre_aggregation ga ON b.id = ga.book_id
WHERE rh.user_id = $1
ORDER BY rh.created_at LIMIT $2
OFFSET $3;

-- name: GetReadingHistoryByUserAndBook :one
WITH genre_aggregation AS (SELECT bg.book_id,
                                  STRING_AGG(g.name, ', ') AS genres
                           FROM book_genres bg
                                    LEFT JOIN genres g ON bg.genre_name = g.name
                           GROUP BY bg.book_id)

SELECT b.id,
       b.title,
       ga.genres,
       b.description,
       b.cover_image_url,
       b.url,
       b.author_name,
       b.publisher_name,
       b.published_date,
       b.isbn,
       rh.status,
       rh.start_date,
       rh.end_date
FROM reading_histories rh
         LEFT JOIN books b ON b.id = rh.book_id
         LEFT JOIN genre_aggregation ga ON b.id = ga.book_id
WHERE rh.user_id = $1
  AND rh.book_id = $2;

-- name: GetReadingHistoryByUserAndStatus :many
WITH genre_aggregation AS (SELECT bg.book_id,
                                  STRING_AGG(g.name, ', ') AS genres
                           FROM book_genres bg
                                    LEFT JOIN genres g ON bg.genre_name = g.name
                           GROUP BY bg.book_id)

SELECT b.id,
       b.title,
       ga.genres,
       b.description,
       b.cover_image_url,
       b.url,
       b.author_name,
       b.publisher_name,
       b.published_date,
       b.isbn,
       rh.status,
       rh.start_date,
       rh.end_date
FROM reading_histories rh
         LEFT JOIN books b ON b.id = rh.book_id
         LEFT JOIN genre_aggregation ga ON b.id = ga.book_id
WHERE rh.user_id = $1
  AND rh.status = $2
ORDER BY rh.created_at LIMIT $3
OFFSET $4;

-- name: UpdateReadingHistory :one
UPDATE reading_histories
SET status     = $3,
    start_date = $4,
    end_date   = $5,
    updated_at = now()
WHERE user_id = $1
  AND book_id = $2 RETURNING *;

-- name: DeleteReadingHistory :execrows
DELETE
FROM reading_histories
WHERE user_id = $1
  AND book_id = $2;