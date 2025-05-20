package pages

import (
	"github.com/gofiber/fiber/v3"
	"sanadam/sandbox-fiber/pkg/tadapter"
	"sanadam/sandbox-fiber/views"
)

type PageHandler struct {
	router fiber.Router
}

func NewPageHandler(router fiber.Router) *PageHandler {
	p := &PageHandler{
		router: router,
	}

	app := p.router
	app.Get("/index", p.index)
	return p
}

func (p *PageHandler) index(c fiber.Ctx) error {
	component := views.Main()
	return tadapter.Render(c, component)
	//return c.Render("base", fiber.Map{
	//	"Title": "Title page",
	//	"Cats":  []string{"Food", "Animals", "Cars", "Sport", "Music", "Technology", "More"},
	//})
}
