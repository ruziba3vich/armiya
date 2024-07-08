package api

import (
	"log"
	"net"

	"github.com/hackathon/army/ammos-service/genprotos"
	"github.com/hackathon/army/ammos-service/internal/config"
	"google.golang.org/grpc"
)

type (
	API struct {
		service genprotos.AmmosServiceServer
	}
)

func New(service genprotos.AmmosServiceServer) *API {
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
	genprotos.RegisterAmmosServiceServer(serverRegisterer, a.service)

	log.Println("server has started running on port", config.Server.Port)

	return serverRegisterer.Serve(listener)
}
