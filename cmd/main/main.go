package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/wibedev-team/datachain-back/internal/config"
	"github.com/wibedev-team/datachain-back/internal/domain"
	"github.com/wibedev-team/datachain-back/pkg/db/postgresql"
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

	if os.Getenv("POSTGRES_HOST") == "" {
		cfg = config.New(args[1])

		pgCfg = postgresql.NewConfig(
			cfg.Postgresql.Username,
			cfg.Postgresql.Password,
			cfg.Postgresql.Host,
			cfg.Postgresql.Port,
			cfg.Postgresql.Database,
		)
	} else {
		pgCfg = postgresql.NewConfig(
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_DB"),
		)
	}

	pgClient := postgresql.New(ctx, pgCfg)

	engine := gin.Default()
	engine.Use(CORSMiddleware())
	engine.Static("/static", "./static")

	domain.NewAuth(engine, pgClient)
	domain.NewAboutUs(engine, pgClient)
	domain.NewStack(engine, pgClient)
	domain.NewSolution(engine, pgClient)
	domain.NewTeam(engine, pgClient)
	domain.NewFooter(engine, pgClient)

	log.Fatal(engine.RunTLS(":8000", "admin.data-chainz.ru.crt", "admin.data-chainz.ru.key"))
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
