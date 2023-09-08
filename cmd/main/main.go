package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"

	"github.com/wibedev-team/datachain-back/internal/config"
	"github.com/wibedev-team/datachain-back/internal/domain"
	"github.com/wibedev-team/datachain-back/pkg/db/postgresql"
)

func main() {
	ctx := context.Background()

	pgCfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	pgClient, err := postgresql.New(ctx, pgCfg)
	if err != nil {
		log.Fatal(err)
	}

	engine := gin.Default()
	engine.Use(CORSMiddleware())
	engine.Static("/static", "./static")

	RegisterRoutes(engine, pgClient)

	log.Fatal(engine.RunTLS(":8000", "admin.data-chainz.ru.crt", "admin.data-chainz.ru.key"))
}

func RegisterRoutes(engine *gin.Engine, pgClient *pgxpool.Pool) {
	domain.NewAuth(engine, pgClient)
	domain.NewAboutUs(engine, pgClient)
	domain.NewStack(engine, pgClient)
	domain.NewSolution(engine, pgClient)
	domain.NewTeam(engine, pgClient)
	domain.NewFooter(engine, pgClient)
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
