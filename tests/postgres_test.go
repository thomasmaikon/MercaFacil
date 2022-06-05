package tests

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"thomas/projeto_mercafacil/db"
	"thomas/projeto_mercafacil/models"
)

func TestConexaoComBancoPOSTGRESQL(t *testing.T) {
	db.GetTestPostgresConnection()
}

func TestCadastroMutiplosUsuariosPOSTGRESQL(t *testing.T) {

	db := db.GetTestPostgresConnection()

	content, _ := ioutil.ReadFile("contacts-varejao.json")

	usuarios := models.ListVarejaoUsers{}
	_ = json.Unmarshal(content, &usuarios)

	for _, usr := range usuarios.Usrs {
		err2 := db.Create(&usr)
		if err2.Error != nil {
			t.Log("Falha no cadastro de usuarios")
			t.Log(usr)
			t.Fail()
		}
	}

}
