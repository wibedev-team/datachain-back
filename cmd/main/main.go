package main

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"

	"github.com/wibedev-team/datachain-back/internal/config"
	"github.com/wibedev-team/datachain-back/internal/domain"
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

	engine := gin.Default()
	engine.Use(cors.Default())

	domain.NewAuth(engine, pgClient)
	domain.NewAboutUs(engine, pgClient, minioClient)
	domain.NewStack(engine, pgClient, minioClient)
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
