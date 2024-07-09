package main

import (
	"log"
	"os"

	"github.com/ruziba3vich/armiya-gateway/api"
	"github.com/ruziba3vich/armiya-gateway/config"
)

func main() {
	var cfg config.Config
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	if err := cfg.Load(); err != nil {
		logger.Fatal(err)
	}
	api := api.New(&cfg, logger)
	logger.Fatal(api.RUN())
}
