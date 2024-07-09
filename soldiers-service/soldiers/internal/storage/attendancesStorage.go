package storage

import (
	"context"
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/ruziba3vich/armiya/soldies-service/genprotos"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type AttendanceStorage struct {
	db     *sql.DB
	sqrl   sq.StatementBuilderType
	logger *log.Logger
}

func NewAttendanceStorage(db *sql.DB, logger *log.Logger, sqrl sq.StatementBuilderType) *AttendanceStorage {
	return &AttendanceStorage{
		db:     db,
		sqrl:   sqrl,
		logger: logger,
	}
}

func (s *AttendanceStorage) CreateAttendance(ctx context.Context, req *genprotos.CreateAttendanceRequest) (*genprotos.Attendance, error) {
	generatedId := uuid.New().String()
	eventTime := req.GetEventTime().AsTime()

	query, args, err := s.sqrl.Insert("attendance").
		Columns(
			"attendance_id",
			"soldier_id",
			"training_id",
			"event_time",
			"used_ammos",
			"used_ammo_type",
			"used_fuels",
			"created_by").
		Values(
			generatedId,
			req.GetSoldierId(),
			req.GetTrainingId(),
			eventTime,
			req.GetUsedAmmos(),
			req.GetUsedAmmoType(),
			req.GetUsedFuels(),
			req.GetCreatedBy()).
		ToSql()

	if err != nil {
		s.logger.Println("ERROR IN GENERATING QUERY TO CREATE ATTENDANCE:", err)
		return nil, err
	}

	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		s.logger.Println("ERROR IN EXECUTING QUERY TO CREATE ATTENDANCE:", err)
		return nil, err
	}

	return &genprotos.Attendance{
		AttendanceId: generatedId,
		SoldierId:    req.GetSoldierId(),
		TrainingId:   req.GetTrainingId(),
		EventTime:    timestamppb.New(eventTime),
		UsedAmmos:    req.GetUsedAmmos(),
		UsedAmmoType: req.GetUsedAmmoType(),
		UsedFuels:    req.GetUsedFuels(),
		CreatedBy:    req.GetCreatedBy(),
	}, nil
}

func (s *AttendanceStorage) AddAmmosToSoldier(ctx context.Context, req *genprotos.AddAmmosRequest) (*genprotos.Attendance, error) {
	query, args, err := s.sqrl.Update("attendance").
		Set("used_ammos", sq.Expr("used_ammos + ?", req.GetAmmosToAdd())).
		Set("used_fuels", sq.Expr("used_fuels + ?", req.GetFuelsToAdd())).
		Where(sq.Eq{"attendance_id": req.GetId()}).
		ToSql()

	if err != nil {
		s.logger.Println("ERROR IN GENERATING QUERY TO ADD AMMOS TO SOLDIER:", err)
		return nil, err
	}

	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		s.logger.Println("ERROR IN EXECUTING QUERY TO ADD AMMOS TO SOLDIER:", err)
		return nil, err
	}

	return s.GetAttendanceById(ctx, req.GetId())
}

func (s *AttendanceStorage) GetAttendanceByDate(ctx context.Context, req *genprotos.GetAttendanceByDateRequest) (*genprotos.GetAttendanceResponse, error) {
	query, args, err := s.sqrl.Select("attendance_id", "soldier_id", "training_id", "event_time", "used_ammos", "used_ammo_type", "used_fuels", "created_by", "deleted", "deleted_by").
		From("attendance").
		Where(sq.Eq{"DATE(event_time)": req.GetDate()}).
		ToSql()

	if err != nil {
		s.logger.Println("ERROR IN GENERATING QUERY TO GET ATTENDANCE BY DATE:", err)
		return nil, err
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		s.logger.Println("ERROR IN EXECUTING QUERY TO GET ATTENDANCE BY DATE:", err)
		return nil, err
	}
	defer rows.Close()

	var attendances []*genprotos.Attendance
	for rows.Next() {
		var a genprotos.Attendance
		var deletedBy sql.NullString

		err := rows.Scan(&a.AttendanceId, &a.SoldierId, &a.TrainingId, &a.EventTime, &a.UsedAmmos, &a.UsedAmmoType, &a.UsedFuels, &a.CreatedBy, &a.Deleted, &deletedBy)
		if err != nil {
			s.logger.Println("ERROR IN SCANNING ATTENDANCE ROW:", err)
			return nil, err
		}

		if deletedBy.Valid {
			a.DeletedBy = wrapperspb.String(deletedBy.String)
		}

		attendances = append(attendances, &a)
	}

	return &genprotos.GetAttendanceResponse{Attendance: attendances}, nil
}

func (s *AttendanceStorage) GetAllAttendanceBySoldierId(ctx context.Context, req *genprotos.GetAllAttendanceBySoldierIdRequest) (*genprotos.GetAttendanceResponse, error) {
	query, args, err := s.sqrl.Select("attendance_id", "soldier_id", "training_id", "event_time", "used_ammos", "used_ammo_type", "used_fuels", "created_by", "deleted", "deleted_by").
		From("attendance").
		Where(sq.Eq{"soldier_id": req.GetSoldierId()}).
		ToSql()

	if err != nil {
		s.logger.Println("ERROR IN GENERATING QUERY TO GET ALL ATTENDANCE BY SOLDIER ID:", err)
		return nil, err
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		s.logger.Println("ERROR IN EXECUTING QUERY TO GET ALL ATTENDANCE BY SOLDIER ID:", err)
		return nil, err
	}
	defer rows.Close()

	var attendances []*genprotos.Attendance
	for rows.Next() {
		var a genprotos.Attendance
		var deletedBy sql.NullString

		err := rows.Scan(&a.AttendanceId, &a.SoldierId, &a.TrainingId, &a.EventTime, &a.UsedAmmos, &a.UsedAmmoType, &a.UsedFuels, &a.CreatedBy, &a.Deleted, &deletedBy)
		if err != nil {
			s.logger.Println("ERROR IN SCANNING ATTENDANCE ROW:", err)
			return nil, err
		}

		if deletedBy.Valid {
			a.DeletedBy = wrapperspb.String(deletedBy.String)
		}

		attendances = append(attendances, &a)
	}

	return &genprotos.GetAttendanceResponse{Attendance: attendances}, nil
}

func (s *AttendanceStorage) GetSoldierAttendanceByDate(ctx context.Context, req *genprotos.GetSoldierAttendanceByDateRequest) (*genprotos.GetAttendanceResponse, error) {
	query, args, err := s.sqrl.Select("attendance_id", "soldier_id", "training_id", "event_time", "used_ammos", "used_ammo_type", "used_fuels", "created_by", "deleted", "deleted_by").
		From("attendance").
		Where(sq.And{
			sq.Eq{"soldier_id": req.GetSoldierId()},
			sq.Expr("DATE(event_time) = ?", req.GetDate()),
		}).
		ToSql()

	if err != nil {
		s.logger.Println("ERROR IN GENERATING QUERY TO GET SOLDIER ATTENDANCE BY DATE:", err)
		return nil, err
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		s.logger.Println("ERROR IN EXECUTING QUERY TO GET SOLDIER ATTENDANCE BY DATE:", err)
		return nil, err
	}
	defer rows.Close()

	var attendances []*genprotos.Attendance
	for rows.Next() {
		var a genprotos.Attendance
		var deletedBy sql.NullString

		err := rows.Scan(&a.AttendanceId, &a.SoldierId, &a.TrainingId, &a.EventTime, &a.UsedAmmos, &a.UsedAmmoType, &a.UsedFuels, &a.CreatedBy, &a.Deleted, &deletedBy)
		if err != nil {
			s.logger.Println("ERROR IN SCANNING ATTENDANCE ROW:", err)
			return nil, err
		}

		if deletedBy.Valid {
			a.DeletedBy = wrapperspb.String(deletedBy.String)
		}

		attendances = append(attendances, &a)
	}

	return &genprotos.GetAttendanceResponse{Attendance: attendances}, nil
}

func (s *AttendanceStorage) UpdateAttendanceBySoldierId(ctx context.Context, req *genprotos.UpdateRequest) (*genprotos.Attendance, error) {
	attendance := req.GetAttendance()
	eventTime := attendance.GetEventTime().AsTime()
	deletedBy := attendance.GetDeletedBy()

	builder := s.sqrl.Update("attendance").
		Set("soldier_id", attendance.GetSoldierId()).
		Set("training_id", attendance.GetTrainingId()).
		Set("event_time", eventTime).
		Set("used_ammos", attendance.GetUsedAmmos()).
		Set("used_ammo_type", attendance.GetUsedAmmoType()).
		Set("used_fuels", attendance.GetUsedFuels()).
		Set("deleted", attendance.GetDeleted())

	if deletedBy != nil {
		builder = builder.Set("deleted_by", deletedBy.GetValue())
	} else {
		builder = builder.Set("deleted_by", nil)
	}

	query, args, err := builder.Where(sq.Eq{"attendance_id": attendance.GetAttendanceId()}).ToSql()
	if err != nil {
		s.logger.Println("ERROR IN GENERATING QUERY TO UPDATE ATTENDANCE:", err)
		return nil, err
	}

	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		s.logger.Println("ERROR IN EXECUTING QUERY TO UPDATE ATTENDANCE:", err)
		return nil, err
	}

	return attendance, nil
}

func (s *AttendanceStorage) DeleteAttendance(ctx context.Context, req *genprotos.DeleteAttendanceRequest) (*genprotos.Attendance, error) {
	query, args, err := s.sqrl.Update("attendance").
		Set("deleted", true).
		Set("deleted_by", req.GetDeletedBy()).
		Where(sq.Eq{"attendance_id": req.GetAttendanceId()}).
		ToSql()

	if err != nil {
		s.logger.Println("ERROR IN GENERATING QUERY TO DELETE ATTENDANCE:", err)
		return nil, err
	}

	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		s.logger.Println("ERROR IN EXECUTING QUERY TO DELETE ATTENDANCE:", err)
		return nil, err
	}

	return s.GetAttendanceById(ctx, req.GetAttendanceId())
}

func (s *AttendanceStorage) GetAttendanceById(ctx context.Context, attendanceId string) (*genprotos.Attendance, error) {
	query, args, err := s.sqrl.Select("attendance_id", "soldier_id", "training_id", "event_time", "used_ammos", "used_ammo_type", "used_fuels", "created_by", "deleted", "deleted_by").
		From("attendance").
		Where(sq.Eq{"attendance_id": attendanceId}).
		ToSql()

	if err != nil {
		s.logger.Println("ERROR IN GENERATING QUERY TO GET ATTENDANCE BY ID:", err)
		return nil, err
	}

	var a genprotos.Attendance
	var deletedBy sql.NullString
	err = s.db.QueryRowContext(ctx, query, args...).Scan(&a.AttendanceId, &a.SoldierId, &a.TrainingId, &a.EventTime, &a.UsedAmmos, &a.UsedAmmoType, &a.UsedFuels, &a.CreatedBy, &a.Deleted, &deletedBy)
	if err != nil {
		s.logger.Println("ERROR IN SCANNING ATTENDANCE ROW:", err)
		return nil, err
	}

	if deletedBy.Valid {
		a.DeletedBy = wrapperspb.String(deletedBy.String)
	}

	return &a, nil
}
