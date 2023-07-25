FROM golang:alpine as builder
ENV POSTGRES_HOST="postgres" \
  POSTGRES_PORT="5432" \
  POSTGRES_DB="datachain" \
  POSTGRES_USER="postgres" \
  POSTGRES_PASSWORD="5432"

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/bin cmd/main/main.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/bin .
COPY --from=builder /app/certs/admin.data-chainz.ru.crt .
COPY --from=builder /app/certs/admin.data-chainz.ru.key .
EXPOSE 8000
CMD ["/app/bin"]