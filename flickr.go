package moul

import (
	"encoding/json"

	"github.com/parnurzeal/gorequest"
	"github.com/patrickmn/go-cache"
)

func init() {
	RegisterAction("flickr-feed", GetFlickrFeedAction)
}

const FlickrFeedURL = "https://api.flickr.com/services/feeds/photos_public.gne?format=json&id=38994875@N06"

type FlickrEntry struct {
	Title string `json:"title"`
	Link  string `json:"link"`
	Media struct {
		M string `json:"m"`
	} `json:"media"`
	DateTaken   string `json:"date_taken"`
	Description string `json:"description"`
	Published   string `json:"published"`
	Author      string `json:"author"`
	AuthorID    string `json:"author_id"`
	Tags        string `json:"tags"`
}

type FlickrResponse struct {
	Title       string        `json:"title"`
	Link        string        `json:"link"`
	Description string        `json:"description"`
	Modified    string        `json:"modified"`
	Generator   string        `json:"generator"`
	Items       []FlickrEntry `json:"items"`
}

func GetFlickrFeedAction(args []string) (interface{}, error) {
	if posts, found := moulCache.Get("flickr-posts"); found {
		return posts, nil
	}
	posts, err := GetFlickrFeed()
	if err != nil {
		return nil, err
	}
	moulCache.Set("flickr-posts", posts, cache.DefaultExpiration)
	return posts, nil
}

func GetFlickrFeed() (*FlickrResponse, error) {
	_, body, errs := gorequest.New().Get(FlickrFeedURL).End()
	if len(errs) > 0 {
		return nil, errs[0]
	}

	body = body[15 : len(body)-1]

	var response FlickrResponse
	err := json.Unmarshal([]byte(body), &response)
	return &response, err
}
