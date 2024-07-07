package storage

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/hackathon/army/fuel-service/genprotos"
	"github.com/hackathon/army/fuel-service/internal/config"

	sq "github.com/Masterminds/squirrel"
)

type (
	FuelSt struct {
		db           *sql.DB
		queryBuilder sq.StatementBuilderType
	}
)

func New(config *config.Config) (*FuelSt, error) {

	db, err := ConnectDB(*config)
	if err != nil {
		return nil, err
	}

	return &FuelSt{
		db:           db,
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}, nil
}

// fuel_management

func (s *FuelSt) CreateFuel(ctx context.Context, req *genprotos.CreateFuelRequest) (*genprotos.FuelResponse, error) {
	id := uuid.New().String()
	last_update := time.Now()
	query, args, err := s.queryBuilder.Insert("fuel_management").
		Columns("id", "name", "type", "quantity", "last_update").
		Values(
			id,
			req.Name,
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

	response := genprotos.Fuel{
		Id:         id,
		Name:       req.Name,
		Type:       req.Type,
		Quantity:   req.Quantity,
		LastUpdate: last_update.Format("2006-01-02 15:04:05"),
	}

	return &genprotos.FuelResponse{Fuel: &response}, nil
}

func (s *FuelSt) GetFuel(ctx context.Context, req *genprotos.GetFuelRequest) (*genprotos.FuelResponse, error) {
	query, args, err := s.queryBuilder.Select("id", "name", "type", "quantity", "last_update").
		From("fuel_management").
		Where(sq.Eq{"id": req.Id}).
		ToSql()
	if err != nil {
		log.Println("Error generating SQL:", err)
		return nil, err
	}

	row := s.db.QueryRowContext(ctx, query, args...)

	var response genprotos.Fuel

	if err := row.Scan(
		&response.Id,
		&response.Name,
		&response.Type,
		&response.Quantity,
		&response.LastUpdate); err != nil {
		log.Println("Scan row:", err)
		return nil, err
	}

	if err := row.Err(); err != nil {
		log.Println("Row error:", err)
		return nil, err
	}

	return &genprotos.FuelResponse{Fuel: &response}, nil
}

func (s *FuelSt) UpdateFuel(ctx context.Context, req *genprotos.UpdateFuelRequest) (*genprotos.FuelResponse, error) {
	query, args, err := s.queryBuilder.Update("fuel_management").
		Set("name", req.Fuel.Name).
		Set("type", req.Fuel.Type).
		Set("quantity", req.Fuel.Quantity).
		Set("last_update", time.Now()).
		Where(sq.Eq{"id": req.Fuel.Id}).
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

	response := genprotos.Fuel{
		Id:         req.Fuel.Id,
		Name:       req.Fuel.Name,
		Type:       req.Fuel.Type,
		Quantity:   req.Fuel.Quantity,
		LastUpdate: time.Now().Format("2006-01-02 15:04:05"),
	}

	return &genprotos.FuelResponse{Fuel: &response}, nil
}

func (s *FuelSt) DeleteFuel(ctx context.Context, req *genprotos.DeleteFuelRequest) (*genprotos.Empty, error) {
	query, args, err := s.queryBuilder.Delete("fuel_management").
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

	return nil, nil
}

func (s *FuelSt) ListFuelByType(ctx context.Context, req *genprotos.ListFuelsByTypeRequest) (*genprotos.ListFuelsByTypeResponse, error) {
	query, args, err := s.queryBuilder.Select("id", "name", "type", "quantity", "last_update").
		From("fuel_management").
		Where(sq.Eq{"type": req.Type}).
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

	var response []*genprotos.Fuel

	for rows.Next() {
		var fuel genprotos.Fuel
		if err := rows.Scan(
			&fuel.Id,
			&fuel.Name,
			&fuel.Type,
			&fuel.Quantity,
			&fuel.LastUpdate); err != nil {
			log.Println("Scan row:", err)
			return nil, err
		}
		if err := rows.Err(); err != nil {
			log.Println("Row error:", err)
			return nil, err
		}

		response = append(response, &fuel)
	}

	return &genprotos.ListFuelsByTypeResponse{Fuels: response}, nil
}

func (s *FuelSt) ListFuels(ctx context.Context, req *genprotos.Empty) (*genprotos.ListFuelsResponse, error) {
	query, _, err := s.queryBuilder.Select("id", "name", "type", "quantity", "last_update").
		From("fuel_management").
		ToSql()
	if err != nil {
		log.Println("Error generating SQL:", err)
		return nil, err
	}

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		log.Println("Error getting rows:", err)
		return nil, err
	}

	var response []*genprotos.Fuel

	for rows.Next() {
		var fuel genprotos.Fuel
		if err := rows.Scan(
			&fuel.Id,
			&fuel.Name,
			&fuel.Type,
			&fuel.Quantity,
			&fuel.LastUpdate); err != nil {
			log.Println("Scan row:", err)
			return nil, err
		}

		if err := rows.Err(); err != nil {
			log.Println("Row error:", err)
			return nil, err
		}
		response = append(response, &fuel)
	}

	return &genprotos.ListFuelsResponse{Fuels: response}, nil
}
