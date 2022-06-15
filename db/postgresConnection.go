package db

import (
	"fmt"
	"log"
	"os"
	"thomas/projeto_mercafacil/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbPostgres *gorm.DB

func SetPostgresConnection(file string) *gorm.DB {

	if dbPostgres == nil {
		godotenv.Load(file)

		dbPostgres = postgresConnection(os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PORT"))
	}
	return dbPostgres
}

func GetPostgresConnection() *gorm.DB {
	return dbPostgres
}

func postgresConnection(host, user, password, dbname, port string) *gorm.DB {

	host = "host=" + host
	user = "user=" + user
	password = "password=" + password
	dbname = "dbname=" + dbname
	port = "port=" + port

	connect := host + " " + user + " " + password + " " + dbname + " " + port + " " + "sslmode=disable TimeZone=America/Fortaleza"
	db, err := gorm.Open(postgres.Open(connect), &gorm.Config{})

	if err != nil {
		fmt.Println("Falha ao se conectar no banco...")
		log.Fatalln(err)
	} else {
		fmt.Println("Conexao:")
		fmt.Println(connect)
	}

	db.AutoMigrate(&models.Varejao{})
	db.AutoMigrate(&models.Login{})

	// inicializa usuario admin para testar cadastro, se ja existir ele informa que usuario ja existe
	db.Create(&models.Login{Email: "admin-varejao", Senha: "admin", Tipo: models.TipoVarejao})

	return db
}
