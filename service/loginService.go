package service

import (
	"fmt"
	"log"
	"thomas/projeto_mercafacil/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var secretKEY []byte = []byte("chavehashsupersecreta")

func CreateToken(user models.Login) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo":   "bar",
		"email": user.Email,
		"tipo":  user.Tipo,
		"nbf":   time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	tokenString, err := token.SignedString(secretKEY)
	if err != nil {
		log.Fatalf("Falha ao criar token")
	}

	return tokenString
}

func validateToken(tk string) (bool, string) {

	token, _ := jwt.Parse(tk, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secretKEY, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		tipo := claims["tipo"].(string)
		return true, tipo
	} else {
		return false, ""
	}
}

func Authorization(c *gin.Context) {

	const BEARER_SCHEMA = "Bearer"

	authHeader := c.GetHeader("Authorization")
	tokenString := authHeader[len(BEARER_SCHEMA)+1:]

	ok, tipo := validateToken(tokenString)

	if !ok {
		c.JSON(401, gin.H{
			"Info": "Usuario nao autorizado",
		})
		c.AbortWithStatus(401)
	}

	c.Request.Header.Add("tipo", string(tipo))
}
