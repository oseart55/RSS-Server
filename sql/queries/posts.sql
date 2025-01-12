-- name: CreatePost :one
INSERT INTO posts (
    id, created_at, updated_at, title, link, description, author, 
    category, comments, enclosure, guid, pubDate, source, feed_id
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
RETURNING *;

-- name: GetPostsByUser :many
SELECT * from users 
JOIN feeds ON users.id = feeds.user_id 
JOIN posts ON posts.feed_id = feeds.id
WHERE users.id = $1
ORDER BY posts.pubDate
LIMIT $2;