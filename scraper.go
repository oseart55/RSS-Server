package main

import (
	"context"
	"database/sql"
	"log"
	"main/internal/database"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

func StartScraping(
	db *database.Queries,
	concurrency int,
	timeBetweenRequest time.Duration,
) {
	log.Printf("Starting Scraping: %v Goroutines. Time between requests: %s\n", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			log.Println(err)
			continue
		}
		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error Marking Feed as Fetched: ", err)
		return
	}
	rssFeed, err := UrlToFeed(feed.Url)
	if err != nil {
		log.Println("Error getting Feed: ", err)
		return
	}

	for _, item := range rssFeed.Channel.Items {
		title := sql.NullString{}
		link := sql.NullString{}
		description := sql.NullString{}
		author := sql.NullString{}
		category := sql.NullString{}
		comments := sql.NullString{}
		guid := sql.NullString{}
		pubdate := sql.NullString{}
		source := sql.NullString{}

		if item.Title != "" {
			title.String = item.Title
			title.Valid = true
		}
		if item.Link != "" {
			link.String = item.Link
			link.Valid = true
		}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}
		if item.Author != "" {
			author.String = item.Author
			author.Valid = true
		}
		if strings.Join(item.Categories, " ") != "" {
			category.String = strings.Join(item.Categories, " ")
			category.Valid = true
		}
		if item.Comments != "" {
			comments.String = item.Comments
			comments.Valid = true
		}
		if item.GUID != "" {
			guid.String = item.GUID
			guid.Valid = true
		}
		if item.PubDate != "" {
			pubdate.String = item.PubDate
			pubdate.Valid = true
		}
		if item.Source != "" {
			source.String = item.Source
			source.Valid = true
		}
		var urls []string
		for _, media := range item.Enclosure {
			urls = append(urls, media.URL)
		}
		_, err := db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       title,
			Link:        link,
			Description: description,
			Author:      author,
			Category:    category,
			Comments:    comments,
			Enclosure:   urls,
			Guid:        guid,
			Pubdate:     pubdate,
			Source:      source,
			FeedID:      feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Println("Error Creating Post: ", err)
		}
	}
	log.Printf("Collected %s, %v posts found", feed.Url, len(rssFeed.Channel.Items))
}
