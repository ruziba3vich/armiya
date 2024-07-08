package api

import (
	"log"
	"net"

	"armiya/ai-service/genprotos"
	"armiya/ai-service/internal/config"

	"google.golang.org/grpc"
)

type (
	API struct {
		service genprotos.AIServiceServer
	}
)

func New(service genprotos.AIServiceServer) *API {
	return &API{
		service: service,
	}
}

func (a *API) RUN(config *config.Config) error {
	listener, err := net.Listen("tcp", config.Port)
	if err != nil {
		return err
	}

	serverRegisterer := grpc.NewServer()
	genprotos.RegisterAIServiceServer(serverRegisterer, a.service)

	log.Println("server has started running on port", config.Port)

	return serverRegisterer.Serve(listener)
}
