package pages

import "github.com/gofiber/fiber/v3"

type PageHandler struct {
	router fiber.Router
}

func NewPageHandler(router fiber.Router) {
	p := &PageHandler{
		router: router,
	}

	app := p.router
	app.Get("/index", p.index)
}

func (p *PageHandler) index(c fiber.Ctx) error {
	return c.SendString("Main page")
	//return fiber.NewError(fiber.StatusNotFound, "Not this page")
}
