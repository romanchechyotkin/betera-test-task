package nasa

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"

	"github.com/romanchechyotkin/betera-test-task/pkg/logger"
	"github.com/romanchechyotkin/betera-test-task/pkg/utils"
)

type storage interface {
	saveAPOD(ctx context.Context, dto *Metadata) error
	getAllAPODs(ctx context.Context) ([]*Metadata, error)
}

type handler struct {
	log         *slog.Logger
	repository  storage
	minioClient *minio.Client
}

func newHandler(logger *slog.Logger, repo storage, minioClient *minio.Client) *handler {
	h := &handler{
		log:         logger,
		repository:  repo,
		minioClient: minioClient,
	}

	go h.parseMetadata()

	return h
}

func (h *handler) RegisterRoutes(engine *gin.Engine) {
	group := engine.Group("/nasa")

	group.GET("/", h.getAllAPODs)
	group.GET("/health", h.index)
}

func (h *handler) getAllAPODs(ctx *gin.Context) {
	apods, err := h.repository.getAllAPODs(ctx)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, apods)
}

func (h *handler) parseMetadata() {
	ticker := time.NewTicker(time.Second * 5)
	day := 11
	for {
		select {
		case <-ticker.C:
			resp, err := http.Get(fmt.Sprintf("https://api.nasa.gov/planetary/apod?api_key=DEMO_KEY&date=2023-10-%d", day))
			if err != nil {
				logger.Error(h.log, "error during request to NASA API", err)
			}

			var dto Metadata
			err = json.NewDecoder(resp.Body).Decode(&dto)
			if err != nil {
				logger.Error(h.log, "error during json decoding", err)
			}

			h.log.Info("got response", slog.Any("metadata", dto))

			resp, err = http.Get(dto.URL)
			if err != nil {
				logger.Error(h.log, "error during request to get image", err)
			}

			file, err := os.Create("tmp/img.jpg")
			if err != nil {
				logger.Error(h.log, "failed to create image file", err)
			}
			defer file.Close()

			_, err = io.Copy(file, resp.Body)
			if err != nil {
				logger.Error(h.log, "failed to create image file", err)
			}

			fileName := uuid.New().String() + ".jpg"

			err = h.saveToMinio(context.Background(), h.minioClient, fileName, "tmp/img.jpg")
			if err != nil {
				logger.Error(h.log, "error during saving to minio", err)
			}

			dto.URL = fileName

			err = h.repository.saveAPOD(context.Background(), &dto)
			if err != nil {
				logger.Error(h.log, "error during json decoding", err)
			}

			err = utils.CleanTmp()
			if err != nil {
				logger.Error(h.log, "error during cleaning tmp dir", err)
			}
			day = day - 1
		}
	}
}

func (h *handler) saveToMinio(ctx context.Context, client *minio.Client, fileName, filePath string) error {
	info, err := client.FPutObject(ctx, "betera", fileName, filePath, minio.PutObjectOptions{ContentType: "image/jpg"})
	if err != nil {
		return err
	}

	h.log.Info("image successfully uploaded", slog.String("info", fmt.Sprintf("%s of size %d", fileName, info.Size)))

	return nil
}

// @Summary Nasa Endpoint Health Check
// @Description Checking health of nasa endpoint
// @Produce application/json
// @Success 200 {string} nasa
// @Router /nasa/health [get]
func (h *handler) index(ctx *gin.Context) {
	ctx.String(http.StatusOK, "nasa")
}
