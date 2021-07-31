package main

import (
	"rest-api/config"
	"rest-api/server"
)

func main() {
	config := config.GetConfig()

	server := &server.Server{}
	server.Initialize(config)
	server.Run(":3000")
}
