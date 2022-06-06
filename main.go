package main

import (
	"thomas/projeto_mercafacil/controller"
	"thomas/projeto_mercafacil/db"
	"thomas/projeto_mercafacil/service"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()

	db.GetMysqlConnection()
	db.GetPostgresConnection()

	router.POST("/logar", controller.Login)

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
