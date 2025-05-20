package template

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/html/v2"
	"log"
	"os"
)

func NewTemplateEngine() fiber.Views {
	if _, err := os.Stat("./html"); os.IsNotExist(err) {
		log.Println("Views folder does not exist", err)
	}
	engine := html.New("./html", ".html")
	if err := engine.Load(); err != nil {
		log.Println("failed load template", err)
	}
	//engine.Templates.Funcs()
	engine.Reload(true)
	return engine
}
