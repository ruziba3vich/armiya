package aihandlers

import (
	"log"

	"github.com/gin-gonic/gin"
	genprotos "github.com/ruziba3vich/armiya-gateway/genprotos/ai_genprotos"
)

// AIHandlers holds the AI service client and logger.
type AIHandlers struct {
	client genprotos.AIServiceClient
	logger *log.Logger
}

// GetEquipmentInfoHandler godoc
// @Summary Get equipment information
// @Description Get detailed information about a specific equipment
// @Tags equipment
// @Accept json
// @Produce json
// @Param request body genprotos.EquipmentRequestAI true "Equipment Request"
// @Success 200 {object} genprotos.EquipmentResponseAI
// @Failure 400 {object} gin.H{"error": "Bad Request"}
// @Failure 500 {object} gin.H{"error": "Internal Server Error"}
// @Router /equipment/info [post]
func (a *AIHandlers) GetEquipmentInfoHandler(ctx *gin.Context) {
	a.logger.Println("Get Equipment Info")

	var req genprotos.EquipmentRequestAI

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := a.client.GetEquipmentInfo(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

// AssessThreat godoc
// @Summary Assess threat level
// @Description Assess the threat based on provided data
// @Tags threat
// @Accept json
// @Produce json
// @Param request body genprotos.ThreatData true "Threat Data"
// @Success 200 {object} genprotos.ThreatAssessment
// @Failure 400 {object} gin.H{"error": "Bad Request"}
// @Failure 500 {object} gin.H{"error": "Internal Server Error"}
// @Router /threat/assess [post]
func (a *AIHandlers) AssessThreat(ctx *gin.Context) {
	a.logger.Println("Assess Threat")

	var req genprotos.ThreatData

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := a.client.AssessThreat(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

// PredictEquipmentMaintenance godoc
// @Summary Predict equipment maintenance
// @Description Predict maintenance schedule for equipment
// @Tags equipment
// @Accept json
// @Produce json
// @Param request body genprotos.EquipmentData true "Equipment Data"
// @Success 200 {object} genprotos.MaintenancePrediction
// @Failure 400 {object} gin.H{"error": "Bad Request"}
// @Failure 500 {object} gin.H{"error": "Internal Server Error"}
// @Router /equipment/maintenance/predict [post]
func (a *AIHandlers) PredictEquipmentMaintenance(ctx *gin.Context) {
	a.logger.Println("Predict Equipment Maintenance")

	var req genprotos.EquipmentData

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := a.client.PredictEquipmentMaintenance(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

// ProvideFirstAidInsturctions godoc
// @Summary Provide first aid instructions
// @Description Provide first aid instructions based on injury details
// @Tags firstaid
// @Accept json
// @Produce json
// @Param request body genprotos.InjuryDetails true "Injury Details"
// @Success 200 {object} genprotos.FirstAidInstructions
// @Failure 400 {object} gin.H{"error": "Bad Request"}
// @Failure 500 {object} gin.H{"error": "Internal Server Error"}
// @Router /firstaid/instructions [post]
func (a *AIHandlers) ProvideFirstAidInsturctions(ctx *gin.Context) {
	a.logger.Println("Provide First Aid Instructions")

	var req genprotos.InjuryDetails

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := a.client.ProvideFirstAidInstructions(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

// FoodRecommendByActivity godoc
// @Summary Recommend food based on activity
// @Description Recommend food intake based on the user's activity level
// @Tags food
// @Accept json
// @Produce json
// @Param request body genprotos.ThreatData true "Activity Data"
// @Success 200 {object} genprotos.FoodRecommendation
// @Failure 400 {object} gin.H{"error": "Bad Request"}
// @Failure 500 {object} gin.H{"error": "Internal Server Error"}
// @Router /food/recommend [post]
func (a *AIHandlers) FoodRecommendByActivity(ctx *gin.Context) {
	a.logger.Println("Food Recommend By Activity")

	var req genprotos.ThreatData

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := a.client.AssessThreat(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}
