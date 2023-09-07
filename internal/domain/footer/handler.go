package footer

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/wibedev-team/datachain-back/pkg/jwt"
	"log"
	"net/http"

	"github.com/wibedev-team/datachain-back/internal/models"
)

type handler struct {
	router  *gin.RouterGroup
	storage Storage
}

type Storage interface {
	CreateFooter(ctx context.Context, footer models.Footer) error
	GetFooter(ctx context.Context) (models.Footer, error)
}

func NewHandler(r *gin.RouterGroup, s Storage) *handler {
	return &handler{
		router:  r,
		storage: s,
	}
}

func (h *handler) Register() {
	h.router.POST("/create", h.createFooter)
	h.router.GET("/get", h.getFooter)
}

func (h *handler) createFooter(c *gin.Context) {
	var dto models.Footer

	adminRole, err := jwt.CheckAdminRole(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	if !adminRole {
		c.JSON(http.StatusUnauthorized, jwt.ErrorNotAdmin.Error())
		return
	}

	err = c.ShouldBindJSON(&dto)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "decode error",
		})
		return
	}

	err = h.storage.CreateFooter(c, dto)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"footer": dto,
	})
}

func (h *handler) getFooter(c *gin.Context) {
	footer, err := h.storage.GetFooter(c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"footer": footer,
	})
}
