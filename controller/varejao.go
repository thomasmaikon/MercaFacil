package controller

import (
	"log"
	"strings"
	"thomas/projeto_mercafacil/db"
	"thomas/projeto_mercafacil/models"
	"thomas/projeto_mercafacil/service"

	"github.com/gin-gonic/gin"
)

func validGroup(str string) bool {
	return strings.Compare(str, models.TipoVarejao) != 0
}

func Cadastro(c *gin.Context) {

	usr := models.ListUsuarios{}

	c.BindJSON(&usr)

	factory := service.FactoryUser{}

	usuarios, repository, erro := factory.GetUserDB(c.GetHeader("tipo"), usr.Usrs)

	if erro != nil {
		log.Fatalln(erro)
	}

	for _, newUser := range usuarios {
		err := repository.Create(newUser)
		if err != nil {
			c.JSON(403, gin.H{
				"info":    "Usuario ja existente",
				"usuario": usuarios,
			})
			return
		}
	}

	c.JSON(201, gin.H{
		"Info":     "Usuarios cadastrados",
		"usuarios": usuarios,
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

func RemoverVarejao(c *gin.Context) {
	conexao := db.GetPostgresConnection()

	// impedir que usuario admin de macapa possa acessar
	if validGroup(c.GetHeader("tipo")) {
		c.JSON(403, gin.H{
			"info": "Usuario nao pertence a esse grupo",
		})

		return
	}

	nome := c.Query("nome")

	if strings.Compare(nome, "") != 0 {
		err := conexao.Where("nome LIKE ?", nome).Delete(&models.Varejao{})
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

func AtualizarVarejao(c *gin.Context) {
	conexao := db.GetPostgresConnection()

	// impedir que usuario admin de macapa possa acessar
	if validGroup(c.GetHeader("tipo")) {
		c.JSON(403, gin.H{
			"info": "Usuario nao pertence a esse grupo",
		})

		return
	}

	nome := c.Query("nome")
	numero := c.Query("numero")

	if strings.Compare(nome, "") != 0 && strings.Compare(nome, "") != 0 {
		err := conexao.Model(models.Varejao{}).Where("nome LIKE ?", nome).Updates(models.Varejao{Nome: nome, Celular: numero})
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
