package routes

import (
	"github.com/antoniodipinto/ikisocket"
	"github.com/gofiber/fiber/v2"
	"github.com/louissaadgo/group-chat/src/controllers"
	"github.com/louissaadgo/group-chat/src/middlewares"
)

func Setup(app *fiber.App) {

	app.Static("/public", "./src/public")

	app.Use(middlewares.IsAuthenticated)

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login", fiber.Map{})
	})

	app.Use("/ws", middlewares.UpgradeToWS)

	controllers.RegisterWSEvents()

	app.Get("/", controllers.Home)

	app.Get("/ws", ikisocket.New(controllers.WS))

}
