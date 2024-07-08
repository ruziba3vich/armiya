package service

import (
	"context"
	"log"
	"os"

	"github.com/hackathon/army/ammos-service/genprotos"
	"github.com/hackathon/army/ammos-service/internal/storage"
)

type (
	AmmosServiceSt struct {
		genprotos.UnimplementedAmmosServiceServer
		service storage.AmmosSt
		logger  *log.Logger
	}
)

func New(service storage.AmmosSt) *AmmosServiceSt {
	return &AmmosServiceSt{
		service: service,
		logger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// ammos

func (s *AmmosServiceSt) CreateAmmo(ctx context.Context, req *genprotos.CreateAmmoRequest) (*genprotos.AmmoResponse, error) {
	s.logger.Println("create ammo request")
	return s.service.CreateAmmo(ctx, req)
}

func (s *AmmosServiceSt) GetAmmoByChoice(ctx context.Context, req *genprotos.GetAmmoByChoiceRequest) (*genprotos.GetAmmoByChoiceResponse, error) {
	s.logger.Println("get ammo by choice request")
	return s.service.GetAmmoByChoice(ctx, req)
}

func (s *AmmosServiceSt) UpdateAmmoById(ctx context.Context, req *genprotos.UpdateAmmoByIdRequest) (*genprotos.AmmoResponse, error) {
	s.logger.Println("update ammo by id request")
	return s.service.UpdateAmmoById(ctx, req)
}

func (s *AmmosServiceSt) DeleteAmmoById(ctx context.Context, req *genprotos.DeleteAmmoByIdRequest) (*genprotos.Empty, error) {
	s.logger.Println("delete ammo by id request")
	return s.service.DeleteAmmoById(ctx, req)
}

func (s *AmmosServiceSt) GetAmmo(ctx context.Context, req *genprotos.Empty) (*genprotos.GetAmmoResponse, error) {
	s.logger.Println("get ammo by id request")
	return s.service.GetAmmo(ctx, req)
}

// ammo_history

func (s *AmmosServiceSt) CreateAmmoHistory(ctx context.Context, req *genprotos.CreateAmmoHistoryRequest) (*genprotos.AmmoHistory, error) {
	s.logger.Println("create ammo history request")
	return s.service.CreateAmmoHistory(ctx, req)
}

func (s *AmmosServiceSt) GetAmmoHistoryByChoice(ctx context.Context, req *genprotos.GetAmmoHistoryByChoiceRequest) (*genprotos.GetAmmoHistoryByChoiceResponse, error) {
	s.logger.Println("get ammo history by choice request")
	return s.service.GetAmmoHistoryByChoice(ctx, req)
}

func (s *AmmosServiceSt) GetAmmoHistoryById(ctx context.Context, req *genprotos.GetAmmoHistoryByIdRequest) (*genprotos.AmmoHistory, error) {
	s.logger.Println("get ammo by id request")
	return s.service.GetAmmoHistoryById(ctx, req)
}

func (s *AmmosServiceSt) GetAmmoHistoryByDate(ctx context.Context, req *genprotos.GetAmmoHistoryByDateRequest) (*genprotos.GetAmmoHistoryByDateResponse, error) {
	s.logger.Println("get ammo history by date request")
	return s.service.GetAmmoHistoryByDate(ctx, req)
}

func (s *AmmosServiceSt) UpdateAmmoHistoryById(ctx context.Context, req *genprotos.UpdateAmmoHistoryByIdRequest) (*genprotos.AmmoHistory, error) {
	s.logger.Println("update ammo history by id request")
	return s.service.UpdateAmmoHistoryById(ctx, req)
}

func (s *AmmosServiceSt) DeleteAmmoHistoryById(ctx context.Context, req *genprotos.DeleteAmmoHistoryByIdRequest) (*genprotos.Empty, error) {
	s.logger.Println("delete ammo history by id request")
	return s.service.DeleteAmmoHistoryById(ctx, req)
}

func (s *AmmosServiceSt) GetAmmoHistory(ctx context.Context, req *genprotos.Empty) (*genprotos.GetAmmoHistoryResponse, error) {
	s.logger.Println("get ammo by id request")
	return s.service.GetAmmoHistory(ctx, req)
}