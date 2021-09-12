package controllers

import "github.com/gofiber/fiber/v2"

func RenderLogin(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{})
}

func RenderRegister(c *fiber.Ctx) error {
	return c.Render("register", fiber.Map{})
}
