package main

import (
	"main/internal/database"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       *string   `json:"title"`
	Link        *string   `json:"link"`
	Description *string   `json:"description"`
	Author      *string   `json:"author"`
	Category    *string   `json:"category"`
	Comments    *string   `json:"comment"`
	Enclosure   []string  `json:"enclosure"`
	Guid        *string   `json:"guid"`
	Pubdate     *string   `json:"published_at"`
	Source      *string   `json:"source"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func databaseUsertoUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		ApiKey:    dbUser.ApiKey,
	}
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}

func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(dbFeed))
	}
	return feeds
}

func databasePostToPost(dbPost database.GetPostsByUserRow) Post {
	var title *string
	if dbPost.Title.Valid {
		title = &dbPost.Title.String
	}
	var link *string
	if dbPost.Link.Valid {
		link = &dbPost.Link.String
	}
	var description *string
	if dbPost.Description.Valid {
		description = &dbPost.Description.String
	}
	var author *string
	if dbPost.Author.Valid {
		author = &dbPost.Author.String
	}
	var category *string
	if dbPost.Category.Valid {
		category = &dbPost.Category.String
	}
	var comments *string
	if dbPost.Comments.Valid {
		comments = &dbPost.Comments.String
	}
	var guid *string
	if dbPost.Guid.Valid {
		guid = &dbPost.Guid.String
	}
	var pubdate *string
	if dbPost.Pubdate.Valid {
		pubdate = &dbPost.Pubdate.String
	}
	var source *string
	if dbPost.Source.Valid {
		source = &dbPost.Source.String
	}

	return Post{
		ID:          dbPost.ID_3,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   dbPost.UpdatedAt,
		Title:       title,
		Link:        link,
		Description: description,
		Author:      author,
		Category:    category,
		Comments:    comments,
		Enclosure:   dbPost.Enclosure,
		Guid:        guid,
		Pubdate:     pubdate,
		Source:      source,
		FeedID:      dbPost.FeedID,
	}
}

func databasePostsToPosts(dbPosts []database.GetPostsByUserRow) []Post {
	posts := []Post{}
	for _, post := range dbPosts {
		posts = append(posts, databasePostToPost(post))
	}
	return posts
}
