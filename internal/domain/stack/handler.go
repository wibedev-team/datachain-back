package stack

import (
	"context"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/wibedev-team/datachain-back/internal/models"
)

type handler struct {
	router  *gin.RouterGroup
	storage Storage
	minio   *minio.Client
}

type Storage interface {
	CreateStackImage(ctx context.Context, img string) error
	GetAllStackImages(ctx context.Context) ([]models.Stack, error)
	RemoveStack(ctx context.Context, id string) error
}

func NewHandler(r *gin.RouterGroup, s Storage, m *minio.Client) *handler {
	return &handler{
		router:  r,
		storage: s,
		minio:   m,
	}
}

func (h *handler) Register() {
	h.router.POST("/create", h.createImage)
	h.router.GET("/all", h.getAllImages)
	h.router.Handle(http.MethodDelete, "/:id", h.removeStack)

}

func (h *handler) createImage(c *gin.Context) {
	img, _ := c.FormFile("image")
	img.Filename = uuid.New().String() + ".png"

	err := c.SaveUploadedFile(img, "static/"+img.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	info, err := h.minio.FPutObject(c, "datachain", img.Filename, "static/"+img.Filename, minio.PutObjectOptions{ContentType: "image/png"})
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Printf("Successfully uploaded %s of size %d\n", img.Filename, info.Size)

	err = os.Remove("static/" + img.Filename)
	if err != nil {
		log.Println(err)
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

	err := h.storage.RemoveStack(c, id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "server error",
		})
		return
	}
}
