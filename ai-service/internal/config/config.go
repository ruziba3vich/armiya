package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	API_KEY string
	Port    string
}

func (c *Config) Load() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	c.API_KEY = os.Getenv("API_KEY")
	c.Port = ":" + os.Getenv("SERVER_PORT")

	return nil
}

func New() (*Config, error) {
	config := &Config{}
	err := config.Load()
	if err != nil {
		return nil, err
	}
	return config, nil
}
