package storage

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	sq "github.com/Masterminds/squirrel"
	"github.com/go-playground/assert/v2"
	"github.com/hackathon/army/ammos-service/genprotos"
	"github.com/stretchr/testify/require"
)

func TestCreateAmmo(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	queryBuilder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	ammosSt := &AmmosSt{db: db, queryBuilder: queryBuilder}

	req := &genprotos.CreateAmmoRequest{
		Name:        "Ammo M",
		Caliber:     "9mm",
		Description: "Standard 9mm ammo",
		Type:        "Handgun",
		Quantity:    1000,
	}

	mock.ExpectExec(`INSERT INTO ammos \(id,name,caliber,description,type,quantity,last_update\) VALUES \(\$1,\$2,\$3,\$4,\$5,\$6,\$7\)`).
		WithArgs(sqlmock.AnyArg(), req.Name, req.Caliber, req.Description, req.Type, req.Quantity, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := ammosSt.CreateAmmo(ctx, req)
	require.NoError(t, err)

	expectedResponse := &genprotos.AmmoResponse{
		Ammo: &genprotos.Ammo{
			Id:          resp.Ammo.Id,
			Name:        req.Name,
			Caliber:     req.Caliber,
			Description: req.Description,
			Type:        req.Type,
			Quantity:    req.Quantity,
			LastUpdate:  resp.Ammo.LastUpdate,
		},
	}

	assert.Equal(t, expectedResponse.Ammo.Id, resp.Ammo.Id)
	assert.Equal(t, expectedResponse.Ammo.Name, resp.Ammo.Name)
	assert.Equal(t, expectedResponse.Ammo.Caliber, resp.Ammo.Caliber)
	assert.Equal(t, expectedResponse.Ammo.Description, resp.Ammo.Description)
	assert.Equal(t, expectedResponse.Ammo.Type, resp.Ammo.Type)
	assert.Equal(t, expectedResponse.Ammo.Quantity, resp.Ammo.Quantity)
	assert.Equal(t, expectedResponse.Ammo.LastUpdate, resp.Ammo.LastUpdate)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetAmmoByChoice(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	queryBuilder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	ammosSt := &AmmosSt{db: db, queryBuilder: queryBuilder}

	req := &genprotos.GetAmmoByChoiceRequest{
		Choice:  "name",
		Message: "Ammo A",
	}

	columns := []string{"id", "name", "caliber", "description", "type", "quantity", "last_update"}
	rows := sqlmock.NewRows(columns).
		AddRow("550e8400-e29b-41d4-a716-446655440000", "Ammo A", "9mm", "Standard 9mm ammo", "Handgun", 1000, "2023-10-01 10:00:00")

	mock.ExpectQuery(`SELECT id, name, caliber, description, type, quantity, last_update FROM ammos WHERE name = \$1`).
		WithArgs(req.Message).
		WillReturnRows(rows)

	resp, err := ammosSt.GetAmmoByChoice(ctx, req)
	require.NoError(t, err)

	expectedResponse := &genprotos.GetAmmoByChoiceResponse{
		Ammo: []*genprotos.Ammo{
			{
				Id:          "550e8400-e29b-41d4-a716-446655440000",
				Name:        "Ammo A",
				Caliber:     "9mm",
				Description: "Standard 9mm ammo",
				Type:        "Handgun",
				Quantity:    1000,
				LastUpdate:  "2023-10-01 10:00:00",
			},
		},
	}

	assert.Equal(t, expectedResponse, resp)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestUpdateAmmoById(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	queryBuilder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	ammosSt := &AmmosSt{db: db, queryBuilder: queryBuilder}

	req := &genprotos.UpdateAmmoByIdRequest{
		Id:          "550e8400-e29b-41d4-a716-446655440000",
		Name:        "Updated Ammo A",
		Caliber:     "9mm",
		Description: "Updated description",
		Type:        "Handgun",
		Quantity:    2000,
	}

	lastUpdate := time.Now()

	mock.ExpectExec(`UPDATE ammos SET name = \$1, caliber = \$2, description = \$3, type = \$4, quantity = \$5, last_update = \$6 WHERE id = \$7`).
		WithArgs(req.Name, req.Caliber, req.Description, req.Type, req.Quantity, sqlmock.AnyArg(), req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := ammosSt.UpdateAmmoById(ctx, req)
	require.NoError(t, err)

	expectedResponse := &genprotos.AmmoResponse{
		Ammo: &genprotos.Ammo{
			Id:          req.Id,
			Name:        req.Name,
			Caliber:     req.Caliber,
			Description: req.Description,
			Type:        req.Type,
			Quantity:    req.Quantity,
			LastUpdate:  lastUpdate.Format("2006-01-02 15:04:05"),
		},
	}

	assert.Equal(t, expectedResponse.Ammo.Id, resp.Ammo.Id)
	assert.Equal(t, expectedResponse.Ammo.Name, resp.Ammo.Name)
	assert.Equal(t, expectedResponse.Ammo.Caliber, resp.Ammo.Caliber)
	assert.Equal(t, expectedResponse.Ammo.Description, resp.Ammo.Description)
	assert.Equal(t, expectedResponse.Ammo.Type, resp.Ammo.Type)
	assert.Equal(t, expectedResponse.Ammo.Quantity, resp.Ammo.Quantity)
	assert.Equal(t, expectedResponse.Ammo.LastUpdate, resp.Ammo.LastUpdate)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestDeleteAmmoById(t *testing.T) {
    db, mock, err := sqlmock.New()
    require.NoError(t, err)
    defer db.Close()

    ctx := context.Background()
    queryBuilder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
    ammosSt := &AmmosSt{db: db, queryBuilder: queryBuilder}

    req := &genprotos.DeleteAmmoByIdRequest{
        Id: "550e8400-e29b-41d4-a716-446655440000",
    }

    mock.ExpectExec(`DELETE FROM ammos WHERE id = \$1`).
        WithArgs(req.Id).
        WillReturnResult(sqlmock.NewResult(1, 1))

    resp, err := ammosSt.DeleteAmmoById(ctx, req)
    require.NoError(t, err)

    expectedResponse := &genprotos.Empty{}

    assert.Equal(t, expectedResponse, resp)

    err = mock.ExpectationsWereMet()
    require.NoError(t, err)
}

func TestGetAmmo(t *testing.T) {
    db, mock, err := sqlmock.New()
    require.NoError(t, err)
    defer db.Close()

    ctx := context.Background()
    queryBuilder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
    ammosSt := &AmmosSt{db: db, queryBuilder: queryBuilder}

    req := &genprotos.Empty{}

    columns := []string{"id", "name", "caliber", "description", "type", "quantity", "last_update"}
    rows := sqlmock.NewRows(columns).
        AddRow("550e8400-e29b-41d4-a716-446655440000", "Ammo A", "9mm", "Standard 9mm ammo", "Handgun", 1000, "2023-10-01 10:00:00").
        AddRow("550e8400-e29b-41d4-a716-446655440001", "Ammo B", "5.56mm", "Standard 5.56mm ammo", "Rifle", 500, "2023-10-01 10:01:00")

    mock.ExpectQuery(`SELECT id, name, caliber, description, type, quantity, last_update FROM ammos`).
        WillReturnRows(rows)

    resp, err := ammosSt.GetAmmo(ctx, req)
    require.NoError(t, err)

    expectedResponse := &genprotos.GetAmmoResponse{
        Ammo: []*genprotos.Ammo{
            {
                Id:          "550e8400-e29b-41d4-a716-446655440000",
                Name:        "Ammo A",
                Caliber:     "9mm",
                Description: "Standard 9mm ammo",
                Type:        "Handgun",
                Quantity:    1000,
                LastUpdate:  "2023-10-01 10:00:00",
            },
            {
                Id:          "550e8400-e29b-41d4-a716-446655440001",
                Name:        "Ammo B",
                Caliber:     "5.56mm",
                Description: "Standard 5.56mm ammo",
                Type:        "Rifle",
                Quantity:    500,
                LastUpdate:  "2023-10-01 10:01:00",
            },
        },
    }

    assert.Equal(t, expectedResponse, resp)

    err = mock.ExpectationsWereMet()
    require.NoError(t, err)
}
