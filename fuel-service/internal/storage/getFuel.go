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
	query, args, err := s.queryBuilder.Insert("fuel_management").
		Columns("id", "name", "type", "quantity", "last_update").
		Values(
			uuid.New().String(),
			req.Name,
			req.Type,
			req.Quantity,
			time.Now()).
		Suffix("RETURNING id, name, type, quantity, last_update").
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
		Suffix("RETURNING id, name, type, quantity, last_update").
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

// fuel_history

func (s *FuelSt) CreateFuelHistory(ctx context.Context, req *genprotos.CreateFuelHistoryRequest) (*genprotos.FuelHistoryResponse, error) {
	query, args, err := s.queryBuilder.Insert("fuel_history").
		Columns("id", "fuel_id", "action", "actior_id", "action_timestamp").
		Values(
			uuid.New().String(),
			req.FuelId,
			req.Action,
			req.ActiorId,
			time.Now()).
		Suffix("RETURNING id, fuel_id, action, actior_id, action_timestamp").
		ToSql()
	if err != nil {
		log.Println("Error generating SQL:", err)
		return nil, err
	}

	row := s.db.QueryRowContext(ctx, query, args...)

	var response genprotos.FuelHistory

	if err := row.Scan(
		&response.Id,
		&response.FuelId,
		&response.Action,
		&response.ActiorId,
		&response.ActionTimestamp); err != nil {
		log.Println("Scan row:", err)
		return nil, err
	}

	if err := row.Err(); err != nil {
		log.Println("Row error:", err)
		return nil, err
	}

	return &genprotos.FuelHistoryResponse{FuelHistory: &response}, nil
}

func (s *FuelSt) GetFuelHistory(ctx context.Context, req *genprotos.GetFuelHistoryRequest) (*genprotos.FuelHistoryResponse, error) {
	query, args, err := s.queryBuilder.Select("id", "fuel_id", "action", "actior_id", "action_timestamp").
		From("fuel_history").
		Where(sq.Eq{"id": req.Id}).
		ToSql()
	if err != nil {
		log.Println("Error generating SQL:", err)
		return nil, err
	}
	row := s.db.QueryRowContext(ctx, query, args...)

	var response genprotos.FuelHistory

	if err := row.Scan(
		&response.Id,
		&response.FuelId,
		&response.Action,
		&response.ActiorId,
		&response.ActionTimestamp); err != nil {
		log.Println("Scan row:", err)
		return nil, err
	}

	if err := row.Err(); err != nil {
		log.Println("Row error:", err)
		return nil, err
	}

	return &genprotos.FuelHistoryResponse{FuelHistory: &response}, nil
}

func (s *FuelSt) ListFuelHistoriesByFuelID(ctx context.Context, req *genprotos.ListFuelHistoriesByFuelIDRequest) (*genprotos.ListFuelHistoriesByFuelIDResponse, error) {
	query, args, err := s.queryBuilder.Select("id", "fuel_id", "action", "actior_id", "action_timestamp").
		From("fuel_history").
		Where(sq.Eq{"fuel_id": req.FuelId}).
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

	var response []*genprotos.FuelHistory

	for rows.Next() {
		var fuel genprotos.FuelHistory
		if err := rows.Scan(
			&fuel.Id,
			&fuel.FuelId,
			&fuel.Action,
			&fuel.ActiorId,
			&fuel.ActionTimestamp); err != nil {
			log.Println("Scan row:", err)
			return nil, err
		}
		if err := rows.Err(); err != nil {
			log.Println("Row error:", err)
			return nil, err
		}
		response = append(response, &fuel)
	}

	return &genprotos.ListFuelHistoriesByFuelIDResponse{FuelHistories: response}, nil
}

func (s *FuelSt) ListFuelHistories(ctx context.Context, req *genprotos.Empty) (*genprotos.ListFuelHistoriesResponse, error) {
	query, _, err := s.queryBuilder.Select("id", "fuel_id", "action", "actior_id", "action_timestamp").
		From("fuel_history").
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

	var response []*genprotos.FuelHistory

	for rows.Next() {
		var fuel genprotos.FuelHistory
		if err := rows.Scan(
			&fuel.Id,
			&fuel.FuelId,
			&fuel.Action,
			&fuel.ActiorId,
			&fuel.ActionTimestamp); err != nil {
			log.Println("Scan row:", err)
			return nil, err
		}
		if err := rows.Err(); err != nil {
			log.Println("Row error:", err)
			return nil, err
		}
		response = append(response, &fuel)
	}

	return &genprotos.ListFuelHistoriesResponse{FuelHistories: response}, nil
}
