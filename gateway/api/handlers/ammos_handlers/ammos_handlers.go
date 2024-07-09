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

func (a *AmmosHandlers) CreateAmmoHandler(ctx *gin.Context) {
	a.logger.Println("CreateAmmos")

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

func (a *AmmosHandlers) GetAmmoHandler(ctx *gin.Context) {
	a.logger.Println("GetAmmo")

	resp, err := a.client.GetAmmo(ctx, &genprotos.Empty{})
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

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

func (a *AmmosHandlers) GetAmmoHistoryHandler(ctx *gin.Context) {
	a.logger.Println("GetAmmoHistory")

	resp, err := a.client.GetAmmoHistory(ctx, &genprotos.Empty{})
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}
