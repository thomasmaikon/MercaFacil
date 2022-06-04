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
		dbPostgres = postgresConnection()
	}
	return dbPostgres
}

func postgresConnection() *gorm.DB {

	connect := "host=localhost user=admin password=admin dbname=postgres port=5432 sslmode=disable TimeZone=America/Fortaleza"
	db, err := gorm.Open(postgres.Open(connect), &gorm.Config{})

	if err != nil {
		fmt.Errorf("Falha ao se conectar no banco...")
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Varejao{})
	db.AutoMigrate(&models.Login{})
	// inicializa usuario admin para testar cadastro
	db.Create(&models.Login{Email: "admin-varejao", Senha: "admin", Tipo: 2})

	return db
}
