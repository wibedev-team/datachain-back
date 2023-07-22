package aboutus

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"

	"github.com/wibedev-team/datachain-back/internal/models"
)

type handler struct {
	router  *gin.RouterGroup
	storage Storage
	minio   *minio.Client
}

type Storage interface {
	SaveSection(ctx context.Context, dto models.About) error
	GetSection(ctx context.Context) (models.About, error)
}

func NewHandler(r *gin.RouterGroup, s Storage, m *minio.Client) *handler {
	return &handler{
		router:  r,
		storage: s,
		minio:   m,
	}
}

func (h *handler) Register() {
	h.router.Handle(http.MethodPost, "/create", h.submitHandler)
	h.router.Handle(http.MethodGet, "/get", h.getAboutSection)
	//h.router.Handle(http.MethodGet, "/update", h.getAboutSection)
}

func (h *handler) submitHandler(c *gin.Context) {
	img, _ := c.FormFile("image")
	img.Filename = uuid.New().String() + ".png"

	var dto struct {
		Title       string `form:"title"`
		Description string `form:"description"`
	}
	err := c.ShouldBind(&dto)
	if err != nil {

	}
	fmt.Println(dto.Title, dto.Description, img.Filename)

	var aboutDto models.About

	err = c.SaveUploadedFile(img, "static/"+img.Filename)
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

	aboutDto.Title = dto.Title
	aboutDto.Description = dto.Description
	aboutDto.Img = img.Filename

	err = h.storage.SaveSection(c, aboutDto)
	if err != nil {
		return
	}
	log.Println(err)

	c.JSON(http.StatusOK, gin.H{
		"title":       aboutDto.Title,
		"description": aboutDto.Description,
		"img":         img.Filename,
	})
}

func (h *handler) getAboutSectionForEdit(c *gin.Context) {
	about, err := h.storage.GetSection(c)
	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title":       about.Title,
		"description": template.HTML(about.Description),
		"img":         about.Img,
	})
}

func (h *handler) getAboutSection(c *gin.Context) {
	about, err := h.storage.GetSection(c)
	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title":       about.Title,
		"description": about.Description,
		"img":         about.Img,
	})
}

//func (h *handler) createAboutUsSection(c *gin.Context) {
//	title := c.Request.FormValue("title")
//	description := c.Request.FormValue("description")
//	img, _ := c.FormFile("image")
//	fmt.Println(title, description, img.Filename)
//	c.HTML(http.StatusOK, "aboutedit.html", gin.H{})
//}
