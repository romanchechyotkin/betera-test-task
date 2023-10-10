package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"

	_ "github.com/romanchechyotkin/betera-test-task/docs"
	"github.com/romanchechyotkin/betera-test-task/pkg/logger"
	"github.com/romanchechyotkin/betera-test-task/pkg/minio"
	"github.com/romanchechyotkin/betera-test-task/pkg/postgresql"
)

func main() {
	log := logger.New(os.Stdout)
	log.Debug("app running")

	minioClient := minio.New(log)
	_ = minioClient

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
	_ = pgClient

	go everyMinute()

	engine := gin.Default()
	engine.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))
	engine.GET("/health", health)

	err := engine.Run(":8080")
	if err != nil {
		logger.Error(log, "http server init failed", err)
		os.Exit(1)
	}
}

func everyMinute() {
	ticker := time.NewTicker(time.Second * 5)
	for {
		select {
		case <-ticker.C:
			log.Println("minute")
		}
	}
}

// @Summary Health Check
// @Description Checking health of backend
// @Produce application/json
// @Success 200 {string} health
// @Router /health [get]
func health(ctx *gin.Context) {
	ctx.String(http.StatusOK, "health")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, DELETE, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
