package service

import (
	"armiya/ai-service/genprotos"
	"armiya/ai-service/internal/gemini"
	"context"
	"log"
	"os"
)

type (
	AIService struct {
		genprotos.UnimplementedAIServiceServer
		aiService gemini.AI
		logger    *log.Logger
	}
)

func New(service gemini.AI) *AIService {
	return &AIService{
		aiService: service,
		logger:    log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (s *AIService) GetEquipmentInfo(ctx context.Context, req *genprotos.EquipmentRequestAI) (*genprotos.EquipmentAI, error) {
	s.logger.Println("Get Equipment Info")
	return s.aiService.GetEquipmentInfo(ctx, req)
}

func (s *AIService) AssessThreat(ctx context.Context, req *genprotos.ThreatData) (*genprotos.ThreatAssessmentResponse, error) {
	s.logger.Println("Assess Threat request")
	return s.aiService.AssessThreat(ctx, req)
}

func (s *AIService) PredictEquipmentMaintenance(ctx context.Context, req *genprotos.EquipmentData) (*genprotos.EquipmentMaintenanceResponse, error) {
	s.logger.Println("Predict Equipment  Maintenance")
	return s.aiService.EquipmentMaintenance(ctx, req)
}
