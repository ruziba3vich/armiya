package storage

import (
	"context"
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/ruziba3vich/armiya/soldies-service/genprotos"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type GroupsStorage struct {
	db     *sql.DB
	sqrl   sq.StatementBuilderType
	logger *log.Logger
}

func NewGroupsStorage(db *sql.DB, logger *log.Logger, sqrl sq.StatementBuilderType) *GroupsStorage {
	return &GroupsStorage{
		db:     db,
		sqrl:   sqrl,
		logger: logger,
	}
}

func (s *GroupsStorage) CreateGroup(ctx context.Context, req *genprotos.CreateGroupRequest) (*genprotos.Group, error) {
	generatedId := uuid.New().String()
	createdAt := req.GetCreatedAt().AsTime()

	query, args, err := s.sqrl.Insert("groups").
		Columns(
			"group_id",
			"group_name",
			"created_by",
			"created_at").
		Values(
			generatedId,
			req.GetGroupName(),
			req.GetCreatedBy(),
			createdAt).
		ToSql()

	if err != nil {
		s.logger.Println("ERROR IN GENERATING QUERY TO CREATE A NEW GROUP:", err)
		return nil, err
	}

	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		s.logger.Println("ERROR IN EXECUTING QUERY TO CREATE A NEW GROUP:", err)
		return nil, err
	}

	return &genprotos.Group{
		GroupId:   generatedId,
		GroupName: req.GetGroupName(),
		CreatedBy: req.GetCreatedBy(),
		CreatedAt: timestamppb.New(createdAt),
	}, nil
}

func (s *GroupsStorage) DeleteGroup(ctx context.Context, req *genprotos.DeleteGroupRequest) (*genprotos.Group, error) {
	groupId := req.GetGroupId()
	deletedAt := req.GetDeletedAt().AsTime()
	deletedBy := req.GetDeletedBy().GetValue()

	query, args, err := s.sqrl.Update("groups").
		Set("deleted", true).
		Set("deleted_at", deletedAt).
		Set("deleted_by", deletedBy).
		Where(sq.Eq{"group_id": groupId}).
		ToSql()

	if err != nil {
		s.logger.Println("ERROR IN GENERATING QUERY TO DELETE GROUP:", err)
		return nil, err
	}

	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		s.logger.Println("ERROR IN EXECUTING QUERY TO DELETE GROUP:", err)
		return nil, err
	}

	return &genprotos.Group{
		GroupId:   groupId,
		Deleted:   true,
		DeletedAt: timestamppb.New(deletedAt),
		DeletedBy: wrapperspb.String(deletedBy),
	}, nil
}

func (s *GroupsStorage) DesignateTeacherToGroup(ctx context.Context, req *genprotos.DesignateTrainerRequest) (*genprotos.DesignateTrainerResponse, error) {
	designationId := uuid.New().String()

	query, args, err := s.sqrl.Insert("group_trainers").
		Columns("designation_id", "trainer_id", "group_id").
		Values(designationId, req.GetTrainerId(), req.GetGroupId()).
		ToSql()

	if err != nil {
		s.logger.Println("ERROR IN GENERATING QUERY TO DESIGNATE TRAINER TO GROUP:", err)
		return nil, err
	}

	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		s.logger.Println("ERROR IN EXECUTING QUERY TO DESIGNATE TRAINER TO GROUP:", err)
		return nil, err
	}

	return &genprotos.DesignateTrainerResponse{
		DesignationId: designationId,
		TrainerId:     req.GetTrainerId(),
		GroupId:       req.GetGroupId(),
	}, nil
}
