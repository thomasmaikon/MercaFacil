package db

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type setup struct {
	postgres *gorm.DB
	mysql    *gorm.DB
}

var connections *setup

func Setup() {
	if connections != nil {

		godotenv.Load("../utils/db.env")

		postgresConnection(os.Getenv("POSTGRES_NAME"), os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PORT"))

		//mysqlConnection(os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"),
		//os.Getenv("MYSQL_NAME"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))

		connections = &setup{postgres: dbPostgres,
			mysql: nil,
		}
	}
}
