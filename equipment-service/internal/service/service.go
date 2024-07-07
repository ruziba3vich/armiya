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
		equipmentService storage.Equipment
		logger           *log.Logger
	}
)

func New(service storage.Equipment) *EquipmentService {
	return &EquipmentService{
		equipmentService: service,
		logger:           log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (s *EquipmentService) CreateEquipment(ctx context.Context, req *genprotos.Equipment) (*genprotos.Equipment, error) {
	s.logger.Println("Create equipment request")
	return s.equipmentService.CreateEquipment(ctx, req)
}

func (s *EquipmentService) GetEquipment(ctx context.Context, req *genprotos.GetRequest) (*genprotos.Equipment, error) {
	s.logger.Println("Get equipment request")
	return s.equipmentService.GetEquipment(ctx, req)
}

func (s *EquipmentService) GetAllEquipments(ctx context.Context, req *genprotos.GetAllRequest) (*genprotos.GetAllResponse, error) {
	s.logger.Println("Get All Equipments request")
	return s.equipmentService.GetAllEquipments(ctx, req)
}

func (s *EquipmentService) CreateEquipmentHistory(ctx context.Context, req *genprotos.EquipmentHistory) (*genprotos.EquipmentHistory, error) {
	s.logger.Println("Create equipment history request")
	return s.equipmentService.CreateEquipmentHistory(ctx, req)
}

func (s *EquipmentService) GetEquipmentHistory(ctx context.Context, req *genprotos.GetHistoryRequest) (*genprotos.EquipmentHistory, error) {
	s.logger.Println("Get equipment history request")
	return s.equipmentService.GetEquipmentHistory(ctx, req)
}

func (s *EquipmentService) GetAllEquipmentHistories(ctx context.Context, req *genprotos.GetAllHistoryRequest) (*genprotos.GetAllHistoryResponse, error) {
	s.logger.Println("Get All equipment histories request")
	return s.equipmentService.GetAllEquipmentHistories(ctx, req)
}

func (s *EquipmentService) UpdateEquipmentHistory(ctx context.Context, req *genprotos.EquipmentHistory) (*genprotos.EquipmentHistory, error) {
	s.logger.Println("Update equipment history request")
	return s.equipmentService.UpdateEquipmentHistory(ctx, req)
}

func (s *EquipmentService) DeleteEquipmentHistory(ctx context.Context, req *genprotos.GetRequest) (*genprotos.Empty, error) {
	s.logger.Println("Delete equipment history request")
	return &genprotos.Empty{}, s.equipmentService.DeleteEquipmentHistory(ctx, req)
}

func (s *EquipmentService) UpdateEquipment(ctx context.Context, req *genprotos.Equipment) (*genprotos.Equipment, error) {
	s.logger.Println("Update equipment request")
	return s.equipmentService.UpdateEquipment(ctx, req)
}

func (s *EquipmentService) DeleteEquipment(ctx context.Context, req *genprotos.GetRequest) (*genprotos.Empty, error) {
	s.logger.Println("Delete equipment request")
	return &genprotos.Empty{}, s.equipmentService.DeleteEquipment(ctx, req)
}
