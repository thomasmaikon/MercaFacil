package tests

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"thomas/projeto_mercafacil/db"
	"thomas/projeto_mercafacil/models"
)

func TestConexaoComBancoMYSQL(t *testing.T) {
	db.GetTestMySqlConnection()
}

func TestCadastrarUsuariosMYSQL(t *testing.T) {

	db := db.GetTestMySqlConnection()
	user := models.Macapa{
		Nome:    "PessoaTESTE",
		Celular: "5584999999999",
	}
	err := db.Create(&user)
	if err.Error != nil {
		t.Log("Falha no cadastro de usuarios")
		t.Log(err)
		t.Fail()
	}
}
func TestCadastroMutiplosUsuariosMYSQL(t *testing.T) {
	db := db.GetTestMySqlConnection()

	content, _ := ioutil.ReadFile("contacts-macapa.json")

	usuarios := models.ListMacapaUsers{}
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

func TestBuscarUsuarioMYSQL(t *testing.T) {
	db := db.GetTestMySqlConnection()

	var contato models.Macapa
	err2 := db.Where("nome = ?", "Maria Fernanda Almeida").Find(&contato)
	if err2.Error != nil {
		t.Log("Falha ao encontrar usuario")
		t.Log(contato)
		t.Fail()
	}
}
