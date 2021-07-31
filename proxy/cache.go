package proxy

import (
	"log"
	"rest-api/ethclient"
	"rest-api/models"
	"time"

	"github.com/patrickmn/go-cache"
)

type Proxy interface {
	GetBlock(key string) (models.Block, error)
}

var Cache = cache.New(5*time.Minute, 5*time.Minute)

func setCache(key string, block models.Block) bool {
	Cache.Set(key, block, cache.NoExpiration)
	return true
}

func getCache(key string) (models.Block, bool) {
	var block models.Block
	var found bool
	data, found := Cache.Get(key)
	if found {
		block = data.(models.Block)
	}
	return block, found
}

type ProxyImpl struct {
	ethClient *ethclient.CloudflareEthGateway
}

func NewProxyImpl(ethClient *ethclient.CloudflareEthGateway) *ProxyImpl {
	return &ProxyImpl{
		ethClient: ethClient,
	}
}

func (p *ProxyImpl) GetBlock(key string) (*models.Block, error) {
	block, found := getCache(key)
	log.Print(key)

	log.Print(found)
	if !found {
		//get from eth client
		data, err := p.ethClient.GetBlock(key)
		if err == nil {
			setCache(key, *data)
			return data, nil
		}
	}

	return &block, nil
}
