package storage

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	sq "github.com/Masterminds/squirrel"
	"github.com/hackathon/army/fuel-service/genprotos"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateFuel(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	queryBuilder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	fuelSt := &FuelSt{db: db, queryBuilder: queryBuilder}

	req := &genprotos.CreateFuelRequest{
		Name:     "New Fuel",
		Type:     "Petrol",
		Quantity: 5000.00,
	}

	mock.ExpectExec(`INSERT INTO fuel_management \(id,name,type,quantity,last_update\) VALUES \(\$1,\$2,\$3,\$4,\$5\)`).
		WithArgs(sqlmock.AnyArg(), req.Name, req.Type, req.Quantity, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := fuelSt.CreateFuel(ctx, req)
	require.NoError(t, err)

	expectedResponse := &genprotos.FuelResponse{
		Fuel: &genprotos.Fuel{
			Id:         resp.Fuel.Id, 
			Name:       req.Name,
			Type:       req.Type,
			Quantity:   req.Quantity,
			LastUpdate: resp.Fuel.LastUpdate, 
		},
	}

	assert.Equal(t, expectedResponse.Fuel.Id, resp.Fuel.Id)
	assert.Equal(t, expectedResponse.Fuel.Name, resp.Fuel.Name)
	assert.Equal(t, expectedResponse.Fuel.Type, resp.Fuel.Type)
	assert.Equal(t, expectedResponse.Fuel.Quantity, resp.Fuel.Quantity)
	// We can't directly compare LastUpdate due to time.Now(), so we skip it here

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetFuel(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := &FuelSt{
		db:           db,
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}

	ctx := context.Background()

	testCases := []struct {
		name        string
		request     *genprotos.GetFuelRequest
		mockRows    *sqlmock.Rows
		mockError   error
		expectedErr error
		expectedRes *genprotos.FuelResponse
	}{
		{
			name: "Successful retrieval",
			request: &genprotos.GetFuelRequest{
				Id: "550e8400-e29b-41d4-a716-446655440000",
			},
			mockRows: sqlmock.NewRows([]string{"id", "name", "type", "quantity", "last_update"}).
				AddRow("550e8400-e29b-41d4-a716-446655440000", "Fuel A", "Diesel", 1000.00, time.Date(2023, 10, 1, 10, 0, 0, 0, time.UTC)),
			expectedRes: &genprotos.FuelResponse{
				Fuel: &genprotos.Fuel{
					Id:         "550e8400-e29b-41d4-a716-446655440000",
					Name:       "Fuel A",
					Type:       "Diesel",
					Quantity:   1000.00,
					LastUpdate: time.Date(2023, 10, 1, 10, 0, 0, 0, time.UTC).Format(time.RFC3339),
				},
			},
		},
		{
			name: "SQL error",
			request: &genprotos.GetFuelRequest{
				Id: "550e8400-e29b-41d4-a716-446655440001",
			},
			mockError:   sql.ErrConnDone,
			expectedErr: sql.ErrConnDone,
		},
		{
			name: "No rows found",
			request: &genprotos.GetFuelRequest{
				Id: "550e8400-e29b-41d4-a716-446655440002",
			},
			mockRows:    sqlmock.NewRows([]string{"id", "name", "type", "quantity", "last_update"}),
			expectedErr: sql.ErrNoRows,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			query := "SELECT id, name, type, quantity, last_update FROM fuel_management WHERE id = \\$1"

			if tc.mockError == nil {
				mock.ExpectQuery(query).
					WithArgs(tc.request.Id).
					WillReturnRows(tc.mockRows)
			} else {
				mock.ExpectQuery(query).
					WithArgs(tc.request.Id).
					WillReturnError(tc.mockError)
			}

			res, err := s.GetFuel(ctx, tc.request)

			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedErr, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, res)
				assert.Equal(t, tc.expectedRes, res)
			}

			err = mock.ExpectationsWereMet()
			assert.NoError(t, err)
		})
	}
}

func TestUpdateFuel(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	queryBuilder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	fuelSt := &FuelSt{db: db, queryBuilder: queryBuilder}

	req := &genprotos.UpdateFuelRequest{
		Id:       "550e8400-e29b-41d4-a716-446655440000",
		Name:     "Updated Fuel A",
		Type:     "Updated Diesel",
		Quantity: 2000.00,
	}

	mock.ExpectExec(`UPDATE fuel_management SET name = \$1, type = \$2, quantity = \$3, last_update = \$4 WHERE id = \$5`).
		WithArgs(req.Name, req.Type, req.Quantity, sqlmock.AnyArg(), req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := fuelSt.UpdateFuel(ctx, req)
	require.NoError(t, err)

	expectedResponse := &genprotos.FuelResponse{
		Fuel: &genprotos.Fuel{
			Id:         req.Id,
			Name:       req.Name,
			Type:       req.Type,
			Quantity:   req.Quantity,
			LastUpdate: time.Now().Format("2006-01-02 15:04:05"),
		},
	}

	assert.Equal(t, expectedResponse.Fuel.Id, resp.Fuel.Id)
	assert.Equal(t, expectedResponse.Fuel.Name, resp.Fuel.Name)
	assert.Equal(t, expectedResponse.Fuel.Type, resp.Fuel.Type)
	assert.Equal(t, expectedResponse.Fuel.Quantity, resp.Fuel.Quantity)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestDeleteFuel(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	queryBuilder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	fuelSt := &FuelSt{db: db, queryBuilder: queryBuilder}

	req := &genprotos.DeleteFuelRequest{
		Id: "550e8400-e29b-41d4-a716-446655440000",
	}

	mock.ExpectExec(`DELETE FROM fuel_management WHERE id = \$1`).
		WithArgs(req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := fuelSt.DeleteFuel(ctx, req)
	require.NoError(t, err)

	expectedResponse := &genprotos.Empty{}

	assert.Equal(t, expectedResponse, resp)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetFuelByChoice(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := &FuelSt{
		db:           db,
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}

	ctx := context.Background()

	testCases := []struct {
		name        string
		choice      string
		message     string
		mockRows    *sqlmock.Rows
		mockError   error
		expectedErr error
		expectedRes *genprotos.GetFuelsByChoiceResponse
	}{
		{
			name:    "Valid type choice",
			choice:  "type",
			message: "Diesel",
			mockRows: sqlmock.NewRows([]string{"id", "name", "type", "quantity", "last_update"}).
				AddRow("550e8400-e29b-41d4-a716-446655440000", "Fuel A", "Diesel", 1000.00, time.Date(2023, 10, 1, 10, 0, 0, 0, time.UTC)),
			expectedRes: &genprotos.GetFuelsByChoiceResponse{
				Fuels: []*genprotos.Fuel{
					{
						Id:         "550e8400-e29b-41d4-a716-446655440000",
						Name:       "Fuel A",
						Type:       "Diesel",
						Quantity:   1000.00,
						LastUpdate: time.Date(2023, 10, 1, 10, 0, 0, 0, time.UTC).Format(time.RFC3339),
					},
				},
			},
		},
		{
			name:    "Valid name choice",
			choice:  "name",
			message: "Fuel B",
			mockRows: sqlmock.NewRows([]string{"id", "name", "type", "quantity", "last_update"}).
				AddRow("550e8400-e29b-41d4-a716-446655440001", "Fuel B", "Petrol", 1500.00, time.Date(2023, 10, 2, 11, 0, 0, 0, time.UTC)),
			expectedRes: &genprotos.GetFuelsByChoiceResponse{
				Fuels: []*genprotos.Fuel{
					{
						Id:         "550e8400-e29b-41d4-a716-446655440001",
						Name:       "Fuel B",
						Type:       "Petrol",
						Quantity:   1500.00,
						LastUpdate: time.Date(2023, 10, 2, 11, 0, 0, 0, time.UTC).Format(time.RFC3339),
					},
				},
			},
		},
		{
			name:        "Invalid choice",
			choice:      "invalid",
			message:     "Fuel C",
			expectedErr: errors.New("invalid choice: invalid"),
		},
		{
			name:        "SQL error",
			choice:      "type",
			message:     "Diesel",
			mockError:   errors.New("SQL error"),
			expectedErr: errors.New("SQL error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.mockRows != nil {
				mock.ExpectQuery("SELECT id, name, type, quantity, last_update FROM fuel_management WHERE " + tc.choice + " = \\$1").
					WithArgs(tc.message).
					WillReturnRows(tc.mockRows)
			}

			if tc.mockError != nil {
				mock.ExpectQuery("SELECT id, name, type, quantity, last_update FROM fuel_management WHERE " + tc.choice + " = \\$1").
					WithArgs(tc.message).
					WillReturnError(tc.mockError)
			}

			req := &genprotos.GetFuelsByChoiceRequest{
				Choice:  tc.choice,
				Message: tc.message,
			}

			res, err := s.GetFuelByChoice(ctx, req)

			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedRes, res)
			}

			err = mock.ExpectationsWereMet()
			assert.NoError(t, err)
		})
	}
}

func TestGetFuels(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	queryBuilder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	fuelSt := &FuelSt{db: db, queryBuilder: queryBuilder}

	columns := []string{"id", "name", "type", "quantity", "last_update"}
	rows := sqlmock.NewRows(columns).
		AddRow("660e8400-e29b-41d4-a716-446655440000", "Diesel", "Type1", 1000.0, time.Now().Format("2006-01-02 15:04:05"))

	mock.ExpectQuery(`SELECT id, name, type, quantity, last_update FROM fuel_management`).
		WillReturnRows(rows)

	resp, err := fuelSt.GetFuels(ctx, &genprotos.Empty{})
	require.NoError(t, err)

	expectedResponse := &genprotos.GetFuelsResponse{
		Fuels: []*genprotos.Fuel{
			{
				Id:         "660e8400-e29b-41d4-a716-446655440000",
				Name:       "Diesel",
				Type:       "Type1",
				Quantity:   1000.0,
				LastUpdate: time.Now().Format("2006-01-02 15:04:05"),
			},
		},
	}

	assert.Equal(t, expectedResponse.Fuels[0].Id, resp.Fuels[0].Id)
	assert.Equal(t, expectedResponse.Fuels[0].Name, resp.Fuels[0].Name)
	assert.Equal(t, expectedResponse.Fuels[0].Type, resp.Fuels[0].Type)
	assert.Equal(t, expectedResponse.Fuels[0].Quantity, resp.Fuels[0].Quantity)
	// We can't directly compare LastUpdate due to time.Now(), so we skip it here

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
