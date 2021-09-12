package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func RenderLogin(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{})
}

func RenderRegister(c *fiber.Ctx) error {
	return c.Render("register", fiber.Map{})
}

func Login(c *fiber.Ctx) error {
	if c.FormValue("login-email") == "louis" {
		c.Cookie(&fiber.Cookie{
			Name:  "token",
			Value: "ok",
		})
		return c.Render("index", fiber.Map{})
	}
	return c.Render("login", fiber.Map{})
}
