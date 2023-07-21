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
	h.router.Handle(http.MethodGet, "/logout", h.logout)
	h.router.Handle(http.MethodGet, "/refresh", h.refresh)
}

func (h *handler) login(c *gin.Context) {
	var dto struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "decode error",
		})
		return
	}

	user, err := h.storage.FindUserByLogin(c, dto.Login)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
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
			"error": "internal error",
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

func (h *handler) logout(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "127.0.01", false, false)
	c.JSON(http.StatusOK, gin.H{
		"message": "log out",
	})
}

func (h *handler) refresh(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	if refreshToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	log.Println(refreshToken)

	token, err := jwt.ParseRefreshTokenToken(refreshToken)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	log.Println(token)
	login := token["login"]

	user, err := h.storage.FindUserByLogin(c, login.(string))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	generateAccessToken, err := jwt.GenerateAccessToken(user.Login, user.Role)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	generateRefreshToken, err := jwt.GenerateRefreshToken(user.Login)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	c.SetCookie("refresh_token", generateRefreshToken, 24*60*60*1000, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"login": user.Login,
		"token": generateAccessToken,
	})
}
