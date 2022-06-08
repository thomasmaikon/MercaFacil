package mysql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"thomas/projeto_mercafacil/models"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/ory/dockertest"
)

/* func TestXxx(t *testing.T) {
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
} */

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

var db *sql.DB

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Maquina nao tem Docker: %s", err)
	}

	resource, err := pool.Run("mysql", "5.7", []string{"MYSQL_ROOT_PASSWORD=admin",
		"MYSQL_DATABASE=admin", "MYSQL_USER=admin", "MYSQL_PASSWORD=admin", "MYSQL_ROOT_HOST=%"})
	if err != nil {
		log.Fatalf("Nao foi possivel inicializar o container: %s", err)
	}

	if err := pool.Retry(func() error {
		var err error
		db, err = sql.Open("mysql", fmt.Sprintf("admin:admin@(localhost:%s)/admin?parseTime=true", resource.GetPort("3306/tcp")))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Nao foi possivel se conectar ao container: %s", err)
	}

	// criar sistema para realizar migrations semelhante ao flyway
	db.Exec(`CREATE table contacts (id serial PRIMARY KEY, nome VARCHAR ( 200 ) NOT NULL, celular VARCHAR ( 20 ) NOT NULL);`)

	code := m.Run()

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Nao foi possivel remover o container: %s", err)
	}

	os.Exit(code)
}
func TestCadastroUsuarios(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	content, _ := ioutil.ReadFile("contacts-macapa.json")

	usuarios := models.ListMacapaUsers{}
	_ = json.Unmarshal(content, &usuarios)

	var size int

	for i, usr := range usuarios.Usrs {
		// tomar cuidado com '?' pode causar problema de sql injection, como e um test ent e controloado pelo desenvolvedor
		_, err2 := db.Exec(`INSERT INTO contacts (nome, celular) VALUES (?, ?);`, usr.NameFormat(), usr.NumberFormat())
		if err2 != nil {
			t.Log("Falha no cadastro de usuarios")
			t.Log(err2)
			t.Fail()
		}
		size = i + 1
	}

	var qtd int
	db.QueryRow(`SELECT count(*) FROM contacts`).Scan(&qtd)

	if size != qtd {
		t.Log("Nao foram cadastrados todos os usuarios")
		t.Fail()
	}

}

/* func TestFindUser(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	t.Run("Pegando usuarios", func(t *testing.T) {
		pessoa := &models.Macapa{Nome: "thomas", Celular: "5584920896635"}
		banco.Create(pessoa)
		var users []models.Macapa
		banco.Find(&users)
		fmt.Println(users)
	})
} */
