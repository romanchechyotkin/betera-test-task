package main

import (
	"context"
	"github.com/romanchechyotkin/betera-test-task/internal/httpserver"
	"github.com/romanchechyotkin/betera-test-task/internal/nasa"
	"github.com/romanchechyotkin/betera-test-task/pkg/logger"
	"github.com/romanchechyotkin/betera-test-task/pkg/minio"
	"github.com/romanchechyotkin/betera-test-task/pkg/postgresql"
	"os"
)

// @title Swagger Documentation
// @version 1.0
// @description Betera test task in Gin Framework
// @host localhost:8080
func main() {
	log := logger.New(os.Stdout)
	log.Debug("app running")

	minioClient := minio.New(log)

	//pgConfig := postgresql.NewPgConfig(
	//	os.Getenv("POSTGRES_USER"),
	//	os.Getenv("POSTGRES_PASSWORD"),
	//	os.Getenv("POSTGRES_HOST"),
	//	os.Getenv("POSTGRES_PORT"),
	//	os.Getenv("POSTGRES_DB"),
	//)

	pgConfig := postgresql.NewPgConfig(
		"chechyotka",
		"5432",
		"localhost",
		"5432",
		"betera",
	)
	pgClient := postgresql.NewClient(context.Background(), log, pgConfig)

	nasaDomain := nasa.RegisterNasaDomain(log, pgClient, minioClient)

	httpserver.Run(log, nasaDomain)
}
