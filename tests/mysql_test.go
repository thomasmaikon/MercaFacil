package tests

import (
	"fmt"
	"log"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestXxx(t *testing.T) {
	connect := "admin:admin@tcp(localhost:3306)/admin?charset=utf8mb4&parseTime=True&loc=Local"

	_, err := gorm.Open(mysql.Open(connect), &gorm.Config{})
	if err != nil {
		fmt.Println("Falha ao se conecctar com MYSQL")
		fmt.Println(connect)
		log.Fatalln(err)
	} else {
		fmt.Println("Conexao:")
		fmt.Println(connect)
	}
}

/* func TestConexaoComBancoMYSQL(t *testing.T) {
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
*/
