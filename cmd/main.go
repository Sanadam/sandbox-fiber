package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
	slogfiber "github.com/samber/slog-fiber"
	"log"
	"sanadam/sandbox-fiber/config"
	"sanadam/sandbox-fiber/internal/derror"
	"sanadam/sandbox-fiber/internal/pages"
	"sanadam/sandbox-fiber/pkg/logger"
)

func main() {
	config.Init()

	dErr := derror.NewDefaultErrorFiberConfig()
	app := fiber.New(dErr)
	customLogger := logger.NewLogger()
	app.Use(slogfiber.New(customLogger))
	app.Use(recover.New())

	_ = pages.NewPageHandler(app)

	log.Println(app.Listen(":8080"))
}
