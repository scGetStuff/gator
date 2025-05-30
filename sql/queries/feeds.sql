-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeeds :many
SELECT f.name, f.url, u.name AS user
FROM feeds f
INNER JOIN users u
ON f.user_id = u.id;
