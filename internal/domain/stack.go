package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/minio/minio-go/v7"

	"github.com/wibedev-team/datachain-back/internal/domain/stack"
)

func NewStack(engine *gin.Engine, pgClient *pgxpool.Pool, minio *minio.Client) {
	stackGroup := engine.Group("/stack")
	stackStorage := stack.NewStorage(pgClient)
	stackHandler := stack.NewHandler(stackGroup, stackStorage, minio)
	stackHandler.Register()
}
