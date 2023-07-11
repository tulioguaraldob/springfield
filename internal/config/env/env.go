package env

import (
	"fmt"
	"log"
	"os"

	"github.com/TulioGuaraldoB/springfield/utils/log/infologger"
	"github.com/joho/godotenv"
)

var (
	MySqlUser     string
	MySqlPassword string
	MySqlHost     string
	MySqlPort     string
	MySqlDbName   string
	SecretString  string
	Port          string
)

func GetConfigEnv() {
	if err := godotenv.Load(".env"); err != nil {
		errorMessage := fmt.Sprintf("failed to load .env ERROR: %s", err.Error())
		infologger.Error(errorMessage)
	}

	envNil := 0

	infologger.Info("Loading .env")

	MySqlUser = os.Getenv("MYSQL_USER")
	if MySqlUser == "" {
		infologger.Error("MYSQL_USER")
		envNil += 1
	} else {
		infologger.Success("MYSQL_USER")
	}

	MySqlPassword = os.Getenv("MYSQL_PASSWORD")
	if MySqlPassword == "" {
		infologger.Error("MYSQL_PASSWORD")
		envNil += 1
	} else {
		infologger.Success("MYSQL_PASSWORD")
	}

	MySqlHost = os.Getenv("MYSQL_HOST")
	if MySqlHost == "" {
		infologger.Error("MYSQL_HOST")
		envNil += 1
	} else {
		infologger.Success("MYSQL_HOST")
	}

	MySqlPort = os.Getenv("MYSQL_PORT")
	if MySqlPort == "" {
		infologger.Error("MYSQL_PORT")
		envNil += 1
	} else {
		infologger.Success("MYSQL_PORT")
	}

	MySqlDbName = os.Getenv("MYSQL_DB_NAME")
	if MySqlDbName == "" {
		infologger.Error("MYSQL_DB_NAME")
		envNil += 1
	} else {
		infologger.Success("MYSQL_DB_NAME")
	}

	SecretString = os.Getenv("SECRET")
	if SecretString == "" {
		infologger.Error("SECRET")
	} else {
		infologger.Success("SECRET")
	}

	Port = os.Getenv("PORT")
	if Port == "" {
		infologger.Error("PORT")
	} else {
		infologger.Success("PORT")
	}

	infologger.Info("Loaded .env")

	if envNil > 0 {
		envNilInfo := fmt.Sprintf("empty environment variables = %d", envNil)
		infologger.Error(envNilInfo)
		infologger.Error("Stopping...")
		log.Fatal()
	}
}
