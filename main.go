package main

import (
	"log"
	"net/http"
	"os"
	"rest-api/ethclient"
	"rest-api/proxy"
	"rest-api/server"
)

func init() {

}

func main() {

	ethClient := ethclient.NewCloudflareEthGateway("https://cloudflare-eth.com", "10")
	proxy := proxy.NewProxyImpl(&ethClient)

	s := server.New(proxy)

	http.HandleFunc("/", s.Router.ServeHTTP)

	err := http.ListenAndServe(":9000", nil)
	check(err)

}

func check(e error) {
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}
}
