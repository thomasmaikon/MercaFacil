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

// var postg *sql.DB

/* rodar testes no banco
func TestMain(m *testing.M) {

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	postgres, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "11",
		Env: []string{
			"POSTGRES_USER=admin",
			"POSTGRES_PASSWORD=admin",
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "postg",
		}
	})

	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err := pool.Retry(func() error {
		var err error
		postg, err = sql.Open("postgres", "host=localhost user=admin password=admin dbname=postgres port=5432 sslmode=disable TimeZone=America/Fortaleza")
		if err != nil {
			return err
		}
		return postg.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	code := m.Run()

	if err := pool.Purge(postgres); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
} */
