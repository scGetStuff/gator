-- name: CreateFeedFollow :one
WITH insert_rec AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
)
SELECT
    insert_rec.*,
    f.name AS feed_name,
    u.name AS user_name
FROM insert_rec
INNER JOIN feeds f ON insert_rec.feed_id = f.id
INNER JOIN users u ON insert_rec.user_id = u.id;

-- name: GetFeedFollowsForUser :many
SELECT name FROM feeds WHERE id IN (SELECT feed_id FROM feed_follows WHERE feed_follows.user_id = $1);

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows WHERE user_id = $1 AND feed_id = $2;