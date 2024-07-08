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
	s.logger.Println("Create equipment request")
	return s.aiService.GetEquipmentInfo(ctx, req)
}
