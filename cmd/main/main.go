package main

import (
	"context"
	"log"
	"os"

	"github.com/wibedev-team/datachain-back/internal/config"
	"github.com/wibedev-team/datachain-back/pkg/db/postgresql"
	"github.com/wibedev-team/datachain-back/pkg/minio"
)

func main() {
	ctx := context.Background()

	args := os.Args
	if len(args) != 2 {
		log.Fatalf("provide path to congig file")
	}

	cfg := config.New(args[1])

	pgCfg := postgresql.NewConfig(
		cfg.Postgresql.Username,
		cfg.Postgresql.Password,
		cfg.Postgresql.Host,
		cfg.Postgresql.Port,
		cfg.Postgresql.Database,
	)

	pgClient := postgresql.New(ctx, pgCfg)
	_ = pgClient

	minioCfg := minio.NewConfig(
		cfg.Minio.Host,
		cfg.Minio.Port,
		cfg.Minio.AccessKeyID,
		cfg.Minio.SecretAccessKey,
		cfg.Minio.BucketName,
	)

	minioClient := minio.New(ctx, minioCfg)
	_ = minioClient
}
