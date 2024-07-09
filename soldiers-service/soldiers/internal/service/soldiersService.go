package service

import (
	"context"
	"log"

	"github.com/ruziba3vich/armiya/soldies-service/genprotos"
	"github.com/ruziba3vich/armiya/soldies-service/internal/storage"
)

type SoldiersService struct {
	logger  *log.Logger
	storage *storage.SoldiersStorage
	genprotos.UnimplementedSoldierServiceServer
}

func NewSoldiersService(storage *storage.SoldiersStorage, logger *log.Logger) *SoldiersService {
	return &SoldiersService{
		logger:  logger,
		storage: storage,
	}
}

func (s *SoldiersService) CreateSoldier(ctx context.Context, req *genprotos.CreateSoldierRequest) (*genprotos.Soldier, error) {
	s.logger.Println("SERVER RECEIVED A REQUEST TO CreateSoldier SERVICE")
	return s.storage.CreateSoldier(ctx, req)
}

func (s *SoldiersService) UpdateSoldier(ctx context.Context, req *genprotos.UpdateSoldierRequest) (*genprotos.UpdateOrGetSoldierResponse, error) {
	s.logger.Println("SERVER RECEIVED A REQUEST TO UpdateSoldier SERVICE")
	return s.storage.UpdateSoldier(ctx, req)
}

func (s *SoldiersService) GetSoldierById(ctx context.Context, req *genprotos.GetByIdRequest) (*genprotos.UpdateOrGetSoldierResponse, error) {
	s.logger.Println("SERVER RECEIVED A REQUEST TO GetSoldierById SERVICE")
	return s.storage.GetSoldierById(ctx, req)
}

func (s *SoldiersService) GetSoldiersByName(ctx context.Context, req *genprotos.GetByName) (*genprotos.GetSoldiersResponse, error) {
	s.logger.Println("SERVER RECEIVED A REQUEST TO GetSoldiersByName SERVICE")
	return s.storage.GetSoldiersByName(ctx, req)
}

func (s *SoldiersService) GetSoldiersBySurname(ctx context.Context, req *genprotos.GetBySurname) (*genprotos.GetSoldiersResponse, error) {
	s.logger.Println("SERVER RECEIVED A REQUEST TO GetSoldiersBySurname SERVICE")
	return s.storage.GetSoldiersBySurname(ctx, req)
}

func (s *SoldiersService) GetSoldiersByGroupName(ctx context.Context, req *genprotos.GetByGroupName) (*genprotos.GetSoldiersResponse, error) {
	s.logger.Println("SERVER RECEIVED A REQUEST TO GetSoldiersByGroupName SERVICE")
	return s.storage.GetSoldiersByGroupName(ctx, req)
}

func (s *SoldiersService) GetAllSoldiers(ctx context.Context, req *genprotos.GetAllSoldiersRequest) (*genprotos.GetSoldiersResponse, error) {
	s.logger.Println("SERVER RECEIVED A REQUEST TO GetAllSoldiers SERVICE")
	return s.storage.GetAllSoldiers(ctx, req)
}

func (s *SoldiersService) GetSoldiersByAge(ctx context.Context, req *genprotos.GetByAgeRequest) (*genprotos.GetSoldiersResponse, error) {
	s.logger.Println("SERVER RECEIVED A REQUEST TO GetSoldiersByAge SERVICE")
	return s.storage.GetSoldiersByAge(ctx, req)
}

func (s *SoldiersService) DeleteSoldier(ctx context.Context, req *genprotos.DeleteRequest) (*genprotos.DeleteSoldierResponse, error) {
	s.logger.Println("SERVER RECEIVED A REQUEST TO DeleteSoldier SERVICE")
	return s.storage.DeleteSoldier(ctx, req)
}

func (s *SoldiersService) MoveSoldierFromGroupAToGroupB(ctx context.Context, req *genprotos.MoveSoldierRequest) (*genprotos.MoveSoldierResponse, error) {
	s.logger.Println("SERVER RECEIVED A REQUEST TO MoveSoldierFromGroupAToGroupB SERVICE")
	return s.storage.MoveSoldierFromGroupAToGroupB(ctx, req)
}

/*
	type SoldierServiceServer interface {
	CreateSoldier(context.Context, *CreateSoldierRequest) (*Soldier, error)
	UpdateSoldier(context.Context, *UpdateSoldierRequest) (*UpdateOrGetSoldierResponse, error)
	GetSoldierById(context.Context, *GetByIdRequest) (*UpdateOrGetSoldierResponse, error)
	GetSoldiersByName(context.Context, *GetByName) (*GetSoldiersResponse, error)
	GetSoldiersBySurname(context.Context, *GetBySurname) (*GetSoldiersResponse, error)
	GetSoldiersByGroupName(context.Context, *GetByGroupName) (*GetSoldiersResponse, error)
	GetAllSoldiers(context.Context, *GetAllSoldiersRequest) (*GetSoldiersResponse, error)
	GetSoldiersByAge(context.Context, *GetByAgeRequest) (*GetSoldiersResponse, error)
	DeleteSoldier(context.Context, *DeleteRequest) (*DeleteSoldierResponse, error)
	MoveSoldierFromGroupAToGroupB(context.Context, *MoveSoldierRequest) (*MoveSoldierResponse, error)
	mustEmbedUnimplementedSoldierServiceServer()
}

*/
