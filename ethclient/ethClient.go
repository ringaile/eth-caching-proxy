package ethclient

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"rest-api/models"
	"time"
)

type EthClient interface {
	GetBlock(param string) (*models.Block, error)
}

type CloudflareEthGateway struct {
	Url    string
	Client *http.Client
}

func NewCloudflareEthGateway(url string, timeout int32) *CloudflareEthGateway {
	return &CloudflareEthGateway{
		Url: url,
		Client: &http.Client{
			Timeout: time.Duration(rand.Int31n(timeout)) * time.Second,
		},
	}
}

func (c *CloudflareEthGateway) GetBlock(param string) (*models.Block, error) {
	request := &models.GetBlockRequest{
		Jsonrpc: "2.0",
		Method:  "eth_getBlockByNumber",
		Params:  []interface{}{param, true},
		ID:      1,
	}

	data, _ := json.Marshal(request)

	req, err := http.NewRequest("POST", c.Url, bytes.NewBuffer(data))

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, nil
	}
	defer resp.Body.Close()

	var response *models.Result
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, nil
	}

	return &response.Block, nil
}
