package controller

import (
	"log"
	"strings"
	"thomas/projeto_mercafacil/models"
	"thomas/projeto_mercafacil/service"

	"github.com/gin-gonic/gin"
)

func validGroup(str string) bool {
	return strings.Compare(str, models.TipoVarejao) != 0
}

var factory service.FactoryUser = service.FactoryUser{}

func Cadastro(c *gin.Context) {

	usr := models.ListUsuarios{}

	c.BindJSON(&usr)

	repository, erro := factory.GetUserDB(c.GetHeader("tipo"))

	if erro != nil {
		log.Fatalln(erro)
	}

	for _, newUser := range usr.Usrs {
		err := repository.Create(newUser)
		if err != nil {
			c.JSON(403, gin.H{
				"info":    "Usuario ja existente",
				"usuario": newUser,
			})
		}
	}

	c.JSON(201, gin.H{
		"Info":     "Usuarios cadastrados",
		"usuarios": usr,
	})
}

func Consulta(c *gin.Context) {

	repository, erro := factory.GetUserDB(c.GetHeader("tipo"))

	if erro != nil {
		log.Fatalln(erro)
		return
	}

	users, err := repository.Find()

	if erro != nil {
		log.Fatalln(err)
		return
	}

	c.JSON(200, gin.H{
		"result": users,
	})
}

func Remover(c *gin.Context) {

	repository, erro := factory.GetUserDB(c.GetHeader("tipo"))

	if erro != nil {
		log.Fatalln(erro)
	}

	nome := c.Query("nome")
	err := repository.Delete(nome)

	if err == nil {
		c.JSON(200, gin.H{
			"Info": "Usuario Deletado",
		})
		return
	}

	c.JSON(400, gin.H{
		"Info": "Houve um erro durante a delecao",
	})
}

func Atualizar(c *gin.Context) {
	repository, erro := factory.GetUserDB(c.GetHeader("tipo"))

	if erro != nil {
		log.Fatalln(erro)
	}

	var usr models.Usuario
	c.BindJSON(&usr)
	id := c.Query("id")

	newUser, err := repository.Update(id, usr)

	if err == nil {
		c.JSON(200, gin.H{
			"Usuario": newUser,
			"Info":    "Atualizado com sucesso",
		})
		return
	}

	c.JSON(400, gin.H{
		"Info": "dados enviados invalidos",
	})
}
