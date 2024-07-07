package main

import (
	"log"

	"github.com/hackathon/army/fuel-service/api"
	"github.com/hackathon/army/fuel-service/internal/config"
	"github.com/hackathon/army/fuel-service/internal/service"
	"github.com/hackathon/army/fuel-service/internal/storage"
)

func main() {
	configs, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	storage, err := storage.New(configs)
	if err != nil {
		log.Fatal(err)
	}

	api := api.New(service.New(*storage))

	log.Fatal(api.RUN(configs))
}
