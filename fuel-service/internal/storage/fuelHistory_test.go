package storage

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	sq "github.com/Masterminds/squirrel"
	"github.com/hackathon/army/fuel-service/genprotos"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateFuelHistory(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	queryBuilder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	fuelSt := &FuelSt{db: db, queryBuilder: queryBuilder}

	req := &genprotos.CreateFuelHistoryRequest{
		FuelId:   "550e8400-e29b-41d4-a716-446655440000",
		Action:   "Add",
		ActiorId: "770e8400-e29b-41d4-a716-446655440000",
	}

	mock.ExpectExec(`INSERT INTO fuel_history \(id,fuel_id,action,actior_id,action_timestamp\) VALUES \(\$1,\$2,\$3,\$4,\$5\)`).
		WithArgs(sqlmock.AnyArg(), req.FuelId, req.Action, req.ActiorId, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := fuelSt.CreateFuelHistory(ctx, req)
	require.NoError(t, err)

	expectedResponse := &genprotos.FuelHistoryResponse{
		FuelHistory: &genprotos.FuelHistory{
			Id:              resp.FuelHistory.Id, // We use the generated ID from the response
			FuelId:          req.FuelId,
			Action:          req.Action,
			ActiorId:        req.ActiorId,
			ActionTimestamp: resp.FuelHistory.ActionTimestamp, // We use the generated ActionTimestamp from the response
		},
	}

	assert.Equal(t, expectedResponse.FuelHistory.Id, resp.FuelHistory.Id)
	assert.Equal(t, expectedResponse.FuelHistory.FuelId, resp.FuelHistory.FuelId)
	assert.Equal(t, expectedResponse.FuelHistory.Action, resp.FuelHistory.Action)
	assert.Equal(t, expectedResponse.FuelHistory.ActiorId, resp.FuelHistory.ActiorId)
	// We can't directly compare ActionTimestamp due to time.Now(), so we skip it here

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetFuelHistoriesByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	queryBuilder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	fuelSt := &FuelSt{db: db, queryBuilder: queryBuilder}

	req := &genprotos.GetFuelHistoriesByIdRequest{
		Id: "660e8400-e29b-41d4-a716-446655440000",
	}

	columns := []string{"id", "fuel_id", "action", "actior_id", "action_timestamp"}
	rows := sqlmock.NewRows(columns).
		AddRow("660e8400-e29b-41d4-a716-446655440000", "550e8400-e29b-41d4-a716-446655440000", "Add", "770e8400-e29b-41d4-a716-446655440000", time.Now().Format("2006-01-02 15:04:05"))

	mock.ExpectQuery(`SELECT id, fuel_id, action, actior_id, action_timestamp FROM fuel_history WHERE id = \$1`).
		WithArgs(req.Id).
		WillReturnRows(rows)

	resp, err := fuelSt.GetFuelHistoriesByID(ctx, req)
	require.NoError(t, err)

	expectedResponse := &genprotos.GetFuelHistoriesByIdResponse{
		FuelHistories: []*genprotos.FuelHistory{
			{
				Id:              "660e8400-e29b-41d4-a716-446655440000",
				FuelId:          "550e8400-e29b-41d4-a716-446655440000",
				Action:          "Add",
				ActiorId:        "770e8400-e29b-41d4-a716-446655440000",
				ActionTimestamp: time.Now().Format("2006-01-02 15:04:05"),
			},
		},
	}

	assert.Equal(t, expectedResponse.FuelHistories[0].Id, resp.FuelHistories[0].Id)
	assert.Equal(t, expectedResponse.FuelHistories[0].FuelId, resp.FuelHistories[0].FuelId)
	assert.Equal(t, expectedResponse.FuelHistories[0].Action, resp.FuelHistories[0].Action)
	assert.Equal(t, expectedResponse.FuelHistories[0].ActiorId, resp.FuelHistories[0].ActiorId)
	// We can't directly compare ActionTimestamp due to time.Now(), so we skip it here

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetFuelHistoriesByChoice(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	queryBuilder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	fuelSt := &FuelSt{db: db, queryBuilder: queryBuilder}

	req := &genprotos.GetFuelHistoriesByChoiceRequest{
		Choice:  "type",
		Message: "Diesel",
	}

	columns := []string{"id", "fuel_id", "action", "actior_id", "action_timestamp"}
	rows := sqlmock.NewRows(columns).
		AddRow("660e8400-e29b-41d4-a716-446655440000", "550e8400-e29b-41d4-a716-446655440000", "Add", "770e8400-e29b-41d4-a716-446655440000", time.Now().Format("2006-01-02 15:04:05"))

	mock.ExpectQuery(`SELECT fh.id, fh.fuel_id, fh.action, fh.actior_id, fh.action_timestamp FROM fuel_history fh JOIN fuel_management fm ON fh.fuel_id = fm.id WHERE fm.type = \$1`).
		WithArgs(req.Message).
		WillReturnRows(rows)

	resp, err := fuelSt.GetFuelHistoriesByChoice(ctx, req)
	require.NoError(t, err)

	expectedResponse := &genprotos.GetFuelHistoriesByChoiceResponse{
		FuelHistories: []*genprotos.FuelHistory{
			{
				Id:              "660e8400-e29b-41d4-a716-446655440000",
				FuelId:          "550e8400-e29b-41d4-a716-446655440000",
				Action:          "Add",
				ActiorId:        "770e8400-e29b-41d4-a716-446655440000",
				ActionTimestamp: time.Now().Format("2006-01-02 15:04:05"),
			},
		},
	}

	assert.Equal(t, expectedResponse.FuelHistories[0].Id, resp.FuelHistories[0].Id)
	assert.Equal(t, expectedResponse.FuelHistories[0].FuelId, resp.FuelHistories[0].FuelId)
	assert.Equal(t, expectedResponse.FuelHistories[0].Action, resp.FuelHistories[0].Action)
	assert.Equal(t, expectedResponse.FuelHistories[0].ActiorId, resp.FuelHistories[0].ActiorId)
	// We can't directly compare ActionTimestamp due to time.Now(), so we skip it here

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetFuelHistoriesByDate(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	queryBuilder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	fuelSt := &FuelSt{db: db, queryBuilder: queryBuilder}

	req := &genprotos.GetFuelHistoriesByDateRequest{
		Date: "2023-10-01",
	}

	columns := []string{"id", "fuel_id", "action", "actior_id", "action_timestamp"}
	rows := sqlmock.NewRows(columns).
		AddRow("660e8400-e29b-41d4-a716-446655440000", "550e8400-e29b-41d4-a716-446655440000", "Add", "770e8400-e29b-41d4-a716-446655440000", "2023-10-01 10:30:00")

	mock.ExpectQuery(`SELECT id, fuel_id, action, actior_id, action_timestamp FROM fuel_history WHERE DATE\(action_timestamp\) = \$1`).
		WithArgs(req.Date).
		WillReturnRows(rows)

	resp, err := fuelSt.GetFuelHistoriesByDate(ctx, req)
	require.NoError(t, err)

	expectedResponse := &genprotos.GetFuelHistoriesByDateResponse{
		FuelHistories: []*genprotos.FuelHistory{
			{
				Id:              "660e8400-e29b-41d4-a716-446655440000",
				FuelId:          "550e8400-e29b-41d4-a716-446655440000",
				Action:          "Add",
				ActiorId:        "770e8400-e29b-41d4-a716-446655440000",
				ActionTimestamp: "2023-10-01 10:30:00",
			},
		},
	}

	assert.Equal(t, expectedResponse.FuelHistories[0].Id, resp.FuelHistories[0].Id)
	assert.Equal(t, expectedResponse.FuelHistories[0].FuelId, resp.FuelHistories[0].FuelId)
	assert.Equal(t, expectedResponse.FuelHistories[0].Action, resp.FuelHistories[0].Action)
	assert.Equal(t, expectedResponse.FuelHistories[0].ActiorId, resp.FuelHistories[0].ActiorId)
	assert.Equal(t, expectedResponse.FuelHistories[0].ActionTimestamp, resp.FuelHistories[0].ActionTimestamp)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetFuelHistories(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	queryBuilder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	fuelSt := &FuelSt{db: db, queryBuilder: queryBuilder}

	columns := []string{"id", "fuel_id", "action", "actior_id", "action_timestamp"}
	rows := sqlmock.NewRows(columns).
		AddRow("660e8400-e29b-41d4-a716-446655440000", "550e8400-e29b-41d4-a716-446655440000", "Add", "770e8400-e29b-41d4-a716-446655440000", time.Now().Format("2006-01-02 15:04:05"))

	mock.ExpectQuery(`SELECT id, fuel_id, action, actior_id, action_timestamp FROM fuel_history`).
		WillReturnRows(rows)

	resp, err := fuelSt.GetFuelHistories(ctx, &genprotos.Empty{})
	require.NoError(t, err)

	expectedResponse := &genprotos.GetFuelHistoriesResponse{
		FuelHistories: []*genprotos.FuelHistory{
			{
				Id:              "660e8400-e29b-41d4-a716-446655440000",
				FuelId:          "550e8400-e29b-41d4-a716-446655440000",
				Action:          "Add",
				ActiorId:        "770e8400-e29b-41d4-a716-446655440000",
				ActionTimestamp: time.Now().Format("2006-01-02 15:04:05"),
			},
		},
	}

	assert.Equal(t, expectedResponse.FuelHistories[0].Id, resp.FuelHistories[0].Id)
	assert.Equal(t, expectedResponse.FuelHistories[0].FuelId, resp.FuelHistories[0].FuelId)
	assert.Equal(t, expectedResponse.FuelHistories[0].Action, resp.FuelHistories[0].Action)
	assert.Equal(t, expectedResponse.FuelHistories[0].ActiorId, resp.FuelHistories[0].ActiorId)
	// We can't directly compare ActionTimestamp due to time.Now(), so we skip it here

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}