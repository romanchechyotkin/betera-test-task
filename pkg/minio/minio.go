package minio

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/romanchechyotkin/betera-test-task/pkg/logger"
)

func New(log *slog.Logger) *minio.Client {
	host := os.Getenv("MINIO_HOST")
	port := os.Getenv("MINIO_PORT")
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY")
	secretAccessKey := os.Getenv("MINIO_SECRET_KEY")
	bucketName := os.Getenv("MINIO_BUCKET_NAME")

	endpoint := fmt.Sprintf("%s:%s", host, port)
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	})
	if err != nil {
		logger.Error(log, "error during minio init", err)
		os.Exit(1)
	}
	log.Debug("got minio client", slog.String("client", fmt.Sprintf("%#v", minioClient)))

	location := "BLR"
	ctx := context.Background()

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Debug("we already own bucket", slog.String("bucket name", bucketName))
		} else {
			logger.Error(log, "error during checking bucket for existing", err)
			os.Exit(1)
		}
	} else {
		log.Debug("bucket created successfully", slog.String("bucket name", bucketName))
	}

	return minioClient
}
