package soldiershandler

import (
	"log"
	"net/http"
	"strconv"


	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/armiya-gateway/api/handlers/auth"
	genprotos "github.com/ruziba3vich/armiya-gateway/genprotos/soldiers_genprotos"
)

type SoldiersHandler struct {
	client genprotos.SoldierServiceClient
	logger *log.Logger
	auth   *auth.TokenManager
}

// CreateSoldierHandler creates a new soldier.
// @Summary Create a new soldier
// @Tags Soldiers
// @Accept json
// @Produce json
// @Param soldier body genprotos.CreateSoldierRequest true "Soldier"
// @Success 200 {object} genprotos.CreateSoldierResponse
// @Failure 400 {object} gin.H{"error": string}
// @Failure 401 {object} gin.H{"error": string}
// @Router /soldiers [post]
func (s *SoldiersHandler) CreateSoldierHandler(c *gin.Context) {
	s.logger.Println("REQUEST RECEIVED INTO CreateSoldierHandler")
	var req genprotos.CreateSoldierRequest
	id, err := s.auth.ExtractIDFromToken(c.GetHeader("Authorization")[8:])
	if err != nil {
		s.logger.Println("ERROR WHILE EXTRACTING ID IN CreateSoldierHandler :", err)
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}
	req.CreatedBy = id
	if err := c.ShouldBindJSON(&req); err != nil {
		s.logger.Println("ERROR WHILE BINDING REQUEST IN CreateSoldierHandler :", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	response, err := s.client.CreateSoldier(c, &req)
	if err != nil {
		s.logger.Println("ERROR SENDING REQUEST TO THE SERVER IN CreateSoldierHandler :", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"response": response})
}

// UpdateSoldier updates an existing soldier.
// @Summary Update an existing soldier
// @Tags Soldiers
// @Accept json
// @Produce json
// @Param soldier body genprotos.UpdateSoldierRequest true "Soldier"
// @Success 200 {object} genprotos.UpdateOrGetSoldierResponse
// @Failure 400 {object} gin.H{"error": string}
// @Router /soldiers/{soldier_id} [put]
func (s *SoldiersHandler) UpdateSoldier(c *gin.Context) {
	s.logger.Println("REQUEST RECEIVED INTO UpdateSoldier")
	var req genprotos.UpdateSoldierRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		s.logger.Println("ERROR WHILE BINDING REQUEST IN UpdateSoldier :", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	response, err := s.client.UpdateSoldier(c, &req)
	if err != nil {
		s.logger.Println("ERROR WHILE SENDING REQUEST TO THE SERVER IN UpdateSoldier :", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"response": response})
}

// GetSoldierByIdHandler retrieves a soldier by ID.
// @Summary Get a soldier by ID
// @Tags Soldiers
// @Produce json
// @Param soldier_id path string true "Soldier ID"
// @Success 200 {object} genprotos.UpdateOrGetSoldierResponse
// @Failure 400 {object} gin.H{"error": string}
// @Router /soldiers/{soldier_id} [get]
func (s *SoldiersHandler) GetSoldierByIdHandler(c *gin.Context) {
	s.logger.Println("REQUEST RECEIVED INTO GetSoldierByIdHandler")
	var req genprotos.GetByIdRequest
	req.SoldierId = c.Param("soldier_id")
	response, err := s.client.GetSoldierById(c, &req)
	if err != nil {
		s.logger.Println("ERROR WHILE SENDING REQUEST TO THE SERVER IN GetSoldierByIdHandler :", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"response": response})
}

// GetSoldiersByNameHandler retrieves soldiers by name.
// @Summary Get soldiers by name
// @Tags Soldiers
// @Produce json
// @Param soldier_name path string true "Soldier Name"
// @Success 200 {object} genprotos.GetSoldiersResponse
// @Failure 400 {object} gin.H{"error": string}
// @Router /soldiers/name/{soldier_name} [get]
func (s *SoldiersHandler) GetSoldiersByNameHandler(c *gin.Context) {
	s.logger.Println("REQUEST RECEIVED INTO GetSoldiersByNameHandler")
	var req genprotos.GetByName
	req.SoldierName = c.Param("soldier_name")
	response, err := s.client.GetSoldiersByName(c, &req)
	if err != nil {
		s.logger.Println("ERROR WHILE SENDING REQUEST TO THE SERVER IN GetSoldiersByNameHandler :", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"response": response})
}

// GetSoldiersBySurnameHandler retrieves soldiers by surname.
// @Summary Get soldiers by surname
// @Tags Soldiers
// @Produce json
// @Param soldier_surname path string true "Soldier Surname"
// @Success 200 {object} genprotos.GetSoldiersResponse
// @Failure 400 {object} gin.H{"error": string}
// @Router /soldiers/surname/{soldier_surname} [get]
func (s *SoldiersHandler) GetSoldiersBySurnameHandler(c *gin.Context) {
	s.logger.Println("REQUEST RECEIVED INTO GetSoldiersBySurname")
	var req genprotos.GetBySurname
	req.SoldierSurname = c.Param("soldier_surname")
	response, err := s.client.GetSoldiersBySurname(c, &req)
	if err != nil {
		s.logger.Println("ERROR WHILE SENDING REQUEST TO THE SERVER IN GetSoldiersBySurname :", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"response": response})
}

// GetSoldiersByGroupNameHandler retrieves soldiers by group name.
// @Summary Get soldiers by group name
// @Tags Soldiers
// @Produce json
// @Param group_name path string true "Group Name"
// @Success 200 {object} genprotos.GetSoldiersResponse
// @Failure 400 {object} gin.H{"error": string}
// @Router /soldiers/group/{group_name} [get]
func (s *SoldiersHandler) GetSoldiersByGroupNameHandler(c *gin.Context) {
	s.logger.Println("REQUEST RECEIVED INTO GetSoldiersByGroupNameHandler")
	var req genprotos.GetByGroupName
	req.GroupName = c.Param("group_name")
	response, err := s.client.GetSoldiersByGroupName(c, &req)
	if err != nil {
		s.logger.Println("ERROR WHILE SENDING REQUEST TO THE SERVER IN GetSoldiersByGroupNameHandler :", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"response": response})
}

// GetAllSoldiersHandler retrieves all soldiers with pagination.
// @Summary Get all soldiers with pagination
// @Tags Soldiers
// @Produce json
// @Param limit query int true "Limit"
// @Param offset query int true "Offset"
// @Success 200 {object} genprotos.GetSoldiersResponse
// @Failure 400 {object} gin.H{"error": string}
// @Router /soldiers [get]
func (s *SoldiersHandler) GetAllSoldiersHandler(c *gin.Context) {
	s.logger.Println("REQUEST RECEIVED INTO GetAllSoldiersHandler")
	var req genprotos.GetAllSoldiersRequest
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		s.logger.Println("ERROR WHILE CATCHING LIMIT IN GetAllSoldiersHandler :", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	req.Limit = int32(limit)
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		s.logger.Println("ERROR WHILE CATCHING OFFSET IN GetAllSoldiersHandler :", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	req.OffsetValue = int32(offset)
	response, err := s.client.GetAllSoldiers(c, &req)
	if err != nil {
		s.logger.Println("ERROR WHILE SENDING REQUEST TO THE SERVER IN GetAllSoldiersHandler :", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"response": response})
}

// GetSoldiersByAgeHandler retrieves soldiers by age.
// @Summary Get soldiers by age
// @Tags Soldiers
// @Produce json
// @Param age path int true "Soldier Age"
// @Success 200 {object} genprotos.GetSoldiersResponse
// @Failure 400 {object} gin.H{"error": string}
// @Router /soldiers/age/{age} [get]
func (s *SoldiersHandler) GetSoldiersByAgeHandler(c *gin.Context) {
	s.logger.Println("REQUEST RECEIVED INTO GetSoldiersByAgeHandler")
	age, err := strconv.Atoi(c.Param("age"))
	if err != nil {
		s.logger.Println("ERROR WHILE CONVERTING THE AGE IN GetSoldiersByAgeHandler SERVICE :", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	response, err := s.client.GetSoldiersByAge(c, &genprotos.GetByAgeRequest{
		SoldierAge: int32(age),
	})
	if err != nil {
		s.logger.Println("ERROR WHILE SENDING REQUEST TO THE SERVER IN GetSoldiersByAgeHandler SERVICE :", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"response": response})
}

// DeleteSoldierHandler deletes a soldier by ID.
// @Summary Delete a soldier by ID
// @Tags Soldiers
// @Produce json
// @Param soldier_id path string true "Soldier ID"
// @Success 200 {object} genprotos.DeleteSoldierResponse
// @Failure 400 {object} gin.H{"error": string}
// @Failure 401 {object} gin.H{"error": string}
// @Router /soldiers/{soldier_id} [delete]
func (s *SoldiersHandler) DeleteSoldierHandler(c *gin.Context) {
	id, err := s.auth.ExtractIDFromToken(c.GetHeader("Authorization")[8:])
	if err != nil {
		s.logger.Println("ERROR WHILE EXTRACTING ID IN CreateSoldierHandler :", err)
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}
	response, err := s.client.DeleteSoldier(c, &genprotos.DeleteRequest{
		SoldierId: c.Param("soldier_id"),
		DeletedBy: id,
	})
	if err != nil {
		s.logger.Println("ERROR WHILE SENDING REQUEST TO THE SERVER IN DeleteSoldierHandler SERVICE :", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"response": response})
}

// MoveSoldierFromGroupAToGroupBHandler moves a soldier from group A to group B.
// @Summary Move a soldier from group A to group B
// @Tags Soldiers
// @Accept json
// @Produce json
// @Param soldier_id path string true "Soldier ID"
// @Param moveSoldier body genprotos.MoveSoldierRequest true "Move Soldier"
// @Success 200 {object} genprotos.MoveSoldierResponse
// @Failure 400 {object} gin.H{"error": string}
// @Router /soldiers/{soldier_id}/move [put]
func (s *SoldiersHandler) MoveSoldierFromGroupAToGroupBHandler(c *gin.Context) {
	var req genprotos.MoveSoldierRequest
	soldierId := c.Param("soldier_id")

	if err := c.ShouldBindJSON(&req); err != nil {
		s.logger.Println("ERROR WHILE BINDING REQUEST IN MoveSoldierFromGroupAToGroupBHandler SERVICE :", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	req.SoldierId = soldierId
	response, err := s.client.MoveSoldierFromGroupAToGroupB(c, &req)
	if err != nil {
		s.logger.Println("ERROR WHILE SENDING REQUEST TO THE SERVER IN DeleteSoldierHandler SERVICE :", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"response": response})
}
