package configs

import (
	"os"
	"strconv"
)

type AppConfig struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOSTNAME string
	DB_PORT     int
	DB_NAME     string
}

func ReadEnv() *AppConfig {
	var app = AppConfig{}
	app.DB_USERNAME = os.Getenv("DBUSER")
	app.DB_PASSWORD = os.Getenv("DBPASS")
	app.DB_HOSTNAME = os.Getenv("DBHOST")
	portConv, errConv := strconv.Atoi(os.Getenv("DBPORT"))
	if errConv != nil {
		panic("error convert dbport")
	}
	app.DB_PORT = portConv
	app.DB_NAME = os.Getenv("DBNAME")
	return &app
}

func InitConfig() *AppConfig {
	return ReadEnv()
}
