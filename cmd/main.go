package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/static"
	slogfiber "github.com/samber/slog-fiber"
	"log"
	"sanadam/sandbox-fiber/config"
	"sanadam/sandbox-fiber/internal/derror"
	"sanadam/sandbox-fiber/internal/pages"
	"sanadam/sandbox-fiber/pkg/logger"
	"sanadam/sandbox-fiber/pkg/template"
)

func main() {
	config.Init()

	dErr := derror.NewDefaultErrorFiberConfig()
	engine := template.NewTemplateEngine()

	app := fiber.New(fiber.Config{
		Views:        engine,
		ErrorHandler: dErr,
	})

	customLogger := logger.NewLogger()
	app.Use(slogfiber.New(customLogger))
	app.Use(recover.New())
	app.Use("/public", static.New("./public"))

	pages.NewPageHandler(app)

	log.Println(app.Listen(":8080"))
}
