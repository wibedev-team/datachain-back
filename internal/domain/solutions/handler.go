package solutions

import (
	"context"
	"github.com/google/uuid"
	"github.com/wibedev-team/datachain-back/pkg/jwt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/wibedev-team/datachain-back/internal/models"
)

type handler struct {
	router  *gin.RouterGroup
	storage Storage
}

type Storage interface {
	SaveSolution(ctx context.Context, dto models.Solution) error
	GetAllSolutions(ctx context.Context) ([]models.Solution, error)
	RemoveSolution(ctx context.Context, title string) error
}

func NewHandler(r *gin.RouterGroup, s Storage) *handler {
	return &handler{
		router:  r,
		storage: s,
	}
}

func (h *handler) Register() {
	h.router.Handle(http.MethodPost, "/create", h.createSolution)
	h.router.Handle(http.MethodGet, "/all", h.getAllSolutions)
	h.router.DELETE("/:id", h.removeSolution)
}

func (h *handler) createSolution(c *gin.Context) {
	img, _ := c.FormFile("file")

	authHeader := c.GetHeader("Authorization")
	headers := strings.Split(authHeader, " ")
	log.Println(headers)
	token, err := jwt.ParseAccessToken(headers[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	role := token["role"]
	if role != "ADMIN" {
		c.JSON(http.StatusUnauthorized, "wrong role")
		return
	}

	log.Println(img.Filename)
	img.Filename = uuid.New().String() + ".png"
	err = c.SaveUploadedFile(img, "static/"+img.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var dto struct {
		Title    string `form:"title"`
		Features string `form:"features"`
		Link     string `form:"link"`
	}
	err = c.ShouldBind(&dto)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "decode error",
		})
		return
	}

	var solution models.Solution
	solution.Title = dto.Title
	solution.Link = dto.Link
	solution.File = img.Filename

	split := strings.Split(dto.Features, "\n")
	features := make([]models.Feature, len(split), len(split))
	for i, s := range split {
		var f models.Feature
		f.Text = s
		features[i] = f
	}
	solution.Features = features

	err = h.storage.SaveSolution(c, solution)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

}

func (h *handler) getAllSolutions(c *gin.Context) {
	solutions, err := h.storage.GetAllSolutions(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, solutions)
}

func (h *handler) removeSolution(c *gin.Context) {
	title := c.Param("id")

	authHeader := c.GetHeader("Authorization")
	headers := strings.Split(authHeader, " ")
	log.Println(headers)
	token, err := jwt.ParseAccessToken(headers[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	role := token["role"]
	if role != "ADMIN" {
		c.JSON(http.StatusUnauthorized, "wrong role")
		return
	}

	err = h.storage.RemoveSolution(c, title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
}
