package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{},
	}
}
