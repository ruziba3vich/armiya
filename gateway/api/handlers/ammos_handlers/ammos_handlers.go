package ammoshandlers

import (
	"log"

	"github.com/gin-gonic/gin"
	genprotos "github.com/ruziba3vich/armiya-gateway/genprotos/ammos_genprotos"
)

type AmmosHandlers struct {
	client genprotos.AmmosServiceClient
	logger *log.Logger
}

// CreateAmmoHandler godoc
// @Summary Create a new ammo
// @Description This endpoint creates a new ammo
// @Accept json
// @Produce json
// @Param request body genprotos.CreateAmmoRequest true "Create Ammo Request"
// @Success 200 {object} genprotos.CreateAmmoResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /ammos [post]
func (a *AmmosHandlers) CreateAmmoHandler(ctx *gin.Context) {
    a.logger.Println("CreateAmmo")

    var req genprotos.CreateAmmoRequest

    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.IndentedJSON(400, gin.H{"error": err.Error()})
        return
    }

    resp, err := a.client.CreateAmmo(ctx, &req)
    if err != nil {
        ctx.IndentedJSON(500, gin.H{"error": err.Error()})
        return
    }

    ctx.IndentedJSON(200, resp)
}

// GetAmmoByChoiceHandler godoc
// @Summary Get ammos by choice
// @Description This endpoint retrieves ammos based on specified choices
// @Accept json
// @Produce json
// @Param request body genprotos.GetAmmoByChoiceRequest true "Get Ammos By Choice Request"
// @Success 200 {object} genprotos.GetAmmoByChoiceResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /ammos/choice [post]
func (a *AmmosHandlers) GetAmmoByChoiceHandler(ctx *gin.Context) {
    a.logger.Println("GetAmmoByChoice")

    var req genprotos.GetAmmoByChoiceRequest

    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.IndentedJSON(400, gin.H{"error": err.Error()})
        return
    }

    resp, err := a.client.GetAmmoByChoice(ctx, &req)
    if err != nil {
        ctx.IndentedJSON(500, gin.H{"error": err.Error()})
        return
    }

    ctx.IndentedJSON(200, resp)
}

// UpdateAmmoByIdHandler godoc
// @Summary Update an ammo by ID
// @Description This endpoint updates an existing ammo by its ID
// @Accept json
// @Produce json
// @Param id path string true "Ammo ID"
// @Param request body genprotos.UpdateAmmoByIdRequest true "Update Ammo By ID Request"
// @Success 200 {object} genprotos.UpdateAmmoByIdResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /ammos/{id} [put]
func (a *AmmosHandlers) UpdateAmmoByIdHandler(ctx *gin.Context) {
    a.logger.Println("UpdateAmmoById")

    id := ctx.Param("id")

    var req = genprotos.UpdateAmmoByIdRequest{
        Id: id,
    }

    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.IndentedJSON(400, gin.H{"error": err.Error()})
        return
    }

    resp, err := a.client.UpdateAmmoById(ctx, &req)
    if err != nil {
        ctx.IndentedJSON(500, gin.H{"error": err.Error()})
        return
    }

    ctx.IndentedJSON(200, resp)
}

// DeleteAmmoByIdHandler godoc
// @Summary Delete an ammo by ID
// @Description This endpoint deletes an ammo by its ID
// @Accept json
// @Produce json
// @Param id path string true "Ammo ID"
// @Success 200 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /ammos/{id} [delete]
func (a *AmmosHandlers) DeleteAmmoByIdHandler(ctx *gin.Context) {
    a.logger.Println("DeleteAmmoById")

    id := ctx.Param("id")

    var req = genprotos.DeleteAmmoByIdRequest{
        Id: id,
    }

    _, err := a.client.DeleteAmmoById(ctx, &req)
    if err != nil {
        ctx.IndentedJSON(500, gin.H{"error": err.Error()})
        return
    }

    ctx.IndentedJSON(200, gin.H{"message": "Ammo deleted"})
}

// GetAmmoHandler godoc
// @Summary Get all ammos
// @Description This endpoint retrieves all ammos
// @Accept json
// @Produce json
// @Success 200 {object} genprotos.GetAmmoResponse
// @Failure 500 {object} gin.H
// @Router /ammos [get]
func (a *AmmosHandlers) GetAmmoHandler(ctx *gin.Context) {
    a.logger.Println("GetAmmo")

    resp, err := a.client.GetAmmo(ctx, &genprotos.Empty{})
    if err != nil {
        ctx.IndentedJSON(500, gin.H{"error": err.Error()})
        return
    }

    ctx.IndentedJSON(200, resp)
}

// CreateAmmoHistoryHandler godoc
// @Summary Create a new ammo history
// @Description This endpoint creates a new ammo history
// @Accept json
// @Produce json
// @Param request body genprotos.CreateAmmoHistoryRequest true "Create Ammo History Request"
// @Success 200 {object} genprotos.CreateAmmoHistoryResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /ammo-histories [post]
func (a *AmmosHandlers) CreateAmmoHistoryHandler(ctx *gin.Context) {
    a.logger.Println("CreateAmmoHistory")

    var req genprotos.CreateAmmoHistoryRequest

    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.IndentedJSON(400, gin.H{"error": err.Error()})
        return
    }

    resp, err := a.client.CreateAmmoHistory(ctx, &req)
    if err != nil {
        ctx.IndentedJSON(500, gin.H{"error": err.Error()})
        return
    }

    ctx.IndentedJSON(200, resp)
}

// GetAmmoHistoryByChoiceHandler godoc
// @Summary Get ammo histories by choice
// @Description This endpoint retrieves ammo histories based on specified choices
// @Accept json
// @Produce json
// @Param request body genprotos.GetAmmoHistoryByChoiceRequest true "Get Ammo Histories By Choice Request"
// @Success 200 {object} genprotos.GetAmmoHistoryByChoiceResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /ammo-histories/choice [post]
func (a *AmmosHandlers) GetAmmoHistoryByChoiceHandler(ctx *gin.Context) {
    a.logger.Println("GetAmmoHistoryByChoice")

    var req genprotos.GetAmmoHistoryByChoiceRequest

    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.IndentedJSON(400, gin.H{"error": err.Error()})
        return
    }

    resp, err := a.client.GetAmmoHistoryByChoice(ctx, &req)
    if err != nil {
        ctx.IndentedJSON(500, gin.H{"error": err.Error()})
        return
    }

    ctx.IndentedJSON(200, resp)
}

// GetAmmoHistoryByIdHandler godoc
// @Summary Get ammo history by ID
// @Description This endpoint retrieves an ammo history by its ID
// @Accept json
// @Produce json
// @Param id path string true "Ammo History ID"
// @Success 200 {object} genprotos.GetAmmoHistoryByIdResponse
// @Failure 500 {object} gin.H
// @Router /ammo-histories/{id} [get]
func (a *AmmosHandlers) GetAmmoHistoryByIdHandler(ctx *gin.Context) {
    a.logger.Println("GetAmmoHistoryById")
    id := ctx.Param("id")

    var req = genprotos.GetAmmoHistoryByIdRequest{
        Id: id,
    }

    resp, err := a.client.GetAmmoHistoryById(ctx, &req)
    if err != nil {
        ctx.IndentedJSON(500, gin.H{"error": err.Error()})
        return
    }

    ctx.IndentedJSON(200, resp)
}

// GetAmmoHistoryByDateHandler godoc
// @Summary Get ammo histories by date
// @Description This endpoint retrieves ammo histories based on specified date
// @Accept json
// @Produce json
// @Param request body genprotos.GetAmmoHistoryByDateRequest true "Get Ammo Histories By Date Request"
// @Success 200 {object} genprotos.GetAmmoHistoryByDateResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /ammo-histories/date [post]
func (a *AmmosHandlers) GetAmmoHistoryByDateHandler(ctx *gin.Context) {
    a.logger.Println("GetAmmoHistoryByDate")

    var req genprotos.GetAmmoHistoryByDateRequest

    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.IndentedJSON(400, gin.H{"error": err.Error()})
        return
    }

    resp, err := a.client.GetAmmoHistoryByDate(ctx, &req)
    if err != nil {
        ctx.IndentedJSON(500, gin.H{"error": err.Error()})
        return
    }

    ctx.IndentedJSON(200, resp)
}

// UpdateAmmoHistoryByIdHandler godoc
// @Summary Update an ammo history by ID
// @Description This endpoint updates an existing ammo history by its ID
// @Accept json
// @Produce json
// @Param id path string true "Ammo History ID"
// @Param request body genprotos.UpdateAmmoHistoryByIdRequest true "Update Ammo History By ID Request"
// @Success 200 {object} genprotos.UpdateAmmoHistoryByIdResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /ammo-histories/{id} [put]
func (a *AmmosHandlers) UpdateAmmoHistoryByIdHandler(ctx *gin.Context) {
    a.logger.Println("UpdateAmmoHistoryById")

    id := ctx.Param("id")

    var req = genprotos.UpdateAmmoHistoryByIdRequest{
        Id: id,
    }

    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.IndentedJSON(400, gin.H{"error": err.Error()})
        return
    }

    resp, err := a.client.UpdateAmmoHistoryById(ctx, &req)
    if err != nil {
        ctx.IndentedJSON(500, gin.H{"error": err.Error()})
        return
    }

    ctx.IndentedJSON(200, resp)
}

// DeleteAmmoHistoryByIdHandler godoc
// @Summary Delete an ammo history by ID
// @Description This endpoint deletes an ammo history by its ID
// @Accept json
// @Produce json
// @Param id path string true "Ammo History ID"
// @Success 200 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /ammo-histories/{id} [delete]
func (a *AmmosHandlers) DeleteAmmoHistoryByIdHandler(ctx *gin.Context) {
    a.logger.Println("DeleteAmmoHistoryById")

    id := ctx.Param("id")

    var req = genprotos.DeleteAmmoHistoryByIdRequest{
        Id: id,
    }

    _, err := a.client.DeleteAmmoHistoryById(ctx, &req)
    if err != nil {
        ctx.IndentedJSON(500, gin.H{"error": err.Error()})
        return
    }

    ctx.IndentedJSON(200, gin.H{"message": "AmmoHistory deleted"})
}

// GetAmmoHistoryHandler godoc
// @Summary Get all ammo histories
// @Description This endpoint retrieves all ammo histories
// @Accept json
// @Produce json
// @Success 200 {object} genprotos.GetAmmoHistoryResponse
// @Failure 500 {object} gin.H
// @Router /ammo-histories [get]
func (a *AmmosHandlers) GetAmmoHistoryHandler(ctx *gin.Context) {
    a.logger.Println("GetAmmoHistory")

    resp, err := a.client.GetAmmoHistory(ctx, &genprotos.Empty{})
    if err != nil {
        ctx.IndentedJSON(500, gin.H{"error": err.Error()})
        return
    }

    ctx.IndentedJSON(200, resp)
}
