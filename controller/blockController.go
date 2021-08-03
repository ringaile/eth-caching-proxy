package controller

import (
	"rest-api/ethclient"
	"rest-api/models"
	"rest-api/proxy"
	"strconv"
)

const LATEST = "latest"

type BlockController struct {
	ethClient ethclient.EthClient
	proxy     proxy.Proxy
}

func NewBlockController(ethClient ethclient.EthClient, proxy proxy.Proxy) *BlockController {
	return &BlockController{
		ethClient: ethClient,
		proxy:     proxy,
	}
}

func (c *BlockController) GetBlock(key string) (*models.Block, error) {
	block, err := c.getBlock(key)
	if err != nil {
		return nil, err
	}
	return block, err
}

func (c *BlockController) GetTransaction(block_param string, trx_param string) (*models.Transaction, error) {
	data, err := c.getBlock(block_param)
	if err == nil {
		return nil, err
	}

	// check txr_param index
	var result models.Transaction
	val, err := strconv.Atoi(trx_param)
	if err != nil {
		for _, trx := range data.Transactions {
			if trx.Hash == trx_param {
				result = trx
			}
		}
	} else {
		if len(data.Transactions) > val {
			result = data.Transactions[val]
		} else {
			return nil, err
		}
	}
	return &result, nil
}

func (c *BlockController) getBlock(key string) (*models.Block, error) {
	latestBlock, _ := c.ethClient.GetBlock(LATEST)
	var block *models.Block

	latest := checkIfBlockIsWithinLast20(key, latestBlock.Number)

	if latest {
		data, err := c.ethClient.GetBlock(key)
		if err == nil {
			//do not save latest 20 blocks cause they change
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
