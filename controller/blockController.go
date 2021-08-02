package controller

import (
	"rest-api/ethclient"
	"rest-api/models"
	"rest-api/proxy"
)

type BlockController struct {
	ethClient *ethclient.CloudflareEthGateway
	proxy     *proxy.ProxyImpl
}

func NewBlockController(ethClient *ethclient.CloudflareEthGateway, proxy *proxy.ProxyImpl) *BlockController {
	return &BlockController{
		ethClient: ethClient,
		proxy:     proxy,
	}
}

func (c *BlockController) GetBlock(key string) (*models.Block, error) {
	latestBlock, _ := c.ethClient.GetBlock("latest")
	var block *models.Block

	latest := checkIfBlockIsWithinLast20(key, latestBlock.Number)

	if latest {
		data, err := c.ethClient.GetBlock(key)
		if err == nil {
			//do not save latest blocks cause they change
			block = data
		}
	} else {
		temp, found := c.proxy.GetCache(key)
		if !found {
			//get from eth client
			data, err := c.ethClient.GetBlock(key)
			if err == nil {
				c.proxy.SetCache(key, data)
				block = data
			}
		} else {
			block = temp
		}
	}

	return block, nil
}
