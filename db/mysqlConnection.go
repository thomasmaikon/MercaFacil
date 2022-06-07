package db

import (
	"fmt"
	"log"
	"thomas/projeto_mercafacil/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbMySql *gorm.DB

func GetMysqlConnection() *gorm.DB {
	return dbMySql
}

func mysqlConnection(user, password, host, port, dbname string) {

	connect := "admin:admin@tcp(mysql:3306)/admin?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{})

	if err != nil {
		fmt.Println("Falha ao se conecctar com MYSQL")
		fmt.Println(connect)
		log.Fatalln(err)
	} else {
		fmt.Println("Conexao:")
		fmt.Println(connect)
	}

	db.AutoMigrate(&models.Macapa{})
	db.AutoMigrate(&models.Login{})

	db.Create(&models.Login{Email: "admin-macapa", Senha: "admin", Tipo: models.TipoMacapa})

	dbMySql = db
}
