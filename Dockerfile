FROM golang:alpine as builder
ENV POSTGRES_HOST="postgres" \
  POSTGRES_PORT="5432" \
  POSTGRES_DB="datachain" \
  POSTGRES_USER="postgres" \
  POSTGRES_PASSWORD="5432" \
  MINIO_HOST="postgres" \
  MINIO_PORT="9000" \
  MINIO_ACCESS="bC2fbyLxLUsUHMtqUvDx" \
  MINIO_SECRET="54rQ0EorX8bTLLo75xLn0lIeu9echhwQXEtwuOxhjA32" \
  MINIO_BUCKET="datachain"

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/bin cmd/main/main.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/bin .
EXPOSE 8000
CMD ["/app/bin"]