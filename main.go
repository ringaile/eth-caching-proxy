package main

import (
	"log"
	"net/http"
	"os"
	"rest-api/ethclient"
	"rest-api/server"
)

func init() {

}

func main() {

	ethClient := ethclient.NewCloudflareEthGateway("https://cloudflare-eth.com", "10")

	s := server.New(&ethClient)

	http.HandleFunc("/", s.Router.ServeHTTP)

	err := http.ListenAndServe(":9000", nil)
	check(err)

	// config := config.GetConfig()
	// ethClient := ethClient.NewCloudflareEthGateway("https://cloudflare-eth.com", "5")

	// server := &server.Server{}
	// server.Initialize(config)
	// server.Run(":3000")
}

func check(e error) {
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}
}
