package storage

import (
	"context"
	"fmt"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/hackathon/army/fuel-service/genprotos"
)

// fuel_history

func (s *FuelSt) CreateFuelHistory(ctx context.Context, req *genprotos.CreateFuelHistoryRequest) (*genprotos.FuelHistoryResponse, error) {
	id := uuid.New().String()
	action_timestamp := time.Now()
	query, args, err := s.queryBuilder.Insert("fuel_history").
		Columns("id", "fuel_id", "action", "actior_id", "action_timestamp").
		Values(
			id,
			req.FuelId,
			req.Action,
			req.ActiorId,
			action_timestamp).
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

	response := genprotos.FuelHistory{
		Id:              id,
		FuelId:          req.FuelId,
		Action:          req.Action,
		ActiorId:        req.ActiorId,
		ActionTimestamp: action_timestamp.Format("2006-01-02 15:04:05"),
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

func (s *FuelSt) ListFuelHistoriesByType(ctx context.Context, req *genprotos.ListFuelHistoriesByTypeRequest) (*genprotos.ListFuelHistoriesByTypeResponse, error) {
	query, args, err := s.queryBuilder.Select(
		"fh.id",
		"fh.fuel_id",
		"fh.action",
		"fh.actior_id",
		"fh.action_timestamp",
	).
		From("fuel_history fh").
		Join("fuel_management fm ON fh.fuel_id = fm.id").
		Where(sq.Eq{"fm.type": req.Type}).
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

	var response []*genprotos.FuelHistory

	for rows.Next() {
		var fuel genprotos.FuelHistory
		if err := rows.Scan(
			&fuel.Id,
			&fuel.FuelId,
			&fuel.Action,
			&fuel.ActiorId,
			&fuel.ActionTimestamp,
		); err != nil {
			log.Println("Scan row:", err)
			return nil, err
		}
		response = append(response, &fuel)
	}

	if err := rows.Err(); err != nil {
		log.Println("Row error:", err)
		return nil, err
	}

	return &genprotos.ListFuelHistoriesByTypeResponse{FuelHistories: response}, nil
}

func (s *FuelSt) ListFuelHistoriesByDate(ctx context.Context, req *genprotos.ListFuelHistoriesByDateRequest) (*genprotos.ListFuelHistoriesByDateResponse, error) {
	if _, err := time.Parse("2006-01-02", req.Date); err != nil {
		log.Println("Invalid date format:", err)
		return nil, fmt.Errorf("invalid date format: %v", err)
	}

	query, args, err := s.queryBuilder.Select("id", "fuel_id", "action", "actior_id", "action_timestamp").
		From("fuel_history").
		Where("DATE(action_timestamp) = ?", req.Date).
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

	var response []*genprotos.FuelHistory

	for rows.Next() {
		var fuel genprotos.FuelHistory
		if err := rows.Scan(
			&fuel.Id,
			&fuel.FuelId,
			&fuel.Action,
			&fuel.ActiorId,
			&fuel.ActionTimestamp,
		); err != nil {
			log.Println("Scan row:", err)
			return nil, err
		}
		response = append(response, &fuel)
	}

	if err := rows.Err(); err != nil {
		log.Println("Row error:", err)
		return nil, err
	}

	return &genprotos.ListFuelHistoriesByDateResponse{FuelHistories: response}, nil
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
