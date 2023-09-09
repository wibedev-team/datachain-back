package stack

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
	CreateStackImage(ctx context.Context, img string) error
	GetAllStackImages(ctx context.Context) ([]models.Stack, error)
	RemoveStack(ctx context.Context, id string) error
}

func NewHandler(r *gin.RouterGroup, s Storage) *handler {
	return &handler{
		router:  r,
		storage: s,
	}
}

func (h *handler) Register() {
	h.router.POST("/create", h.createImage)
	h.router.GET("/all", h.getAllImages)
	h.router.DELETE("/:id", h.removeStack)

}

func (h *handler) createImage(c *gin.Context) {
	img, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	img.Filename = uuid.New().String() + ".png"

	adminRole, err := jwt.CheckAdminRole(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	if !adminRole {
		c.JSON(http.StatusUnauthorized, jwt.ErrorNotAdmin.Error())
		return
	}

	err = c.SaveUploadedFile(img, "static/"+img.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = h.storage.CreateStackImage(c, img.Filename)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"img": img.Filename,
	})
}

func (h *handler) getAllImages(c *gin.Context) {
	images, err := h.storage.GetAllStackImages(c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"stacks": images,
	})
}

func (h *handler) removeStack(c *gin.Context) {
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

	err = h.storage.RemoveStack(c, id)
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
