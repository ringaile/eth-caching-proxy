package proxy

import (
	"rest-api/models"
	"time"

	"github.com/patrickmn/go-cache"
)

var Cache = cache.New(5*time.Minute, 5*time.Minute)

func SetCache(key string, block models.Block) bool {
	Cache.Set(key, block, cache.NoExpiration)
	return true
}

func GetCache(key string) (models.Block, bool) {
	var block models.Block
	var found bool
	data, found := Cache.Get(key)
	if found {
		block = data.(models.Block)
	}
	return block, found
}
