package derror

import (
	"errors"
	"github.com/gofiber/fiber/v3"
)

func NewDefaultErrorFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: func(c fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	}
}
