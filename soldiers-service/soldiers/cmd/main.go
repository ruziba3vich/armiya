package main

import (
	"log"
	"os"

	sq "github.com/Masterminds/squirrel"
	"github.com/ruziba3vich/armiya/soldies-service/api"
	"github.com/ruziba3vich/armiya/soldies-service/internal/config"
	"github.com/ruziba3vich/armiya/soldies-service/internal/service"
	"github.com/ruziba3vich/armiya/soldies-service/internal/storage"

	// swaggerFiles "github.com/swaggo/files"
    // ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Soldiers API
// @version 1.0
// @description This is a sample server for managing soldiers.

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8080
// @BasePath /api/v1
func main() {
	var config config.Config

	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	if err := config.Load(); err != nil {
		logger.Fatal(err)
	}
	db, err := storage.ConnectDB(config)
	if err != nil {
		logger.Fatal(err)
	}

	sqrl := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	api := api.New(
		service.NewAdminsService(
			storage.NewAdminsStorage(db, logger, sqrl),
			logger,
		),
		service.NewSoldiersService(
			storage.NewSoldiersStorage(db, logger, sqrl),
			logger,
		),
		service.NewGroupsService(
			storage.NewGroupsStorage(db, logger, sqrl),
			logger,
		),
		service.NewAttendanceService(
			storage.NewAttendanceStorage(db, logger, sqrl),
			logger,
		),
	)
	logger.Fatal(api.RUN(&config))
}
