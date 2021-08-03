package ethcache

import (
	"math/rand"
	"rest-api/model"
	"time"

	"github.com/patrickmn/go-cache"
	c "github.com/patrickmn/go-cache"
)

type EthCache interface {
	SetCache(key string, block *model.Block)
	GetCache(key string) (*model.Block, bool)
}

type EthCacheImpl struct {
	cache *c.Cache
}

func New(defaultExpiration int32, cleanupExpiration int32) EthCacheImpl {
	return EthCacheImpl{
		cache: c.New(time.Duration(rand.Int31n(defaultExpiration))*time.Minute, time.Duration(rand.Int31n(cleanupExpiration))*time.Minute),
	}
}

func (p *EthCacheImpl) SetCache(key string, block *model.Block) {
	p.cache.Set(key, block, cache.DefaultExpiration)
}

func (p *EthCacheImpl) GetCache(key string) (*model.Block, bool) {
	var block *model.Block
	var found bool
	data, found := p.cache.Get(key)
	if found {
		block = data.(*model.Block)
	}
	return block, found
}
