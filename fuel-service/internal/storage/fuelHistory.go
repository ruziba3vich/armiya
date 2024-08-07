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

func (s *FuelSt) GetFuelHistoriesByID(ctx context.Context, req *genprotos.GetFuelHistoriesByIdRequest) (*genprotos.GetFuelHistoriesByIdResponse, error) {
	query, args, err := s.queryBuilder.Select("id", "fuel_id", "action", "actior_id", "action_timestamp").
		From("fuel_history").
		Where(sq.Eq{"id": req.Id}).
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

	return &genprotos.GetFuelHistoriesByIdResponse{FuelHistories: response}, nil
}

func (s *FuelSt) GetFuelHistoriesByChoice(ctx context.Context, req *genprotos.GetFuelHistoriesByChoiceRequest) (*genprotos.GetFuelHistoriesByChoiceResponse, error) {
	var choice string
	if req.Choice == "type" {
		choice = "type"
	} else if req.Choice == "name" {
		choice = "name"
	} else {
		return nil, fmt.Errorf("invalid choice: %s", req.Choice)
	}
	query, args, err := s.queryBuilder.Select(
		"fh.id",
		"fh.fuel_id",
		"fh.action",
		"fh.actior_id",
		"fh.action_timestamp",
	).
		From("fuel_history fh").
		Join("fuel_management fm ON fh.fuel_id = fm.id").
		Where(sq.Eq{"fm."+choice: req.Message}).
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

	return &genprotos.GetFuelHistoriesByChoiceResponse{FuelHistories: response}, nil
}

func (s *FuelSt) GetFuelHistoriesByDate(ctx context.Context, req *genprotos.GetFuelHistoriesByDateRequest) (*genprotos.GetFuelHistoriesByDateResponse, error) {
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

	return &genprotos.GetFuelHistoriesByDateResponse{FuelHistories: response}, nil
}

func (s *FuelSt) GetFuelHistories(ctx context.Context, req *genprotos.Empty) (*genprotos.GetFuelHistoriesResponse, error) {
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

	return &genprotos.GetFuelHistoriesResponse{FuelHistories: response}, nil
}
