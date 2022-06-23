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

	autorizacao := router.Group("/", service.Authorization)
	autorizacao.Use(service.Authorization)
	{
		autorizacao.POST("/cadastrar", controller.Cadastro)
		autorizacao.GET("/consultar", controller.Consulta)
		autorizacao.DELETE("/remover", controller.Remover)
		autorizacao.PUT("/atualizar", controller.Atualizar)
	}

	router.Run(":8000")

}
