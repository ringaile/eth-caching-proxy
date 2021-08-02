package proxy

import (
	"math/rand"
	"rest-api/models"
	"time"

	"github.com/patrickmn/go-cache"
	c "github.com/patrickmn/go-cache"
)

type Proxy interface {
	SetCache(key string, block models.Block)
	GetCache(key string) (models.Block, bool)
}

type ProxyImpl struct {
	cache *c.Cache
}

func New(defaultExpiration int32, cleanupExpiration int32) *ProxyImpl {
	return &ProxyImpl{
		cache: c.New(time.Duration(rand.Int31n(defaultExpiration))*time.Minute, time.Duration(rand.Int31n(cleanupExpiration))*time.Minute),
	}
}

func (p *ProxyImpl) SetCache(key string, block *models.Block) {
	p.cache.Set(key, block, cache.DefaultExpiration)
}

func (p *ProxyImpl) GetCache(key string) (*models.Block, bool) {
	var block models.Block
	var found bool
	data, found := p.cache.Get(key)
	if found {
		block = data.(models.Block)
	}
	return &block, found
}
