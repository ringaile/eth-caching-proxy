package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Port              string
	Timeout           int32
	DefaultExpiration int32
	CleanupExpiration int32
	EthClientUrl      string
}

func GetConfig() (*Config, error) {
	conf := Config{}
	path, _ := os.Getwd()
	log.Print(path)
	if _, err := toml.DecodeFile("./config/config.toml", &conf); err != nil {
		log.Print(err)
		return nil, nil
	}
	return &conf, nil
}
