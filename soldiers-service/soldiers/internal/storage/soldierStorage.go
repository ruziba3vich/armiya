package storage

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/ruziba3vich/armiya/soldies-service/genprotos"
	"github.com/ruziba3vich/armiya/soldies-service/internal/models"
)

type SoldiersStorage struct {
	db     *sql.DB
	logger *log.Logger
	sqrl   sq.StatementBuilderType
}

func NewSoldiersStorage(db *sql.DB, logger *log.Logger, sqrl sq.StatementBuilderType) *SoldiersStorage {
	return &SoldiersStorage{
		db:     db,
		logger: logger,
		sqrl:   sqrl,
	}
}

func (s *SoldiersStorage) CreateSoldier(ctx context.Context, req *genprotos.CreateSoldierRequest) (*genprotos.Soldier, error) {
	generatedId := uuid.New().String()
	query, args, err := s.sqrl.Insert("soldiers").
		Columns(
			"soldier_id",
			"soldier_name",
			"soldier_surname",
			"birth_date",
			"join_date",
			"leave_date",
			"group_id",
			"completed",
			"created_by").
		Values(
			generatedId,
			req.GetSoldierName(),
			req.GetSoldierSurname(),
			req.GetBirthDate().AsTime(),
			req.GetJoinDate().AsTime(),
			req.GetLeaveDate().AsTime(),
			req.GetGroupId(),
			false,
			req.GetCreatedBy()).
		ToSql()

	if err != nil {
		s.logger.Println("ERROR IN GENERATING QUERY TO CREATE A NEW SOLDIER :", err)
		return nil, err
	}

	res, err := s.db.ExecContext(ctx, query, args...)
	if err != nil {
		s.logger.Println("ERROR IN EXECUTING QUERY TO CREATE A NEW SOLDIER :", err)
		return nil, err
	}
	if ra, err := res.RowsAffected(); err != nil || ra == 0 {
		s.logger.Println("ERROR IN GETTING AFFECTED ROWS TO CREATE A NEW SOLDIER :", err)
		return nil, err
	}
	return &genprotos.Soldier{
		SoldierId:      generatedId,
		SoldierName:    req.GetSoldierName(),
		SoldierSurname: req.GetSoldierSurname(),
		BirthDate:      req.GetBirthDate(),
		JoinDate:       req.GetJoinDate(),
		LeaveDate:      req.GetLeaveDate(),
		GroupId:        req.GetGroupId(),
		Completed:      false,
		CreatedBy:      req.GetCreatedBy(),
	}, nil
}

func (s *SoldiersStorage) UpdateSoldier(ctx context.Context, req *genprotos.UpdateSoldierRequest) (*genprotos.UpdateOrGetSoldierResponse, error) {
	query := "UPDATE soldiers SET "
	var updatedFieldsCount int16
	if len(req.GetSoldierName()) > 0 {
		query += fmt.Sprintf("soldier_name = %s", req.GetSoldierName())
		updatedFieldsCount++
	}
	if len(req.GetSoldierSurname()) > 0 {
		query += fmt.Sprintf("soldier_surname = %s", req.GetSoldierSurname())
		updatedFieldsCount++
	}
	if len(req.BirthDate.String()) > 0 {
		query += fmt.Sprintf("birth_date = %x", req.GetBirthDate())
		updatedFieldsCount++
	}
	if len(req.GetJoinDate().String()) > 0 {
		query += fmt.Sprintf("join_date = %x", req.GetJoinDate())
		updatedFieldsCount++
	}
	if len(req.GetLeaveDate().String()) > 0 {
		query += fmt.Sprintf("leave_date = %x", req.GetLeaveDate())
		updatedFieldsCount++
	}
	if len(req.GetGroupId()) > 0 {
		query += fmt.Sprintf("group_id = %x", req.GetGroupId())
		updatedFieldsCount++
	}

	query += fmt.Sprintf(" WHERE soldier_id = %s", req.GetSoldierId())

	if updatedFieldsCount > 0 {
		res, err := s.db.ExecContext(ctx, query)
		if err != nil {
			s.logger.Println("ERROR WHILE RUNNING UPDATE QUERY FOR SOLDIER :", err)
			return nil, err
		}
		if ra, err := res.RowsAffected(); ra == 0 || err != nil {
			s.logger.Println("ERROR WHILE GETTING NUMBER OF AFFECTED ROWS IN UPDATE QUERY FOR SOLDIER :", err)
			return nil, err
		}
		return s.GetSoldierById(ctx,
			&genprotos.GetByIdRequest{
				SoldierId: req.GetSoldierId(),
			})
	}
	return nil, fmt.Errorf("no values given to update soldier")
}

func (s *SoldiersStorage) GetSoldierById(ctx context.Context, req *genprotos.GetByIdRequest) (*genprotos.UpdateOrGetSoldierResponse, error) {
	query := `
			SELECT
				s.soldier_id,
				s.soldier_name,
				s.soldier_surname,
				s.birth_date,
				s.join_date,
				s.leave_date,
				g.group_name,
				s.completed
				s.created_by
			FROM
				soldiers s
			INNER JOIN
				groups g
			ON
				g.group_id = s.group_id
			WHERE
				soldier_id = $1`

	row := s.db.QueryRowContext(ctx, query, req.GetSoldierId())
	var response genprotos.UpdateOrGetSoldierResponse

	if err := row.Scan(
		&response.SoldierId,
		&response.SoldierName,
		&response.SoldierSurname,
		&response.BirthDate,
		&response.JoinDate,
		&response.LeaveDate,
		&response.GroupName,
		&response.Completed,
		&response.CreatedBy,
	); err != nil {
		s.logger.Println("ERROR WHILE SCANNING VALUES FROM RESPONSE IN GetSoldierById SERVICE")
		return nil, err
	}
	return &response, nil
}

func (s *SoldiersStorage) GetSoldiersByName(ctx context.Context, req *genprotos.GetByName) (*genprotos.GetSoldiersResponse, error) {
	resquest := models.GetByFieldRequest{
		RequiredField: "soldier_name",
		Value:         req.GetSoldierName(),
	}
	return s.getByField(ctx, resquest)
}

func (s *SoldiersStorage) GetSoldiersBySurname(ctx context.Context, req *genprotos.GetBySurname) (*genprotos.GetSoldiersResponse, error) {
	resquest := models.GetByFieldRequest{
		RequiredField: "soldier_surname",
		Value:         req.GetSoldierSurname(),
	}
	return s.getByField(ctx, resquest)
}

func (s *SoldiersStorage) GetSoldiersByGroupName(ctx context.Context, req *genprotos.GetByGroupName) (*genprotos.GetSoldiersResponse, error) {
	resquest := models.GetByFieldRequest{
		RequiredField: "group_name",
		Value:         req.GetGroupName(),
	}
	return s.getByField(ctx, resquest)
}

func (s *SoldiersStorage) GetAllSoldiers(ctx context.Context, req *genprotos.GetAllSoldiersRequest) (*genprotos.GetSoldiersResponse, error) {
	query, args, err := s.sqrl.Select("soldiers s").
		Columns(
			"s.soldier_id",
			"s.soldier_name",
			"s.soldier_surname",
			"s.birth_date",
			"s.join_date",
			"s.leave_date",
			"g.group_name",
			"s.completed",
			"s.created_by",
		).
		InnerJoin("groups g ON g.group_id = s.group_id").
		OrderBy("s.soldier_id").
		Offset(uint64(req.GetOffsetValue())).
		Limit(uint64(req.GetLimit())).
		ToSql()

	if err != nil {
		s.logger.Println("ERROR OCCURED WHILE CREATING QUERY FOR GetAllSoldiers SERVICE :", err)
		return nil, err
	}

	rows, err := s.db.QueryContext(ctx, query, args)
	if err != nil {
		s.logger.Println("ERROR OCCURED WHILE QUERYING THE DATABASE FOR GetAllSoldiers SERVICE :", err)
		return nil, err
	}

	var response genprotos.GetSoldiersResponse

	for rows.Next() {
		var responsecha genprotos.UpdateOrGetSoldierResponse
		if err := rows.Scan(
			&responsecha.SoldierId,
			&responsecha.SoldierName,
			&responsecha.SoldierSurname,
			&responsecha.BirthDate,
			&responsecha.JoinDate,
			&responsecha.LeaveDate,
			&responsecha.GroupName,
			&responsecha.Completed,
			&responsecha.CreatedBy,
		); err != nil {
			s.logger.Println("ERROR WHILE SCANNING VALUES FROM RESPONSE IN GET ALL SOLDIERS SERVICE")
			return nil, err
		}
		response.Soldiers = append(response.Soldiers, &responsecha)
	}

	if err := rows.Err(); err != nil {
		s.logger.Println("ERROR FOUND IN CATCHING ERROR IN GET ALL SOLDIERS SERVICE")
		return nil, err
	}
	return &response, nil
}

func (s *SoldiersStorage) GetSoldiersByAge(ctx context.Context, req *genprotos.GetByAgeRequest) (*genprotos.GetSoldiersResponse, error) {

	return s.getByField(ctx, models.GetByFieldRequest{
		RequiredField: "age",
		Value:         strconv.Itoa(int(req.GetSoldierAge())),
	})
}

func (s *SoldiersStorage) DeleteSoldier(ctx context.Context, req *genprotos.DeleteRequest) (*genprotos.DeleteSoldierResponse, error) {
	query, args, err := s.sqrl.Update("soldiers").
		Set("deleted_by", req.GetDeletedBy()).
		Set("deleted", true).
		Where(sq.Eq{"soldier_id": req.GetSoldierId()}).
		ToSql()
	if err != nil {
		s.logger.Println("ERROR WHILE BUILDING QUERY IN DeleteSoldier SERVICE")
		return nil, err
	}

	res, err := s.db.ExecContext(ctx, query, args...)
	if err != nil {
		s.logger.Println("ERROR OCCURED WHILE EXECUTING QUERY TO THE DB IN DeleteSoldier :", err)
		return nil, err
	}

	if ra, err := res.RowsAffected(); ra == 0 || err != nil {
		s.logger.Println("ERROR OCCURED WHILE GETTING AFFECTED ROWS IN DeleteSoldier :", err)
		return nil, err
	}
	return &genprotos.DeleteSoldierResponse{
		Response: "SOLDIER HAS SUCCESSFULLY BEEN DELETED",
	}, nil
}

func (s *SoldiersStorage) MoveSoldierFromGroupAToGroupB(ctx context.Context, req *genprotos.MoveSoldierRequest) (*genprotos.MoveSoldierResponse, error) {
	query, args, err := s.sqrl.Update("soldiers").
		Set("group_id", req.GetNewGroupId()).
		Where(sq.Eq{"soldier_id": req.GetSoldierId()}).
		ToSql()

	if err != nil {
		s.logger.Println("ERROR WHILE BULDING QUERY IN MoveSoldierFromGroupAToGroupB SERVICE :", err)
		return nil, err
	}
	res, err := s.db.ExecContext(ctx, query, args...)

	if err != nil {
		s.logger.Println("ERROR WHILE QUERYING TO THE DB IN MoveSoldierFromGroupAToGroupB SERVICE :", err)
		return nil, err
	}
	if ra, err := res.RowsAffected(); ra == 0 || err != nil {
		s.logger.Println("ERROR WHILE CATCHING AFFECTED ROWS IN MoveSoldierFromGroupAToGroupB SERVICE :", err)
		return nil, err
	}

	soldier, err := s.GetSoldierById(ctx, &genprotos.GetByIdRequest{
		SoldierId: req.GetSoldierId(),
	})
	if err != nil {
		s.logger.Println("ERROR WHILE GETTING SOLDIER DATA IN MoveSoldierFromGroupAToGroupB SERVICE :", err)
		return nil, err
	}

	return &genprotos.MoveSoldierResponse{
		Message: "soldier has successfully been moved",
		Soldier: soldier,
	}, nil
}

func (s *SoldiersStorage) getByField(ctx context.Context, req models.GetByFieldRequest) (*genprotos.GetSoldiersResponse, error) {
	query := `
			SELECT
				s.soldier_id,
				s.soldier_name,
				s.soldier_surname,
				s.birth_date,
				s.join_date,
				s.leave_date,
				g.group_name,
				s.completed,
				s.created_by,
				EXTRACT(YEAR FROM AGE($1, s.birth_date)) AS years_difference
			FROM
				soldiers s
			INNER JOIN
				groups g
			ON
				g.group_id = s.group_id
			WHERE s.deleted = false AND 
		`
	if req.RequiredField == "group_name" {
		query += fmt.Sprintf("g.%s = $2;", req.RequiredField)
	} else if req.RequiredField == "age" {
		query += fmt.Sprintf("s.%s = $2 AND EXTRACT(YEAR FROM AGE(NOW(), s.birth_date)) = $3;", req.RequiredField)
	} else {
		query += fmt.Sprintf("s.%s = $2;", req.RequiredField)
	}
	rows, err := s.db.QueryContext(ctx, query, time.Now(), req.Value)
	if err != nil {
		s.logger.Printf("ERROR WHILE GETTING RESPONSE FROM DATABASE IN GET BY %s SERVICE\n", req.RequiredField)
		return nil, err
	}

	defer rows.Close()

	var response genprotos.GetSoldiersResponse

	for rows.Next() {
		var responsecha genprotos.UpdateOrGetSoldierResponse
		if err := rows.Scan(
			&responsecha.SoldierId,
			&responsecha.SoldierName,
			&responsecha.SoldierSurname,
			&responsecha.BirthDate,
			&responsecha.JoinDate,
			&responsecha.LeaveDate,
			&responsecha.GroupName,
			&responsecha.Completed,
			&responsecha.CreatedBy,
			&responsecha.SoldierAge,
		); err != nil {
			s.logger.Printf("ERROR WHILE SCANNING VALUES FROM RESPONSE IN GET BY %s SERVICE\n", req.RequiredField)
			return nil, err
		}
		response.Soldiers = append(response.Soldiers, &responsecha)
	}

	if err := rows.Err(); err != nil {
		s.logger.Printf("ERROR FOUND IN CATCHING ERROR IN GET BY %s SERVICE\n", req.RequiredField)
		return nil, err
	}
	return &response, nil
}

/*
	CREATE TABLE IF NOT EXISTS soldiers (
		soldier_id UUID PRIMARY KEY,
		soldier_name VARCHAR(64),
		soldier_surname VARCHAR(64),
		birth_date TIMESTAMP,
		join_date DATE,
		leave_date DATE,
		group_id UUID REFERENCES groups(group_id),
		completed BOOLEAN,
		created_by UUID REFERENCES objects(object_id)
	);
*/

/*
	type SoldierServiceServer interface {
		// CreateSoldier(context.Context, *CreateSoldierRequest) (*Soldier, error)
		// UpdateSoldier(context.Context, *UpdateSoldierRequest) (*UpdateOrGetSoldierResponse, error)
		// GetSoldierById(context.Context, *GetByIdRequest) (*UpdateOrGetSoldierResponse, error)
		// GetSoldiersByName(context.Context, *GetByName) (*GetSoldiersResponse, error)
		// GetSoldiersBySurname(context.Context, *GetBySurname) (*GetSoldiersResponse, error)
		// GetSoldiersByGroupName(context.Context, *GetByGroupName) (*GetSoldiersResponse, error)
		// GetAllSoldiers(context.Context, *GetAllSoldiersRequest) (*GetSoldiersResponse, error)
		// GetSoldiersByAge(context.Context, *GetByAgeRequest) (*GetSoldiersResponse, error)
		// DeleteSoldier(context.Context, *DeleteRequest) (*Soldier, error)
		mustEmbedUnimplementedSoldierServiceServer()
	}
*/
