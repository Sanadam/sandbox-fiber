package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"log"
	"sanadam/sandbox-fiber/config"
	"sanadam/sandbox-fiber/internal/derror"
	"sanadam/sandbox-fiber/internal/pages"
)

func main() {
	config.Init()
	//conf := config.NewConfig()

	dErr := derror.NewDefaultErrorFiberConfig()
	app := fiber.New(dErr)
	app.Use(logger.New())
	app.Use(recover.New())

	_ = pages.NewPageHandler(app)

	log.Println(app.Listen(":8080"))
}
