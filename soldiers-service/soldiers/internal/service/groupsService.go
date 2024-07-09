package service

import (
	"context"
	"log"

	"github.com/ruziba3vich/armiya/soldies-service/genprotos"
	"github.com/ruziba3vich/armiya/soldies-service/internal/storage"
)

type GroupsService struct {
	logger  *log.Logger
	storage *storage.GroupsStorage
	genprotos.UnimplementedGroupsServiceServer
}

func NewGroupsService(storage *storage.GroupsStorage, logger *log.Logger) *GroupsService {
	return &GroupsService{
		logger:  logger,
		storage: storage,
	}
}

func (s *GroupsService) CreateGroup(ctx context.Context, req *genprotos.CreateGroupRequest) (*genprotos.Group, error) {
	s.logger.Println("SERVER RECEIVED A REQUEST TO CreateGroup SERVICE")
	return s.storage.CreateGroup(ctx, req)
}

func (s *GroupsService) DeleteGroup(ctx context.Context, req *genprotos.DeleteGroupRequest) (*genprotos.Group, error) {
	s.logger.Println("SERVER RECEIVED A REQUEST TO DeleteGroup SERVICE")
	return s.storage.DeleteGroup(ctx, req)
}

func (s *GroupsService) DesignateTeacherToGroup(ctx context.Context, req *genprotos.DesignateTrainerRequest) (*genprotos.DesignateTrainerResponse, error) {
	s.logger.Println("SERVER RECEIVED A REQUEST TO DesignateTeacherToGroup SERVICE")
	return s.storage.DesignateTeacherToGroup(ctx, req)
}

/*
   type GroupsServiceServer interface {
       CreateGroup(context.Context, *CreateGroupRequest) (*Group, error)
       DeleteGroup(context.Context, *DeleteGroupRequest) (*Group, error)
       DesignateTeacherToGroup(context.Context, *DesignateTrainerRequest) (*DesignateTrainerResponse, error)
       mustEmbedUnimplementedGroupsServiceServer()
   }
*/
