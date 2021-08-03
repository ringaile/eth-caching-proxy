package controller

import (
	"rest-api/ethcache"
	"rest-api/ethclient"
	"rest-api/model"
	"strconv"
)

const LATEST = "latest"

type BlockController struct {
	ethClient ethclient.EthClient
	ethCache  ethcache.EthCache
}

func NewBlockController(ethClient ethclient.EthClient, ethCache ethcache.EthCache) *BlockController {
	return &BlockController{
		ethClient: ethClient,
		ethCache:  ethCache,
	}
}

func (c *BlockController) GetBlock(key string) (*model.Block, error) {
	block, err := c.getBlock(key)
	if err != nil {
		return nil, err
	}
	return block, err
}

func (c *BlockController) GetTransaction(block_param string, trx_param string) (*model.Transaction, error) {
	data, err := c.getBlock(block_param)
	if err != nil {
		return nil, err
	}

	// check txr_param index
	var result model.Transaction
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

func (c *BlockController) getBlock(key string) (*model.Block, error) {
	latestBlock, err := c.ethClient.GetBlock(LATEST)
	if err != nil {
		return nil, err
	}
	var block *model.Block

	latest := checkIfBlockIsWithinLast20(key, latestBlock.Number)

	if latest {
		data, err := c.ethClient.GetBlock(key)
		if err == nil {
			return nil, err
		}
		//do not save latest 20 blocks cause they change
		block = data
	} else {
		temp, found := c.ethCache.GetCache(key)
		if !found {
			//get from eth client
			data, err := c.ethClient.GetBlock(key)
			if err != nil {
				return nil, err
			}
			c.ethCache.SetCache(key, data)
			block = data
		} else {
			block = temp
		}
	}

	return block, nil
}
