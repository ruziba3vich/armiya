package storage

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/ruziba3vich/armiya/soldies-service/genprotos"
)

type AdminStorage struct {
	db   *sql.DB
	sqrl sq.StatementBuilderType
}

func New(db *sql.DB) *AdminStorage {
	return &AdminStorage{
		db:   db,
		sqrl: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

/*
	type AdminServiceServer interface {
		CreateTrainer(context.Context, *CreateTrainerRequest) (*Trainer, error)
		DeleteTrainer(context.Context, *DeleteTrainerRequest) (*Trainer, error)
		mustEmbedUnimplementedAdminServiceServer()
	}
*/

func (a *AdminStorage) CreateTrainerStorage(ctx context.Context, req *genprotos.CreateTrainerRequest) (*genprotos.Trainer, error) {
	query, args, err := a.sqrl.Insert("trainers").
		Columns(
			"trainer_name",
			"trainer_surname",
			"role",
			"created_at",
			"deleted_at",
			"deleted",
			"deleted_by").
		Values(
			req.GetTrainerName(),
			req.GetTrainerSurname(),
			req.GetRole(),
			req.GetCreatedAt(),
			req.GetDeletedAt(),
			req.GetDeleted(),
			req.GetDeletedBy()).
		Suffix("RETURNING trainer_id, trainer_name, trainer_surname, role, created_at, deleted_at, deleted, deleted_by").
		ToSql()
	if err != nil {
		return nil, err
	}

	row := a.db.QueryRowContext(ctx, query, args...)
	var response genprotos.Trainer
	if err := row.Scan(
		&response.TrainerId,
		&response.TrainerName,
		&response.TrainerSurname,
		&response.Role,
		&response.CreatedAt,
		&response.DeletedAt,
		&response.Deleted,
		&response.DeletedBy,
	); err != nil {
		return nil, err
	}

	return &response, nil
}

func (a *AdminStorage) DeleteTrainerStorage(ctx context.Context, req *genprotos.DeleteTrainerRequest) (*genprotos.Trainer, error) {
	query, args, err := a.sqrl.Update("trainers").
		Set("deleted_at", time.Now()).
		Set("deleted_by", req.GetDletedBy()).
		Where(sq.Eq{"trainer_id": req.GetTrainerId()}).
		Suffix("RETURNING trainer_id, trainer_name, trainer_surname, role, created_at, deleted_at, deleted, deleted_by").
		ToSql()
	if err != nil {
		return nil, err
	}

	row := a.db.QueryRowContext(ctx, query, args...)
	var response genprotos.Trainer
	if err := row.Scan(
		&response.TrainerId,
		&response.TrainerName,
		&response.TrainerSurname,
		&response.Role,
		&response.CreatedAt,
		&response.DeletedAt,
		&response.Deleted,
		&response.DeletedBy,
	); err != nil {
		return nil, err
	}

	return &response, nil
}
