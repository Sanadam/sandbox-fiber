package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/template/html/v2"
	slogfiber "github.com/samber/slog-fiber"
	"log"
	"os"
	"sanadam/sandbox-fiber/config"
	"sanadam/sandbox-fiber/internal/derror"
	"sanadam/sandbox-fiber/internal/pages"
	"sanadam/sandbox-fiber/pkg/logger"
)

func main() {
	config.Init()

	dErr := derror.NewDefaultErrorFiberConfig()
	engine := html.New("./html", ".html")
	//engine.Templates.Funcs()
	engine.Reload(true)

	if _, err := os.Stat("./html"); os.IsNotExist(err) {
		log.Fatal("Views folder does not exist")
	}

	app := fiber.New(fiber.Config{
		Views:        engine,
		ErrorHandler: dErr,
	})
	customLogger := logger.NewLogger()
	app.Use(slogfiber.New(customLogger))
	app.Use(recover.New())

	_ = pages.NewPageHandler(app)

	log.Println(app.Listen(":8080"))
}
