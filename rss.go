package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

// RSS represents the root element of the RSS feed
type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

// Channel represents the channel element in the RSS feed
type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

// Item represents each item in the RSS feed
type Item struct {
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	Author      string   `xml:"author"`
	Categories  []string `xml:"category"`
	Comments    string   `xml:"comments"`
	Enclosure   []Media  `xml:"content"`
	GUID        string   `xml:"guid"`
	PubDate     string   `xml:"pubDate"`
	Source      string   `xml:"source"`
}

// Media represents the media:content element in the RSS feed
type Media struct {
	URL string `xml:"url,attr"`
}

func UrlToFeed(url string) (RSS, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := httpClient.Get(url)
	if err != nil {
		return RSS{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RSS{}, err
	}

	rssFeed := RSS{}
	err = xml.Unmarshal(data, &rssFeed)
	if err != nil {
		return RSS{}, err
	}
	return rssFeed, nil
}
