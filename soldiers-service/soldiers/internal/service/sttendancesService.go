package service

import (
	"context"
	"log"

	"github.com/ruziba3vich/armiya/soldies-service/genprotos"
	"github.com/ruziba3vich/armiya/soldies-service/internal/storage"
)

type AttendanceService struct {
	logger  *log.Logger
	storage *storage.AttendanceStorage
	genprotos.UnimplementedAttendanceServiceServer
}

func NewAttendanceService(storage *storage.AttendanceStorage, logger *log.Logger) *AttendanceService {
	return &AttendanceService{
		logger:  logger,
		storage: storage,
	}
}

func (s *AttendanceService) CreateAttendance(ctx context.Context, req *genprotos.CreateAttendanceRequest) (*genprotos.Attendance, error) {
	s.logger.Println("SERVER RECEIVED A REQUEST TO CreateAttendance SERVICE")
	return s.storage.CreateAttendance(ctx, req)
}

func (s *AttendanceService) AddAmmosToSoldier(ctx context.Context, req *genprotos.AddAmmosRequest) (*genprotos.Attendance, error) {
	s.logger.Println("SERVER RECEIVED A REQUEST TO AddAmmosToSoldier SERVICE")
	return s.storage.AddAmmosToSoldier(ctx, req)
}

func (s *AttendanceService) GetAttendanceByDate(ctx context.Context, req *genprotos.GetAttendanceByDateRequest) (*genprotos.GetAttendanceResponse, error) {
	s.logger.Println("SERVER RECEIVED A REQUEST TO GetAttendanceByDate SERVICE")
	return s.storage.GetAttendanceByDate(ctx, req)
}

func (s *AttendanceService) GetAllAttendanceBySoldierId(ctx context.Context, req *genprotos.GetAllAttendanceBySoldierIdRequest) (*genprotos.GetAttendanceResponse, error) {
	s.logger.Println("SERVER RECEIVED A REQUEST TO GetAllAttendanceBySoldierId SERVICE")
	return s.storage.GetAllAttendanceBySoldierId(ctx, req)
}

func (s *AttendanceService) GetSoldierAttendanceByDate(ctx context.Context, req *genprotos.GetSoldierAttendanceByDateRequest) (*genprotos.GetAttendanceResponse, error) {
	s.logger.Println("SERVER RECEIVED A REQUEST TO GetSoldierAttendanceByDate SERVICE")
	return s.storage.GetSoldierAttendanceByDate(ctx, req)
}

func (s *AttendanceService) UpdateAttendanceBySoldierId(ctx context.Context, req *genprotos.UpdateRequest) (*genprotos.Attendance, error) {
	s.logger.Println("SERVER RECEIVED A REQUEST TO UpdateAttendanceBySoldierId SERVICE")
	return s.storage.UpdateAttendanceBySoldierId(ctx, req)
}

func (s *AttendanceService) DeleteAttendance(ctx context.Context, req *genprotos.DeleteAttendanceRequest) (*genprotos.Attendance, error) {
	s.logger.Println("SERVER RECEIVED A REQUEST TO DeleteAttendance SERVICE")
	return s.storage.DeleteAttendance(ctx, req)
}
