package controller

import (
	"strings"
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

func RemoverMacapa(c *gin.Context) {
	conexao := db.GetMysqlConnection()

	tipo := c.GetHeader("tipo")

	if strings.Compare(tipo, models.TipoMacapa) != 0 {
		c.JSON(403, gin.H{
			"info": "Usuario nao pertence a esse grupo",
		})

		return
	}

	nome := c.Query("nome")

	if strings.Compare(nome, "") != 0 {
		err := conexao.Where("nome LIKE ?", nome).Delete(&models.Macapa{})
		if err.Error == nil {
			c.JSON(200, gin.H{
				"Usuario": nome,
				"Info":    "Removido com sucesso",
			})
			return
		}
	}

	c.JSON(400, gin.H{
		"Info": "dados enviados invalidos",
	})
}

func AtualizarMacapa(c *gin.Context) {
	conexao := db.GetMysqlConnection()

	tipo := c.GetHeader("tipo")

	if strings.Compare(tipo, models.TipoMacapa) != 0 {
		c.JSON(403, gin.H{
			"info": "Usuario nao pertence a esse grupo",
		})

		return
	}

	nome := c.Query("nome")
	numero := c.Query("numero")

	macapa := models.Macapa{
		Nome:    nome,
		Celular: numero,
	}

	if strings.Compare(nome, "") != 0 && strings.Compare(nome, "") != 0 {
		err := conexao.Model(models.Macapa{}).Where("nome LIKE ?", macapa.NameFormat()).Updates(models.Macapa{Nome: macapa.NameFormat(), Celular: macapa.NumberFormat()})
		if err.Error == nil {
			c.JSON(200, gin.H{
				"Usuario": nome,
				"Info":    "Atualizado com sucesso",
			})
			return
		}
	}

	c.JSON(400, gin.H{
		"Info": "dados enviados invalidos",
	})
}
