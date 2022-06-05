package controller

import (
	"thomas/projeto_mercafacil/db"
	"thomas/projeto_mercafacil/models"

	"github.com/gin-gonic/gin"
)

func CadastroVarejao(c *gin.Context) {
	conexao := db.GetPostgresConnection()

	email := c.GetHeader("email")

	var login models.Login
	conexao.Where("email = ?", email).Find(&login)

	// impedir que usuario admin de macapa possa acessar
	if login.Tipo != 2 {
		c.JSON(403, gin.H{
			"info": "Usuario nao pertence a esse grupo",
		})

		return
	}

	usr := models.ListVarejaoUsers{}

	c.BindJSON(&usr)

	for i := 0; i < len(usr.Usrs); i++ {
		err := conexao.Create(usr.Usrs[i].Format())

		if err.Error != nil {
			c.JSON(403, gin.H{
				"info":    "Usuario ja existente",
				"usuario": usr.Usrs[i],
			})
			return
		}
	}

	c.JSON(201, gin.H{
		"result": "Todos os contatos foram cadastrados com sucesso",
	})
}

func ConsultaVarejao(c *gin.Context) {
	conexao := db.GetPostgresConnection()

	var users []models.Varejao

	conexao.Find(&users)

	c.JSON(200, gin.H{
		"result": users,
	})
}
