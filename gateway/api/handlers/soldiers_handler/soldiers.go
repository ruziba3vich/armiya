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

func (s *SoldiersHandler) GetAllSoldiersHandler(c *gin.Context) {
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

/*
	// UpdateSoldier(ctx context.Context, in *UpdateSoldierRequest, opts ...grpc.CallOption) (*UpdateOrGetSoldierResponse, error)
    // GetSoldierById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*UpdateOrGetSoldierResponse, error)
    // GetSoldiersByName(ctx context.Context, in *GetByName, opts ...grpc.CallOption) (*GetSoldiersResponse, error)
    // GetSoldiersBySurname(ctx context.Context, in *GetBySurname, opts ...grpc.CallOption) (*GetSoldiersResponse, error)
    // GetSoldiersByGroupName(ctx context.Context, in *GetByGroupName, opts ...grpc.CallOption) (*GetSoldiersResponse, error)
    // GetAllSoldiers(ctx context.Context, in *GetAllSoldiersRequest, opts ...grpc.CallOption) (*GetSoldiersResponse, error)
    // GetSoldiersByAge(ctx context.Context, in *GetByAgeRequest, opts ...grpc.CallOption) (*GetSoldiersResponse, error)
    // DeleteSoldier(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteSoldierResponse, error)
    MoveSoldierFromGroupAToGroupB(ctx context.Context, in *MoveSoldierRequest, opts ...grpc.CallOption) (*MoveSoldierResponse, error)
}
*/
