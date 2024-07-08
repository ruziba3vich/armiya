package main

import (
	"armiya/ai-service/api"
	"armiya/ai-service/internal/config"
	"armiya/ai-service/internal/gemini"
	"armiya/ai-service/internal/service"
	"log"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	configs, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	gemini, err := gemini.New(configs)
	if err != nil {
		log.Fatal(err)
	}

	api := api.New(service.New(*gemini))

	log.Fatal(api.RUN(configs))
}

// func migrateDB() {
// 	databaseURL, ok := os.LookupEnv("PG_URL")
// 	if !ok || len(databaseURL) == 0 {
// 		log.Fatalf("migrate: environment variable PG_URL not declared or empty")
// 	}

// 	m, err := migrate.New("file://migrations", databaseURL)
// 	if err != nil {
// 		log.Fatalf("Migrate: error creating migrate instance: %v", err)
// 	}

// 	err = m.Up()
// 	if err != nil {
// 		if err.Error() == "no change" {
// 			log.Printf("Migrate: no change")
// 		} else {
// 			log.Fatalf("Migrate: error applying migrations: %v", err)
// 		}
// 	} else {
// 		pp.Printf("Migrate: up success")
// 	}
// }
