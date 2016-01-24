package moul

import (
	"github.com/SlyMarbo/rss"
	"github.com/patrickmn/go-cache"
)

var instagramFeed *rss.Feed

func init() {
	RegisterAction("instagram-activity", GetInstagramActivityAction)
}

func GetInstagramActivityAction(args []string) (interface{}, error) {
	if activity, found := moulCache.Get("instagram-activity"); found {
		return activity, nil
	}
	activity, err := GetInstagramActivity()
	if err != nil {
		return nil, err
	}
	moulCache.Set("instagram-activity", activity, cache.DefaultExpiration)
	return activity, nil
}

func GetInstagramActivity() (*rss.Feed, error) {
	var err error
	if instagramFeed == nil {
		instagramFeed, err = rss.Fetch("http://widget.websta.me/rss/n/m42am")
	} else {
		err = instagramFeed.Update()
	}
	return instagramFeed, err
}
