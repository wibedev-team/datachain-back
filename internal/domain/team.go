package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/minio/minio-go/v7"

	"github.com/wibedev-team/datachain-back/internal/domain/team"
)

func NewTeam(engine *gin.Engine, pgClient *pgxpool.Pool, minio *minio.Client) {
	teamGroup := engine.Group("/team")
	teamStorage := team.NewStorage(pgClient)
	teamHandler := team.NewHandler(teamGroup, teamStorage, minio)
	teamHandler.Register()
}
