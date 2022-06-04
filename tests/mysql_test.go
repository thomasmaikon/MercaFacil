package tests

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"thomas/projeto_mercafacil/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestConexaoComBancoMYSQL(t *testing.T) {
	connect := "admin:admin@tcp(localhost:3306)/admin?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(connect), &gorm.Config{})
	if err != nil {
		t.Log("Falha ao se conectar no banco...")
		t.Fail()
	}
}

func TestCadastrarUsuariosMYSQL(t *testing.T) {
	connect := "admin:admin@tcp(localhost:3306)/admin?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(connect), &gorm.Config{})
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
	var connect string = "admin:admin@tcp(localhost:3306)/admin?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(connect), &gorm.Config{})

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
	var connect string = "admin:admin@tcp(localhost:3306)/admin?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(connect), &gorm.Config{})

	var contato models.Macapa
	err2 := db.Where("nome = ?", "Maria Fernanda Almeida").Find(&contato)
	if err2.Error != nil {
		t.Log("Falha ao encontrar usuario")
		t.Log(contato)
		t.Fail()
	}
}
