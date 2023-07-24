package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/wibedev-team/datachain-back/internal/domain/solutions"
)

func NewSolution(engine *gin.Engine, pgClient *pgxpool.Pool) {
	solutionGroup := engine.Group("/solution")
	solutionStorage := solutions.NewStorage(pgClient)
	solutionHandler := solutions.NewHandler(solutionGroup, solutionStorage)
	solutionHandler.Register()
}
