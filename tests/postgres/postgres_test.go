package postgres

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"thomas/projeto_mercafacil/models"

	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
)

var db *sql.DB

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalln("Nao foi possivel se conectar com docker local: ", err)
	}

	resource, err := pool.Run("postgres", "11", []string{"POSTGRES_PASSWORD=admin",
		"POSTGRES_USER=admin"})

	resource.Expire(60 * 5)
	if err != nil {
		log.Fatalln("Nao foi possivel inicializar o container: ", err)
	}

	if err := pool.Retry(func() error {
		var err error
		db, err = sql.Open("postgres", fmt.Sprintf("postgres://admin:admin@localhost:%s/postgres?sslmode=disable", resource.GetPort("5432/tcp")))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalln("Nao foi possivel se conectar ao container: ", err)
	}

	// utilizar sistema de migrations para importar esses dados
	db.Exec(`CREATE table contacts (id serial PRIMARY KEY,nome VARCHAR ( 100 ) NOT NULL,celular VARCHAR ( 13 ) NOT NULL);`)

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

	content, _ := ioutil.ReadFile("contacts-varejao.json")

	usuarios := models.ListVarejaoUsers{}
	_ = json.Unmarshal(content, &usuarios)

	var size int

	for _, usr := range usuarios.Usrs {
		// Nesse caso, a '?' nao funciona, tem que utilizar o '$', referente a cada variavel
		_, err2 := db.Exec(`INSERT INTO contacts (nome, celular) VALUES ($1, $2);`, usr.NameFormat(), usr.NumberFormat())
		if err2 != nil {
			t.Log("Falha no cadastro de usuarios")
			t.Log(err2)
			t.Log(usr)
			//t.Fail()
			// nesse teste propositalmente existe um dado replicado
			size--
		}
		size++
	}

	var qtd int
	db.QueryRow(`SELECT count(*) FROM contacts`).Scan(&qtd)

	if size != qtd {
		t.Log("Nao foram cadastrados todos os usuarios")
		t.Fail()
	}

}
