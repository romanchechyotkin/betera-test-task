FROM golang:alpine as builder
ENV MINIO_HOST="" \
    MINIO_PORT="" \
    MINIO_ACCESS_KEY="" \
    MINIO_SECRET_KEY="" \
    MINIO_BUCKET_NAME="" \
    POSTGRES_HOST="" \
    POSTGRES_PORT="" \
    POSTGRES_DB="" \
    POSTGRES_USER="" \
    POSTGRES_PASSWORD="" \
    ENVIRONMENT=""

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/bin cmd/main/main.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/bin .
EXPOSE 8080
CMD ["/app/bin"]