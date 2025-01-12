-- +goose Up
CREATE TABLE posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT,
    link TEXT UNIQUE,
    description TEXT,
    author TEXT,
    category TEXT,
    comments TEXT,
    enclosure TEXT[],
    guid TEXT,
    pubDate TEXT,
    source TEXT,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;