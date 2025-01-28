package middleware

import (
	"github.com/anjarmath/01_golang_api_sederhana/utils"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		valid := utils.ValidateSession(c)
		if !valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Error{
				Message: "Siapa lu?",
			})
		}

		return c.Next()
	}
}
