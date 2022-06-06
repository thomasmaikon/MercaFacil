package controller

import (
	"thomas/projeto_mercafacil/db"
	"thomas/projeto_mercafacil/models"

	"github.com/gin-gonic/gin"
)

func CadastroMacapa(c *gin.Context) {
	conexao := db.GetMysqlConnection()

	tipo := c.GetHeader("tipo")

	if tipo != models.TipoMacapa {
		c.JSON(403, gin.H{
			"info": "Usuario nao pertence a esse grupo",
		})
		return
	}

	usr := models.ListMacapaUsers{}

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

func ConsultaMacapa(c *gin.Context) {
	conexao := db.GetMysqlConnection()

	var users []models.Macapa
	conexao.Find(&users)

	c.JSON(200, gin.H{
		"result": users,
	})
}
