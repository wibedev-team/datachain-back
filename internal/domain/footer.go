package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/wibedev-team/datachain-back/internal/domain/footer"
)

func NewFooter(engine *gin.Engine, pgClient *pgxpool.Pool) {
	footerGroup := engine.Group("/footer")
	footerStorage := footer.NewStorage(pgClient)
	footerHandler := footer.NewHandler(footerGroup, footerStorage)
	footerHandler.Register()
}
