package moul

import (
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/patrickmn/go-cache"
)

func init() {
	RegisterAction("twitter-last-tweets", GetTwitterLastTweetsAction)
	RegisterAction("twitter-followers-list", GetTwitterFollowersListAction)
}

func GetTwitterLastTweetsAction(args []string) (interface{}, error) {
	if tweets, found := moulCache.Get("twitter-last-tweets"); found {
		return tweets, nil
	}
	tweets, err := GetTwitterLastTweets()
	if err != nil {
		return nil, err
	}
	moulCache.Set("twitter-last-tweets", tweets, cache.DefaultExpiration)
	return tweets, nil
}

func GetTwitterFollowersListAction(args []string) (interface{}, error) {
	if followers, found := moulCache.Get("twitter-followers"); found {
		return followers, nil
	}
	followers, err := GetTwitterFollowersList()
	if err != nil {
		return nil, err
	}
	moulCache.Set("twitter-followers", followers, cache.DefaultExpiration)
	return followers, nil
}

func initTwitterAPI() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(os.Getenv("TWITTER_CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("TWITTER_CONSUMER_SECRET"))
	api := anaconda.NewTwitterApi(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"))
	return api
}

func GetTwitterFollowersList() (interface{}, error) {
	api := initTwitterAPI()
	return api.GetFollowersList(nil)
}

func GetTwitterLastTweets() (interface{}, error) {
	api := initTwitterAPI()
	return api.GetUserTimeline(nil)
}
