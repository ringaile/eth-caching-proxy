package main

import (
	"log"
	"net/http"
	"os"
	"rest-api/config"
	"rest-api/controller"
	"rest-api/ethcache"
	"rest-api/ethclient"
	"rest-api/server"
)

func main() {

	// Get config
	conf, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize components
	ethClient := ethclient.NewCloudflareEthGateway(conf.EthClientUrl, conf.Timeout)
	ethCache := ethcache.New(conf.DefaultExpiration, conf.CleanupExpiration)

	blockController := controller.NewBlockController(&ethClient, &ethCache)

	// Start server
	s := server.New(blockController)

	http.HandleFunc("/", s.Router.ServeHTTP)

	err = http.ListenAndServe(conf.Port, nil)
	check(err)

}

func check(e error) {
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}
}
