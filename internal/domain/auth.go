package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/wibedev-team/datachain-back/internal/domain/auth"
)

func NewAuth(engine *gin.Engine, pgClient *pgxpool.Pool) {
	authGroup := engine.Group("/auth")
	authStorage := auth.NewStorage(pgClient)
	authHandler := auth.NewHandler(authGroup, authStorage)
	authHandler.Register()
}
