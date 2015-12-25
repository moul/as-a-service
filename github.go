package moul

import "github.com/SlyMarbo/rss"

var feed *rss.Feed

func GetGithubActivity() (*rss.Feed, error) {
	var err error
	if feed == nil {
		feed, err = rss.Fetch("https://github.com/moul.atom")
	} else {
		err = feed.Update()
	}
	return feed, err
}