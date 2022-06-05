package db

import (
	"fmt"
	"log"
	"thomas/projeto_mercafacil/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbMySql *gorm.DB
var testMySql *gorm.DB

func GetTestMySqlConnection() *gorm.DB {
	if testMySql == nil {
		testMySql = testMySqlConnection()
	}
	return testMySql
}

func testMySqlConnection() *gorm.DB {
	connect := "admin:admin@tcp(projeto_mercafacil-mysql-1:3306)/admin?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{})

	if err != nil {
		fmt.Println("Falha ao se conecctar com MYSQL")
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Macapa{})
	db.AutoMigrate(&models.Login{})

	db.Create(&models.Login{Email: "admin-macapa", Senha: "admin", Tipo: 1})

	return db
}

func GetMysqlConnection() *gorm.DB {
	if dbMySql == nil {
		dbMySql = mysqlConnection()
	}
	return dbMySql
}

func mysqlConnection() *gorm.DB {
	connect := "admin:admin@tcp(localhost:3306)/admin?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{})

	if err != nil {
		fmt.Println("Falha ao se conecctar com MYSQL")
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Macapa{})
	db.AutoMigrate(&models.Login{})

	db.Create(&models.Login{Email: "admin-macapa", Senha: "admin", Tipo: 1})

	return db
}
