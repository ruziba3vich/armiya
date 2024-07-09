package storage

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/ruziba3vich/armiya/soldies-service/genprotos"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type AdminStorage struct {
	db     *sql.DB
	sqrl   sq.StatementBuilderType
	logger *log.Logger
}

func NewAdminsStorage(db *sql.DB, logger *log.Logger, sqrl sq.StatementBuilderType) *AdminStorage {
	return &AdminStorage{
		db:     db,
		sqrl:   sqrl,
		logger: logger,
	}
}

/*
	type AdminServiceServer interface {
		CreateTrainer(context.Context, *CreateTrainerRequest) (*Trainer, error)
		DeleteTrainer(context.Context, *DeleteTrainerRequest) (*Trainer, error)
		mustEmbedUnimplementedAdminServiceServer()
	}
*/

func (a *AdminStorage) CreateTrainerStorage(ctx context.Context, req *genprotos.CreateTrainerRequest) (*genprotos.Objects, error) {
	newUUID := uuid.New()
	objectId := newUUID.String()
	createdTime := time.Now()
	query, args, err := a.sqrl.Insert("objects").
		Columns(
			"object_id",
			"object_name",
			"object_surname",
			"position",
			"created_at",
		).
		Values(
			objectId,
			req.GetTrainerName(),
			req.GetTrainerSurname(),
			req.GetRole(),
			createdTime,
		).
		ToSql()
	if err != nil {
		a.logger.Println("ERROR 1 :", err)
		return nil, err
	}

	res, err := a.db.ExecContext(ctx, query, args...)
	response := genprotos.Objects{
		ObjectId:      objectId,
		ObjectName:    req.GetTrainerName(),
		ObjectSurname: req.GetTrainerSurname(),
		Position:      req.GetRole(),
		CreatedAt:     timestamppb.New(createdTime),
	}
	if err != nil {
		a.logger.Println("ERROR 2 :", err)
		return nil, err
	}

	if ra, err := res.RowsAffected(); ra == 0 || err != nil {
		return nil, errors.New("data could not be inserted")
	}

	return &response, nil
}

func (a *AdminStorage) DeleteTrainerStorage(ctx context.Context, req *genprotos.DeleteTrainerRequest) (*genprotos.Objects, error) {
	deletedTime := time.Now()
	query, args, err := a.sqrl.Update("objects").
		Set("deleted_at", deletedTime).
		Set("deleted_by", req.GetDletedBy()).
		Set("deleted", true).
		Where(sq.Eq{"object_id": req.GetTrainerId()}).
		Suffix("RETURNING object_id, object_name, object_surname, position, created_at, deleted_by, deleted").
		ToSql()
	if err != nil {
		a.logger.Println("ERROR 3 :", err)
		return nil, err
	}

	row := a.db.QueryRowContext(ctx, query, args...)
	var response genprotos.Objects
	var tm time.Time

	var deletedByValue string

	if err := row.Scan(
		&response.ObjectId,
		&response.ObjectName,
		&response.ObjectSurname,
		&response.Position,
		&tm,
		&deletedByValue,
		&response.Deleted,
	); err != nil {
		a.logger.Println("ERROR 3 :", err)
		return nil, err
	}
	response.DeletedAt = timestamppb.New(deletedTime)
	response.CreatedAt = timestamppb.New(tm)
	response.DeletedBy = wrapperspb.String(deletedByValue)

	return &response, nil
}
