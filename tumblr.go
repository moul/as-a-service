package moul

import (
	"encoding/json"

	"github.com/parnurzeal/gorequest"
	"github.com/patrickmn/go-cache"
)

func init() {
	RegisterAction("tumblr", GetLatestBlogPostsAction)
}

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

func GetLatestBlogPostsAction(args []string) (interface{}, error) {
	if posts, found := moulCache.Get("tumblr-posts"); found {
		return posts, nil
	}
	posts, err := GetLatestBlogPosts()
	if err != nil {
		return nil, err
	}
	moulCache.Set("tumblr-posts", posts, cache.DefaultExpiration)
	return posts, nil
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
