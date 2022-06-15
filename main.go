package main

import (
	"os"
	"strings"
	"thomas/projeto_mercafacil/controller"
	"thomas/projeto_mercafacil/db"
	"thomas/projeto_mercafacil/service"

	"github.com/gin-gonic/gin"
)

func setup(ambiente string) {

	if strings.Compare(ambiente, "local") == 0 {
		db.SetPostgresConnection("local.env")
		db.SetMysqlConnection("local.env")
	} else {
		db.SetPostgresConnection("producao.env")
		db.SetMysqlConnection("producao.env")
	}

}

func main() {

	setup(os.Args[1])

	router := gin.New()

	router.POST("/logar", controller.Login)

	router.POST("/cadastrar", service.Authorization, controller.Cadastro)
	/* authorized := router.Group("/cadastrar", service.Authorization)
	authorized.Use(service.Authorization)
	{
		//authorized.POST("/macapa", controller.CadastroMacapa)
		authorized.POST("/", controller.Cadastro)
	} */

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
