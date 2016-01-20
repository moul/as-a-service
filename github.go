package moul

import (
	"github.com/SlyMarbo/rss"
	"github.com/patrickmn/go-cache"
)

var feed *rss.Feed

func init() {
	RegisterAction("github-activity", GetGithubActivityAction)
}

func GetGithubActivityAction(args []string) (interface{}, error) {
	if activity, found := moulCache.Get("github-activity"); found {
		return activity, nil
	}
	activity, err := GetGithubActivity()
	if err != nil {
		return nil, err
	}
	moulCache.Set("github-activity", activity, cache.DefaultExpiration)
	return activity, nil
}

func GetGithubActivity() (*rss.Feed, error) {
	var err error
	if feed == nil {
		feed, err = rss.Fetch("https://github.com/moul.atom")
	} else {
		err = feed.Update()
	}
	return feed, err
}
