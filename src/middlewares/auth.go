package middlewares

import "github.com/gofiber/fiber/v2"

func IsAuthenticated(c *fiber.Ctx) error {
	token := c.Cookies("token")
	if token != "" {
		return c.Next()
	}
	return c.Render("redirect", fiber.Map{})
}
