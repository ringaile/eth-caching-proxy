package main

import (
	"log"
	"net/http"
	"os"
	"rest-api/config"
	"rest-api/controller"
	"rest-api/ethclient"
	"rest-api/proxy"
	"rest-api/server"
)

func main() {

	// Get config
	conf, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize components
	cloud := ethclient.NewCloudflareEthGateway(conf.EthClientUrl, conf.Timeout)
	proxy := proxy.New(conf.DefaultExpiration, conf.CleanupExpiration)

	blockController := controller.NewBlockController(&cloud, &proxy)

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
