package api

import (
	"log"

	"github.com/gin-gonic/gin"
	soldiershandler "github.com/ruziba3vich/armiya-gateway/api/handlers/soldiers_handler"
	"github.com/ruziba3vich/armiya-gateway/config"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type API struct {
	logger          *log.Logger
	cfg             *config.Config
	soldiershandler *soldiershandler.SoldiersHandler
}

func New(
	cfg *config.Config,
	logger *log.Logger,
	soldiershandler *soldiershandler.SoldiersHandler) *API {
	return &API{
		logger:          logger,
		cfg:             cfg,
		soldiershandler: soldiershandler,
	}
}

// @title Soldiers API
// @version 1.0
// @description This is a sample server for managing soldiers.

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8080
// @BasePath /api/v1
func (a *API) RUN() error {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/v1")
	{
		api.POST("/soldiers", a.soldiersHandler.CreateSoldierHandler)
		api.PUT("/soldiers", a.soldiersHandler.UpdateSoldier)
		api.GET("/soldiers/:soldier_id", a.soldiersHandler.GetSoldierByIdHandler)
		api.GET("/soldiers/name/:soldier_name", a.soldiersHandler.GetSoldiersByNameHandler)
		api.GET("/soldiers/surname/:soldier_surname", a.soldiersHandler.GetSoldiersBySurnameHandler)
		api.GET("/soldiers/group/:group_name", a.soldiersHandler.GetSoldiersByGroupNameHandler)
		api.GET("/soldiers", a.soldiersHandler.GetAllSoldiersHandler)
		api.GET("/soldiers/age/:age", a.soldiersHandler.GetSoldiersByAgeHandler)
		api.DELETE("/soldiers/:soldier_id", a.soldiersHandler.DeleteSoldierHandler)
		api.POST("/soldiers/move/:soldier_id", a.soldiersHandler.MoveSoldierFromGroupAToGroupBHandler)
	}

	return router.Run(a.cfg.ServerAddress)
}
