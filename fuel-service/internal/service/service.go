package service

import (
	"context"
	"log"
	"os"

	"github.com/hackathon/army/fuel-service/genprotos"
	"github.com/hackathon/army/fuel-service/internal/storage"
)

type (
	FuelServiceSt struct {
		genprotos.UnimplementedFuelServiceServer
		service storage.FuelSt
		logger  *log.Logger
	}
)

func New(service storage.FuelSt) *FuelServiceSt {
	return &FuelServiceSt{
		service: service,
		logger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// fuel_management
func (s *FuelServiceSt) CreateFuel(ctx context.Context, req *genprotos.CreateFuelRequest) (*genprotos.FuelResponse, error) {
	s.logger.Println("get fuel request")
	return s.service.CreateFuel(ctx, req)
}

func (s *FuelServiceSt) GetFuel(ctx context.Context, req *genprotos.GetFuelRequest) (*genprotos.FuelResponse, error) {
	s.logger.Println("get fuel request")
	return s.service.GetFuel(ctx, req)
}

func (s *FuelServiceSt) UpdateFuel(ctx context.Context, req *genprotos.UpdateFuelRequest) (*genprotos.FuelResponse, error) {
	s.logger.Println("update fuel request")
	return s.service.UpdateFuel(ctx, req)
}

func (s *FuelServiceSt) DeleteFuel(ctx context.Context, req *genprotos.DeleteFuelRequest) (*genprotos.Empty, error) {
	s.logger.Println("delete fuel request")
	return s.service.DeleteFuel(ctx, req)
}

func (s *FuelServiceSt) ListFuelByChoice(ctx context.Context, req *genprotos.ListFuelsByChoiceRequest) (*genprotos.ListFuelsByChoiceResponse, error) {
	s.logger.Println("list fuel by name request")
	return s.service.ListFuelByChoice(ctx, req)
}

func (s *FuelServiceSt) ListFuels(ctx context.Context, req *genprotos.Empty) (*genprotos.ListFuelsResponse, error) {
	s.logger.Println("list fuels request")
	return s.service.ListFuels(ctx, req)
}

// fuel_history

func (s *FuelServiceSt) CreateFuelHistory(ctx context.Context, req *genprotos.CreateFuelHistoryRequest) (*genprotos.FuelHistoryResponse, error) {
	s.logger.Println("create fuel history request")
	return s.service.CreateFuelHistory(ctx, req)
}

func (s *FuelServiceSt) ListFuelHistoriesByFuelID(ctx context.Context, req *genprotos.ListFuelHistoriesByFuelIDRequest) (*genprotos.ListFuelHistoriesByFuelIDResponse, error) {
	s.logger.Println("update fuel history request")
	return s.service.ListFuelHistoriesByFuelID(ctx, req)
}

func (s *FuelServiceSt) ListFuelHistoriesByChoice(ctx context.Context, req *genprotos.ListFuelHistoriesByChoiceRequest) (*genprotos.ListFuelHistoriesByChoiceResponse, error) {
	s.logger.Println("list fuel histories by name request")
	return s.service.ListFuelHistoriesByChoice(ctx, req)
}

func (s *FuelServiceSt) ListFuelHistoriesByDate(ctx context.Context, req *genprotos.ListFuelHistoriesByDateRequest) (*genprotos.ListFuelHistoriesByDateResponse, error) {
	s.logger.Println("list fuel histories by date request")
	return s.service.ListFuelHistoriesByDate(ctx, req)
}

func (s *FuelServiceSt) ListFuelHistories(ctx context.Context, req *genprotos.Empty) (*genprotos.ListFuelHistoriesResponse, error) {
	s.logger.Println("list fuel histories request")
	return s.service.ListFuelHistories(ctx, req)
}
