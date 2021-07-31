package ethclient

import (
	"bytes"
	"encoding/json"
	"net/http"
	"rest-api/models"
	"time"
)

type EthClient interface {
	GetBlock(param string) (*models.Block, error)
}

type CloudflareEthGateway struct {
	Url    string
	client *http.Client
}

func NewCloudflareEthGateway(url string, _ string) CloudflareEthGateway {
	cloudflareEthGateway := CloudflareEthGateway{Url: url}

	// timeOutIntv, err := time.ParseDuration(timeout)
	// if err != nil {
	// 	panic("util: Can't parse duration `" + timeout + "`: " + err.Error())
	// }

	cloudflareEthGateway.client = &http.Client{
		Timeout: 5 * time.Second,
	}

	return cloudflareEthGateway
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

	resp, err := c.client.Do(req)
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
