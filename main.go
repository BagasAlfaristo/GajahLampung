package main

import (
	"gajahlampung/app/configs"
	"gajahlampung/app/databases"
	"gajahlampung/app/migrations"
	"gajahlampung/app/routers"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := configs.InitConfig()
	dbMysql := databases.InitDBMysql(cfg)
	migrations.InitMigrations(dbMysql)

	e := echo.New()

	routers.InitRoutuer(e, dbMysql)
	e.Logger.Fatal(e.Start(":8080"))
}
