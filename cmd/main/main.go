package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/wibedev-team/datachain-back/internal/config"
	"github.com/wibedev-team/datachain-back/internal/domain"
	"github.com/wibedev-team/datachain-back/pkg/db/postgresql"
	"github.com/wibedev-team/datachain-back/pkg/minio"
)

func main() {
	ctx := context.Background()

	args := os.Args
	if len(args) != 2 {
		if os.Getenv("POSTGRES_DB") == "" {
			log.Fatalf("provide path to congig file")
		}
	}

	var cfg *config.Config
	var pgCfg *postgresql.PgConfig
	var minioCfg *minio.MinioCfg

	if os.Getenv("POSTGRES_HOST") == "" {
		cfg = config.New(args[1])

		pgCfg = postgresql.NewConfig(
			cfg.Postgresql.Username,
			cfg.Postgresql.Password,
			cfg.Postgresql.Host,
			cfg.Postgresql.Port,
			cfg.Postgresql.Database,
		)

		minioCfg = minio.NewConfig(
			cfg.Minio.Host,
			cfg.Minio.Port,
			cfg.Minio.AccessKeyID,
			cfg.Minio.SecretAccessKey,
			cfg.Minio.BucketName,
		)
	} else {
		pgCfg = postgresql.NewConfig(
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_DB"),
		)

		minioCfg = minio.NewConfig(
			os.Getenv("MINIO_HOST"),
			os.Getenv("MINIO_PORT"),
			os.Getenv("MINIO_ACCESS"),
			os.Getenv("MINIO_SECRET"),
			os.Getenv("MINIO_BUCKET"),
		)
	}

	pgClient := postgresql.New(ctx, pgCfg)
	minioClient := minio.New(ctx, minioCfg)

	engine := gin.Default()
	engine.Use(cors.Default())

	domain.NewAuth(engine, pgClient)
	domain.NewAboutUs(engine, pgClient, minioClient)
	domain.NewStack(engine, pgClient, minioClient)
	//domain.NewStack(engine, pgClient, minioClient)
	domain.NewTeam(engine, pgClient, minioClient)
	domain.NewFooter(engine, pgClient)

	log.Fatal(engine.Run(":8000"))
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
