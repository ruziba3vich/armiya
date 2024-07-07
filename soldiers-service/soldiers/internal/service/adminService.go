package service

import (
	"context"
	"log"
	"os"

	"github.com/ruziba3vich/armiya/soldies-service/genprotos"
	"github.com/ruziba3vich/armiya/soldies-service/internal/storage"
)

type AdminService struct {
	logger  *log.Logger
	storage *storage.AdminStorage
	genprotos.UnimplementedAdminServiceServer
}

func New(storage *storage.AdminStorage) *AdminService {
	return &AdminService{
		logger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		storage: storage,
	}
}

func (a *AdminService) CreateTrainer(ctx context.Context, req *genprotos.CreateTrainerRequest) (*genprotos.Trainer, error) {
	a.logger.Println("SERVER RECEIVED A REQUEST FOR CREATE TRAINER SERVICE")
	return a.storage.CreateTrainerStorage(ctx, req)
}
func (a *AdminService) DeleteTrainer(ctx context.Context, req *genprotos.DeleteTrainerRequest) (*genprotos.Trainer, error) {
	a.logger.Println("SERVER RECEIVED A REQUEST FOR DELETE TRAINER SERVICE")
	return a.storage.DeleteTrainerStorage(ctx, req)
}
