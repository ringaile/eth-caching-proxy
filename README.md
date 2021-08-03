# ETH Caching Proxy
A caching proxy for `eth_getBlockByNumber` written in golang.

## Installation & Run
```bash
# Download this project
go get github.com/ringaile/eth-caching-proxy
```

Before running API server, you can configure the application settings [config.toml](https://github.com/ringaile/eth-caching-proxy/blob/main/config/config.toml)
```toml
	Port = ":9000"
	Timeout = 5
	DefaultExpiration = 5
	CleanupExpiration = 5
	EthClientUrl = "https://cloudflare-eth.com"
```

```bash
# Build and Run
cd eth-caching-proxy
go build
./eth-caching-proxy

# API Endpoint : http://localhost:9000/
```

# Makefile

```bash
# Build and Run
make all

# API Endpoint : http://localhost:9000/
```

```bash
# Build Docker image
make docker_build
```

```bash
# Run Docker image
make docker_run

# API Endpoint : http://localhost:9000/
```


## Structure
```
├── server
│   ├── server.go
│   ├── handlers         	// The API core handlers
│   └── helpers          	// The API core helpers
│── model
│   └── model.go     	 	// Models for the application
│── ethclient
│   └── ethclient.go     	// Ethereum client
│── ethcache
│   └── ethcache.go      	// Cache
│── controller
│   └── blockcontroller.go  // Application logic
├── config
│   └── config.go        	// Configuration
└── main.go
```

## API

#### /block/:block
* `GET` : Get a block [has either "latest" or a number as a param.]

#### /block/:block/txs/:txs
* `GET` : Get a transaction of a block [has either index or tx hash as a param.]