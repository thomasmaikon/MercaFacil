package tests

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"thomas/projeto_mercafacil/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestConexaoComBancoPOSTGRESQL(t *testing.T) {
	// descobrir porque mesmo definindo o host no docker compose, nao funciona
	var connect string = "host=localhost user=admin password=admin dbname=postgres port=5432 sslmode=disable TimeZone=America/Fortaleza"
	_, err := gorm.Open(postgres.Open(connect), &gorm.Config{})

	if err != nil {
		t.Log("Falha ao se conectar no banco...")
		t.Fail()
	}
}

func TestCadastroMutiplosUsuariosPOSTGRESQL(t *testing.T) {
	var connect string = "host=localhost user=admin password=admin dbname=postgres port=5432 sslmode=disable TimeZone=America/Fortaleza"
	db, _ := gorm.Open(postgres.Open(connect), &gorm.Config{})

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
