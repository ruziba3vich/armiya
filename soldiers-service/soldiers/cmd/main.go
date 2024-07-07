package main

import (
	"log"

	"github.com/ruziba3vich/armiya/soldies-service/api"
	"github.com/ruziba3vich/armiya/soldies-service/internal/config"
	"github.com/ruziba3vich/armiya/soldies-service/internal/service"
	"github.com/ruziba3vich/armiya/soldies-service/internal/storage"
)

func main() {
	var config config.Config

	if err := config.Load(); err != nil {
		log.Fatal(err)
	}
	db, err := storage.ConnectDB(config)
	if err != nil {
		log.Fatal(err)
	}
	api := api.New(
		service.New(
			storage.New(db),
		),
	)
	log.Fatal(api.RUN(&config))
}
