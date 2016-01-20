package moul

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var moulCache *cache.Cache

func init() {
	moulCache = cache.New(5*time.Minute, 30*time.Second)
}
