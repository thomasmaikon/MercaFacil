package controller

import (
	"thomas/projeto_mercafacil/models"
	"thomas/projeto_mercafacil/service"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var usuario models.Login
	c.BindJSON(&usuario)

	dbConnection, err := factory.GetUserDB(models.TipoMacapa)
	user, err := dbConnection.FindByUser(usuario)
	if err == nil {
		c.JSON(200, gin.H{
			"Token": service.CreateToken(user),
		})
		return
	}

	dbConnection, err = factory.GetUserDB(models.TipoVarejao)
	user, err = dbConnection.FindByUser(usuario)
	if err == nil {
		c.JSON(200, gin.H{
			"Token": service.CreateToken(user),
		})
		return
	}

	c.JSON(400, gin.H{
		"Info":    "Usuario nao encontrado",
		"Usuario": usuario.Email,
	})
}
