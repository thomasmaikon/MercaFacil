package main

import (
	"thomas/projeto_mercafacil/controller"
	"thomas/projeto_mercafacil/db"
	"thomas/projeto_mercafacil/service"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()

	//db.Setup()

	db.GetPostgresConnection()
	db.GetMysqlConnection()
	// utilizar uma rota para cada operacao e interpretar o tipo de usuario e viavel ate que ponto?
	// proxima analise - feat

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

	remover := router.Group("/remover", service.Authorization)
	remover.Use(service.Authorization)
	{
		remover.DELETE("/macapa", controller.RemoverMacapa)
		remover.DELETE("/varejao", controller.RemoverVarejao)
	}

	atualizar := router.Group("/atualizar", service.Authorization)
	atualizar.Use(service.Authorization)
	{
		atualizar.PUT("/macapa", controller.AtualizarMacapa)
		atualizar.PUT("/varejao", controller.AtualizarVarejao)
	}

	router.Run(":8000")

}
