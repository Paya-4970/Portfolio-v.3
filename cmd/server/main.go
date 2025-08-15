package main

import (
	"github.com/Paya-4970/Portfolio-v.3/config"
	"github.com/Paya-4970/Portfolio-v.3/internal/database"
	"github.com/Paya-4970/Portfolio-v.3/internal/routers"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	cfg := config.LoadConfig()
	database.InitDB(cfg)

	e := echo.New()
	router.InitRouter(e)

	log.Fatal(e.Start(":" + cfg.AppPort))
}
