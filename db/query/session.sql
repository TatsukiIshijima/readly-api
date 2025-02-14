-- name: CreateSession :one
INSERT INTO sessions (id,
                      user_id,
                      refresh_token,
                      expires_at,
                      ip_address,
                      user_agent)
VALUES ($1,
        $2,
        $3,
        $4,
        $5,
        $6) RETURNING *;

-- name: GetSessionByID :one
SELECT *
FROM sessions
WHERE id = $1;

-- name: GetSessionByUserID :many
SELECT *
FROM sessions
WHERE user_id = $1;

-- name: UpdateSession :one
UPDATE sessions
SET refresh_token = $2,
    expires_at    = $3,
    ip_address    = $4,
    user_agent    = $5,
    revoked       = $6,
    revoked_at    = $7
WHERE id = $1 RETURNING *;

-- name: DeleteSessionByUserID :execrows
DELETE
FROM sessions
WHERE id IN (SELECT s.id
             FROM sessions AS s
             WHERE s.user_id = $1
             ORDER BY s.created_at ASC
    LIMIT $2
    );