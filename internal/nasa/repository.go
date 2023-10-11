package nasa

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/romanchechyotkin/betera-test-task/pkg/logger"
	"github.com/romanchechyotkin/betera-test-task/pkg/postgresql"
)

var ErrNotFound = errors.New("not found")

type repository struct {
	log  *slog.Logger
	pool *pgxpool.Pool
}

func newRepository(logger *slog.Logger, pool *pgxpool.Pool) storage {
	return &repository{
		log:  logger,
		pool: pool,
	}
}

func (r *repository) saveAPOD(ctx context.Context, dto *Metadata) error {
	query := `
		INSERT INTO apods (title, explanation, image, media_type, service_version, date) 
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	r.log.Info("database query", slog.String("query", postgresql.FormatQuery(query)))
	exec, err := r.pool.Exec(ctx, query, dto.Title, dto.Explanation, dto.URL, dto.MediaType, dto.ServiceVersion, dto.Date)
	if err != nil {
		logger.Error(r.log, "error during execution", err)
		return err
	}
	r.log.Info("result of execution", slog.String("result", fmt.Sprintf("rows affected %d", exec.RowsAffected())))

	return nil
}

func (r *repository) getAllAPODs(ctx context.Context) ([]*Metadata, error) {
	query := `
		SELECT title, explanation, image, media_type, service_version, date FROM apods
	`

	r.log.Info("database query", slog.String("query", postgresql.FormatQuery(query)))
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		logger.Error(r.log, "error during query", err)
		return nil, err
	}
	defer rows.Close()

	var res []*Metadata
	for rows.Next() {
		var m Metadata

		err = rows.Scan(&m.Title, &m.Explanation, &m.URL, &m.MediaType, &m.ServiceVersion, &m.Date)
		if err != nil {
			logger.Error(r.log, "error during scanning", err)
			return nil, err
		}

		res = append(res, &m)
	}

	return res, nil
}
