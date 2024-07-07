package service

import (
	"armiya/equipment-service/genprotos"
	"armiya/equipment-service/internal/storage"
	"context"
	"log"
	"os"
)

type (
	EquipmentService struct {
		genprotos.UnimplementedEquipmentServiceServer
		service storage.Equipment
		logger  *log.Logger
	}
)

func New(service storage.Equipment) *EquipmentService {
	return &EquipmentService{
		service: service,
		logger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (s *EquipmentService) CreateEquipment(ctx context.Context, req *genprotos.Equipment) (*genprotos.Equipment, error) {
	s.logger.Println("Create equipment request")
	return s.service.CreateEquipment(ctx, req)
}

func (s *EquipmentService) GetEquipment(ctx context.Context, req *genprotos.GetRequest) (*genprotos.Equipment, error) {
	s.logger.Println("Get equipment requset")
	return s.service.GetEquipment(ctx, req)
}

func (s *EquipmentService) GetAllEquipments(ctx context.Context, req *genprotos.GetAllRequest) (*genprotos.GetAllResponse, error) {
	s.logger.Println("Get All Equipments request")
	return s.service.GetAllEquipments(ctx, req)
}
