package db

import (
	"fmt"
	"log"
	"os"
	"thomas/projeto_mercafacil/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbMySql *gorm.DB

func SetMysqlConnection(file string) *gorm.DB {
	if dbMySql == nil {
		godotenv.Load(file)

		dbMySql = mysqlConnection(os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"),
			os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
	}
	return dbMySql
}

func GetMysqlConnection() *gorm.DB {
	return dbMySql
}

func mysqlConnection(user, password, host, port, dbname string) *gorm.DB {

	connect := user + ":" + password + "@(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"

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

	return db
}
