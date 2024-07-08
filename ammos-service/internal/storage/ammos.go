package storage

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/hackathon/army/ammos-service/genprotos"
	"github.com/hackathon/army/ammos-service/internal/config"

	sq "github.com/Masterminds/squirrel"
)

type (
	AmmosSt struct {
		db           *sql.DB
		queryBuilder sq.StatementBuilderType
	}
)

func New(config *config.Config) (*AmmosSt, error) {

	db, err := ConnectDB(*config)
	if err != nil {
		return nil, err
	}

	return &AmmosSt{
		db:           db,
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}, nil
}

// ammos

func (s *AmmosSt) CreateAmmo(ctx context.Context, req *genprotos.CreateAmmoRequest) (*genprotos.AmmoResponse, error) {
	id := uuid.New().String()
	last_update := time.Now()
	query, args, err := s.queryBuilder.Insert("ammos").
		Columns("id", "name", "caliber", "description", "type", "quantity", "last_update").
		Values(
			id,
			req.Name,
			req.Caliber,
			req.Description,
			req.Type,
			req.Quantity,
			last_update).
		ToSql()
	if err != nil {
		log.Println("Error generating SQL:", err)
		return nil, err
	}

	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		log.Println("Error creating row:", err)
		return nil, err
	}

	response := genprotos.Ammo{
		Id:          id,
		Name:        req.Name,
		Caliber:     req.Caliber,
		Description: req.Description,
		Type:        req.Type,
		Quantity:    req.Quantity,
		LastUpdate:  last_update.Format("2006-01-02 15:04:05"),
	}

	return &genprotos.AmmoResponse{Ammo: &response}, nil
}

func (s *AmmosSt) GetAmmoByChoice(ctx context.Context, req *genprotos.GetAmmoByChoiceRequest) (*genprotos.GetAmmoByChoiceResponse, error) {
	var choice string

	switch req.Choice {
	case "id", "name", "caliber", "type":
		choice = req.Choice
	default:
		return nil, fmt.Errorf("invalid choice: %s", req.Choice)
	}

	query, args, err := s.queryBuilder.Select("id", "name", "caliber", "description", "type", "quantity", "last_update").
		From("ammos").
		Where(sq.Eq{choice: req.Message}).
		ToSql()
	if err != nil {
		log.Println("Error generating SQL:", err)
		return nil, err
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		log.Println("Error getting rows:", err)
		return nil, err
	}
	defer rows.Close()

	var response []*genprotos.Ammo

	for rows.Next() {
		var ammo genprotos.Ammo
		err = rows.Scan(
			&ammo.Id,
			&ammo.Name,
			&ammo.Caliber,
			&ammo.Description,
			&ammo.Type,
			&ammo.Quantity,
			&ammo.LastUpdate,
		)
		if err != nil {
			log.Println("Scan row:", err)
			return nil, err
		}
		response = append(response, &ammo)
	}
	return &genprotos.GetAmmoByChoiceResponse{Ammo: response}, nil
}

func (s *AmmosSt) UpdateAmmoById(ctx context.Context, req *genprotos.UpdateAmmoByIdRequest) (*genprotos.AmmoResponse, error) {
	last_update := time.Now()
	query, args, err := s.queryBuilder.Update("ammos").
		Set("name", req.Name).
		Set("caliber", req.Caliber).
		Set("description", req.Description).
		Set("type", req.Type).
		Set("quantity", req.Quantity).
		Set("last_update", last_update).
		Where(sq.Eq{"id": req.Id}).
		ToSql()
	if err != nil {
		log.Println("Error generating SQL:", err)
		return nil, err
	}

	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		log.Println("Error updating row:", err)
		return nil, err
	}

	response := genprotos.Ammo{
		Id:          req.Id,
		Name:        req.Name,
		Caliber:     req.Caliber,
		Description: req.Description,
		Type:        req.Type,
		Quantity:    req.Quantity,
		LastUpdate:  last_update.Format("2006-01-02 15:04:05"),
	}

	return &genprotos.AmmoResponse{Ammo: &response}, nil
}

func (s *AmmosSt) DeleteAmmoById(ctx context.Context, req *genprotos.DeleteAmmoByIdRequest) (*genprotos.Empty, error) {
	query, args, err := s.queryBuilder.Delete("ammos").
		Where(sq.Eq{"id": req.Id}).
		ToSql()
	if err != nil {
		log.Println("Error generating SQL:", err)
		return nil, err
	}

	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		log.Println("Error deleting row:", err)
		return nil, err
	}

	return &genprotos.Empty{}, nil
}

func (s *AmmosSt) GetAmmo(ctx context.Context, req *genprotos.Empty) (*genprotos.GetAmmoResponse, error) {
	query, args, err := s.queryBuilder.Select("id", "name", "caliber", "description", "type", "quantity", "last_update").
		From("ammos").
		ToSql()
	if err != nil {
		log.Println("Error generating SQL:", err)
		return nil, err
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		log.Println("Error getting rows:", err)
		return nil, err
	}
	defer rows.Close()

	var response []*genprotos.Ammo
	
	for rows.Next() {
		var ammo genprotos.Ammo
		err = rows.Scan(
			&ammo.Id,
			&ammo.Name,
			&ammo.Caliber,
			&ammo.Description,
			&ammo.Type,
			&ammo.Quantity,
			&ammo.LastUpdate,
		)
		if err != nil {
			log.Println("Scan row:", err)
			return nil, err
		}
		response = append(response, &ammo)
	}

	return &genprotos.GetAmmoResponse{Ammo: response}, nil
}
