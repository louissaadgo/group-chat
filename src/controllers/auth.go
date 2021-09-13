package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/louissaadgo/group-chat/src/models"
)

func RenderLogin(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{})
}

func RenderRegister(c *fiber.Ctx) error {
	return c.Render("register", fiber.Map{})
}

func Login(c *fiber.Ctx) error {
	user := models.User{}
	c.BodyParser(&user)
	for _, i := range Users {
		if i.Username == user.Username {
			if i.Password == user.Password {
				c.Cookie(&fiber.Cookie{
					Name:  "token",
					Value: user.Username,
				})
				return c.JSON(user)
			}
		}
	}
	c.Status(400)
	return c.JSON(user)
}

func Register(c *fiber.Ctx) error {
	user := models.User{}
	c.BodyParser(&user)
	create := true
	for _, i := range Users {
		if i.Username == user.Username {
			create = false
		}
	}
	if create {
		Users = append(Users, user)
		c.Cookie(&fiber.Cookie{
			Name:  "token",
			Value: user.Username,
		})
		return c.JSON(user)
	}
	c.Status(400)
	return c.JSON(user)
}
