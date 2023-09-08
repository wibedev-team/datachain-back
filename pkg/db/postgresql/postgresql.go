package postgresql

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PgConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

func NewConfig(username, password, host, port, database string) *PgConfig {
	return &PgConfig{
		Username: username,
		Password: password,
		Host:     host,
		Port:     port,
		Database: database,
	}
}

func New(ctx context.Context, cfg *PgConfig) (*pgxpool.Pool, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	log.Println(connString)

	log.Println("postgresql client init")
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to connect to postgresql; err: %v", err))
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to connect to postgresql; err: %v", err))
	}

	return pool, nil
}
