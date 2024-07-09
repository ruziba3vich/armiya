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

func TestCreateAmmoHistory(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	queryBuilder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	ammosSt := &AmmosSt{db: db, queryBuilder: queryBuilder}

	req := &genprotos.CreateAmmoHistoryRequest{
		AmmoId:  "550e8400-e29b-41d4-a716-446655440000",
		Action:  "Created",
		ActorId: "770e8400-e29b-41d4-a716-446655440000",
	}

	mock.ExpectExec(`INSERT INTO ammo_history \(id,ammo_id,action,actor_id,action_timestamp\) VALUES \(\$1,\$2,\$3,\$4,\$5\)`).
		WithArgs(sqlmock.AnyArg(), req.AmmoId, req.Action, req.ActorId, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := ammosSt.CreateAmmoHistory(ctx, req)
	require.NoError(t, err)

	expectedResponse := &genprotos.AmmoHistory{
		Id:              resp.Id, // Biz bu yerda haqiqiy IDni ishlatamiz
		AmmoId:          req.AmmoId,
		Action:          req.Action,
		ActorId:         req.ActorId,
		ActionTimestamp: resp.ActionTimestamp, // Biz bu yerda haqiqiy timestampni ishlatamiz
	}

	assert.Equal(t, expectedResponse.Id, resp.Id)
	assert.Equal(t, expectedResponse.AmmoId, resp.AmmoId)
	assert.Equal(t, expectedResponse.Action, resp.Action)
	assert.Equal(t, expectedResponse.ActorId, resp.ActorId)
	assert.Equal(t, expectedResponse.ActionTimestamp, resp.ActionTimestamp)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetAmmoHistoryByChoice(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	queryBuilder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	ammosSt := &AmmosSt{db: db, queryBuilder: queryBuilder}

	req := &genprotos.GetAmmoHistoryByChoiceRequest{
		Choice:  "name",
		Message: "Ammo A",
	}

	columns := []string{"id", "ammo_id", "action", "actor_id", "action_timestamp"}
	rows := sqlmock.NewRows(columns).
		AddRow("660e8400-e29b-41d4-a716-446655440001", "550e8400-e29b-41d4-a716-446655440000", "Updated", "770e8400-e29b-41d4-a716-446655440001", "2023-10-01 11:00:00").
		AddRow("660e8400-e29b-41d4-a716-446655440000", "550e8400-e29b-41d4-a716-446655440000", "Created", "770e8400-e29b-41d4-a716-446655440000", "2023-10-01 10:00:00")

	mock.ExpectQuery(`SELECT ah.id, ah.ammo_id, ah.action, ah.actor_id, ah.action_timestamp FROM ammo_history as ah JOIN ammos as am ON ah.ammo_id = am.id WHERE am.name = \$1 ORDER BY ah.action_timestamp DESC`).
		WithArgs(req.Message).
		WillReturnRows(rows)

	resp, err := ammosSt.GetAmmoHistoryByChoice(ctx, req)
	require.NoError(t, err)

	expectedResponse := &genprotos.GetAmmoHistoryByChoiceResponse{
		AmmoHistory: []*genprotos.AmmoHistory{
			{
				Id:              "660e8400-e29b-41d4-a716-446655440001",
				AmmoId:          "550e8400-e29b-41d4-a716-446655440000",
				Action:          "Updated",
				ActorId:         "770e8400-e29b-41d4-a716-446655440001",
				ActionTimestamp: "2023-10-01 11:00:00",
			},
			{
				Id:              "660e8400-e29b-41d4-a716-446655440000",
				AmmoId:          "550e8400-e29b-41d4-a716-446655440000",
				Action:          "Created",
				ActorId:         "770e8400-e29b-41d4-a716-446655440000",
				ActionTimestamp: "2023-10-01 10:00:00",
			},
		},
	}

	assert.Equal(t, expectedResponse, resp)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetAmmoHistoryById(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	queryBuilder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	ammosSt := &AmmosSt{db: db, queryBuilder: queryBuilder}

	req := &genprotos.GetAmmoHistoryByIdRequest{
		Id: "660e8400-e29b-41d4-a716-446655440000",
	}

	columns := []string{"id", "ammo_id", "action", "actor_id", "action_timestamp"}
	row := sqlmock.NewRows(columns).
		AddRow("660e8400-e29b-41d4-a716-446655440000", "550e8400-e29b-41d4-a716-446655440000", "Created", "770e8400-e29b-41d4-a716-446655440000", "2023-10-01 10:00:00")

	mock.ExpectQuery(`SELECT id, ammo_id, action, actor_id, action_timestamp FROM ammo_history WHERE id = \$1`).
		WithArgs(req.Id).
		WillReturnRows(row)

	resp, err := ammosSt.GetAmmoHistoryById(ctx, req)
	require.NoError(t, err)

	expectedResponse := &genprotos.AmmoHistory{
		Id:              "660e8400-e29b-41d4-a716-446655440000",
		AmmoId:          "550e8400-e29b-41d4-a716-446655440000",
		Action:          "Created",
		ActorId:         "770e8400-e29b-41d4-a716-446655440000",
		ActionTimestamp: "2023-10-01 10:00:00",
	}

	assert.Equal(t, expectedResponse, resp)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestGetAmmoHistoryByDate(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	queryBuilder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	ammosSt := &AmmosSt{db: db, queryBuilder: queryBuilder}

	req := &genprotos.GetAmmoHistoryByDateRequest{
		Date: "2023-10-01",
	}

	columns := []string{"id", "ammo_id", "action", "actor_id", "action_timestamp"}
	rows := sqlmock.NewRows(columns).
		AddRow("660e8400-e29b-41d4-a716-446655440000", "550e8400-e29b-41d4-a716-446655440000", "Created", "770e8400-e29b-41d4-a716-446655440000", "2023-10-01 10:00:00").
		AddRow("660e8400-e29b-41d4-a716-446655440001", "550e8400-e29b-41d4-a716-446655440000", "Updated", "770e8400-e29b-41d4-a716-446655440001", "2023-10-01 11:00:00")

	mock.ExpectQuery(`SELECT id, ammo_id, action, actor_id, action_timestamp FROM ammo_history WHERE DATE\(action_timestamp\) = \$1`).
		WithArgs(req.Date).
		WillReturnRows(rows)

	resp, err := ammosSt.GetAmmoHistoryByDate(ctx, req)
	require.NoError(t, err)

	expectedResponse := &genprotos.GetAmmoHistoryByDateResponse{
		AmmoHistory: []*genprotos.AmmoHistory{
			{
				Id:              "660e8400-e29b-41d4-a716-446655440000",
				AmmoId:          "550e8400-e29b-41d4-a716-446655440000",
				Action:          "Created",
				ActorId:         "770e8400-e29b-41d4-a716-446655440000",
				ActionTimestamp: "2023-10-01 10:00:00",
			},
			{
				Id:              "660e8400-e29b-41d4-a716-446655440001",
				AmmoId:          "550e8400-e29b-41d4-a716-446655440000",
				Action:          "Updated",
				ActorId:         "770e8400-e29b-41d4-a716-446655440001",
				ActionTimestamp: "2023-10-01 11:00:00",
			},
		},
	}

	assert.Equal(t, expectedResponse, resp)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestUpdateAmmoHistoryById(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	queryBuilder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	ammosSt := &AmmosSt{db: db, queryBuilder: queryBuilder}

	req := &genprotos.UpdateAmmoHistoryByIdRequest{
		Id:      "660e8400-e29b-41d4-a716-446655440000",
		AmmoId:  "550e8400-e29b-41d4-a716-446655440000",
		Action:  "Updated",
		ActorId: "770e8400-e29b-41d4-a716-446655440001",
	}

	actionTimestamp := time.Now()

	mock.ExpectExec(`UPDATE ammo_history SET ammo_id = \$1, action = \$2, actor_id = \$3, action_timestamp = \$4 WHERE id = \$5`).
		WithArgs(req.AmmoId, req.Action, req.ActorId, sqlmock.AnyArg(), req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := ammosSt.UpdateAmmoHistoryById(ctx, req)
	require.NoError(t, err)

	expectedResponse := &genprotos.AmmoHistory{
		Id:              req.Id,
		AmmoId:          req.AmmoId,
		Action:          req.Action,
		ActorId:         req.ActorId,
		ActionTimestamp: actionTimestamp.Format("2006-01-02 15:04:05"),
	}

	assert.Equal(t, expectedResponse.Id, resp.Id)
	assert.Equal(t, expectedResponse.AmmoId, resp.AmmoId)
	assert.Equal(t, expectedResponse.Action, resp.Action)
	assert.Equal(t, expectedResponse.ActorId, resp.ActorId)
	assert.Equal(t, expectedResponse.ActionTimestamp, resp.ActionTimestamp)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestDeleteAmmoHistoryById(t *testing.T) {
    db, mock, err := sqlmock.New()
    require.NoError(t, err)
    defer db.Close()

    ctx := context.Background()
    queryBuilder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
    ammosSt := &AmmosSt{db: db, queryBuilder: queryBuilder}

    req := &genprotos.DeleteAmmoHistoryByIdRequest{
        Id: "660e8400-e29b-41d4-a716-446655440000",
    }

    mock.ExpectExec(`DELETE FROM ammo_history WHERE id = \$1`).
        WithArgs(req.Id).
        WillReturnResult(sqlmock.NewResult(1, 1))

    resp, err := ammosSt.DeleteAmmoHistoryById(ctx, req)
    require.NoError(t, err)

    expectedResponse := &genprotos.Empty{}

    assert.Equal(t, expectedResponse, resp)

    err = mock.ExpectationsWereMet()
    require.NoError(t, err)
}

func TestGetAmmoHistory(t *testing.T) {
    db, mock, err := sqlmock.New()
    require.NoError(t, err)
    defer db.Close()

    ctx := context.Background()
    queryBuilder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
    ammosSt := &AmmosSt{db: db, queryBuilder: queryBuilder}

    req := &genprotos.Empty{}

    columns := []string{"id", "ammo_id", "action", "actor_id", "action_timestamp"}
    rows := sqlmock.NewRows(columns).
        AddRow("660e8400-e29b-41d4-a716-446655440000", "550e8400-e29b-41d4-a716-446655440000", "Created", "770e8400-e29b-41d4-a716-446655440000", "2023-10-01 10:00:00").
        AddRow("660e8400-e29b-41d4-a716-446655440001", "550e8400-e29b-41d4-a716-446655440000", "Updated", "770e8400-e29b-41d4-a716-446655440001", "2023-10-01 11:00:00")

    mock.ExpectQuery(`SELECT id, ammo_id, action, actor_id, action_timestamp FROM ammo_history`).
        WillReturnRows(rows)

    resp, err := ammosSt.GetAmmoHistory(ctx, req)
    require.NoError(t, err)

    expectedResponse := &genprotos.GetAmmoHistoryResponse{
        AmmoHistory: []*genprotos.AmmoHistory{
            {
                Id:              "660e8400-e29b-41d4-a716-446655440000",
                AmmoId:          "550e8400-e29b-41d4-a716-446655440000",
                Action:          "Created",
                ActorId:         "770e8400-e29b-41d4-a716-446655440000",
                ActionTimestamp: "2023-10-01 10:00:00",
            },
            {
                Id:              "660e8400-e29b-41d4-a716-446655440001",
                AmmoId:          "550e8400-e29b-41d4-a716-446655440000",
                Action:          "Updated",
                ActorId:         "770e8400-e29b-41d4-a716-446655440001",
                ActionTimestamp: "2023-10-01 11:00:00",
            },
        },
    }

    assert.Equal(t, expectedResponse, resp)

    err = mock.ExpectationsWereMet()
    require.NoError(t, err)
}
