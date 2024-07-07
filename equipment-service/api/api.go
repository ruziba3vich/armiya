package api

import (
	"log"
	"net"

	"armiya/equipment-service/genprotos"
	"armiya/equipment-service/internal/config"
	"google.golang.org/grpc"
)

type (
	API struct {
		service genprotos.EquipmentServiceServer
	}
)

func New(service genprotos.EquipmentServiceServer) *API {
	return &API{
		service: service,
	}
}

func (a *API) RUN(config *config.Config) error {
	listener, err := net.Listen("tcp", config.Server.Port)
	if err != nil {
		return err
	}

	serverRegisterer := grpc.NewServer()
	genprotos.RegisterEquipmentServiceServer(serverRegisterer, a.service)

	log.Println("server has started running on port", config.Server.Port)

	return serverRegisterer.Serve(listener)
}
