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
	if dbMySql == nil {
		dbMySql = mysqlConnection()
	}
	return dbMySql
}

func mysqlConnection() *gorm.DB {
	connect := "admin:admin@tcp(localhost:3306)/admin?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{})

	if err != nil {
		fmt.Errorf("Falha ao se conecctar com MYSQL")
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Macapa{})
	db.AutoMigrate(&models.Login{})

	db.Create(&models.Login{Email: "admin-macapa", Senha: "admin", Tipo: 1})

	return db
}
