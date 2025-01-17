// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: posts.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (
    id, created_at, updated_at, title, link, description, author, 
    category, comments, enclosure, guid, pubDate, source, feed_id
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
RETURNING id, created_at, updated_at, title, link, description, author, category, comments, enclosure, guid, pubdate, source, feed_id
`

type CreatePostParams struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       sql.NullString
	Link        sql.NullString
	Description sql.NullString
	Author      sql.NullString
	Category    sql.NullString
	Comments    sql.NullString
	Enclosure   []string
	Guid        sql.NullString
	Pubdate     sql.NullString
	Source      sql.NullString
	FeedID      uuid.UUID
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Title,
		arg.Link,
		arg.Description,
		arg.Author,
		arg.Category,
		arg.Comments,
		pq.Array(arg.Enclosure),
		arg.Guid,
		arg.Pubdate,
		arg.Source,
		arg.FeedID,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		&i.Link,
		&i.Description,
		&i.Author,
		&i.Category,
		&i.Comments,
		pq.Array(&i.Enclosure),
		&i.Guid,
		&i.Pubdate,
		&i.Source,
		&i.FeedID,
	)
	return i, err
}

const getPostsByUser = `-- name: GetPostsByUser :many
SELECT users.id, users.created_at, users.updated_at, name, api_key, feeds.id, feeds.created_at, feeds.updated_at, url, user_id, last_fetched_at, posts.id, posts.created_at, posts.updated_at, title, link, description, author, category, comments, enclosure, guid, pubdate, source, feed_id from users 
JOIN feeds ON users.id = feeds.user_id 
JOIN posts ON posts.feed_id = feeds.id
WHERE users.id = $1
ORDER BY posts.pubDate
LIMIT $2
`

type GetPostsByUserParams struct {
	ID    uuid.UUID
	Limit int32
}

type GetPostsByUserRow struct {
	ID            uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Name          string
	ApiKey        string
	ID_2          uuid.UUID
	CreatedAt_2   time.Time
	UpdatedAt_2   time.Time
	Url           string
	UserID        uuid.UUID
	LastFetchedAt sql.NullTime
	ID_3          uuid.UUID
	CreatedAt_3   time.Time
	UpdatedAt_3   time.Time
	Title         sql.NullString
	Link          sql.NullString
	Description   sql.NullString
	Author        sql.NullString
	Category      sql.NullString
	Comments      sql.NullString
	Enclosure     []string
	Guid          sql.NullString
	Pubdate       sql.NullString
	Source        sql.NullString
	FeedID        uuid.UUID
}

func (q *Queries) GetPostsByUser(ctx context.Context, arg GetPostsByUserParams) ([]GetPostsByUserRow, error) {
	rows, err := q.db.QueryContext(ctx, getPostsByUser, arg.ID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPostsByUserRow
	for rows.Next() {
		var i GetPostsByUserRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.ApiKey,
			&i.ID_2,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
			&i.Url,
			&i.UserID,
			&i.LastFetchedAt,
			&i.ID_3,
			&i.CreatedAt_3,
			&i.UpdatedAt_3,
			&i.Title,
			&i.Link,
			&i.Description,
			&i.Author,
			&i.Category,
			&i.Comments,
			pq.Array(&i.Enclosure),
			&i.Guid,
			&i.Pubdate,
			&i.Source,
			&i.FeedID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
