-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, url, user_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;


-- name: GetFeeds :many
SELECT * FROM feeds WHERE user_id = $1;