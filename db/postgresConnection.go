package db

import (
	"fmt"
	"log"
	"thomas/projeto_mercafacil/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbPostgres *gorm.DB

func GetPostgresConnection() *gorm.DB {

	if dbPostgres == nil {
		dbPostgres = postgresConnection("postgres", "admin", "admin", "postgres", "5432")
	}
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
