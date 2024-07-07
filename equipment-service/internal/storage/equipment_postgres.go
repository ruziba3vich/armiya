package storage

import (
	"armiya/equipment-service/genprotos"
	"armiya/equipment-service/internal/config"
	"context"
	sql2 "database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type (
	Equipment struct {
		db           *sql2.DB
		queryBuilder sq.StatementBuilderType
	}
)

func New(config *config.Config) (*Equipment, error) {

	db, err := ConnectDB(*config)
	if err != nil {
		return nil, err
	}

	return &Equipment{
		db:           db,
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}, nil
}

// equipmentsSelectQuery constructs a SQL select query for the "equipments" table.
// It selects multiple columns from the table.
//
// Returns:
//
//	sq.SelectBuilder - A builder for constructing SQL select queries.
func equipmentsSelectQuery() sq.SelectBuilder {
	query := sq.Select(
		"name",
		"description",
		"origin_country",
		"classification",
		"quantity",
		"main_armament",
		"crew_size",
		"weight_kg",
		"length_cm",
		"width_cm",
		"height_cm",
		"max_speed_kmh",
		"operational_range_km",
		"year_of_introduction",
		"created_at",
	).From("equipments")

	return query
}

// CreateEquipment inserts a new equipment record into the database.
// It takes a context and a request containing equipment details, and returns the created equipment or an error.
//
// Parameters:
//
//	ctx - The context for managing request-scoped values, cancellation, and deadlines.
//	req - A pointer to a genprotos.Equipment struct containing the equipment details to be inserted.
//
// Returns:
//
//	*genprotos.Equipment - A pointer to the created equipment record.
//	error - An error if the operation fails, otherwise nil.
func (e *Equipment) CreateEquipment(ctx context.Context, req *genprotos.Equipment) (*genprotos.Equipment, error) {
	data := map[string]interface{}{
		"id":                   uuid.NewString(),
		"name":                 req.Name,
		"description":          req.Description,
		"origin_country":       req.OriginCountry,
		"classification":       req.Classification,
		"quantity":             req.Quantity,
		"main_armament":        req.MainArmament,
		"crew_size":            req.CrewSize,
		"weight_kg":            req.WeightKg,
		"length_cm":            req.LengthCm,
		"width_cm":             req.WidthCm,
		"height_cm":            req.HeightCm,
		"max_speed_kmh":        req.MaxSpeedKm,
		"operational_range_km": req.OperationalRangeKm,
		"year_of_introduction": req.YearOfIntroduction,
		"created_at":           time.Now(),
	}
	query, args, err := e.queryBuilder.Insert("equipments").
		SetMap(data).
		ToSql()
	if err != nil {
		return nil, err
	}

	if _, err = e.db.ExecContext(ctx, query, args...); err != nil {
		return nil, err
	}

	return &genprotos.Equipment{
		Id:                 data["id"].(string),
		Name:               req.Name,
		Description:        req.Description,
		OriginCountry:      req.OriginCountry,
		Classification:     req.Classification,
		Quantity:           req.Quantity,
		MainArmament:       req.MainArmament,
		CrewSize:           req.CrewSize,
		WeightKg:           req.WeightKg,
		LengthCm:           req.LengthCm,
		WidthCm:            req.WidthCm,
		HeightCm:           req.HeightCm,
		MaxSpeedKm:         req.MaxSpeedKm,
		OperationalRangeKm: req.OperationalRangeKm,
		YearOfIntroduction: req.YearOfIntroduction,
		CreatedAt:          data["created_at"].(time.Time).Format(time.RFC1123),
	}, nil
}

// GetEquipment retrieves an equipment record from the database based on the provided request.
// It takes a context and a request containing the equipment ID, and returns the retrieved equipment or an error.
//
// Parameters:
//
//	ctx - The context for managing request-scoped values, cancellation, and deadlines.
//	req - A pointer to a genprotos.GetRequest struct containing the equipment ID to be retrieved.
//
// Returns:
//
//	*genprotos.Equipment - A pointer to the retrieved equipment record.
//	error - An error if the operation fails, otherwise nil.
func (e *Equipment) GetEquipment(ctx context.Context, req *genprotos.GetRequest) (*genprotos.Equipment, error) {
	sql, args, err := equipmentsSelectQuery().Where(sq.Eq{
		"id": req.Id,
	}).ToSql()
	if err != nil {
		return nil, err
	}

	row := e.db.QueryRowContext(ctx, sql, args...)
	if row.Err() != nil {
		return nil, row.Err()
	}

	equipment := &genprotos.Equipment{}
	err = row.Scan(
		&equipment.Name,
		&equipment.Description,
		&equipment.OriginCountry,
		&equipment.Classification,
		&equipment.Quantity,
		&equipment.MainArmament,
		&equipment.CrewSize,
		&equipment.WeightKg,
		&equipment.LengthCm,
		&equipment.WidthCm,
		&equipment.HeightCm,
		&equipment.MaxSpeedKm,
		&equipment.OperationalRangeKm,
		&equipment.YearOfIntroduction,
		&equipment.CreatedAt,
	)
	if err != nil {
		if err == sql2.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return equipment, nil
}

// GetAllMessages retrieves paginated and ordered equipment records from the database.
// It takes a context and a GetAllRequest, and returns a GetAllResponse or an error.
//
// Parameters:
//
//	ctx - The context for managing request-scoped values, cancellation, and deadlines.
//	req - The GetAllRequest containing pagination and ordering information.
//
// Returns:
//
//	*genprotos.GetAllResponse - The response containing the retrieved equipment records and the total count.
//	error - An error if the operation fails, otherwise nil.
func (e *Equipment) GetAllEquipments(ctx context.Context, req *genprotos.GetAllRequest) (*genprotos.GetAllResponse, error) {
	sql, args, err := equipmentsSelectQuery().
		OrderBy(req.OrderBy).
		Limit(uint64(req.Limit)).
		Offset(uint64((req.Page - 1) * req.Limit)).
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := e.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var equipments []*genprotos.Equipment

	for rows.Next() {
		equipment := &genprotos.Equipment{}
		err := rows.Scan(
			&equipment.Name,
			&equipment.Description,
			&equipment.OriginCountry,
			&equipment.Classification,
			&equipment.Quantity,
			&equipment.MainArmament,
			&equipment.CrewSize,
			&equipment.WeightKg,
			&equipment.LengthCm,
			&equipment.WidthCm,
			&equipment.HeightCm,
			&equipment.MaxSpeedKm,
			&equipment.OperationalRangeKm,
			&equipment.YearOfIntroduction,
			&equipment.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		equipments = append(equipments, equipment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	response := &genprotos.GetAllResponse{
		Equipments: equipments,
		Count:      int64(len(equipments)),
	}

	return response, nil
}
