package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/wibedev-team/datachain-back/internal/domain/stack"
)

func NewStack(engine *gin.Engine, pgClient *pgxpool.Pool) {
	stackGroup := engine.Group("/stack")
	stackStorage := stack.NewStorage(pgClient)
	stackHandler := stack.NewHandler(stackGroup, stackStorage)
	stackHandler.Register()
}
