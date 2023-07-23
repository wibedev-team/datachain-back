package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/wibedev-team/datachain-back/internal/domain/aboutus"
)

func NewAboutUs(engine *gin.Engine, pgClient *pgxpool.Pool) {
	aboutGroup := engine.Group("/about")
	aboutStorage := aboutus.NewStorage(pgClient)
	aboutHandler := aboutus.NewHandler(aboutGroup, aboutStorage)
	aboutHandler.Register()
}
