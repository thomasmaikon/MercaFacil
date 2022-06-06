package controller

import (
	"strings"
	"thomas/projeto_mercafacil/db"
	"thomas/projeto_mercafacil/models"
	"thomas/projeto_mercafacil/service"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var usuario models.Login
	c.BindJSON(&usuario)

	var dbUser models.Login
	db.GetMysqlConnection().Where(&usuario).Find(&dbUser)

	if strings.Compare(dbUser.Email, usuario.Email) == 0 {
		c.JSON(200, gin.H{
			"token": service.CreateToken(dbUser),
		})
		return
	}

	db.GetPostgresConnection().Where(&usuario).Find(&dbUser)

	if strings.Compare(dbUser.Email, usuario.Email) == 0 {
		c.JSON(200, gin.H{
			"token": service.CreateToken(dbUser),
		})
		return
	}
	c.JSON(400, gin.H{
		"Info":    "Usuario nao encontrado",
		"Usuario": usuario.Email,
	})
}
