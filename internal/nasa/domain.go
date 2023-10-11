package nasa

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/minio/minio-go/v7"
)

func RegisterNasaDomain(logger *slog.Logger, pool *pgxpool.Pool, client *minio.Client) *handler {
	repo := newRepository(logger, pool)
	h := newHandler(logger, repo, client)
	return h
}
