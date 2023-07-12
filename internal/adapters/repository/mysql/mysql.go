package mysql

import (
	"fmt"
	"log"

	"github.com/TulioGuaraldoB/springfield/internal/adapters/repository/mysql/migration"
	"github.com/TulioGuaraldoB/springfield/internal/adapters/repository/mysql/seed"
	"github.com/TulioGuaraldoB/springfield/internal/config/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func startDatabase() {
	database, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		env.MySqlUser,
		env.MySqlPassword,
		env.MySqlHost,
		env.MySqlPort,
		env.MySqlDbName,
	)), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to MySQL Server. %s", err.Error())
	}

	db = database
	log.Println("Connected to MySQL Server")

	migration.Run(db)
	seed.Run(db)
}

func OpenConnection() *gorm.DB {
	if db == nil {
		startDatabase()
	}

	return db
}
