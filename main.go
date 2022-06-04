package main

import (
	"strings"
	"thomas/projeto_mercafacil/controller"
	"thomas/projeto_mercafacil/db"
	"thomas/projeto_mercafacil/models"
	"thomas/projeto_mercafacil/service"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()

	db.GetMysqlConnection()
	db.GetPostgresConnection()

	router.POST("/logar", func(c *gin.Context) {
		var usuario models.Login
		c.BindJSON(&usuario)

		var dbUser models.Login
		db.GetMysqlConnection().Where(&usuario).Find(&dbUser)

		if strings.Compare(dbUser.Email, usuario.Email) == 0 {
			c.JSON(200, gin.H{
				"token": service.CreateToken(dbUser.Email),
			})
			return
		}

		db.GetPostgresConnection().Where(&usuario).Find(&dbUser)

		if strings.Compare(dbUser.Email, usuario.Email) == 0 {
			c.JSON(200, gin.H{
				"token": service.CreateToken(dbUser.Email),
			})
			return
		}
		c.JSON(400, gin.H{
			"Info":    "Usuario nao encontrado",
			"Usuario": usuario.Email,
		})
	})

	authorized := router.Group("/cadastrar", service.Authorization)
	authorized.Use(service.Authorization)
	{
		authorized.POST("/macapa", controller.CadastroMacapa)
		authorized.POST("/varejao", controller.CadastroVarejao)
	}

	consulta := router.Group("/consultar")
	{
		consulta.GET("/macapa", controller.ConsultaMacapa)
		consulta.GET("/varejao", controller.ConsultaVarejao)
	}

	router.Run(":8000")

}
