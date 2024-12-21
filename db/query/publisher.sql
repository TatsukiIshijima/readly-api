-- name: GetAllPublishers :many
SELECT *
FROM publishers;

-- name: ListPublishers :many
SELECT *
FROM publishers LIMIT $1
OFFSET $2;

-- name: GetPublisherByName :one
SELECT *
FROM publishers
WHERE name = $1;

-- name: InsertPublisher :one
INSERT INTO publishers (name)
VALUES ($1) RETURNING *;

-- name: DeletePublisher :exec
DELETE
FROM publishers
WHERE name = $1;