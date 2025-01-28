package utils

import (
	"fmt"

	"github.com/anjarmath/01_golang_api_sederhana/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func GenerateSession(c *fiber.Ctx, user model.User) error {
	sess := c.Locals("session").(*session.Session)

	sess.Set("username", user.Username)
	sess.Set("user_id", user.Id.String())

	// Save session
	if err := sess.Save(); err != nil {
		return fmt.Errorf("failed to create session")
	}

	return nil
}

func ValidateSession(c *fiber.Ctx) bool {
	sess := c.Locals("session").(*session.Session)

	username := sess.Get("username")
	id := sess.Get("user_id")
	return !(username == nil || id == nil)
}
