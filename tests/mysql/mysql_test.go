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
	"github.com/ory/dockertest/v3"
)

var db *sql.DB

func TestMain(m *testing.M) {
	// conectando com o docker da maquina
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Maquina nao tem Docker: %s", err)
	}

	// definindo container
	resource, err := pool.Run("mysql", "5.7", []string{"MYSQL_ROOT_PASSWORD=admin",
		"MYSQL_DATABASE=admin", "MYSQL_USER=admin", "MYSQL_PASSWORD=admin", "MYSQL_ROOT_HOST=%"})

	resource.Expire(60 * 5)
	if err != nil {
		log.Fatalf("Nao foi possivel inicializar o container: %s", err)
	}

	// tentando conectar com o container
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

	// Os testes tem ate 5 minutos para serem executados
	resource.Expire(60 * 5)

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
