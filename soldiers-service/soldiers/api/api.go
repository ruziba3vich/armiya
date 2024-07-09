package api

import (
	"log"
	"net"

	"github.com/ruziba3vich/armiya/soldies-service/genprotos"
	"github.com/ruziba3vich/armiya/soldies-service/internal/config"
	"google.golang.org/grpc"
)

type API struct {
	adminService   genprotos.AdminServiceServer
	soldierService genprotos.SoldierServiceServer
}

func New(
	adminService genprotos.AdminServiceServer,
	soldierService genprotos.SoldierServiceServer,
) *API {
	return &API{
		adminService:   adminService,
		soldierService: soldierService,
	}
}

func (a *API) RUN(config *config.Config) error {
	listener, err := net.Listen("tcp", config.Server.Port)
	if err != nil {
		return err
	}

	serverRegisterer := grpc.NewServer()
	genprotos.RegisterAdminServiceServer(serverRegisterer, a.adminService)
	genprotos.RegisterSoldierServiceServer(serverRegisterer, a.soldierService)
	log.Println("SERVER HAS STARTED ON PORT", config.Server.Port)
	return serverRegisterer.Serve(listener)
}
