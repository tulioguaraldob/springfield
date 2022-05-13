package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB(
	MySql_User string,
	MySql_Password string,
	MySql_Host string,
	MySql_Port string,
	MySql_Db_Name string,
) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		MySql_User,
		MySql_Password,
		MySql_Host,
		MySql_Port,
		MySql_Db_Name,
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error to connect to database: ", err.Error())
	}

	db = database
}

func ConnectDB() *gorm.DB {
	StartDB(
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASS"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DB_NAME"),
	)

	return db
}
