package moul

import (
	"github.com/SlyMarbo/rss"
	"github.com/google/go-github/github"
	"github.com/patrickmn/go-cache"
)

var feed *rss.Feed

func init() {
	RegisterAction("github-activity", GetGithubActivityAction)
	RegisterAction("github-repos", GetGithubReposAction)
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

func GetGithubReposAction(args []string) (interface{}, error) {
	if repos, found := moulCache.Get("github-repos"); found {
		return repos, nil
	}
	repos, err := GetGithubRepos()
	if err != nil {
		return nil, err
	}
	moulCache.Set("github-repos", repos, cache.DefaultExpiration)
	return repos, nil
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

func GetGithubRepos() (interface{}, error) {
	client := github.NewClient(nil)
	opt := &github.RepositoryListOptions{Type: "owner", Sort: "updated", Direction: "desc"}
	repos, _, err := client.Repositories.List("moul", opt)
	return repos, err
}
