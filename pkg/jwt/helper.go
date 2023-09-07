package jwt

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var ErrorNotAdmin = errors.New("not admin")

func GenerateAccessToken(login, role string) (token string, err error) {
	t := jwt.New(jwt.SigningMethodHS256)
	mapClaims := t.Claims.(jwt.MapClaims)
	mapClaims["login"] = login
	mapClaims["role"] = role
	mapClaims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	token, err = t.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return token, err
}

func ParseAccessToken(token string) (jwt.MapClaims, error) {
	claims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing")
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, err
	}

	return claims.Claims.(jwt.MapClaims), err
}

func GenerateRefreshToken(login string) (token string, err error) {
	t := jwt.New(jwt.SigningMethodHS256)
	mapClaims := t.Claims.(jwt.MapClaims)
	mapClaims["login"] = login
	mapClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token, err = t.SignedString([]byte("refresh_token"))
	if err != nil {
		return "", err
	}

	return token, err
}

func ParseRefreshTokenToken(token string) (jwt.MapClaims, error) {
	claims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing")
		}
		return []byte("refresh_token"), nil
	})
	if err != nil {
		return nil, err
	}

	return claims.Claims.(jwt.MapClaims), err
}

func CheckAdminRole(ctx *gin.Context) (bool, error) {
	authHeader := ctx.GetHeader("Authorization")
	headers := strings.Split(authHeader, " ")
	log.Println(headers)
	token, err := ParseAccessToken(headers[1])
	if err != nil {
		return false, err
	}

	role := token["role"]
	if role != "ADMIN" {
		return false, ErrorNotAdmin
	}

	return true, nil
}
