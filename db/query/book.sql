-- name: GetBooksByID :one
SELECT b.id,
       b.title,
       STRING_AGG(g.name, ', ') AS genres,
       b.description,
       b.cover_image_url,
       b.url,
       b.author_name,
       b.publisher_name,
       b.published_date,
       b.isbn,
       b.created_at,
       b.updated_at
FROM books b
         LEFT JOIN book_genres bg ON b.id = bg.book_id
         LEFT JOIN genres g ON bg.genre_name = g.name
WHERE b.id = $1
GROUP BY b.id;

-- name: GetBooksByTitle :many
SELECT b.id,
       b.title,
       STRING_AGG(g.name, ', ') AS genres,
       b.description,
       b.cover_image_url,
       b.url,
       b.author_name,
       b.publisher_name,
       b.published_date,
       b.isbn,
       b.created_at,
       b.updated_at
FROM books b
         LEFT JOIN book_genres bg ON b.id = bg.book_id
         LEFT JOIN genres g ON bg.genre_name = g.name
WHERE b.title LIKE $1
GROUP BY b.id
ORDER BY b.created_at;

-- name: GetBooksByISBN :many
SELECT b.id,
       b.title,
       STRING_AGG(g.name, ', ') AS genres,
       b.description,
       b.cover_image_url,
       b.url,
       b.author_name,
       b.publisher_name,
       b.published_date,
       b.isbn,
       b.created_at,
       b.updated_at
FROM books b
         LEFT JOIN book_genres bg ON b.id = bg.book_id
         LEFT JOIN genres g ON bg.genre_name = g.name
WHERE b.isbn = $1
GROUP BY b.id
ORDER BY b.created_at;

-- name: GetBooksByAuthor :many
SELECT b.id,
       b.title,
       STRING_AGG(g.name, ', ') AS genres,
       b.description,
       b.cover_image_url,
       b.url,
       b.author_name,
       b.publisher_name,
       b.published_date,
       b.isbn,
       b.created_at,
       b.updated_at
FROM books b
         LEFT JOIN book_genres bg ON b.id = bg.book_id
         LEFT JOIN genres g ON bg.genre_name = g.name
WHERE b.author_name LIKE $1
GROUP BY b.id
ORDER BY b.created_at;

-- name: GetBooksByPublisher :many
SELECT b.id,
       b.title,
       STRING_AGG(g.name, ', ') AS genres,
       b.description,
       b.cover_image_url,
       b.url,
       b.author_name,
       b.publisher_name,
       b.published_date,
       b.isbn,
       b.created_at,
       b.updated_at
FROM books b
         LEFT JOIN book_genres bg ON b.id = bg.book_id
         LEFT JOIN genres g ON bg.genre_name = g.name
WHERE b.publisher_name LIKE $1
GROUP BY b.id
ORDER BY b.created_at;

-- name: CreateBook :one
INSERT
INTO books (title,
            description,
            cover_image_url,
            url,
            author_name,
            publisher_name,
            published_date,
            isbn)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *;

-- name: UpdateBook :one
UPDATE books
SET title           = $2,
    description     = $3,
    cover_image_url = $4,
    url             = $5,
    author_name     = $6,
    publisher_name  = $7,
    published_date  = $8,
    isbn            = $9,
    updated_at      = now()
WHERE id = $1 RETURNING *;

-- name: DeleteBook :execrows
DELETE
FROM books
WHERE id = $1;