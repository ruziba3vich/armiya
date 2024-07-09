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

// CreateFuelHandler godoc
// @Summary Create a new fuel
// @Description This endpoint creates a new fuel
// @Accept json
// @Produce json
// @Param request body genprotos.CreateFuelRequest true "Create Fuel Request"
// @Success 200 {object} genprotos.CreateFuelResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /fuels [post]
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

// GetFuelHandler godoc
// @Summary Get a fuel by ID
// @Description This endpoint retrieves a fuel by its ID
// @Accept json
// @Produce json
// @Param id path string true "Fuel ID"
// @Success 200 {object} genprotos.GetFuelResponse
// @Failure 500 {object} gin.H
// @Router /fuels/{id} [get]
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

// UpdateFuelHandler godoc
// @Summary Update a fuel
// @Description This endpoint updates an existing fuel
// @Accept json
// @Produce json
// @Param id path string true "Fuel ID"
// @Param request body genprotos.UpdateFuelRequest true "Update Fuel Request"
// @Success 200 {object} genprotos.UpdateFuelResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /fuels/{id} [put]
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

// DeleteFuelHandler godoc
// @Summary Delete a fuel
// @Description This endpoint deletes a fuel by its ID
// @Accept json
// @Produce json
// @Param id path string true "Fuel ID"
// @Success 200 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /fuels/{id} [delete]
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

// GetFuelByChoiceHandler godoc
// @Summary Get fuels by choice
// @Description This endpoint retrieves fuels based on specified choices
// @Accept json
// @Produce json
// @Param request body genprotos.GetFuelsByChoiceRequest true "Get Fuels By Choice Request"
// @Success 200 {object} genprotos.GetFuelsByChoiceResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /fuels/choice [post]
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

// GetFuelsHandler godoc
// @Summary Get all fuels
// @Description This endpoint retrieves all fuels
// @Accept json
// @Produce json
// @Success 200 {object} genprotos.GetFuelsResponse
// @Failure 500 {object} gin.H
// @Router /fuels [get]
func (f *FuelsHandlers) GetFuelsHandler(ctx *gin.Context) {
    f.logger.Println("GetFuels")

    resp, err := f.client.GetFuels(ctx, &genprotos.Empty{})
    if err != nil {
        ctx.IndentedJSON(500, gin.H{"error": err.Error()})
        return
    }

    ctx.IndentedJSON(200, resp)
}

// CreateFuelHistoryHandler godoc
// @Summary Create a new fuel history
// @Description This endpoint creates a new fuel history
// @Accept json
// @Produce json
// @Param request body genprotos.CreateFuelHistoryRequest true "Create Fuel History Request"
// @Success 200 {object} genprotos.CreateFuelHistoryResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /fuel-histories [post]
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

// GetFuelHistoriesByIDHandler godoc
// @Summary Get fuel histories by ID
// @Description This endpoint retrieves fuel histories by fuel ID
// @Accept json
// @Produce json
// @Param id path string true "Fuel ID"
// @Success 200 {object} genprotos.GetFuelHistoriesByIdResponse
// @Failure 500 {object} gin.H
// @Router /fuel-histories/{id} [get]
func (f *FuelsHandlers) GetFuelHistoriesByIDHandler(ctx *gin.Context) {
    f.logger.Println("GetFuelHistoryByID")
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

// DeleteFuelHistoriesByIDHandler godoc
// @Summary Delete fuel histories by ID
// @Description This endpoint deletes fuel histories by fuel ID
// @Accept json
// @Produce json
// @Param id path string true "Fuel ID"
// @Success 200 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /fuel-histories/{id} [delete]
func (f *FuelsHandlers) DeleteFuelHistoriesByIDHandler(ctx *gin.Context) {
    f.logger.Println("DeleteFuelHistoryByID")
    id := ctx.Param("id")

    var req = genprotos.DeleteFuelHistoriesByIdRequest{
        Id: id,
    }

    _, err := f.client.DeleteFuelHistoriesByID(ctx, &req)
    if err != nil {
        ctx.IndentedJSON(500, gin.H{"error": err.Error()})
        return
    }

    ctx.IndentedJSON(200, gin.H{"message": "Fuel history deleted"})
}

// GetFuelHistoriesByChoiceHandler godoc
// @Summary Get fuel histories by choice
// @Description This endpoint retrieves fuel histories based on specified choices
// @Accept json
// @Produce json
// @Param request body genprotos.GetFuelHistoriesByChoiceRequest true "Get Fuel Histories By Choice Request"
// @Success 200 {object} genprotos.GetFuelHistoriesByChoiceResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /fuel-histories/choice [post]
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

// GetFuelHistoriesByDateHandler godoc
// @Summary Get fuel histories by date
// @Description This endpoint retrieves fuel histories based on specified date
// @Accept json
// @Produce json
// @Param request body genprotos.GetFuelHistoriesByDateRequest true "Get Fuel Histories By Date Request"
// @Success 200 {object} genprotos.GetFuelHistoriesByDateResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /fuel-histories/date [post]
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

// GetFuelHistoriesHandler godoc
// @Summary Get all fuel histories
// @Description This endpoint retrieves all fuel histories
// @Accept json
// @Produce json
// @Success 200 {object} genprotos.GetFuelHistoriesResponse
// @Failure 500 {object} gin.H
// @Router /fuel-histories [get]
func (f *FuelsHandlers) GetFuelHistoriesHandler(ctx *gin.Context) {
    f.logger.Println("GetFuelHistories")

    resp, err := f.client.GetFuelHistories(ctx, &genprotos.Empty{})
    if err != nil {
        ctx.IndentedJSON(500, gin.H{"error": err.Error()})
        return
    }

    ctx.IndentedJSON(200, resp)
}
