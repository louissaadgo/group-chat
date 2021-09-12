package middlewares

import "github.com/gofiber/fiber/v2"

func IsAuthenticated(c *fiber.Ctx) error {
	if true {
		return c.Next()
	}
	return c.Render("login", fiber.Map{})
}
