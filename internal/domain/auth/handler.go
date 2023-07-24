package auth

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/wibedev-team/datachain-back/internal/models"
	"github.com/wibedev-team/datachain-back/pkg/jwt"
)

type handler struct {
	router  *gin.RouterGroup
	storage Storage
}

type Storage interface {
	FindUserByLogin(ctx context.Context, login string) (models.User, error)
}

func NewHandler(r *gin.RouterGroup, s Storage) *handler {
	return &handler{
		router:  r,
		storage: s,
	}
}

func (h *handler) Register() {
	h.router.Handle(http.MethodPost, "/login", h.login)
}

func (h *handler) login(c *gin.Context) {
	var dto struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "decode error",
		})
		return
	}

	log.Println(dto.Login, dto.Password)

	user, err := h.storage.FindUserByLogin(c, dto.Login)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "wrong login",
		})
		return
	}

	if dto.Password != user.Password {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "wrong password",
		})
		return
	}

	accessToken, err := jwt.GenerateAccessToken(user.Login, user.Role)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	refreshToken, err := jwt.GenerateRefreshToken(user.Login)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	c.SetCookie("refresh_token", refreshToken, 24*60*60*1000, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"login": user.Login,
		"token": accessToken,
	})
}
