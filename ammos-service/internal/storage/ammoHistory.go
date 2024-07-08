package storage

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/hackathon/army/ammos-service/genprotos"

	sq "github.com/Masterminds/squirrel"
)

// ammosHistory

func (s *AmmosSt) CreateAmmoHistory(ctx context.Context, req *genprotos.CreateAmmoHistoryRequest) (*genprotos.AmmoHistory, error) {
	id := uuid.New().String()
	action_timestamp := time.Now()
	query, args, err := s.queryBuilder.Insert("ammo_history").
		Columns("id", "ammo_id", "action", "actor_id", "action_timestamp").
		Values(
			id,
			req.AmmoId,
			req.Action,
			req.ActorId,
			action_timestamp,
		).
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

	response := genprotos.AmmoHistory{
		Id:              id,
		AmmoId:          req.AmmoId,
		Action:          req.Action,
		ActorId:         req.ActorId,
		ActionTimestamp: action_timestamp.Format("2006-01-02 15:04:05"),
	}

	return &response, nil
}

func (s *AmmosSt) GetAmmoHistoryByChoice(ctx context.Context, req *genprotos.GetAmmoHistoryByChoiceRequest) (*genprotos.GetAmmoHistoryByChoiceResponse, error) {
	var choice string

	switch req.Choice {
	case "id", "name", "caliber", "type":
		choice = req.Choice
	default:
		return nil, fmt.Errorf("invalid choice: %s", req.Choice)
	}

	query, args, err := s.queryBuilder.Select(
		"ah.id",
		"ah.ammo_id",
		"ah.action",
		"ah.actor_id",
		"ah.action_timestamp",
	).
		From("ammo_history as ah").
		Join("ammos as am ON ah.ammo_id = am.id").
		Where(sq.Eq{"am." + choice: req.Message}).
		OrderBy("ah.action_timestamp DESC").
		ToSql()

	if err != nil {
		log.Println("Error generating SQL:", err)
		return nil, err
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		log.Println("Error querying rows:", err)
		return nil, err
	}
	defer rows.Close()

	var response []*genprotos.AmmoHistory

	for rows.Next() {
		var ammoHistory genprotos.AmmoHistory
		err = rows.Scan(
			&ammoHistory.Id,
			&ammoHistory.AmmoId,
			&ammoHistory.Action,
			&ammoHistory.ActorId,
			&ammoHistory.ActionTimestamp,
		)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		response = append(response, &ammoHistory)
	}

	if err = rows.Err(); err != nil {
		log.Println("Error with rows:", err)
		return nil, err
	}

	return &genprotos.GetAmmoHistoryByChoiceResponse{AmmoHistory: response}, nil
}

func (s *AmmosSt) GetAmmoHistoryById(ctx context.Context, req *genprotos.GetAmmoHistoryByIdRequest) (*genprotos.AmmoHistory, error) {
	query, args, err := s.queryBuilder.Select("id", "ammo_id", "action", "actor_id", "action_timestamp").
		From("ammo_history").
		Where(sq.Eq{"id": req.Id}).
		ToSql()
	if err != nil {
		log.Println("Error generating SQL:", err)
		return nil, err
	}

	var response genprotos.AmmoHistory

	row := s.db.QueryRowContext(ctx, query, args...)
	if err := row.Scan(
		&response.Id,
		&response.AmmoId,
		&response.Action,
		&response.ActorId,
		&response.ActionTimestamp,
	); err != nil {
		log.Println("Error scanning row:", err)
		return nil, err
	}

	return &response, nil
}

func (s *AmmosSt) GetAmmoHistoryByDate(ctx context.Context, req *genprotos.GetAmmoHistoryByDateRequest) (*genprotos.GetAmmoHistoryByDateResponse, error) {
	if _, err := time.Parse("2006-01-02", req.Date); err != nil {
		log.Println("Invalid date format:", err)
		return nil, fmt.Errorf("invalid date format: %v", err)
	}

	query, args, err := s.queryBuilder.Select("id", "ammo_id", "action", "actor_id", "action_timestamp").
		From("ammo_history").
		Where("DATE(action_timestamp) = ?", req.Date).
		ToSql()

	if err != nil {
		log.Println("Error generating SQL:", err)
		return nil, err
	}
	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		log.Println("Error querying rows:", err)
		return nil, err
	}
	defer rows.Close()

	var response []*genprotos.AmmoHistory

	for rows.Next() {
		var ammoHistory genprotos.AmmoHistory
		err = rows.Scan(
			&ammoHistory.Id,
			&ammoHistory.AmmoId,
			&ammoHistory.Action,
			&ammoHistory.ActorId,
			&ammoHistory.ActionTimestamp,
		)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		response = append(response, &ammoHistory)
	}
	if err = rows.Err(); err != nil {
		log.Println("Error with rows:", err)
		return nil, err
	}

	return &genprotos.GetAmmoHistoryByDateResponse{AmmoHistory: response}, nil
}

func (s *AmmosSt) UpdateAmmoHistoryById(ctx context.Context, req *genprotos.UpdateAmmoHistoryByIdRequest) (*genprotos.AmmoHistory, error) {
	action_timestamp := time.Now()
	query, args, err := s.queryBuilder.Update("ammo_history").
		Set("ammo_id", req.AmmoId).
		Set("action", req.Action).
		Set("actor_id", req.ActorId).
		Set("action_timestamp", action_timestamp).
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

	response := genprotos.AmmoHistory{
		Id:              req.Id,
		AmmoId:          req.AmmoId,
		Action:          req.Action,
		ActorId:         req.ActorId,
		ActionTimestamp: action_timestamp.Format("2006-01-02 15:04:05"),
	}

	return &response, nil
}

func (s *AmmosSt) DeleteAmmoHistoryById(ctx context.Context, req *genprotos.DeleteAmmoHistoryByIdRequest) (*genprotos.Empty, error) {
	query, args, err := s.queryBuilder.Delete("ammo_history").
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

func (s AmmosSt) GetAmmoHistory(ctx context.Context, req *genprotos.Empty) (*genprotos.GetAmmoHistoryResponse, error) {
	query, args, err := s.queryBuilder.Select("id", "ammo_id", "action", "actor_id", "action_timestamp").
		From("ammo_history").
		ToSql()
	if err != nil {
		log.Println("Error generating SQL:", err)
		return nil, err
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		log.Println("Error querying rows:", err)
		return nil, err
	}
	defer rows.Close()

	var response []*genprotos.AmmoHistory

	for rows.Next() {
		var ammoHistory genprotos.AmmoHistory
		err = rows.Scan(
			&ammoHistory.Id,
			&ammoHistory.AmmoId,
			&ammoHistory.Action,
			&ammoHistory.ActorId,
			&ammoHistory.ActionTimestamp,
		)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		response = append(response, &ammoHistory)
	}

	return &genprotos.GetAmmoHistoryResponse{AmmoHistory: response}, nil
}
