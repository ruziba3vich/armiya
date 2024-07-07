package storage

import (
	"context"
	"database/sql"
	"log"

	"github.com/ruziba3vich/countries/genprotos"
	"github.com/ruziba3vich/countries/internal/config"

	sq "github.com/Masterminds/squirrel"
)

type (
	CountrySt struct {
		db           *sql.DB
		queryBuilder sq.StatementBuilderType
	}
)

func New(config *config.Config) (*CountrySt, error) {

	db, err := ConnectDB(*config)
	if err != nil {
		return nil, err
	}

	return &CountrySt{
		db:           db,
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}, nil
}

func (s *CountrySt) CreateEquipment(ctx context.Context, req *genprotos.Equipment) error {
	query, args, err := s.queryBuilder.Insert("countries").
		Columns("country_name", "latitude", "longitude").
		Values(
			req.CountryName,
			req.Latitude,
			req.Longitude).
		Suffix("RETURNING country_id, country_name, latitude, longitude").
		ToSql()
	if err != nil {
		log.Println("Error generating SQL:", err)
		return nil, err
	}

	row := s.db.QueryRowContext(ctx, query, args...)

	var response genprotos.Country

	if err := row.Scan(
		&response.CountryId,
		&response.CountryName,
		&response.Latitude,
		&response.Longitude); err != nil {
		log.Println("Scan error:", err)
		return nil, err
	}
	if err := row.Err(); err != nil {
		log.Println("Row error:", err)
		return nil, err
	}
	return &response, nil
}
