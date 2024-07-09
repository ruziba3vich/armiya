package fuleshandlers

import (
	"log"

	"github.com/gin-gonic/gin"
	genprotos "github.com/ruziba3vich/armiya-gateway/genprotos/fuels_genprotos"
)

type FuelsHandlers struct {
	client genprotos.FuelServiceClient
	logger *log.Logger
}

func (f *FuelsHandlers) CreateFuelHandler(ctx *gin.Context) {
	f.logger.Println("CreateFuel")

	var req genprotos.CreateFuelRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := f.client.CreateFuel(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

func (f *FuelsHandlers) GetFuelHandler(ctx *gin.Context) {
	f.logger.Println("GetFuel")

	id := ctx.Param("id")

	var req = genprotos.GetFuelRequest{
		Id: id,
	}

	resp, err := f.client.GetFuel(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

func (f *FuelsHandlers) UpdateFuelHandler(ctx *gin.Context) {
	f.logger.Println("UpdateFuel")

	id := ctx.Param("id")

	var req = genprotos.UpdateFuelRequest{
		Id: id,
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := f.client.UpdateFuel(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

func (f *FuelsHandlers) DeleteFuelHandler(ctx *gin.Context) {
	f.logger.Println("DeleteFuel")

	id := ctx.Param("id")

	var req = genprotos.DeleteFuelRequest{
		Id: id,
	}

	_, err := f.client.DeleteFuel(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, gin.H{"message": "Fuel deleted"})
}

func (f *FuelsHandlers) GetFuelByChoiceHandler(ctx *gin.Context) {
	f.logger.Println("GetFuelByChoice")

	var req genprotos.GetFuelsByChoiceRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := f.client.GetFuelByChoice(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

func (f *FuelsHandlers) GetFuelsHandler(ctx *gin.Context) {
	f.logger.Println("GetFuels")

	resp, err := f.client.GetFuels(ctx, &genprotos.Empty{})
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

func (f *FuelsHandlers) CreateFuelHistoryHandler(ctx *gin.Context) {
	f.logger.Println("CreateFuelHistory")

	var req genprotos.CreateFuelHistoryRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := f.client.CreateFuelHistory(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

func (f *FuelsHandlers) GetFuelHistoriesByIDHandler(ctx *gin.Context) {
	f.logger.Println("GetFuelHistory")
	id := ctx.Param("id")

	var req = genprotos.GetFuelHistoriesByIdRequest{
		Id: id,
	}

	resp, err := f.client.GetFuelHistoriesByID(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

func (f *FuelsHandlers) GetFuelHistoriesByChoiceHandler(ctx *gin.Context) {
	f.logger.Println("GetFuelHistoriesByChoice")

	var req genprotos.GetFuelHistoriesByChoiceRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := f.client.GetFuelHistoriesByChoice(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

func (f *FuelsHandlers) GetFuelHistoriesByDateHandler(ctx *gin.Context) {
	f.logger.Println("GetFuelHistoriesByChoiceAndDate")

	var req genprotos.GetFuelHistoriesByDateRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := f.client.GetFuelHistoriesByDate(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

func(f *FuelsHandlers) GetFuelHistoriesHandler(ctx *gin.Context) {
	f.logger.Println("GetFuelHistories")

	resp, err := f.client.GetFuelHistories(ctx, &genprotos.Empty{})
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}