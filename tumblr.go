package moul

import (
	"encoding/json"

	"github.com/parnurzeal/gorequest"
)

const TumblrFeedURL = "https://ajax.googleapis.com/ajax/services/feed/load?v=1.0&num=10&q=http://manfredtouron.tumblr.com/rss"

type TumblrEntry struct {
	Title          string
	Link           string
	Author         string
	PublishedDate  string
	ContentSnippet string
	Content        string
	Categories     []string
}

type TumblrResponse struct {
	ResponseData struct {
		Feed struct {
			FeedURL     string
			Title       string
			Link        string
			Author      string
			Description string
			Type        string
			Entries     []TumblrEntry
		}
	}
	// ResponseDetails
	ResponseStatus int
}

func GetLatestBlogPosts() (*TumblrResponse, error) {
	_, body, errs := gorequest.New().Get(TumblrFeedURL).End()
	if len(errs) > 0 {
		return nil, errs[0]
	}

	var response TumblrResponse
	err := json.Unmarshal([]byte(body), &response)
	return &response, err
}