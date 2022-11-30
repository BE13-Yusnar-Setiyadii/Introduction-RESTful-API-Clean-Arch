package main

import (
	"fmt"
	"yusnar/clean-arch/config"
	"yusnar/clean-arch/factory"
	"yusnar/clean-arch/middlewares"
	"yusnar/clean-arch/utils/database/mysql"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)

	e := echo.New()

	factory.InitFactory(db, e)

	middlewares.LogMiddlewares(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVER_PORT)))
}
