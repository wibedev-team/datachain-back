package team

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wibedev-team/datachain-back/pkg/jwt"
	"log"
	"net/http"
	"os"

	"github.com/wibedev-team/datachain-back/internal/models"
)

type handler struct {
	router  *gin.RouterGroup
	storage Storage
}

type Storage interface {
	SaveTeammate(ctx context.Context, dto models.Team) error
	GetAllTeammates(ctx context.Context) ([]models.Team, error)
	RemoveTeammate(ctx context.Context, id string) error
}

func NewHandler(r *gin.RouterGroup, s Storage) *handler {
	return &handler{
		router:  r,
		storage: s,
	}
}

func (h *handler) Register() {
	h.router.Handle(http.MethodPost, "/create", h.createTeammate)
	h.router.Handle(http.MethodGet, "/get", h.getTeammates)
	h.router.Handle(http.MethodDelete, "/:id", h.removeTeammate)
}

func (h *handler) createTeammate(c *gin.Context) {
	img, err := c.FormFile("image")
	img.Filename = uuid.New().String() + ".png"
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	adminRole, err := jwt.CheckAdminRole(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	if !adminRole {
		c.JSON(http.StatusUnauthorized, jwt.ErrorNotAdmin.Error())
		return
	}

	var dto struct {
		Name     string `form:"name"`
		Position string `form:"position"`
		Link     string `form:"link"`
	}
	err = c.ShouldBind(&dto)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "decode error",
		})
	}
	log.Println(dto.Name, dto.Position, dto.Link)

	var teamDto models.Team

	err = c.SaveUploadedFile(img, "static/"+img.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	teamDto.Name = dto.Name
	teamDto.Position = dto.Position
	teamDto.Link = dto.Link
	teamDto.Img = img.Filename

	err = h.storage.SaveTeammate(c, teamDto)
	if err != nil {
		return
	}
	log.Println(err)

	c.JSON(http.StatusOK, gin.H{
		"teammate": teamDto,
	})
}

func (h *handler) getTeammates(c *gin.Context) {
	teammates, err := h.storage.GetAllTeammates(c)
	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"teammates": teammates,
	})
}

func (h *handler) removeTeammate(c *gin.Context) {
	id := c.Param("id")

	adminRole, err := jwt.CheckAdminRole(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	if !adminRole {
		c.JSON(http.StatusUnauthorized, jwt.ErrorNotAdmin.Error())
		return
	}

	err = h.storage.RemoveTeammate(c, id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "server error",
		})
		return
	}

	err = os.Remove("static/" + id)
	if err != nil {
		log.Println(err)
	}
}
