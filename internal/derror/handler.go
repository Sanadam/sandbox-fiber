package derror

import (
	"errors"
	"github.com/gofiber/fiber/v3"
	"log"
)

func NewDefaultErrorFiberConfig() fiber.ErrorHandler {
	return func(c fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError

		var e *fiber.Error
		if errors.As(err, &e) {
			code = e.Code
			log.Println(e.Message)
		}
		return c.Status(code).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
}
